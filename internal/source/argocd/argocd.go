package argocd

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/kubeshop/botkube-cloud-plugins/internal/auth"

	"github.com/avast/retry-go/v4"
	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/formatx"
	"github.com/kubeshop/botkube/pkg/loggerx"
	pluginx "github.com/kubeshop/botkube/pkg/plugin"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var _ source.Source = (*Source)(nil)

var (
	//go:embed config_schema.json
	configJSONSchema string

	//go:embed req-jsonschema.json
	requestJSONSchema string

	//go:embed default-config.yaml
	defaultConfigYAML string
)

const (
	// PluginName is the name of the source plugin.
	PluginName = "argocd"

	description = "Argo source plugin is used to get ArgoCD trigger-based notifications."
)

var htmlDescription = fmt.Sprintf(`%s<br/><strong>Prerequisite:</strong> Configure RBAC resources with specific permissions. <a href="https://docs.botkube.io/configuration/source/argocd#prerequisite-elevated-rbac-permissions" target="_blank">Read more</a>`, description)

type sourceInstance struct {
	cfg    Config
	srcCtx source.CommonSourceContext
}

// Source defines ArgoCD source plugin.
type Source struct {
	pluginVersion string
	log           logrus.FieldLogger
	cfgs          sync.Map
}

// NewSource returns a new instance of Source.
func NewSource(version string) source.Source {
	src := &Source{
		pluginVersion: version,
		cfgs:          sync.Map{},
	}
	return auth.NewProtectedSource(src)
}

type subscription struct {
	TriggerName string
	WebhookName string
	Application config.K8sResourceRef
}

// Stream set-ups ArgoCD notifications.
func (s *Source) Stream(ctx context.Context, input source.StreamInput) (source.StreamOutput, error) {
	if err := pluginx.ValidateKubeConfigProvided(PluginName, input.Context.KubeConfig); err != nil {
		return source.StreamOutput{}, err
	}

	cfg, err := mergeConfigs(input.Configs)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while merging input configs: %w", err)
	}

	s.log = loggerx.New(cfg.Log)

	sourceName := input.Context.SourceName
	s.cfgs.Store(sourceName, sourceInstance{
		cfg:    cfg,
		srcCtx: input.Context.CommonSourceContext,
	})

	k8sCli, err := s.getK8sClient(input.Context.KubeConfig)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while creating K8s clientset: %w", err)
	}

	s.log.Info("Preparing configuration...")

	err = retry.Do(
		func() error {
			return s.setupArgoNotifications(ctx, k8sCli, sourceInstance{
				cfg:    cfg,
				srcCtx: input.Context.CommonSourceContext,
			})
		},
		retry.OnRetry(func(n uint, err error) {
			s.log.WithField("error", err).Errorf("Error setting up Argo notifications for %q. Retrying...", sourceName)
		}),
		retry.DelayType(retry.RandomDelay), // Randomize the retry time as ConfigMap is updated and there might be conflicts when there are multiple plugin configurations
		retry.MaxJitter(5*time.Second),
		retry.Attempts(5),
		retry.LastErrorOnly(false),
	)
	if err != nil {
		return source.StreamOutput{}, fmt.Errorf("while configuring Argo notifications: %w", err)
	}

	s.log.Infof("Setup successful for source configuration %q", input.Context.SourceName)
	return source.StreamOutput{}, nil
}

// HandleExternalRequest handles external requests from ArgoCD.
func (s *Source) HandleExternalRequest(_ context.Context, input source.ExternalRequestInput) (source.ExternalRequestOutput, error) {
	payload := formatx.StructDumper().Sdump(string(input.Payload))
	s.log.WithField("payload", payload).Debug("Handling external request...")
	fallbackTimestamp := time.Now()

	var reqBody IncomingRequestBody
	err := json.Unmarshal(input.Payload, &reqBody)
	if err != nil {
		return source.ExternalRequestOutput{}, fmt.Errorf("while unmarshalling payload: %w", err)
	}

	srcCfg, ok := s.cfgs.Load(input.Context.SourceName)
	if !ok {
		return source.ExternalRequestOutput{}, fmt.Errorf("source configuration not found")
	}

	srcInstance, ok := srcCfg.(sourceInstance)
	if !ok {
		return source.ExternalRequestOutput{}, fmt.Errorf("source configuration is invalid")
	}

	msg := reqBody.Message
	if msg.Timestamp.IsZero() {
		msg.Timestamp = fallbackTimestamp
	}

	if input.Context.IsInteractivitySupported {
		section := s.generateInteractivitySection(reqBody, srcInstance.cfg)
		if section != nil {
			msg.Sections = append(msg.Sections, *section)
		}
	} else {
		msg.Type = api.NonInteractiveSingleSection
		lastSectionIdx := len(msg.Sections) - 1
		if lastSectionIdx != -1 {
			msg.Sections[lastSectionIdx].TextFields = append(msg.Sections[lastSectionIdx].TextFields, s.generateNonInteractiveFields(reqBody, srcInstance.cfg)...)
		}
	}

	return source.ExternalRequestOutput{
		Event: source.Event{
			Message:   msg,
			RawObject: reqBody,
		},
	}, nil
}

// Metadata returns metadata of the ArgoCD configuration.
func (s *Source) Metadata(_ context.Context) (api.MetadataOutput, error) {
	return api.MetadataOutput{
		Version:          s.pluginVersion,
		Description:      htmlDescription,
		DocumentationURL: "https://docs.botkube.io/configuration/source/argocd",
		JSONSchema: api.JSONSchema{
			Value: configJSONSchema,
		},
		ExternalRequest: api.ExternalRequestMetadata{
			Payload: api.ExternalRequestPayload{
				JSONSchema: api.JSONSchema{
					Value: requestJSONSchema,
				},
			},
		},
		Recommended: false,
	}, nil
}

func (s *Source) getK8sClient(k8sBytes []byte) (*dynamic.DynamicClient, error) {
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
