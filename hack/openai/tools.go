package main

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var openAITools = []openai.AssistantTool{
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name: "kubectlGetResource",
			Description: heredoc.Doc(`
			Get a Kubernetes resources. If resource_name is specified, returns the resource in a YAML definition. 
			Useful for debugging and troubleshooting. Examples:
				  # List all pods in ps output format
				  kubectl get pods
				
				  # List a single replication controller with specified NAME in ps output format
				  kubectl get replicationcontroller web
				
				  # List all replication controllers and services together in ps output format
				  kubectl get rc,services
				
				  # List resource by type and name
				  kubectl get pods web-pod-13je7`),

			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"resource_type": {
						Type:        jsonschema.String,
						Description: "Kubernetes resource type, e.g. pod, deployment, replicationcontroller.",
					},
					"resource_name": {
						Type:        jsonschema.String,
						Description: "Kubernetes resource name, e.g. botkube-6c6fd8b4d6-f559q",
					},
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system",
					},
					"all-namespaces": {
						Type:        jsonschema.Boolean,
						Description: "If present, list the requested resource(s) across all namespaces. Cannot be used together with the namespace property.",
					},
				},
				Required: []string{"resource_type"},
			},
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name:        "kubectlGetEvents",
			Description: "Returns a table of the most important information about events. You can request events for a namespace, for all namespace, or filtered to only those pertaining to a specified resource.",
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system",
					},
					"all-namespaces": {
						Type:        jsonschema.Boolean,
						Description: "If present, list the requested object(s) across all namespaces. Cannot be used together with the namespace property.",
					},
					"types": {
						Type:        jsonschema.String,
						Description: "Output only events of given types. Accepts comma separated values.",
					},
					"for": {
						Type:        jsonschema.String,
						Description: "Filter events to only those pertaining to the specified resource.",
					},
				},
			},
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name: "kubectlDescribeResource",
			Description: heredoc.Doc(`
				Returns detailed description of the selected resources, 
				including related resources such as events or controllers. Examples:
					  # Describe a node
					  kubectl describe nodes kubernetes-node-emt8.c.myproject.internal
					
					  # Describe a pod
					  kubectl describe pods nginx`),
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"resource_type": {
						Type:        jsonschema.String,
						Description: "Kubernetes resource type, e.g. pod, deployment, replicationcontroller.",
					},
					"resource_name": {
						Type:        jsonschema.String,
						Description: "Kubernetes resource name, e.g. botkube-6c6fd8b4d6-f559q",
					},
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system",
					},
					"all-namespaces": {
						Type:        jsonschema.Boolean,
						Description: "If present, list the requested resource(s) across all namespaces. Cannot be used together with the namespace property.",
					},
				},
				Required: []string{"resource_type"},
			},
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name: "kubectlGetResourceConsumption",
			Description: heredoc.Doc(`
					Returns the CPU and memory usage for nodes or pods. Examples:
					  # Show metrics for all pods in the default namespace
					  kubectl top pod
					  # Show metrics for all pods in the given namespace
					  kubectl top pod --namespace=NAMESPACE
					  # Show metrics for a given pod and its containers
					  kubectl top pod POD_NAME --containers
					  # Show metrics for all nodes
					  kubectl top node
					  # Show metrics for a given node
					  kubectl top node NODE_NAME`),
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{

					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system. It can be used only for pods.",
					},
					"resource_type": {
						Type:        jsonschema.String,
						Description: "Kubernetes pod name, e.g. botkube-6c6fd8b4d6-f559q",
						Enum:        []string{"node", "pod"},
					},
					"resource_name": {
						Type:        jsonschema.String,
						Description: "The pod or node name, e.g. botkube-api-server.",
					},
				},
				Required: []string{"resource_type"},
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
