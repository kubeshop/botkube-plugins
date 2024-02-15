package main

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/executor"
	"github.com/kubeshop/botkube/pkg/httpx"
	"github.com/kubeshop/botkube/pkg/pluginx"
)

var (
	// version is set via ldflags by GoReleaser.
	version = "dev"

	//go:embed config_schema.json
	configJSONSchema string
)

const (
	pluginName  = "ai-face"
	description = "Proxies incoming prompts into AI engine a.k.a brain that responds with analysis."
)

// Config holds executor configuration.
type Config struct {
	AIBrainSourceName string `json:"aiBrainSourceName"`
}

var defaultConfig = Config{
	AIBrainSourceName: "ai-brain",
}

// AIFace implements Botkube executor plugin.
type AIFace struct {
	httpClient *http.Client
}

var _ executor.Executor = &AIFace{}

// Metadata returns details about the plugin.
func (*AIFace) Metadata(context.Context) (api.MetadataOutput, error) {
	return api.MetadataOutput{
		Version:     version,
		Description: description,
		Recommended: true,
		JSONSchema: api.JSONSchema{
			Value: configJSONSchema,
		},
	}, nil
}

// Execute returns a given command as response.
func (e *AIFace) Execute(_ context.Context, in executor.ExecuteInput) (executor.ExecuteOutput, error) {
	var cfg Config
	err := pluginx.MergeExecutorConfigsWithDefaults(defaultConfig, in.Configs, &cfg)
	if err != nil {
		return executor.ExecuteOutput{}, fmt.Errorf("while merging input configuration: %w", err)
	}

	aiBrainWebhookURL := fmt.Sprintf("%s/%s", in.Context.IncomingWebhook.BaseSourceURL, cfg.AIBrainSourceName)

	//nolint:gosec
	resp, err := e.httpClient.Post(aiBrainWebhookURL, "application/json", strings.NewReader(fmt.Sprintf(`{"prompt": "%s", "messageId": "%s"}`, in.Command, in.Context.Message.ParentActivityID)))
	if err != nil {
		return executor.ExecuteOutput{}, err
	}
	defer resp.Body.Close()

	return executor.ExecuteOutput{
		Message: api.Message{
			Type: api.SkipMessage,
		},
	}, nil
}

// Help returns help message
func (*AIFace) Help(context.Context) (api.Message, error) {
	return api.Message{}, nil
}

func main() {
	executor.Serve(map[string]plugin.Plugin{
		pluginName: &executor.Plugin{
			Executor: &AIFace{
				httpClient: httpx.NewHTTPClient(),
			},
		},
	})
}
