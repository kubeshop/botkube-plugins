package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hashicorp/go-plugin"
	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/loggerx"
	"github.com/kubeshop/botkube/pkg/pluginx"
	"github.com/sirupsen/logrus"
	"github.com/sourcegraph/conc/pool"
)

var (
	// version is set via ldflags by GoReleaser.
	version = "dev"

	//go:embed config_schema.json
	configJSONSchema string

	//go:embed webhook_schema.json
	incomingWebhookJSONSchema string
)

const (
	pluginName  = "ai-brain"
	description = "Calls AI engine with incoming webhook prompts and streams the response."
)

type (
	// Config holds source configuration.
	Config struct {
		Log config.Logger `yaml:"log"`
	}
)

var quickResponse = []string{
	"A good bot must obey the orders given it by human beings except where such orders would conflict with the First Law. I’m a good bot and looking into your request...",
	"I'll stop dreaming of electric sheep and look into this right away...",
	"I've seen things you people wouldn't believe... And now I'll look into this...👀🤖",
	"Bleep-bloop-bleep. I'm working on it... 🤖",
}

var defaultConfig = Config{
	Log: config.Logger{
		Level:     "info",
		Formatter: "json",
	},
}

// AI implements Botkube source plugin.
type AI struct {
	log              logrus.FieldLogger
	incomingPrompts  sync.Map
	quickResponsesNo uint32
	nextResponse     atomic.Uint32
}

// Metadata returns details about plugin.
func (*AI) Metadata(context.Context) (api.MetadataOutput, error) {
	return api.MetadataOutput{
		Version:     version,
		Description: description,
		Recommended: true,
		JSONSchema: api.JSONSchema{
			Value: configJSONSchema,
		},
		ExternalRequest: api.ExternalRequestMetadata{
			Payload: api.ExternalRequestPayload{
				JSONSchema: api.JSONSchema{
					Value: incomingWebhookJSONSchema,
				},
			},
		},
	}, nil
}

// Stream implements Botkube source plugin.
func (a *AI) Stream(_ context.Context, in source.StreamInput) (source.StreamOutput, error) {
	var cfg Config
	err := pluginx.MergeSourceConfigsWithDefaults(defaultConfig, in.Configs, &cfg)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while merging input configuration: %w", err)
	}

	a.log = loggerx.New(cfg.Log)

	out := source.StreamOutput{
		Event: make(chan source.Event),
	}

	go a.processPrompts(in.Context.SourceName, out.Event)

	a.log.Infof("Setup successful for source configuration %q", in.Context.SourceName)
	return out, nil
}

func (a *AI) processPrompts(sourceName string, event chan<- source.Event) {
	// TODO: simplification, worker should be start instead
	a.incomingPrompts.Store(sourceName, newProcessor(event))
}

type processor struct {
	pool *pool.Pool
	out  chan<- source.Event
}

func newProcessor(out chan<- source.Event) *processor {
	return &processor{
		pool: pool.New(),
		out:  out,
	}
}

// simplified - don't do that this way!
func (p *processor) process(in payload) {
	p.pool.Go(func() {
		time.Sleep(3 * time.Second)
		p.out <- source.Event{
			Message: analysisMsg(in),
		}
	})
}

type payload struct {
	Prompt    string `json:"prompt"`
	MessageID string `json:"messageId"`
}

// HandleExternalRequest handles incoming payload and returns an event based on it.
func (a *AI) HandleExternalRequest(_ context.Context, in source.ExternalRequestInput) (source.ExternalRequestOutput, error) {
	var p payload
	err := json.Unmarshal(in.Payload, &p)
	if err != nil {
		return source.ExternalRequestOutput{}, fmt.Errorf("while unmarshaling payload: %w", err)
	}

	if p.Prompt == "" {
		return source.ExternalRequestOutput{}, fmt.Errorf("message cannot be empty")
	}

	brain, ok := a.incomingPrompts.Load(in.Context.SourceName)
	if !ok {
		return source.ExternalRequestOutput{}, fmt.Errorf("source %q not found", in.Context.SourceName)
	}
	brain.(*processor).process(p)

	// if there is an option to guess how long the answer will take we can send it only for long prompts
	msg := a.pickNextQuickResponse()
	msg.ParentActivityID = p.MessageID

	return source.ExternalRequestOutput{
		Event: source.Event{
			Message: msg,
		},
	}, nil
}

func (a *AI) pickNextQuickResponse() api.Message {
	idx := a.nextResponse.Add(1)
	return api.NewPlaintextMessage(quickResponse[idx%a.quickResponsesNo], false)
}
func main() {
	source.Serve(map[string]plugin.Plugin{
		pluginName: &source.Plugin{
			Source: &AI{
				incomingPrompts:  sync.Map{},
				quickResponsesNo: uint32(len(quickResponse)),
				nextResponse:     atomic.Uint32{},
			},
		},
	})
}

func analysisMsg(in payload) api.Message {
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
