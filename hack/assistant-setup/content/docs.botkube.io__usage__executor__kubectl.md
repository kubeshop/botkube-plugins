Title: Kubectl | Botkube

URL Source: https://docs.botkube.io/usage/executor/kubectl

Markdown Content:
Botkube allows you to execute `kubectl` commands on your Kubernetes cluster. By default, kubectl command execution is disabled. See the [**Enabling plugin**](https://docs.botkube.io/configuration/executor/kubectl#enabling-plugin) section from the `kubectl` configuration documentation.

To execute the `kubectl` commands, send message in following format in the channel where Botkube is already added:

    @Botkube kubectl [verb] [resource] [flags]

By default, `k` and `kc` are configured as aliases for the `kubectl` command. You can use them on par with the `kubectl` command. To read more about aliases configuration, see the [Alias](https://docs.botkube.io/configuration/alias) section.

This command needs to be executed from configured channel.

Interactive kubectl commands builder[​](#interactive-kubectl-commands-builder "Direct link to Interactive kubectl commands builder")
------------------------------------------------------------------------------------------------------------------------------------

Use the interactive `kubectl` command builder to construct a `kubectl` command just by selecting items from dropdowns. This is especially useful on mobile when typing the command is harder.

The builder includes a resource name dropdown list. This is pre-populated with all the relevant resource names. It's great for discovering resources with the option to select them. E.g. Just grab a Pod name without needing to type or copy-and-paste.

To start the interactive `kubectl` command builder, run `@Botkube kubectl` from the configured channel where Botkube is added. You can also use any of the [configured aliases](https://docs.botkube.io/configuration/alias) - for example, the default `k` one, as illustrated below:

![Image 1: kubectl command builder](https://docs.botkube.io/assets/images/kc-cmd-builder-90ea740becbf2c0f126436c4a6c013bd.gif)

### Limitations[​](#limitations "Direct link to Limitations")

Keep in mind that the interactive command builder may not support all the commands that you can run just by typing them directly in a chat window. The following policies are applied:

*   Under verbs' dropdown, we display verbs that are defined under the `interactiveBuilder.allowed.verbs` configuration.
    
*   Under resources' dropdown, we display resources that are defined under the `interactiveBuilder.allowed.resources` configuration and are allowed for already selected verb. For example, for the `logs` verb we display only `pods`, `statefulsets`, `deployments`, and `daemonsets`.
    
*   For resources that are namespace-scoped, under namespaces' dropdown, we display `interactiveBuilder.allowed.namespaces` if defined. If static namespaces are not specified, plugin needs to have access to fetch all Namespaces, otherwise Namespace dropdown won't be visible at all.
    
*   The `kubectl` command preview is displayed only if the command that you built is valid, and you have permission to run it.
