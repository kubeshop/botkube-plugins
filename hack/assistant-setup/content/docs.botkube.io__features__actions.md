Title: Automated Actions | Botkube

URL Source: https://docs.botkube.io/features/actions

Markdown Content:
The action configuration allows you to automate your workflow by defining commands that are executed on an incoming event. Similar to communication platforms, actions use source bindings to define the event sources (command triggers) and executor bindings to run a given command.

The actions can be disabled and enabled using dedicated commands.

Predefined actions[​](https://docs.botkube.io/features/actions/#predefined-actions "Direct link to Predefined actions")
-----------------------------------------------------------------------------------------------------------------------

There are two predefined actions available:

*   `describe-created-resource` - describes a newly created resource,
*   `show-logs-on-error` - prints logs on Pod/StatefulSet/DaemonSet/Deployment error.

Both of them are disabled by default.

Action output[​](https://docs.botkube.io/features/actions/#action-output "Direct link to Action output")
--------------------------------------------------------------------------------------------------------

The command output is sent to the communication platforms that define the same source bindings. If there is no communication platform defined with the same source bindings, the action is still executed, but the output is ignored.

Command templates[​](https://docs.botkube.io/features/actions/#command-templates "Direct link to Command templates")
--------------------------------------------------------------------------------------------------------------------

Each action defines the `command` property which defines a command to be executed. It supports [Go template](https://golang.org/pkg/text/template/) syntax. The command template is rendered before actual execution.

The following variables are available:

*   `{{ .Event }}` - the event object that triggered the action. See all available event properties in the [`event.go`](https://github.com/kubeshop/botkube/blob/main/internal/source/kubernetes/event/event.go) file.

There are multiple helper functions supported by the templating engine. To learn more, read the documentation on the [`slim-sprig`](https://go-task.github.io/slim-sprig/) library page.

### Examples[​](https://docs.botkube.io/features/actions/#examples "Direct link to Examples")

*   Run `kubectl describe` for each event:
    
    ```
    kubectl describe {{ .Event.TypeMeta.Kind | lower }}{{ if .Event.Namespace }} -n {{ .Event.Namespace }}{{ end }} {{ .Event.Name }}
    ```
    
*   Run `kubectl logs` for each event for specific resources (`v1/pods`, `apps/v1/deployments`, `apps/v1/statefulsets`, `apps/v1/daemonsets`, `batch/v1/jobs`:
    
    ```
    kubectl logs {{ .Event.TypeMeta.Kind | lower }}/{{ .Event.Name }} -n {{ .Event.Namespace }}
    ```
    

RBAC[​](https://docs.botkube.io/features/actions/#rbac "Direct link to RBAC")
-----------------------------------------------------------------------------

The `ChannelName` RBAC policy is not supported for automated actions at this time. Use the `Static` RBAC policy instead. For more information read the [RBAC](https://docs.botkube.io/features/rbac) documentation.

Manage Actions with Botkube commands[​](https://docs.botkube.io/features/actions/#manage-actions-with-botkube-commands "Direct link to Manage Actions with Botkube commands")
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Botkube allows you to manage actions using `@Botkube` commands.

### List available actions[​](https://docs.botkube.io/features/actions/#list-available-actions "Direct link to List available actions")

Run `@Botkube list actions` to get list of configured actions and their running status:

![Image 1: List available actions](https://docs.botkube.io/assets/images/list-actions-e1d1d86e622d7a10077d5347958a3559.png)

### Disable action[​](https://docs.botkube.io/features/actions/#disable-action "Direct link to Disable action")

Run `@Botkube disable action {action-name}` to disable an action named `{action-name}`. The action settings are persisted across Botkube app restarts.

![Image 2: Disable action](https://docs.botkube.io/assets/images/disable-action-414dd23e8a7bcb9efc1d52251f68999c.png)

When you disable an action, changes are applied once the Botkube is restarted. It is an automated process which usually takes a few seconds.

### Enable action[​](https://docs.botkube.io/features/actions/#enable-action "Direct link to Enable action")

Run `@Botkube enable action {action-name}` to enable an action named `{action-name}`. The action settings are persisted across Botkube app restarts.

![Image 3: Enable action](https://docs.botkube.io/assets/images/enable-action-08c9232d0d21939ec91201abdcb70a50.png)

When you enable an action, changes are applied once the Botkube is restarted. It is an automated process which usually takes a few seconds.

Self-hosted Configuration Syntax[​](https://docs.botkube.io/features/actions/#self-hosted-configuration-syntax "Direct link to Self-hosted Configuration Syntax")
-----------------------------------------------------------------------------------------------------------------------------------------------------------------

The following snippet contains an example configuration for self-hosted Botkube installation:

```
# Map of actions. Action contains configuration for automation based on observed events.# The property name under `actions` object is an alias for a given configuration. You can define multiple actions configuration with different names.## Format: actions.{alias}actions:  "describe-created-resource":    # If true, enables the action.    enabled: false    # Action display name posted in the channels bound to the same source bindings.    displayName: "Describe created resource"    # Command to execute when the action is triggered. You can use Go template (https://pkg.go.dev/text/template) together with all helper functions defined by Slim-Sprig library (https://go-task.github.io/slim-sprig).    # You can use the `{{ .Event }}` variable, which contains the event object that triggered the action. See all available event properties on https://github.com/kubeshop/botkube/blob/main/internal/source/kubernetes/event/event.go.    command: "kubectl describe {{ .Event.TypeMeta.Kind | lower }}{{ if .Event.Namespace }} -n {{ .Event.Namespace }}{{ end }} {{ .Event.Name }}"    # Bindings for a given action.    bindings:      # Event sources that trigger a given action.      sources:        - k8s-create-events      # Executors configuration used to execute a configured command.      executors:        - k8s-default-tools  "show-logs-on-error":    # If true, enables the action.    enabled: false    # Action display name posted in the channels bound to the same source bindings.    displayName: "Show logs on error"    # Command to execute when the action is triggered. You can use Go template (https://pkg.go.dev/text/template) together with all helper functions defined by Slim-Sprig library (https://go-task.github.io/slim-sprig).    # You can use the `{{ .Event }}` variable, which contains the event object that triggered the action. See all available event properties on https://github.com/kubeshop/botkube/blob/main/pkg/event/event.go.    command: "kubectl logs {{ .Event.TypeMeta.Kind | lower }}/{{ .Event.Name }} -n {{ .Event.Namespace }}"    # Bindings for a given action.    bindings:      # Event sources that trigger a given action.      sources:        - k8s-err-with-logs-events      # Executors configuration used to execute a configured command.      executors:        - k8s-default-tools
```
