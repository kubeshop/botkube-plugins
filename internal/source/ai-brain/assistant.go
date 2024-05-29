package aibrain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kubeshop/botkube/pkg/ptr"

	"github.com/kubeshop/botkube-cloud-plugins/internal/otelx"
	"github.com/kubeshop/botkube-cloud-plugins/internal/remote"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	cacheTTL                  = 8 * time.Hour
	openAIPollInterval        = 2 * time.Second
	maxToolExecutionRetries   = 3
	quotaExceededErrCode      = "quota_exceeded"
	tracerName                = "source.aibrain"
	serviceName               = "botkube-plugins-source-ai-brain"
	clusterScanSubcommandName = "scan"

	clusterScanPrompt = `
Scan the Kubernetes cluster for critical issues that could significantly impact the cluster's health, stability, or security.
Focus on problems that may not be immediately apparent through events or standard monitoring.

Provide a concise overview of the scan results, including the total number of
critical issues found. If there were no issues found for a specific check, do
not include that section in the report. List the Kubernetes objects directly
affected by the issue. Make sure that your checks are relevant to the current
state of the cluster, do not include resources that no longer exist.

Summary section needs to be at the top of the report, followed by specific checks.

Specific Checks:

Pod Health:
Identify pods in a crash-loop backoff state with a high restart count.
Identify pods that have been OOMKilled (Out of Memory Killed) multiple times.
Look for pods stuck in a pending state for an extended period.
Resource Utilization:
Identify nodes or pods with critically high CPU or memory usage (e.g., above 90% of limits).
Check for critical resource starvation issues affecting multiple pods or namespaces.
Configuration:
Look for pods running with very insecure capabilities (e.g., ALL, NET_RAW, SYS_ADMIN).
Identify pods using deprecated or insecure container images.
Check for misconfigured network policies that could expose sensitive services.
Networking:
Identify pods or services experiencing significant network latency or packet loss.
Check for network partitions or connectivity issues between critical components.

Additional Guidance for the LLM Agent:

Prioritize issues that pose the most immediate threat to the cluster's stability, performance, or security.
Filter out informational or low-severity issues that are unlikely to cause major problems.
Be as specific as possible in the descriptions. Do not exceed 2000 characters in your response.
`
)

var temperature float32 = 0.1

type tool func(ctx context.Context, args []byte, p *Payload) (string, error)

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
	tracer       trace.Tracer
}

func newAssistant(cfg *Config, log logrus.FieldLogger, out chan source.Event, kubeConfigPath string) (*assistant, error) {
	if cfg.HoneycombAPIKey != "" {
		log.Debug("Setting up opentelemetry with honeycomb")
		_, err := otelx.Init(cfg.HoneycombAPIKey, serviceName, cfg.HoneycombSampleRate, cfg.Version)
		if err != nil {
			log.WithError(err).Error("failed to initilise otel with honeycomb")
		}
	}
	tracer := otel.Tracer(serviceName)

	kcRunner := NewKubectlRunner(kubeConfigPath, tracer)
	bkRunner, err := NewBotkubeRunner(tracer)
	if err != nil {
		return nil, fmt.Errorf("while creating Botkube runner: %w", err)
	}

	config := openai.DefaultConfig("")
	config.HTTPClient = &http.Client{
		Transport: newAPIKeySecuredTransport(),
	}
	config.BaseURL = cfg.OpenAICloudServiceURL
	config.AssistantVersion = "v2"

	return &assistant{
		log:          log,
		out:          out,
		tracer:       tracer,
		openaiClient: openai.NewClientWithConfig(config),
		assistID:     cfg.OpenAIAssistantID,
		cache:        newCache(cacheTTL),
		tools: map[string]tool{
			"kubectlGetResource":                  kcRunner.GetResource,
			"kubectlDescribeResource":             kcRunner.DescribeResource,
			"kubectlGetEvents":                    kcRunner.GetEvents,
			"kubectlTopPods":                      kcRunner.TopPods,
			"kubectlTopNodes":                     kcRunner.TopNodes,
			"kubectlLogs":                         kcRunner.Logs,
			"botkubeGetStartupAgentConfiguration": bkRunner.GetStartupAgentConfiguration,
			"botkubeGetAgentStatus":               bkRunner.GetAgentStatus,
		},
	}, nil
}

