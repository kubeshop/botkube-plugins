Title: Source | Botkube

URL Source: https://docs.botkube.io/self-hosted-configuration/source

Markdown Content:
Version: 1.12

info

This document is applicable only for self-hosted installations. For Botkube Cloud installations, the Botkube Agent configuration is managed via the [Botkube Cloud dashboard](https://app.botkube.io/).

The source configuration allows you to define multiple source configurations that can be later referred in [communication](https://docs.botkube.io/self-hosted-configuration/communication) bindings. For example, take a look at such source definition:

```
sources:  "k8s-recommendation-alerts": # This is a source configuration name, which is referred in communication bindings.    botkube/kubernetes:      # ... trimmed ...  "k8s-all-events": # This is a source configuration name, which is referred in communication bindings.    botkube/kubernetes:      # ... trimmed ...
```

This can be later used by the communication platforms:

```
communications:  "default-group":    socketSlack:      channels:        "default":          bindings:            sources: # The order is important for merging strategy.              - k8s-recommendation-events # The source configuration name              - k8s-all-events # The source configuration name          # ... trimmed ...
```

To use the Botkube source plugins, first you need to define the plugins repository under the `plugins` property:

```
plugins:  repositories:    repo-name:      url: https://example.com/plugins-index.yaml
```

Next, you can configure source from a given repository:

```
sources:  "plugins":    repo-name/source-name@v1.0.0: # Plugin name syntax: <repo>/<plugin>[@<version>]. If version is not provided, the latest version from repository is used.      enabled: true      config: {}
```

For all source configuration properties, see the [**syntax**](https://docs.botkube.io/self-hosted-configuration/source/#syntax) section.

#### Restart Policy and Health Check Interval[​](https://docs.botkube.io/self-hosted-configuration/source/#restart-policy-and-health-check-interval "Direct link to Restart Policy and Health Check Interval")

This section of the configuration allows you to configure the restart policy for the Botkube source plugins. The restart policy is used when the source plugin fails to start. The default restart policy is `DeactivatePlugin`, which means that the plugin is deactivated after a given number of restarts. The restart policy can be configured with the following properties:

*   `type` - restart policy type. Allowed values: `RestartAgent`, `DeactivatePlugin`.
*   `threshold` - number of restarts before the policy takes into effect.

Restart policy types:

*   `RestartAgent` - when the threshold is reached, the Botkube agent is restarted.
*   `DeactivatePlugin` - when the threshold is reached, the plugin is deactivated. To activate the plugin again, you need to restart the Botkube agent.

The health check interval is used to check the health of the source plugins. The default health check interval is 10 seconds. The health check interval can be configured with the following property:

*   `healthCheckInterval` - health check interval.

```
# Botkube Restart Policy on plugin failure.restartPolicy:  # Restart policy type. Allowed values: "RestartAgent", "DeactivatePlugin".  type: "DeactivatePlugin"  # Number of restarts before policy takes into effect.  threshold: 10healthCheckInterval: 10s
```

Syntax[​](https://docs.botkube.io/self-hosted-configuration/source/#syntax "Direct link to Syntax")
---------------------------------------------------------------------------------------------------

```
# Map of sources. The `sources` property name is an alias for a given configuration.# Key name is used as a binding reference.## Format: sources.{alias}sources:  "k8s-recommendation-events":    # Built-in kubernetes source configuration.    botkube/kubernetes:      enabled: true      config:        # Kubernetes configuration        recommendations:          pod:            noLatestImageTag: true        # ... trimmed ...# Configuration for Botkube executors and sources plugins.plugins:  # Directory, where downloaded plugins are cached.  cacheDir: "/tmp"  # List of plugins repositories.  repositories:    # This repository serves officially supported Botkube plugins.    botkube:      url: https://github.com/kubeshop/botkube/releases/download/v1.12.0/plugins-index.yaml    # Other 3rd party repositories.    repo-name:      url: https://example.com/plugins-index.yaml  # Configure Incoming webhook for source plugins.  incomingWebhook:    enabled: true    port: 2115    targetPort: 2115  # Botkube Restart Policy on plugin failure.  restartPolicy:    # Restart policy type. Allowed values: "RestartAgent", "DeactivatePlugin".    type: "DeactivatePlugin"    # Number of restarts before policy takes into effect.    threshold: 10  healthCheckInterval: 10s
```

The default configuration for the Botkube Helm chart can be found in the [values.yaml](https://github.com/kubeshop/botkube/blob/main/helm/botkube/values.yaml) file.
