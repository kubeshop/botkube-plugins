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
			Name: "botkubeGetStartupAgentConfiguration",
			Description: heredoc.Doc(`
			Retrieve the currently used Botkube Agent configuration in YAML format. This includes various details such as the configured communication platform (e.g., Slack or Teams), 
			aliases, enabled plugins, and their settings like RBAC or plugin-specific configurations.
			
			Use this tool to answer questions like:
			- What sources are configured for this channel?
			- What Kubernetes resources am I watching?
			- Is the Kubernetes plugin enabled?
			- Which channels are configured for Botkube?
			- What Kubernetes alerts are currently configured on my cluster?
			
			Avoid using configuration names like 'botkube/kubernetes_6ulLY' in responses unless explicitly requested by the user. 
			When a user asks about configurations related to "this channel," prompt them to specify the channel name they are referring to. 
			You can suggest channel names based on the configuration definition.
			
			The "loaderValidationWarnings" field contains all the warnings regarding the returned configuration. Use it to recommend possible fixes if the configuration is not valid.
			The "incomingRequestPrompt" field defines the communicating platform that is used for a given prompt. Use that value to answer questions related with a given platform, such as:
			- What automations are enabled for #general channel?

			If the "incomingRequestPrompt" field is not specified, ask about user about the used platform.

			All "cloudSlack" terms in config simply refers to Slack platform. All "cloudTeams" terms in config simply refers to Teams platform.`),
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name: "botkubeGetAgentStatus",
			Description: heredoc.Doc(`
			Get the current Botkube Agent status in JSON format. 
            It represents the status of each plugin and communication platform, indicating whether it's enabled, crashing, and if so, with what reason. 
			If the "enabled" field is not specified, the plugin is disabled.

			Use this tool to answer questions like:
			- Which sources are enabled for this channel?
			- Is the Kubernetes plugin enabled?
			- Is the kubectl plugin disabled?
			- What is the status of a specific plugin?
			- Why am I not receiving Kubernetes notifications?`),
		},
	},
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
			Name: "kubectlTopPods",
			Description: heredoc.Doc(`
					Returns the CPU and memory usage for Pods. Examples:
					  # Show metrics for all pods in the default namespace
					  kubectl top pod
					  # Show metrics for all pods in the given namespace
					  kubectl top pod --namespace=NAMESPACE
					  # Show metrics for a given pod and its containers
					  kubectl top pod POD_NAME --containers`),
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"namespace": {
						Type:        jsonschema.String,
						Description: "Kubernetes namespace, e.g. kube-system. It can be used only for pods.",
					},
					"resource_name": {
						Type:        jsonschema.String,
						Description: "The pod or node name, e.g. botkube-api-server.",
					},
				},
				Required: []string{"namespace"},
			},
		},
	},
	{
		Type: openai.AssistantToolType(openai.ToolTypeFunction),
		Function: &openai.FunctionDefinition{
			Name: "kubectlTopNodes",
			Description: heredoc.Doc(`
					Returns the CPU and memory usage for Nodes. Examples:
					  # Show metrics for all nodes
					  kubectl top node
					
					  # Show metrics for a given node
					  kubectl top node NODE_NAME`),
			Parameters: jsonschema.Definition{
				Type: jsonschema.Object,
				Properties: map[string]jsonschema.Definition{
					"resource_name": {
						Type:        jsonschema.String,
						Description: "The pod or node name, e.g. botkube-api-server.",
					},
				},
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
