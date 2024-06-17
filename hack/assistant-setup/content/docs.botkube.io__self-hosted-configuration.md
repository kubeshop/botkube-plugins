Title: Self-hosted Configuration | Botkube

URL Source: https://docs.botkube.io/self-hosted-configuration/

Markdown Content:
info

This document is applicable only for self-hosted installations. For Botkube Cloud installations, the Botkube Agent configuration is managed via the [Botkube Cloud dashboard](https://app.botkube.io/).

Overview[​](#overview "Direct link to Overview")
------------------------------------------------

Botkube self-hosted installation allows you to specify [source](https://docs.botkube.io/self-hosted-configuration/source), [executor](https://docs.botkube.io/self-hosted-configuration/executor), [communication](https://docs.botkube.io/self-hosted-configuration/communication), and [general](https://docs.botkube.io/self-hosted-configuration/general) Botkube settings as a part of YAML configuration. Check the related documents for more detailed explanation.

Configuration source[​](#configuration-source "Direct link to Configuration source")
------------------------------------------------------------------------------------

The configuration settings are read from two sources:

*   the configuration files specified by the `BOTKUBE_CONFIG_PATHS` environment variable or `--config/-c` flag. For example:
    
    ```
    export BOTKUBE_CONFIG_PATHS="global.yaml,team-b-specific.yaml"# or./botkube --config "global.yaml,team-b-specific.yaml"
    ```
    
    You can split individual settings into multiple configuration files. The priority will be given to the last (right-most) file specified. Files with `_` name prefix are always read as the last ones. See the [merging strategy](#merging-strategy) section for more details.
    
    note
    
    For Helm installation, Botkube uses `_runtime_state.yaml` and `_startup_state.yaml` files to store its internal state. Remember to keep these files in the `BOTKUBE_CONFIG_PATHS` environment variable.
    
*   the exported [environment variables](#environment-variables) that overrides the configuration specified in the files.
    

Helm install options[​](#helm-install-options "Direct link to Helm install options")
------------------------------------------------------------------------------------

Advanced Helm install options are documented [here](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters).

Updating the configuration[​](#updating-the-configuration "Direct link to Updating the configuration")
------------------------------------------------------------------------------------------------------

To update Botkube configuration, you can either:

*   upgrade Botkube installation with Helm,
*   use dedicated `@Botkube` commands, to e.g. toggle notifications or edit Source Bindings. See the [Notification-related commands](https://docs.botkube.io/features/executing-commands#notification-related-commands) for more details,
*   or, [migrate to Botkube Cloud](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud) and manage configuration through the Botkube Cloud Dashboard.

If you wish to change the configuration with Helm, create a `/tmp/values.yaml` file that contains the new values and run:

*   Botkube CLI
*   Helm CLI

```
botkube install -f /tmp/values.yaml
```

As both Helm release upgrade and some of the `@Botkube` commands modify the same configuration, it is merged during command execution. Whenever you specify a new value in the `/tmp/values.yaml` file, it will override the existing value in the configuration.

### Preventing overrides by default Helm chart values[​](#preventing-overrides-by-default-helm-chart-values "Direct link to Preventing overrides by default Helm chart values")

Keep in mind that even if you don't specify custom values in the `/tmp/values.yaml` file, Helm can override the existing values with the default ones.

Consider the following config:

```
communications:  "default-group":    socketSlack:      enabled: true      botToken: "{botToken}"      appToken: "{appToken}"      channels:        "default":          name: general          notification:            disabled: false # default from the Helm chart          bindings:            sources:              - k8s-all-events # default from the Helm chart# (...)
```

Assume that users ran the following commands:

```
@Botkube edit SourceBindings k8s-err-events, k8s-recommendation-events@Botkube disable notifications
```

Which effectively result in the following config that Botkube sees:

```
communications:  "default-group":    socketSlack:      enabled: true      botToken: "{botToken}"      appToken: "{appToken}"      channels:        "default":          name: general          notification:            disabled: true # set by user command          bindings:            sources:              - k8s-err-events # set by user command              - k8s-recommendation-events # set by user command# (...)
```

To persist the configuration that users provided, and not overwrite notification and source bindings values, run Helm upgrade with:

```
communications:  "default-group":    socketSlack:      channels:        "default":          name: general          notification: null # explicitly not use defaults from Helm chart          bindings:            sources: null # explicitly not use defaults from Helm chart# (...) other values
```

The following properties need such `null` value during upgrade, if you want to keep the previous configuration:

*   `communications.default-group.{communication-platform}.channels.default.notifications`, where `{communication-platform}` is any communication platform supported except Microsoft Teams,
*   `communications.default-group.{communication-platform}.channels.default.bindings.sources`, where `{communication-platform}` is any communication platform supported except Microsoft Teams,
*   `communications.default-group.teams.bindings.sources`.

To learn more, read the [Deleting a default key](https://helm.sh/docs/chart_template_guide/values_files/#deleting-a-default-key) paragraph in Helm documentation.

Environment variables[​](#environment-variables "Direct link to Environment variables")
---------------------------------------------------------------------------------------

The individual communication settings can be specified via environment variables. They take priority and override the configuration specified in the file.

To construct the environment variable name, take any property from the configuration file and make it uppercase. Use the underscore for properties that are nested. Use the double underscore for all camelCase properties. Finally, add the `BOTKUBE_` prefix.

For example, such configuration property from YAML:

```
settings:  kubectl:    defaultNamespace: "NAMESPACE"
```

is mapped to the `BOTKUBE_SETTINGS_KUBECTL_DEFAULT__NAMESPACE` environment variable.

This is a useful feature that allows you to store the overall configuration in a file, where sensitive data, such as tokens, can be put in environment variables. See the [**Tokens from Vault via CSI driver**](https://docs.botkube.io/self-hosted-configuration/communication/vault-csi/) tutorial for an example use-case.

Merging strategy[​](#merging-strategy "Direct link to Merging strategy")
------------------------------------------------------------------------

Botkube allows you to split individual settings into multiple configuration files. The following rules apply:

*   The priority will be given to the last (right-most) file specified.
    
*   Files with `_` name prefix are always read as the last ones following the initial order.
    
*   Objects are merged together and primitive fields are overridden. For example:
    
    ```
    # a.yaml - first filesettings:  clusterName: dev-cluster  kubectl:    enabled: false
    ```
    
    ```
    # _a.yaml - second file with `_` prefixsettings:  clusterName: demo-cluster
    ```
    
    ```
    # b.yaml - third filesettings:  kubectl:    enabled: true
    ```
    
    ```
    # resultsettings:  clusterName: demo-cluster  kubectl:    enabled: true
    ```
    
*   The arrays items are not merged, they are overridden. For example:
    
    ```
    # a.yamlsettings:  kubectl:    enabled: true    commands:      verbs:        ["api-resources", "api-versions", "cluster-info", "describe", "diff", "explain", "get", "logs", "top", "auth"]
    ```
    
    ```
    # b.yamlsettings:  kubectl:    commands:      verbs: ["get", "logs", "top", "auth"]
    ```
    
    ```
    # resultsettings:  kubectl:    enabled: true    commands:      verbs: ["get", "logs", "top", "auth"]
    ```
