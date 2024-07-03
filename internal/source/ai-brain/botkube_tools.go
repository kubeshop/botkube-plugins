package aibrain

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kubeshop/botkube-cloud-plugins/internal/remote"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/httpx"
	"github.com/kubeshop/botkube/pkg/ptr"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/yaml.v3"
)

const (
	healthURLFmt = "http://127.0.0.1:%s/healthz"
)

// BotkubeRunner is a runner that executes Botkube related commands.
type BotkubeRunner struct {
	tracer           trace.Tracer
	deployCli        *remote.DeploymentClient
	httpCli          *http.Client
	rawStartupConfig string
	startupConfig    *ConfigWithDetails
}

// ConfigWithDetails represents Botkube configuration with additional details.
type ConfigWithDetails struct {
	*config.Config           `yaml:",inline"`
	LoaderValidationWarnings string `yaml:"loaderValidationWarnings"`
	IncomingRequestPrompt    string `yaml:"incomingRequestPrompt,omitempty"`
}

// NewBotkubeRunner creates new runner instance.
func NewBotkubeRunner(tracer trace.Tracer) (*BotkubeRunner, error) {
	cfg, ok := remote.GetConfig()
	if !ok {
		return nil, fmt.Errorf("cannot get Botkube cloud related environment variables")
	}

	r := &BotkubeRunner{
		tracer:    tracer,
		deployCli: remote.NewDeploymentClient(cfg),
		httpCli:   httpx.NewHTTPClient(),
	}

	err := r.initStartupConfig()
	if err != nil {
		return nil, fmt.Errorf("while initializing startup config: %w", err)
	}
	return r, nil
}

// GetStartupAgentConfiguration returns Botkube startup configuration.
func (r *BotkubeRunner) GetStartupAgentConfiguration(ctx context.Context, _ []byte, p *Payload) (string, error) {
	_, span := r.tracer.Start(ctx, "aibrain.BotkubeRunner.GetStartupAgentConfiguration")
	defer span.End()

	platform := getPlatform(p.MessageID)
	if platform == nil {
		return r.rawStartupConfig, nil
	}

	out := ConfigWithDetails{
		Config:                   r.startupConfig.Config,
		LoaderValidationWarnings: r.startupConfig.LoaderValidationWarnings,
		IncomingRequestPrompt:    ptr.ToValue(platform),
	}
	raw, err := yaml.Marshal(out)
	if err != nil {
		return "", fmt.Errorf("while marshaling Botkube configuration: %w", err)
	}

	return string(raw), nil
}

// GetAgentStatus returns Botkube Agent health status.
func (r *BotkubeRunner) GetAgentStatus(ctx context.Context, _ []byte, _ *Payload) (string, error) {
	_, span := r.tracer.Start(ctx, "aibrain.BotkubeRunner.GetAgentStatus")
	defer span.End()

	url := fmt.Sprintf(healthURLFmt, r.startupConfig.Settings.HealthPort)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("while creating Botkube health request: %w", err)
	}

	resp, err := r.httpCli.Do(req)
	if err != nil {
		return "", fmt.Errorf("while checking Botkube health: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("while reading Botkube health response: %w", err)
	}

	return string(body), nil
}

func (r *BotkubeRunner) initStartupConfig() error {
	rawCfg, cfg, err := r.fetchConfig(context.Background())
	if err != nil {
		return err
	}

	r.rawStartupConfig = rawCfg
	r.startupConfig = cfg

	return nil
}
func (r *BotkubeRunner) fetchConfig(ctx context.Context) (string, *ConfigWithDetails, error) {
	out, err := r.deployCli.GetConfig(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("while getting Botkube configuration: %w", err)
	}

	// Load with defaults to make sure Agent's ENVs are also taken into account
	cfg, details, err := config.LoadWithDefaults([][]byte{
		[]byte(out.YAMLConfig),
	})
	if err != nil {
		return "", nil, fmt.Errorf("while merging app configuration: %w", err)
	}
	if cfg == nil {
		return "", nil, fmt.Errorf("configuration cannot be nil")
	}

	redactedCfg := config.HideSensitiveInfo(*cfg)
	cfgWithDetails := ConfigWithDetails{
		Config: &redactedCfg,
	}
	if details.ValidateWarnings != nil {
		cfgWithDetails.LoaderValidationWarnings = details.ValidateWarnings.Error()
	}

	rawCfg, err := yaml.Marshal(cfgWithDetails)
	if err != nil {
		return "", nil, fmt.Errorf("while marshaling Botkube configuration: %w", err)
	}

	return string(rawCfg), &cfgWithDetails, nil
}

func getPlatform(msgID string) *string {
	if msgID == "" {
		return nil
	}

	if strings.Contains(msgID, teamsMessageIDSubstr) {
		return ptr.FromType("Teams")
	}
	return ptr.FromType("Slack")
}
