Title: Alias | Botkube

URL Source: https://docs.botkube.io/configuration/alias

Markdown Content:
Botkube can define multiple aliases for arbitrary commands. The aliases are replaced with the underlying command before executing it. Aliases can replace a single word or multiple ones. For example, you can define a `k` alias for `kubectl`, or `kgp` for `kubectl get pods`.

Aliases work for all commands, including executor plugins and Botkube built-in ones. To learn more about how to configure Executors, see the [Executor](https://docs.botkube.io/configuration/executor/) section.

Aliases are defined globally for the whole Botkube installation. Once they are configured, read the [Aliases](https://docs.botkube.io/usage/executor/#aliases) section in Usage document.

You can configure aliases in the Botkube Cloud dashboard.

To configure the aliases for the self-hosted Botkube installation, use the following syntax:

# Custom aliases for given commands.# The aliases are replaced with the underlying command before executing it.# Aliases can replace a single word or multiple ones. For example, you can define a `k` alias for `kubectl`, or `kgp` for `kubectl get pods`.### Format: aliases.{alias}aliases:  kc:    command: kubectl    displayName: "Kubectl alias"  k:    command: kubectl    displayName: "Kubectl alias"## Multi-word alias example:#  kgp:#    command: kubectl get pods#    displayName: "Get pods"
