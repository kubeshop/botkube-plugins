package aibrain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kubeshop/botkube-cloud-plugins/internal/remote"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	cacheTTL                = 8 * time.Hour
	openAIPollInterval      = 2 * time.Second
	maxToolExecutionRetries = 3
)

type tool func(ctx context.Context, args []byte) (string, error)

// Payload represents incoming webhook payload.
type Payload struct {
	Prompt    string `json:"prompt"`
	MessageID string `json:"messageId"`
}

type assistant struct {
	log          logrus.FieldLogger
	out          chan<- source.Event
	openaiClient *openai.Client
	assistID     string
	tools        map[string]tool
	cache        *cache
}

func newAssistant(cfg *Config, log logrus.FieldLogger, out chan source.Event, kubeConfigPath string) *assistant {
	kcRunner := NewKubectlRunner(kubeConfigPath)

	config := openai.DefaultConfig("")
	config.HTTPClient = &http.Client{
		Transport: newAPIKeySecuredTransport(),
	}
	config.BaseURL = cfg.OpenAICloudServiceURL

	return &assistant{
		log:          log,
		out:          out,
		openaiClient: openai.NewClientWithConfig(config),
		assistID:     cfg.OpenAIAssistantID,
		cache:        newCache(cacheTTL),
		tools: map[string]tool{
			"kubectlGetPods":     kcRunner.GetPods,
			"kubectlGetSecrets":  kcRunner.GetSecrets,
			"kubectlDescribePod": kcRunner.DescribePod,
			"kubectlLogs":        kcRunner.Logs,
		},
	}
}

func (i *assistant) handle(in source.ExternalRequestInput) (api.Message, error) {
	var p Payload

	err := json.Unmarshal(in.Payload, &p)
	if err != nil {
		return api.Message{}, fmt.Errorf("while unmarshalling payload: %w", err)
	}

	if p.Prompt == "" {
		return api.NewPlaintextMessage("Please clarify your question.", false), nil
	}

	go func() {
		if err := i.handleThread(context.Background(), &p); err != nil {
			// TODO: It would be great to send the user prompt and error message
			// back to us for analysis and potential fixing, enhancing our prompt.
			// can we do that @Blair?
			i.out <- source.Event{Message: msgUnableToHelp(p.MessageID)}
			i.log.WithError(err).Error("failed to handle request")
		}
	}()

	return pickQuickResponse(p.MessageID), nil
}

// handleThread creates a new OpenAI assistant thread and handles the conversation.
func (i *assistant) handleThread(ctx context.Context, p *Payload) error {
	var thread openai.Thread
	var err error

	// First we check if we have a cached thread ID for the given message ID.
	threadID, ok := i.cache.Get(p.MessageID)
	if ok {
		err = i.createNewMessage(ctx, threadID, p)
		if err != nil {
			return fmt.Errorf("while creating a new message on a thread: %w", err)
		}
	} else {
		thread, err = i.createNewThread(ctx, p)
		if err != nil {
			return fmt.Errorf("while creating a new thread: %w", err)
		}
		threadID = thread.ID
	}

	i.cache.Set(p.MessageID, threadID)

	run, err := i.openaiClient.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: i.assistID,
	})
	if err != nil {
		return fmt.Errorf("while creating a thread run: %w", err)
	}

	toolsRetries := 0
	return wait.PollUntilContextCancel(ctx, openAIPollInterval, false, func(ctx context.Context) (bool, error) {
		run, err = i.openaiClient.RetrieveRun(ctx, run.ThreadID, run.ID)
		if err != nil {
			return false, fmt.Errorf("while retrieving assistant thread run: %w", err)
		}

		i.log.WithFields(logrus.Fields{
			"messageId": p.MessageID,
			"runStatus": run.Status,
		}).Debug("retrieved assistant thread run")

		switch run.Status {
		case openai.RunStatusCancelling, openai.RunStatusFailed:
			return false, fmt.Errorf("got unexpected status: %s", run.Status)

		case openai.RunStatusExpired:
			i.cache.Delete(p.MessageID)
			return true, nil

		case openai.RunStatusQueued, openai.RunStatusInProgress:
			return false, nil // continue

		case openai.RunStatusCompleted:
			if err = i.handleStatusCompleted(ctx, run, p); err != nil {
				return false, fmt.Errorf("while handling completed case: %w", err)
			}
			return true, nil // success

		case openai.RunStatusRequiresAction:
			if err = i.handleStatusRequiresAction(ctx, run); err != nil {
				toolsRetries++
				return toolsRetries >= maxToolExecutionRetries, fmt.Errorf("while handling requires action: %w", err)
			}
		}
		return false, nil
	})
}

