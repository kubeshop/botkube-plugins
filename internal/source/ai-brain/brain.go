package aibrain

import (
	"context"
	_ "embed"
	"fmt"
	"sync"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/kubeshop/botkube/pkg/loggerx"
	pluginx "github.com/kubeshop/botkube/pkg/plugin"
	"github.com/sirupsen/logrus"
)

const (
	PluginName  = "ai-brain"
	description = "Calls AI engine with incoming webhook prompts and streams the response."
)

var (
	// Compile-time check for interface impl.
	_ source.Source = (*Source)(nil)

	//go:embed config_schema.json
	ConfigJSONSchema string

	//go:embed webhook_schema.json
	IncomingWebhookJSONSchema string
)

// Source implements AI source plugin.
type Source struct {
	version   string
	instances sync.Map
	log       logrus.FieldLogger
}

// NewSource creates new source plugin instance.
func NewSource(version string) *Source {
	return &Source{
		version:   version,
		instances: sync.Map{},
	}
}

// Metadata returns details about plugin.
func (s *Source) Metadata(context.Context) (api.MetadataOutput, error) {
	return api.MetadataOutput{
		Version:      s.version,
		Description:  description,
		Recommended:  true,
		Dependencies: binaryDependencies(),
		JSONSchema: api.JSONSchema{
			Value: ConfigJSONSchema,
		},
		ExternalRequest: api.ExternalRequestMetadata{
			Payload: api.ExternalRequestPayload{
				JSONSchema: api.JSONSchema{
					Value: IncomingWebhookJSONSchema,
				},
			},
		},
	}, nil
}

// Stream implements Botkube source plugin.
func (s *Source) Stream(ctx context.Context, in source.StreamInput) (source.StreamOutput, error) {
	if err := pluginx.ValidateKubeConfigProvided(PluginName, in.Context.KubeConfig); err != nil {
		return source.StreamOutput{}, err
	}

	kubeConfigPath, _, err := pluginx.PersistKubeConfig(ctx, in.Context.KubeConfig)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while writing kubeconfig file: %w", err)
	}

	cfg, err := mergeConfigs(in.Configs)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while merging input configs: %w", err)
	}
	s.log = loggerx.New(cfg.Log)

	sourceName := in.Context.SourceName

	streamOutput := source.StreamOutput{
		Event: make(chan source.Event),
	}

	instance := newAssistant(cfg, s.log, streamOutput.Event, kubeConfigPath)
	s.instances.Store(sourceName, instance)

	s.log.Infof("Setup successful for source configuration %q", sourceName)
	return streamOutput, nil
}

// HandleExternalRequest handles incoming payload and returns an event based on it.
func (s *Source) HandleExternalRequest(_ context.Context, in source.ExternalRequestInput) (source.ExternalRequestOutput, error) {
	s.log.Infof("Handling external request for source: %s", in.Context.SourceName)
	instance, ok := s.instances.Load(in.Context.SourceName)
	if !ok {
		return source.ExternalRequestOutput{}, fmt.Errorf("source %q not found", in.Context.SourceName)
	}

	resp, err := instance.(*assistant).handle(in)
	if err != nil {
		return source.ExternalRequestOutput{}, fmt.Errorf("while processing payload: %w", err)
	}

	return source.ExternalRequestOutput{
		Event: source.Event{
			Message: resp,
		},
	}, nil
}
