package aibrain

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"sync"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/loggerx"
	pluginx "github.com/kubeshop/botkube/pkg/plugin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	PluginName  = "ai-brain"
	description = "Calls AI engine with incoming webhook prompts and streams the response."
	// TODO this should come from cloud config?
	assistantID = "asst_eMM9QaWLi6cajHE4PdG1yU53" // Botkube
)

var (
	// Compile-time check for interface impl.
	_ source.Source = (*Source)(nil)

	//go:embed config_schema.json
	ConfigJSONSchema string

	//go:embed webhook_schema.json
	IncomingWebhookJSONSchema string
)

// AI implements Botkube source plugin.
type Source struct {
	version   string
	instances sync.Map
	log       logrus.FieldLogger
}

// Config holds source configuration.
type Config struct {
	Log          config.Logger `yaml:"log"`
	OpenAIAPIKey string        `yaml:"openaiApiKey"`
	AssistantID  string        `yaml:"assistantId"`
}

type sourceInstance struct {
	cfg    *Config
	srcCtx source.CommonSourceContext
	log    logrus.FieldLogger

	out          chan<- source.Event
	openaiClient *openai.Client
	kubeClient   *dynamic.DynamicClient
	assistID     string
}

func NewSource(version string) *Source {
	return &Source{
		version:   version,
		instances: sync.Map{},
	}
}

// Metadata returns details about plugin.
func (s *Source) Metadata(context.Context) (api.MetadataOutput, error) {
	return api.MetadataOutput{
		Version:     s.version,
		Description: description,
		Recommended: true,
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
func (s *Source) Stream(_ context.Context, in source.StreamInput) (source.StreamOutput, error) {
	if err := pluginx.ValidateKubeConfigProvided(PluginName, in.Context.KubeConfig); err != nil {
		return source.StreamOutput{}, err
	}

	cfg, err := mergeConfigs(in.Configs)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while merging input configs: %w", err)
	}
	s.log = loggerx.New(cfg.Log)

	// Get kube client.
	kubeClient, err := getK8sClient(in.Context.KubeConfig)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while creating K8s clientset: %w", err)
	}

	sourceName := in.Context.SourceName

	streamOutput := source.StreamOutput{
		Event: make(chan source.Event),
	}

	s.instances.Store(sourceName, &sourceInstance{
		cfg:          cfg,
		log:          s.log,
		srcCtx:       in.Context.CommonSourceContext,
		out:          streamOutput.Event,
		openaiClient: openai.NewClient(cfg.OpenAIAPIKey),
		kubeClient:   kubeClient,
		assistID:     cfg.AssistantID,
	})

	s.log.Infof("Setup successful for source configuration %q", sourceName)
	return streamOutput, nil
}

// HandleExternalRequest handles incoming payload and returns an event based on it.
func (s *Source) HandleExternalRequest(ctx context.Context, in source.ExternalRequestInput) (source.ExternalRequestOutput, error) {
	s.log.Infof("handling external request for source: %s", in.Context.SourceName)
	instance, ok := s.instances.Load(in.Context.SourceName)
	if !ok {
		return source.ExternalRequestOutput{}, fmt.Errorf("source %q not found", in.Context.SourceName)
	}

	resp, err := instance.(*sourceInstance).handle(in)
	if err != nil {
		return source.ExternalRequestOutput{}, fmt.Errorf("while processing payload: %w", err)
	}

	return source.ExternalRequestOutput{
		Event: source.Event{
			Message: resp,
		},
	}, nil
}

// mergeConfigs merges the configuration.
func mergeConfigs(configs []*source.Config) (*Config, error) {
	defaults := &Config{
		AssistantID:  assistantID,
		OpenAIAPIKey: os.Getenv("OPENAI_API_KEY"),
		Log: config.Logger{
			Level:     "info",
			Formatter: "json",
		},
	}
	var cfg *Config
	if err := pluginx.MergeSourceConfigsWithDefaults(defaults, configs, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func getK8sClient(k8sBytes []byte) (*dynamic.DynamicClient, error) {
	kubeConfig, err := clientcmd.RESTConfigFromKubeConfig(k8sBytes)
	if err != nil {
		return nil, fmt.Errorf("while reading kube config: %v", err)
	}

	dynamicK8sCli, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("while creating dynamic K8s client: %w", err)
	}

	return dynamicK8sCli, nil
}
