package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/go-plugin"

	"github.com/kubeshop/botkube-cloud-plugins/internal/auth"
	aibrain "github.com/kubeshop/botkube-cloud-plugins/internal/source/ai-brain"
	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/executor"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/httpx"
	"github.com/kubeshop/botkube/pkg/loggerx"
	pluginx "github.com/kubeshop/botkube/pkg/plugin"
)

var (
	// version is set via ldflags by GoReleaser.
	version = "dev"

	//go:embed config_schema.json
	configJSONSchema string
)

const (
	pluginName      = "ai-face"
	description     = "Proxies incoming prompts into AI engine a.k.a brain that responds with analysis."
	maxRespBodySize = 5 * 1024 * 1024 // 5 MB
)

// Config holds executor configuration.
type Config struct {
	AIBrainSourceName string        `yaml:"aiBrainSourceName"`
	Log               config.Logger `yaml:"log"`
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

	log := loggerx.New(cfg.Log)

	aiBrainWebhookURL := fmt.Sprintf("%s/%s", in.Context.IncomingWebhook.BaseSourceURL, cfg.AIBrainSourceName)

	body, err := json.Marshal(aibrain.Payload{
		Prompt:    strings.TrimPrefix(in.Command, pluginName),
		MessageID: in.Context.Message.ParentActivityID,
	})
	if err != nil {
		return executor.ExecuteOutput{}, fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := e.httpClient.Post(aiBrainWebhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return executor.ExecuteOutput{}, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Failed to dispatch event. Response code: %d", resp.StatusCode)
		withLimit := io.LimitReader(resp.Body, maxRespBodySize)
		responseBody, err := io.ReadAll(withLimit)
		if err != nil {
			return executor.ExecuteOutput{}, fmt.Errorf("failed to read response body: %v", err)
		}
		return executor.ExecuteOutput{}, fmt.Errorf("failed to dispatch event. Response code: %d, Body: %s", resp.StatusCode, responseBody)
	}

	return executor.ExecuteOutput{
		Message: api.Message{
			Type: api.SkipMessage,
		},
	}, nil
}

// Help returns help message
func (*AIFace) Help(context.Context) (api.Message, error) {
	btnBuilder := api.NewMessageButtonBuilder()

	return api.Message{
		Sections: []api.Section{
			{
				Base: api.Base{
					Header: "Botkube: An AI-powered plugin for diagnosing your Kubernetes clusters in natural language.",
				},
				Buttons: []api.Button{
					btnBuilder.ForCommandWithDescCmd("Ask a question", "ai are there any failing pods in my cluster?"),
				},
			},
		},
	}, nil
}

func main() {
	executor.Serve(map[string]plugin.Plugin{
		pluginName: &executor.Plugin{
			Executor: auth.NewProtectedExecutor(&AIFace{
				httpClient: httpx.NewHTTPClient(),
			}),
		},
	})
}
