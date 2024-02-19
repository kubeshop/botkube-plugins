package aibrain

import (
	_ "embed"

	"github.com/kubeshop/botkube/pkg/api/source"
	"github.com/kubeshop/botkube/pkg/config"
	pluginx "github.com/kubeshop/botkube/pkg/plugin"
)

var (
	//go:embed config_schema.json
	ConfigJSONSchema string

	//go:embed webhook_schema.json
	IncomingWebhookJSONSchema string
)

// Config holds source configuration.
type Config struct {
	Log config.Logger `yaml:"log"`
}

// MergeConfigs merges the configuration.
func MergeConfigs(configs []*source.Config) (Config, error) {
	defaults := Config{
		Log: config.Logger{
			Level:     "info",
			Formatter: "json",
		},
	}
	var cfg Config
	err := pluginx.MergeSourceConfigsWithDefaults(defaults, configs, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
