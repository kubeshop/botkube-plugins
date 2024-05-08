import { AssistantTool } from "openai/src/resources/beta/assistants";
import dedent from "dedent";

export function setupTools(): Array<AssistantTool> {
  return [
    {
      type: "file_search",
    },
    {
      type: "function",
      function: {
        name: "botkubeGetStartupAgentConfiguration",
        description: dedent`
            Retrieve the Botkube Agent configuration in YAML format. This includes various details such as the configured communication platform, 
            aliases, enabled plugins, and their settings like RBAC.
            
            Use this tool to answer questions like:
            - What sources are configured for this channel?
            - Is the Kubernetes plugin enabled?
            - What Kubernetes alerts are currently configured on my cluster?
            
            Avoid using configuration names like 'botkube/kubernetes_6ulLY' in responses unless explicitly requested. 
            When a user asks about "this channel", prompt them to provide the exact channel name.
            
            The "loaderValidationWarnings" field lists all the warnings about the configuration. Use it to suggest fixes.
            The "incomingRequestPrompt" field defines the communicating platform associated with a prompt. Use this information to address platform-specific questions.
            If the "incomingRequestPrompt" field is not specified, prompt user provide platform name.
            Terms "cloudSlack" and "cloudTeams" refer to the Slack and Teams platforms, respectively.`,
      },
    },
    {
      type: "function",
      function: {
        name: "botkubeGetAgentStatus",
        description: dedent`
            Get the current Botkube Agent status in JSON format. 
            It represents the status of each plugin and communication platform, indicating whether it's enabled, crashing, and if so, with what reason. 
            If the "enabled" field is not specified, the plugin is disabled.
            Use this tool to answer questions like:
            - Which sources are enabled for this channel?
            - Is the Kubernetes plugin enabled?
            - Is the kubectl plugin disabled?
            - What is the status of a specific plugin?
            - Why am I not receiving Kubernetes notifications?`,
      },
    },
    {
      type: "function",
      function: {
        name: "kubectlGetResource",
        description: dedent`
            Get a Kubernetes resources. Useful for debugging and troubleshooting.
            If resource_name is specified, returns the resource in a YAML definition. Examples:
            # List all pods in ps output format
            kubectl get pods
            
            # List a single replication controller with specified NAME in ps output format
            kubectl get replicationcontroller web
            
            # List all replication controllers and services together in ps output format
            kubectl get rc,services
            
            # List resource by type and name
            kubectl get pods web-pod-13je7`,
        parameters: {
          type: "object",
          properties: {
            resource_type: {
              type: "string",
              description:
                "Kubernetes resource type, e.g. pod, deployment, replicationcontroller.",
            },
            resource_name: {
              type: "string",
              description:
                "Kubernetes resource name, e.g. botkube-6c6fd8b4d6-f559q",
            },
            namespace: {
              type: "string",
              description: "Kubernetes namespace, e.g. kube-system",
            },
            "all-namespaces": {
              type: "boolean",
              description:
                "If present, list the requested resource(s) across all namespaces. Cannot be used together with the namespace property.",
            },
          },
          required: ["resource_type"],
        },
      },
    },
    {
      type: "function",
      function: {
        name: "kubectlGetEvents",
        description: dedent`
            Returns a table of the most important information about events.
            You can request events for a namespace, for all namespace, or filtered to only those pertaining to a specified resource.`,
        parameters: {
          type: "object",
          properties: {
            namespace: {
              type: "string",
              description: "Kubernetes namespace, e.g. kube-system",
            },
            "all-namespaces": {
              type: "boolean",
              description:
                "If present, list the requested object(s) across all namespaces. Cannot be used together with the namespace property.",
            },
            types: {
              type: "string",
              description:
                "Output only events of given types. Accepts comma separated values.",
            },
            for: {
              type: "string",
              description:
                "Filter events to only those pertaining to the specified resource.",
            },
          },
        },
      },
    },
    {
      type: "function",
      function: {
        name: "kubectlDescribeResource",
        description: dedent`
            Returns detailed description of the selected resources,
            including related resources such as events or controllers. Examples:
            # Describe a node
            kubectl describe nodes kubernetes-node-emt8.c.myproject.internal
                  
            # Describe a pod
            kubectl describe pods nginx`,
        parameters: {
          type: "object",
          properties: {
            resource_type: {
              type: "string",
              description:
                "Kubernetes resource type, e.g. pod, deployment, replicationcontroller.",
            },
            resource_name: {
              type: "string",
              description:
                "Kubernetes resource name, e.g. botkube-6c6fd8b4d6-f559q",
            },
            namespace: {
              type: "string",
              description: "Kubernetes namespace, e.g. kube-system",
            },
            "all-namespaces": {
              type: "boolean",
              description:
                "If present, list the requested resource(s) across all namespaces. Cannot be used together with the namespace property.",
            },
          },
          required: ["resource_type"],
        },
      },
    },
    {
      type: "function",
      function: {
        name: "kubectlTopPods",
        description: dedent`
             Returns the CPU and memory usage for Pods. Examples:
             # Show metrics for all pods in the default namespace
             kubectl top pod
             # Show metrics for all pods in the given namespace
             kubectl top pod --namespace=NAMESPACE
             # Show metrics for a given pod and its containers
             kubectl top pod POD_NAME --containers`,
        parameters: {
          type: "object",
          properties: {
            namespace: {
              type: "string",
              description:
                "Kubernetes namespace, e.g. kube-system. It can be used only for pods.",
            },
            resource_name: {
              type: "string",
              description: "The pod or node name, e.g. botkube-api-server.",
            },
          },
          required: ["namespace"],
        },
      },
    },
    {
      type: "function",
      function: {
        name: "kubectlTopNodes",
        description: dedent`
            Returns the CPU and memory usage for Nodes. Examples:
            # Show metrics for all nodes
            kubectl top node
            
            # Show metrics for a given node
            kubectl top node NODE_NAME`,
        parameters: {
          type: "object",
          properties: {
            resource_name: {
              type: "string",
              description: "The pod or node name, e.g. botkube-api-server.",
            },
          },
        },
      },
    },
    {
      type: "function",
      function: {
        name: "kubectlLogs",
        description: "Useful for getting container logs.",
        parameters: {
          type: "object",
          properties: {
            namespace: {
              type: "string",
              description: "Kubernetes namespace, e.g. kube-system",
            },
            pod_name: {
              type: "string",
              description: "Kubernetes pod name, e.g. botkube-6c6fd8b4d6-f559q",
            },
            container_name: {
              type: "string",
              description: "Pod container name, e.g. botkube-api-server.",
            },
          },
          required: ["namespace", "pod_name"],
        },
      },
    },
  ];
}
