package aibrain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

const openAIPollInterval = 2 * time.Second

// Payload represents incoming webhook payload.
type Payload struct {
	Prompt    string `json:"prompt"`
	MessageID string `json:"messageId"`
}

func (i *sourceInstance) handle(in source.ExternalRequestInput) (api.Message, error) {
	p := new(Payload)
	err := json.Unmarshal(in.Payload, p)
	if err != nil {
		return api.Message{}, fmt.Errorf("while unmarshalling payload: %w", err)
	}

	if p.Prompt == "" {
		return api.NewPlaintextMessage("Please clarify your question.", false), nil
	}

	go func() {
		if err := i.handleThread(context.Background(), p); err != nil {
			i.log.WithError(err).Error("failed to handle request")
		}
	}()

	return pickQuickResponse(p.MessageID), nil
}

// handleThread creates a new OpenAI assistant thread and handles the conversation.
func (i *sourceInstance) handleThread(ctx context.Context, p *Payload) error {
	// Start a new thread and run it.
	run, err := i.openaiClient.CreateThreadAndRun(ctx, openai.CreateThreadAndRunRequest{
		RunRequest: openai.RunRequest{
			AssistantID: i.assistID,
		},
		Thread: openai.ThreadRequest{
			Metadata: map[string]any{
				"messageId": p.MessageID,
			},
			Messages: []openai.ThreadMessage{
				{
					Role:    openai.ThreadMessageRoleUser,
					Content: p.Prompt,
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("while creating thread and run: %w", err)
	}

	for {
		// Wait a little bit before polling. OpenAI assistant api does not support streaming yet.
		time.Sleep(openAIPollInterval)

		// Get the run.
		run, err = i.openaiClient.RetrieveRun(ctx, run.ThreadID, run.ID)
		if err != nil {
			i.out <- source.Event{Message: msgUnableToHelp(p.MessageID)}
			return fmt.Errorf("while retrieving assistant thread run: %w", err)
		}

		i.log.WithFields(logrus.Fields{
			"messageId": p.MessageID,
			"runStatus": run.Status,
		}).Debug("retrieved assistant thread run")

		switch run.Status {
		case openai.RunStatusCancelling, openai.RunStatusFailed, openai.RunStatusExpired:
			i.out <- source.Event{Message: msgUnableToHelp(p.MessageID)}
			return nil

		case openai.RunStatusQueued, openai.RunStatusInProgress:
			continue

		// Fetch and return the response.
		case openai.RunStatusCompleted:
			if err = i.handleStatusCompleted(ctx, run, p); err != nil {
				i.out <- source.Event{Message: msgUnableToHelp(p.MessageID)}
				return fmt.Errorf("while handling completed case: %w", err)
			}
			return nil

		// The assistant is attempting to call a function.
		case openai.RunStatusRequiresAction:
			if err = i.handleStatusRequiresAction(ctx, run); err != nil {
				i.out <- source.Event{Message: msgUnableToHelp(p.MessageID)}
				return fmt.Errorf("while handling requires action: %w", err)
			}
		}
	}
}

func (i *sourceInstance) handleStatusCompleted(ctx context.Context, run openai.Run, p *Payload) error {
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
			Message: api.Message{
				ParentActivityID: p.MessageID,
				Sections: []api.Section{
					{
						Base: api.Base{
							Body: api.Body{Plaintext: "I am sorry, but I don't have a good answer."},
						},
					},
				},
			},
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
			Message: api.Message{
				ParentActivityID: p.MessageID,
				Sections: []api.Section{
					{
						Base: api.Base{
							Body: api.Body{Plaintext: c.Text.Value},
						},
						Context: []api.ContextItem{
							{Text: "AI-generated content may be incorrect."},
						},
					},
				},
			},
		}
	}

	return nil
}

func (i *sourceInstance) handleStatusRequiresAction(ctx context.Context, run openai.Run) error {
	// That should never happen, unless there is a bug or something	is wrong with OpenAI APIs.
	if run.RequiredAction == nil || run.RequiredAction.SubmitToolOutputs == nil {
		return errors.New("run.RequiredAction or run.RequiredAction.SubmitToolOutputs is nil, that should not happen")
	}

	toolOutputs := []openai.ToolOutput{}

	for _, t := range run.RequiredAction.SubmitToolOutputs.ToolCalls {
		if t.Type != openai.ToolTypeFunction {
			continue
		}

		switch t.Function.Name {
		case "kubectlGetPods":
			args := &kubectlGetPodsArgs{}
			if err := json.Unmarshal([]byte(t.Function.Arguments), args); err != nil {
				return err
			}

			out, err := kubectlGetPods(args)
			if err != nil {
				return err
			}

			toolOutputs = append(toolOutputs, openai.ToolOutput{
				ToolCallID: t.ID,
				Output:     string(out),
			})

		case "kubectlGetSecrets":
			args := &kubectlGetSecretsArgs{}
			if err := json.Unmarshal([]byte(t.Function.Arguments), args); err != nil {
				return err
			}

			out, err := kubectlGetSecrets(args)
			if err != nil {
				return err
			}

			toolOutputs = append(toolOutputs, openai.ToolOutput{
				ToolCallID: t.ID,
				Output:     string(out),
			})

		case "kubectlDescribePod":
			args := &kubectlDescribePodArgs{}
			if err := json.Unmarshal([]byte(t.Function.Arguments), args); err != nil {
				return err
			}

			out, err := kubectlDescribePod(args)
			if err != nil {
				return err
			}

			toolOutputs = append(toolOutputs, openai.ToolOutput{
				ToolCallID: t.ID,
				Output:     string(out),
			})

		case "kubectlLogs":
			args := &kubectlLogsArgs{}
			if err := json.Unmarshal([]byte(t.Function.Arguments), args); err != nil {
				return err
			}

			out, err := kubectlLogs(args)
			if err != nil {
				return err
			}

			toolOutputs = append(toolOutputs, openai.ToolOutput{
				ToolCallID: t.ID,
				Output:     string(out),
			})
		}
	}

	_, err := i.openaiClient.SubmitToolOutputs(ctx, run.ThreadID, run.ID, openai.SubmitToolOutputsRequest{
		ToolOutputs: toolOutputs,
	})
	if err != nil {
		return fmt.Errorf("while submitting tool outputs: %w", err)
	}

	return nil
}
