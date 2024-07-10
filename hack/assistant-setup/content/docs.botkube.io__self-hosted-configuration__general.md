Title: General | Botkube

URL Source: https://docs.botkube.io/self-hosted-configuration/general

Markdown Content:
General | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/self-hosted-configuration/general#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/self-hosted-configuration/general)
*   [1.12](https://docs.botkube.io/self-hosted-configuration/general)
*   [1.11](https://docs.botkube.io/1.11/self-hosted-configuration/general)
*   [1.10](https://docs.botkube.io/1.10/)
*   [1.9](https://docs.botkube.io/1.9/)
*   [1.8](https://docs.botkube.io/1.8/)
*   [1.7](https://docs.botkube.io/1.7/)
*   * * *
    
*   [All versions](https://docs.botkube.io/versions)

[GitHub](https://github.com/kubeshop/botkube)[Slack](https://join.botkube.io/)

Search

*   [Installation](https://docs.botkube.io/)
    
*   [Tutorials and examples](https://docs.botkube.io/examples-and-tutorials/)
    
*   [Features](https://docs.botkube.io/features/event-notifications)
    
*   [Plugins](https://docs.botkube.io/plugins/)
    
*   [CLI](https://docs.botkube.io/cli/getting-started)
    
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
    
    *   [Communication](https://docs.botkube.io/self-hosted-configuration/communication/)
        
    *   [Source](https://docs.botkube.io/self-hosted-configuration/source)
    *   [Executor](https://docs.botkube.io/self-hosted-configuration/executor)
    *   [General](https://docs.botkube.io/self-hosted-configuration/general)
    *   [Helm chart parameters](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters)
*   [Troubleshooting](https://docs.botkube.io/troubleshooting/common-problems)
    
*   [Architecture](https://docs.botkube.io/architecture/)
    

*   [](https://docs.botkube.io/)
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
*   General

Version: 1.12

On this page

General
=======

info

This document is applicable only for self-hosted installations. For Botkube Cloud installations, the Botkube Agent configuration is managed via the [Botkube Cloud dashboard](https://app.botkube.io/).

The general settings holds a general configuration for the Botkube Agent. For example, log level, disabling config watcher and similar.

Syntax[​](https://docs.botkube.io/self-hosted-configuration/general#syntax "Direct link to Syntax")
---------------------------------------------------------------------------------------------------

```
# General Botkube configuration.settings:  # Cluster name to differentiate incoming messages.  clusterName: not-configured  # If true, notifies about new Botkube releases.  upgradeNotifier: true  # Botkube logging settings.  log:    # Sets one of the log levels. Allowed values: `info`, `warn`, `debug`, `error`, `fatal`, `panic`.    level: info    # Configures log format. Allowed values: `text`, `json`.    formatter: json    # If true, disable ANSI colors in logging. Ignored when `json` formatter is used.    disableColors: false# Parameters for the Config Watcher component which reloads Botkube on ConfigMap changes.# It restarts Botkube when configuration data change is detected. It watches ConfigMaps and/or Secrets with the `botkube.io/config-watch: "true"` label from the namespace where Botkube is installed.configWatcher:  # If true, restarts the Botkube Pod on config changes.  enabled: true  # In-cluster Config Watcher configuration. It is used when remote configuration is not provided.  inCluster:    # Resync period for the Config Watcher informers.    informerResyncPeriod: 10m
```

The default configuration for Helm chart can be found in the [values.yaml](https://github.com/kubeshop/botkube/blob/main/helm/botkube/values.yaml) file.

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/self-hosted-configuration/general.md)

[Previous Executor](https://docs.botkube.io/self-hosted-configuration/executor)[Next Helm chart parameters](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters)

*   [Syntax](https://docs.botkube.io/self-hosted-configuration/general#syntax)

Community

*   [Contribute](https://docs.botkube.io/community/contribute)
*   [GitHub](https://github.com/kubeshop/botkube)
*   [Community Slack](https://join.botkube.io/)
*   [Support](https://docs.botkube.io/support)

Legal

*   [License](https://docs.botkube.io/license)
*   [Privacy & Legal](https://botkube.io/privacy-policy)
*   [Telemetry](https://docs.botkube.io/telemetry)

Learn

*   [Installation](https://docs.botkube.io/)

Social

*   [Twitter](https://twitter.com/Botkube_io)

Copyright © 2024 Kubeshop, Inc. Built with Docusaurus.
