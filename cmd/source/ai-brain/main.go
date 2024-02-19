package main

import (
	"context"
	_ "embed"
	"fmt"
	"sync"

	"github.com/hashicorp/go-plugin"
	aibrain "github.com/kubeshop/botkube-cloud-plugins/internal/source/ai-brain"
	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/kubeshop/botkube/pkg/loggerx"
	"github.com/sirupsen/logrus"
)

// version is set via ldflags by GoReleaser.
var version = "dev"

const (
	pluginName  = "ai-brain"
	description = "Calls AI engine with incoming webhook prompts and streams the response."
)

// AI implements Botkube source plugin.
type AI struct {
	incomingPrompts sync.Map
}

// Metadata returns details about plugin.
func (*AI) Metadata(context.Context) (api.MetadataOutput, error) {
	return api.MetadataOutput{
		Version:     version,
		Description: description,
		Recommended: true,
		JSONSchema: api.JSONSchema{
			Value: aibrain.ConfigJSONSchema,
		},
		ExternalRequest: api.ExternalRequestMetadata{
			Payload: api.ExternalRequestPayload{
				JSONSchema: api.JSONSchema{
					Value: aibrain.IncomingWebhookJSONSchema,
				},
			},
		},
	}, nil
}

// Stream implements Botkube source plugin.
func (a *AI) Stream(_ context.Context, in source.StreamInput) (source.StreamOutput, error) {
	cfg, err := aibrain.MergeConfigs(in.Configs)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while merging configuration: %w", err)
	}

	log := loggerx.New(cfg.Log)
	out := source.StreamOutput{
		Event: make(chan source.Event),
	}
	go a.processPrompts(in.Context.SourceName, out.Event, log)

	log.Infof("Setup successful for source configuration %q", in.Context.SourceName)
	return out, nil
}

func (a *AI) processPrompts(sourceName string, event chan<- source.Event, log logrus.FieldLogger) {
	a.incomingPrompts.Store(sourceName, aibrain.NewProcessor(log, event))
}

// HandleExternalRequest handles incoming payload and returns an event based on it.
func (a *AI) HandleExternalRequest(_ context.Context, in source.ExternalRequestInput) (source.ExternalRequestOutput, error) {
	brain, ok := a.incomingPrompts.Load(in.Context.SourceName)
	if !ok {
		return source.ExternalRequestOutput{}, fmt.Errorf("source %q not found", in.Context.SourceName)
	}
	quickResponse, err := brain.(*aibrain.Processor).Process(in.Payload)
	if err != nil {
		return source.ExternalRequestOutput{}, fmt.Errorf("while processing payload: %w", err)
	}

	return source.ExternalRequestOutput{
		Event: source.Event{
			Message: quickResponse,
		},
	}, nil
}

func main() {
	source.Serve(map[string]plugin.Plugin{
		pluginName: &source.Plugin{
			Source: &AI{
				incomingPrompts: sync.Map{},
			},
		},
	})
}
