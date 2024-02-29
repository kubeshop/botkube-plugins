package aibrain

import (
	"errors"
	"fmt"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/multierror"
	pluginx "github.com/kubeshop/botkube/pkg/plugin"
)

const assistantID = "asst_eMM9QaWLi6cajHE4PdG1yU53"

// Config holds source configuration.
type Config struct {
	Log                   config.Logger `yaml:"log"`
	OpenAICloudServiceURL string        `yaml:"openAICloudServiceURL"`
	OpenAIAssistantID     string        `yaml:"openAIAssistantId"`
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	issues := multierror.New()
	if c.OpenAIAssistantID == "" {
		issues = multierror.Append(issues, errors.New("the Open AI Assistant ID cannot be empty"))
	}
	if c.OpenAICloudServiceURL == "" {
		issues = multierror.Append(issues, errors.New("the Open AI API key cannot be empty"))
	}
	return issues.ErrorOrNil()
}

func mergeConfigs(configs []*source.Config) (*Config, error) {
	defaults := &Config{
		OpenAIAssistantID: assistantID,
		Log: config.Logger{
			Level:     "info",
			Formatter: "json",
		},
	}
	var cfg *Config
	if err := pluginx.MergeSourceConfigsWithDefaults(defaults, configs, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func binaryDependencies() map[string]api.Dependency {
	return map[string]api.Dependency{
		kubectlBinaryName: {
			URLs: map[string]string{
				"windows/amd64": fmt.Sprintf("https://dl.k8s.io/release/%s/bin/windows/amd64/kubectl.exe", kubectlVersion),
				"darwin/amd64":  fmt.Sprintf("https://dl.k8s.io/release/%s/bin/darwin/amd64/kubectl", kubectlVersion),
				"darwin/arm64":  fmt.Sprintf("https://dl.k8s.io/release/%s/bin/darwin/arm64/kubectl", kubectlVersion),
				"linux/amd64":   fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/amd64/kubectl", kubectlVersion),
				"linux/s390x":   fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/s390x/kubectl", kubectlVersion),
				"linux/ppc64le": fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/ppc64le/kubectl", kubectlVersion),
				"linux/arm64":   fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/arm64/kubectl", kubectlVersion),
				"linux/386":     fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/386/kubectl", kubectlVersion),
			},
		},
	}
}
