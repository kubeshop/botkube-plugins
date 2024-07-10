Title: Prometheus | Botkube

URL Source: https://docs.botkube.io/plugins/prometheus

Markdown Content:
Prometheus | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/plugins/prometheus#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/plugins/prometheus)
*   [1.12](https://docs.botkube.io/plugins/prometheus)
*   [1.11](https://docs.botkube.io/1.11/plugins/prometheus)
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
    
    *   [Kubernetes](https://docs.botkube.io/plugins/kubernetes)
    *   [Kubectl](https://docs.botkube.io/plugins/kubectl)
    *   [AI Assistant](https://docs.botkube.io/plugins/ai-assistant)
    *   [Exec](https://docs.botkube.io/plugins/exec)
    *   [Helm](https://docs.botkube.io/plugins/helm)
    *   [Prometheus](https://docs.botkube.io/plugins/prometheus)
    *   [GitHub Events](https://docs.botkube.io/plugins/github-events)
    *   [Flux](https://docs.botkube.io/plugins/flux)
    *   [ArgoCD](https://docs.botkube.io/plugins/argocd)
    *   [Keptn](https://docs.botkube.io/plugins/keptn)
    *   [Plugin development](https://docs.botkube.io/plugins/development/)
        
*   [CLI](https://docs.botkube.io/cli/getting-started)
    
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
    
*   [Troubleshooting](https://docs.botkube.io/troubleshooting/common-problems)
    
*   [Architecture](https://docs.botkube.io/architecture/)
    

*   [](https://docs.botkube.io/)
*   [Plugins](https://docs.botkube.io/plugins/)
*   Prometheus

Version: 1.12

On this page

Prometheus
==========

info

**This plugin is hosted by the [Botkube Cloud](https://app.botkube.io/) plugin repository and requires active Botkube Cloud account.**

The Botkube Prometheus source plugin allows you to fetch alerts from AlertManager of Prometheus deployment and notify in configured platforms.

Get started[​](https://docs.botkube.io/plugins/prometheus#get-started "Direct link to Get started")
---------------------------------------------------------------------------------------------------

### Enable the plugin[​](https://docs.botkube.io/plugins/prometheus#enable-the-plugin "Direct link to Enable the plugin")

You can enable the plugin as a part of Botkube instance configuration.

1.  If you don't have an existing Botkube instance, create a new one, according to the [Installation](https://docs.botkube.io/) docs.
2.  From the [Botkube Cloud homepage](https://app.botkube.io/), click on a card of a given Botkube instance.
3.  Navigate to the platform tab which you want to configure.
4.  Click **Add plugin** button.
5.  Select the Prometheus plugin.
6.  Provide the Prometheus endpoint according to the [Configuration](https://docs.botkube.io/plugins/prometheus#configuration) section.
7.  Click **Save** button.

Usage[​](https://docs.botkube.io/plugins/prometheus#usage "Direct link to Usage")
---------------------------------------------------------------------------------

Once it is enabled, Botkube Prometheus plugin will consume Prometheus alerts and send them to configured platforms as shown below.

![Image 3: Prometheus Alerts](https://docs.botkube.io/assets/images/prometheus-alerts-6ebb9e0ebbaa0f2e19530612da120e6b.png)

Configuration[​](https://docs.botkube.io/plugins/prometheus#configuration "Direct link to Configuration")
---------------------------------------------------------------------------------------------------------

This plugin supports the following configuration:

```
# Prometheus endpoint without api version and resource.url: "http://localhost:9090"# If set as true, Prometheus source plugin will not send alerts that is created before plugin start time.ignoreOldAlerts: true# Only the alerts that have state provided in this config will be sent as notification. https://pkg.go.dev/github.com/prometheus/prometheus/rules#AlertStatealertStates: ["firing", "pending", "inactive"]# Logging configurationlog:  # Log level  level: info
```

Merging strategy[​](https://docs.botkube.io/plugins/prometheus#merging-strategy "Direct link to Merging strategy")
------------------------------------------------------------------------------------------------------------------

For all collected `prometheus` sources bindings, configuration properties are overridden based on the order of the binding list for a given channel. The priority is given to the last binding specified on the list. Empty properties are omitted.

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/plugins/prometheus.md)

[Previous Helm](https://docs.botkube.io/plugins/helm)[Next GitHub Events](https://docs.botkube.io/plugins/github-events)

*   [Get started](https://docs.botkube.io/plugins/prometheus#get-started)
    *   [Enable the plugin](https://docs.botkube.io/plugins/prometheus#enable-the-plugin)
*   [Usage](https://docs.botkube.io/plugins/prometheus#usage)
*   [Configuration](https://docs.botkube.io/plugins/prometheus#configuration)
*   [Merging strategy](https://docs.botkube.io/plugins/prometheus#merging-strategy)

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
