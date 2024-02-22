package main

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var openAITools = []openai.AssistantTool{
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name:        "kubectlGetPods",
			Description: "Useful for getting one pod or all pods in a kubernetes namespace.",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system",
					},
					"pod_name": {
						Type:        jsonschema.String,
						Description: "Kubernetes pod name, e.g. botkube-6c6fd8b4d6-f559q",
					},
				},
				Required: []string{"namespace"},
			},
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name:        "kubectlGetSecrets",
			Description: "Useful for getting one secret or all secrets in a kubernetes namespace.",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system",
					},
					"secret_name": {
						Type:        jsonschema.String,
						Description: "Kubernetes secret name, e.g. api-key",
					},
				},
				Required: []string{"namespace"},
			},
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name:        "kubectlDescribePod",
			Description: "Useful for describing a pod in a kubernetes namespace.",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system",
					},
					"pod_name": {
						Type:        jsonschema.String,
						Description: "Kubernetes pod name, e.g. botkube-6c6fd8b4d6-f559q",
					},
				},
				Required: []string{"namespace", "pod_name"},
			},
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name:        "kubectlLogs",
			Description: "Useful for getting container logs.",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system",
					},
					"pod_name": {
						Type:        jsonschema.String,
						Description: "Kubernetes pod name, e.g. botkube-6c6fd8b4d6-f559q",
					},
					"container_name": {
						Type:        jsonschema.String,
						Description: "Pod container name, e.g. botkube-api-server.",
					},
				},
				Required: []string{"namespace", "pod_name"},
			},
		},
	},
}
