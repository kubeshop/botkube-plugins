Title: Executor | Botkube

URL Source: https://docs.botkube.io/usage/executor/

Markdown Content:
While deploying Botkube, you can specify which [executors](https://docs.botkube.io/configuration/executor) you want to enable.

To check which executors are enabled, and get the aliases configured for them, run `@Botkube list executors`.

Aliases[​](#aliases "Direct link to Aliases")
---------------------------------------------

Alias is a shortcut for a longer command or just a part of it. It can be defined for all commands, including executor plugins and built-in Botkube commands. When you use an alias, the command is replaced with the underlying command before executing it. For example, `@Botkube k get pods` is replaced with `@Botkube kubectl get pods` before executing it.

Once you configured aliases, you can use them interchangeably with a full command. For example:

*   `k` as `kubectl`,
*   `kgp` as `kubectl get pods`,
*   `kgpa` as `kubectl get pods -A`,
*   `hh` as `helm history`,
*   `a` as `list actions`, the built-in Botkube command,

and so on. Your imagination is the limit!

Aliases are defined globally for the whole Botkube installation. To see which aliases are available for current conversation, run `@Botkube list aliases`.

To learn how to configure aliases and see the default configuration, see the [Alias](https://docs.botkube.io/configuration/alias) section.

Specify cluster name[​](#specify-cluster-name "Direct link to Specify cluster name")
------------------------------------------------------------------------------------

danger

Multi-cluster approach is supported only for the Mattermost and Discord integrations. Slack and Microsoft Teams integrations require separate Slack or Microsoft Teams apps for each Botkube installation.

If you have installed Botkube backend on multiple clusters, you can pass `--cluster-name` flag to execute kubectl command on specific cluster.

To get the list of all clusters configured in botkube, you can use the ping command.

For cluster-specific response, use `--cluster-name` flag to specify the cluster's name on which command needs to be executed. Use of this flag allows you to get response from any channel or group where Botkube is added. The flag is ignored in notifier commands as they can be executed from the configured channel only.

Filtering text output[​](#filtering-text-output "Direct link to Filtering text output")
---------------------------------------------------------------------------------------

Use the `--filter` flag to filter the output of Botkube executor commands. This returns any lines matching the flag's provided value.

The `--filter` flag uses simple string matching. And, only works for Botkube executor commands that return text output, e.g. `kubectl` or `list executors` commands.

### Filtering `kubectl` output[​](#filtering-kubectl-output "Direct link to filtering-kubectl-output")

![Image 1: flag_filter_kubectl_get_nodes](https://docs.botkube.io/assets/images/flag_filter_kubectl_get_nodes-16a8e267dffa5c29c4f0cb53d549bbe9.png) ![Image 2: flag_filter_kubectl_logs](https://docs.botkube.io/assets/images/flag_filter_kubectl_logs-3740479b57c49363f90ba706f516d758.png)
