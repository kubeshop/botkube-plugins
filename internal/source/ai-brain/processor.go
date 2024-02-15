package aibrain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/sirupsen/logrus"
	"github.com/sourcegraph/conc/pool"
)

// Payload represents incoming webhook payload.
type Payload struct {
	Prompt    string `json:"prompt"`
	MessageID string `json:"messageId"`
}

type Processor struct {
	log                 logrus.FieldLogger
	pool                *pool.Pool
	out                 chan<- source.Event
	quickResponsePicker *QuickResponsePicker
}

func NewProcessor(log logrus.FieldLogger, out chan<- source.Event) *Processor {
	return &Processor{
		pool:                pool.New(),
		out:                 out,
		log:                 log.WithField("service", "ai-brain"),
		quickResponsePicker: NewQuickResponsePicker(),
	}
}

// Process is simplified - don't do that this way!
func (p *Processor) Process(rawPayload []byte) (api.Message, error) {
	var payload Payload
	p.log.WithField("req", string(rawPayload)).Debug("Handle external request...")

	err := json.Unmarshal(rawPayload, &payload)
	if err != nil {
		return api.Message{}, fmt.Errorf("while unmarshalling payload: %w", err)
	}

	if payload.Prompt == "" {
		return api.NewPlaintextMessage("Please specify your prompt", false), nil
	}

	p.pool.Go(func() {
		time.Sleep(3 * time.Second)
		p.out <- source.Event{
			Message: analysisMsg(payload),
		}
	})

	// If there is an option to guess how long the answer will take we can send it only for long prompts
	//
	// Consider: use channels to wait for a quick response from started goroutine with `time.After`
	// when time elapsed send quick response, otherwise sent the final response. Pseudocode:
	//
	//  select {
	//	case <-quickResponse:
	//	case <-time.After(5 * time.Second):
	//	   // here sent quick response and either inform started goroutine that event should be sent to p.out instead
	//	}
	//
	return p.quickResponsePicker.Pick(payload), nil
}

func analysisMsg(in Payload) api.Message {
	btnBldr := api.NewMessageButtonBuilder()
	return api.Message{
		ParentActivityID: in.MessageID,
		Sections: []api.Section{
			{
				Base: api.Base{
					Header: ":warning: Detected Issues",
					Body: api.Body{
						Plaintext: `I looks like a Pod named "nginx" in the "botkube" namespace is failing to pull the "nginx2" image due to an "ErrImagePull" error. The error indicates insufficient scope or authorization failure for the specified image repository. Ensure the image exists, check authentication, and verify network connectivity.`,
					},
				},
			},
			{
				Base: api.Base{
					Body: api.Body{
						Plaintext: "After resolving, delete the pod with:",
					},
				},
				Buttons: []api.Button{
					btnBldr.ForCommandWithDescCmd("Restart pod", "kubectl delete pod -n botkube nginx", api.ButtonStyleDanger),
				},
			},
			{
				Context: []api.ContextItem{
					{
						Text: "AI-generated content may be incorrect.",
					},
				},
			},
		},
	}
}