func (i *assistant) handle(ctx context.Context, in source.ExternalRequestInput) (api.Message, error) {
	var p Payload

	err := json.Unmarshal(in.Payload, &p)
	if err != nil {
		return api.Message{}, fmt.Errorf("while unmarshalling payload: %w", err)
	}

	if p.Prompt == "" {
		return api.NewPlaintextMessage("Please clarify your question.", false), nil
	}

	if strings.ToLower(p.Prompt) == clusterScanSubcommandName {
		p.Prompt = clusterScanPrompt
	}

	go func() {
		if err := i.handleThread(ctx, &p); err != nil {
			i.out <- source.Event{Message: i.handleThreadError(p.MessageID, err)}
		}
	}()

	return pickQuickResponse(p.MessageID), nil
}

// TODO: Send the user prompt and error message back to us for analysis and potential fixing, enhancing our prompt.
// Our privacy policy allows us to do so.
func (i *assistant) handleThreadError(messageID string, err error) api.Message {
	log := i.log.WithError(err).WithField("messageID", messageID)

	var openAIErr *openai.APIError
	ok := errors.As(err, &openAIErr)
	if ok && openAIErr != nil && openAIErr.Code == quotaExceededErrCode && openAIErr.Type == quotaExceededErrCode {
		log.Info("Handling quota exceeded error")
		return msgQuotaExceeded(messageID)
	}

	log.Error("Failed to handle user prompt")
	return msgUnableToHelp(messageID)
}

// handleThread creates a new OpenAI assistant thread and handles the conversation.
func (i *assistant) handleThread(ctx context.Context, p *Payload) (err error) {
	ctx, span := i.tracer.Start(ctx, "aibrain.assistant.handleThread")
	defer span.End()

	span.SetAttributes(attribute.String("payload.prompt", p.Prompt))

	var thread openai.Thread

	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
	}()

	// First we check if we have a cached thread ID for the given message ID.
	threadID, ok := i.cache.Get(p.MessageID)
	if ok {
		span.AddEvent("Found an existing thread in the internal cache.")
		err = i.createNewMessage(ctx, threadID, p)
		if err != nil {
			return fmt.Errorf("while creating a new message on a thread: %w", err)
		}
	} else {
		span.AddEvent("Existing thread was not found in the internal cache.")
		thread, err = i.createNewThread(ctx, p)
		if err != nil {
			return fmt.Errorf("while creating a new thread: %w", err)
		}
		threadID = thread.ID
	}

	span.SetAttributes(
		attribute.String("openai.threadId", threadID),
		attribute.String("payload.messageId", p.MessageID),
	)

	i.cache.Set(p.MessageID, threadID)

	log := i.log.WithFields(logrus.Fields{
		"messageId": p.MessageID,
		"threadId":  threadID,
		"prompt":    p.Prompt,
	})
	run, err := i.openaiClient.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: i.assistID,
		Temperature: &temperature,
	})
	if err != nil {
		return fmt.Errorf("while creating a thread run: %w", err)
	}

	log = log.WithField("runId", run.ID)
	log.Info("created a new assistant run")

	toolsRetries := 0
	return wait.PollUntilContextCancel(ctx, openAIPollInterval, false, func(ctx context.Context) (bool, error) {
		run, err = i.openaiClient.RetrieveRun(ctx, run.ThreadID, run.ID)
		if err != nil {
			log.WithError(err).Error("failed to retrieve assistant thread run")
			return false, fmt.Errorf("while retrieving assistant thread run: %w", err)
		}

		log.WithFields(logrus.Fields{
			"messageId": p.MessageID,
			"runStatus": run.Status,
		}).Debug("retrieved assistant thread run")

		switch run.Status {
		case openai.RunStatusCancelling, openai.RunStatusFailed:
			return true, fmt.Errorf("got unexpected status: %s", run.Status)

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
			if err = i.handleStatusRequiresAction(ctx, run, p); err != nil {
				toolsRetries++
				return toolsRetries >= maxToolExecutionRetries, fmt.Errorf("while handling requires action: %w", err)
			}
		}
		return false, nil
	})
}

