Title: General | Botkube

URL Source: https://docs.botkube.io/self-hosted-configuration/general

Markdown Content:
The general settings holds a general configuration for the Botkube Agent. For example, log level, disabling config watcher and similar.

    # General Botkube configuration.settings:  # Cluster name to differentiate incoming messages.  clusterName: not-configured  # If true, notifies about new Botkube releases.  upgradeNotifier: true  # Botkube logging settings.  log:    # Sets one of the log levels. Allowed values: `info`, `warn`, `debug`, `error`, `fatal`, `panic`.    level: info    # Configures log format. Allowed values: `text`, `json`.    formatter: json    # If true, disable ANSI colors in logging. Ignored when `json` formatter is used.    disableColors: false# Parameters for the Config Watcher component which reloads Botkube on ConfigMap changes.# It restarts Botkube when configuration data change is detected. It watches ConfigMaps and/or Secrets with the `botkube.io/config-watch: "true"` label from the namespace where Botkube is installed.configWatcher:  # If true, restarts the Botkube Pod on config changes.  enabled: true  # In-cluster Config Watcher configuration. It is used when remote configuration is not provided.  inCluster:    # Resync period for the Config Watcher informers.    informerResyncPeriod: 10m