func (i *assistant) createNewThread(ctx context.Context, p *Payload) (openai.Thread, error) {
	return i.openaiClient.CreateThread(ctx, openai.ThreadRequest{
		Metadata: map[string]any{
			"messageId":  p.MessageID,
			"instanceId": os.Getenv(remote.ProviderIdentifierEnvKey),
		},
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleUser,
				Content: p.Prompt,
			},
		},
	})
}

func (i *assistant) createNewMessage(ctx context.Context, threadID string, p *Payload) error {
	_, err := i.openaiClient.CreateMessage(ctx, threadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: p.Prompt,
	})
	return err
}

func (i *assistant) handleStatusCompleted(ctx context.Context, run openai.Run, p *Payload) error {
	limit := 1
	msgList, err := i.openaiClient.ListMessage(ctx, run.ThreadID, &limit, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("while getting assistant messages response: %w", err)
	}

	// We're listing messages in a thread. They are ordered in desc order. If
	// there are no messages in the entire thread, we imply that something went
	// wrong on the OpenAI side, could be a bug.
	if len(msgList.Messages) == 0 {
		i.log.Debug("no response messages were found, that seems like an edge case.")
		i.out <- source.Event{
			Message: msgNoAIAnswer(p.MessageID),
		}
		return nil
	}

	// Iterate over text content to build messages. We're only interested in text
	// context, since the assistant is instructed to return text only.
	for _, c := range msgList.Messages[0].Content {
		if c.Text == nil {
			continue
		}

		i.out <- source.Event{
			Message: msgAIAnswer(p.MessageID, c.Text.Value),
		}
	}

	return nil
}

func (i *assistant) handleStatusRequiresAction(ctx context.Context, run openai.Run) error {
	// That should never happen, unless there is a bug or something	is wrong with OpenAI APIs.
	if run.RequiredAction == nil || run.RequiredAction.SubmitToolOutputs == nil {
		return errors.New("run.RequiredAction or run.RequiredAction.SubmitToolOutputs is nil, that should not happen")
	}

	var toolOutputs []openai.ToolOutput
	for _, t := range run.RequiredAction.SubmitToolOutputs.ToolCalls {
		if t.Type != openai.ToolTypeFunction {
			continue
		}

		doer, found := i.tools[t.Function.Name]
		if !found {
			continue
		}

		out, err := doer(ctx, []byte(t.Function.Arguments))
		if err != nil {
			return err
		}

		toolOutputs = append(toolOutputs, openai.ToolOutput{
			ToolCallID: t.ID,
			Output:     out,
		})
	}

	_, err := i.openaiClient.SubmitToolOutputs(ctx, run.ThreadID, run.ID, openai.SubmitToolOutputsRequest{
		ToolOutputs: toolOutputs,
	})
	if err != nil {
		return fmt.Errorf("while submitting tool outputs: %w", err)
	}

	return nil
}

type apiKeySecuredTransport struct {
	transport *http.Transport
}

func newAPIKeySecuredTransport() *apiKeySecuredTransport {
	return &apiKeySecuredTransport{
		transport: http.DefaultTransport.(*http.Transport).Clone(),
	}
}

func (t *apiKeySecuredTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-Key", os.Getenv(remote.ProviderAPIKeyEnvKey))
	return t.transport.RoundTrip(req)
}