func (i *assistant) createNewThread(ctx context.Context, p *Payload) (openai.Thread, error) {
	ctx, span := i.tracer.Start(ctx, "aibrain.assistant.createNewThread")
	defer span.End()

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
	ctx, span := i.tracer.Start(ctx, "aibrain.assistant.createNewMessage")
	defer span.End()

	_, err := i.openaiClient.CreateMessage(ctx, threadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: p.Prompt,
	})
	return err
}

func (i *assistant) handleStatusCompleted(ctx context.Context, run openai.Run, p *Payload) error {
	ctx, span := i.tracer.Start(ctx, "aibrain.assistant.handleStatusCompleted")
	defer span.End()

	limit := 1
	msgList, err := i.openaiClient.ListMessage(ctx, run.ThreadID, &limit, nil, nil, nil)
	if err != nil {
		err = fmt.Errorf("while getting assistant messages response: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
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

	toolCalls, err := i.listToolCalls(ctx, run.ThreadID, run.ID)
	if err != nil {
		err = fmt.Errorf("while listing tool calls: %w", err)
		span.RecordError(err)
		i.log.WithError(err).Error("failed to get tool calls")
		// not returning as the tool usage is optional
	}

	// Iterate over text content to build messages. We're only interested in text
	// context, since the assistant is instructed to return text only.
	for _, c := range msgList.Messages[0].Content {
		if c.Text == nil {
			continue
		}

		textValue := i.trimCitationsIfPresent(i.log, c.Text)

		i.out <- source.Event{
			Message: msgAIAnswer(run, p, textValue, toolCalls),
		}
	}

	return nil
}

func (i *assistant) handleStatusRequiresAction(ctx context.Context, run openai.Run, p *Payload) error {
	ctx, span := i.tracer.Start(ctx, "aibrain.assistant.handleStatusRequiresAction")
	defer span.End()

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
			i.log.WithField("toolName", t.Function.Name).Warn("AI tool not found")
			continue
		}

		out, err := doer(ctx, []byte(t.Function.Arguments), p)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
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
		err = fmt.Errorf("while submitting tool outputs: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}

const fileCitationAnnotation = "file_citation"

func (i *assistant) trimCitationsIfPresent(log logrus.FieldLogger, in *openai.MessageText) string {
	if in == nil {
		return ""
	}

	if len(in.Annotations) == 0 {
		return in.Value
	}

	outValue := in.Value
	for _, a := range in.Annotations {
		annotation, ok := a.(map[string]interface{})
		if !ok {
			continue
		}

		if annotation["type"] != fileCitationAnnotation {
			continue
		}

		text, ok := annotation["text"].(string)
		if !ok {
			continue
		}

		log.Debugf("Citation found. Trimming %q...", text)
		outValue = strings.Replace(outValue, text, "", 1)
	}

	return outValue
}

func (i *assistant) listToolCalls(ctx context.Context, threadID, runID string) (map[string]struct{}, error) {
	const maxLimit = 100
	var (
		runSteps []openai.RunStep
		after    *string
	)

	for {
		res, err := i.openaiClient.ListRunSteps(ctx, threadID, runID, openai.Pagination{Limit: ptr.FromType(maxLimit), After: after})
		if err != nil {
			return nil, fmt.Errorf("while getting assistant run steps: %w", err)
		}

		runSteps = append(runSteps, res.RunSteps...)
		if !res.HasMore {
			break
		}

		after = &res.LastID
		i.log.WithFields(logrus.Fields{
			"lastID": res.LastID,
			"count":  len(runSteps),
		}).Debug("Paginating assistant run steps...")
	}
	i.log.WithField("count", len(runSteps)).Debug("Finished paginating assistant run steps")

	return getFriendlyToolCallsFromRunSteps(runSteps), nil
}

type apiKeySecuredTransport struct {
	transport http.RoundTripper
}

func newAPIKeySecuredTransport() *apiKeySecuredTransport {
	return &apiKeySecuredTransport{
		transport: otelhttp.NewTransport(http.DefaultTransport),
	}
}

func (t *apiKeySecuredTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-Key", os.Getenv(remote.ProviderAPIKeyEnvKey))
	return t.transport.RoundTrip(req)
}

func instanceID() string {
	return os.Getenv(remote.ProviderIdentifierEnvKey)
}
