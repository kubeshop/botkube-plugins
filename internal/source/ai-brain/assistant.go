package aibrain

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

// Payload represents incoming webhook payload.
type Payload struct {
	Prompt    string `json:"prompt"`
	MessageID string `json:"messageId"`
}

// handle is simplified - don't do that this way!
func (i *sourceInstance) handle(in source.ExternalRequestInput) (api.Message, error) {
	p := new(Payload)
	err := json.Unmarshal(in.Payload, p)
	if err != nil {
		return api.Message{}, fmt.Errorf("while unmarshalling payload: %w", err)
	}

	// TODO: why is the Prompt prefixed with `ai-face`?
	if p.Prompt == "ai-face" {
		return api.NewPlaintextMessage("Please clarify your question.", false), nil
	}

	// Cleanup the prompt.
	p.Prompt = strings.TrimPrefix(p.Prompt, "ai-face")

	// TODO: needs better goroutine management with persistent thread mapping.
	go func() {
		_ = i.handleThread(context.Background(), p)
	}()

	return api.Message{
		ParentActivityID: p.MessageID,
		Sections: []api.Section{
			{
				// TODO: remove?
				Base: api.Base{
					Body: api.Body{Plaintext: "Let me figure this out.."},
				},
			},
		},
	}, nil
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
		time.Sleep(2 * time.Second)

		// Get the run.
		run, err = i.openaiClient.RetrieveRun(ctx, run.ThreadID, run.ID)
		if err != nil {
			i.log.WithError(err).Error("while retrieving assistant thread run")
			continue
		}

		i.log.WithFields(logrus.Fields{
			"messageId": p.MessageID,
			"runStatus": run.Status,
		}).Debug("retrieved assistant thread run")

		switch run.Status {
		case openai.RunStatusCancelling, openai.RunStatusFailed, openai.RunStatusExpired:
			// TODO tell the user that the assistant has stopped processing the request.
			continue

		// We have to wait. Here we could tell the user that we are waiting.
		case openai.RunStatusQueued, openai.RunStatusInProgress:
			continue

		// Fetch and return the response.
		case openai.RunStatusCompleted:
			if err = i.handleStatusCompleted(ctx, run, p); err != nil {
				i.log.WithError(err).Error("while handling completed case")
				continue
			}
			return nil

		// The assistant is attempting to call a function.
		case openai.RunStatusRequiresAction:
			if err = i.handleStatusRequiresAction(ctx, run); err != nil {
				return fmt.Errorf("while handling requires action: %w", err)
			}
		}
	}
}

func (i *sourceInstance) handleStatusCompleted(ctx context.Context, run openai.Run, p *Payload) error {
	msgList, err := i.openaiClient.ListMessage(ctx, run.ThreadID, nil, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("while getting assistant messages response")
	}

	if len(msgList.Messages) == 0 {
		i.log.Debug("no response messages were found")
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

	i.out <- source.Event{
		Message: api.Message{
			ParentActivityID: p.MessageID,
			Sections: []api.Section{
				{
					Base: api.Base{
						Body: api.Body{Plaintext: msgList.Messages[0].Content[0].Text.Value},
					},
				},
			},
		},
	}
	return nil
}

func (i *sourceInstance) handleStatusRequiresAction(ctx context.Context, run openai.Run) error {
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
			// Submit tool output.
			_, err = i.openaiClient.SubmitToolOutputs(ctx, run.ThreadID, run.ID, openai.SubmitToolOutputsRequest{
				ToolOutputs: []openai.ToolOutput{
					{
						ToolCallID: t.ID,
						Output:     string(out),
					},
				},
			})
			if err != nil {
				return err
			}

		case "kubectlGetSecrets":
			args := &kubectlGetSecretsArgs{}
			if err := json.Unmarshal([]byte(t.Function.Arguments), args); err != nil {
				return err
			}

			out, err := kubectlGetSecrets(args)
			if err != nil {
				return err
			}
			// Submit tool output.
			_, err = i.openaiClient.SubmitToolOutputs(ctx, run.ThreadID, run.ID, openai.SubmitToolOutputsRequest{
				ToolOutputs: []openai.ToolOutput{
					{
						ToolCallID: t.ID,
						Output:     string(out),
					},
				},
			})
			if err != nil {
				return err
			}

		case "kubectlDescribePod":
			args := &kubectlDescribePodArgs{}
			if err := json.Unmarshal([]byte(t.Function.Arguments), args); err != nil {
				return err
			}

			out, err := kubectlDescribePod(args)
			if err != nil {
				return err
			}
			// Submit tool output.
			_, err = i.openaiClient.SubmitToolOutputs(ctx, run.ThreadID, run.ID, openai.SubmitToolOutputsRequest{
				ToolOutputs: []openai.ToolOutput{
					{
						ToolCallID: t.ID,
						Output:     string(out),
					},
				},
			})
			if err != nil {
				return err
			}

		case "kubectlLogs":
			args := &kubectlLogsArgs{}
			if err := json.Unmarshal([]byte(t.Function.Arguments), args); err != nil {
				return err
			}

			out, err := kubectlLogs(args)
			if err != nil {
				return err
			}
			// Submit tool output.
			_, err = i.openaiClient.SubmitToolOutputs(ctx, run.ThreadID, run.ID, openai.SubmitToolOutputsRequest{
				ToolOutputs: []openai.ToolOutput{
					{
						ToolCallID: t.ID,
						Output:     string(out),
					},
				},
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
