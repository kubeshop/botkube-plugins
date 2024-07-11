Title: Keptn | Botkube

URL Source: https://docs.botkube.io/plugins/keptn

Markdown Content:
Keptn | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/plugins/keptn#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/plugins/keptn)
*   [1.12](https://docs.botkube.io/plugins/keptn)
*   [1.11](https://docs.botkube.io/1.11/plugins/keptn)
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
*   Keptn

Version: 1.12

On this page

Keptn
=====

info

**This plugin is hosted by the [Botkube Cloud](https://app.botkube.io/) plugin repository and requires active Botkube Cloud account.**

The Botkube Keptn source plugin allows you to consume events from Keptn deployment and notify in configured platforms.

Get started[​](https://docs.botkube.io/plugins/keptn#get-started "Direct link to Get started")
----------------------------------------------------------------------------------------------

### Enable the plugin[​](https://docs.botkube.io/plugins/keptn#enable-the-plugin "Direct link to Enable the plugin")

You can enable the plugin as a part of Botkube instance configuration.

1.  If you don't have an existing Botkube instance, create a new one, according to the [Installation](https://docs.botkube.io/) docs.
2.  From the [Botkube Cloud homepage](https://app.botkube.io/), click on a card of a given Botkube instance.
3.  Navigate to the platform tab which you want to configure.
4.  Click **Add plugin** button.
5.  Select the Keptn plugin.
6.  Click **Save** button.

Usage[​](https://docs.botkube.io/plugins/keptn#usage "Direct link to Usage")
----------------------------------------------------------------------------

Once it is enabled, Botkube Keptn plugin will consume Keptn events and send them to configured platforms as shown below.

![Image 3: Keptn Events](https://docs.botkube.io/assets/images/keptn-events-0758acae069a3caf19ba8b0451a5c7cc.png)

Configuration[​](https://docs.botkube.io/plugins/keptn#configuration "Direct link to Configuration")
----------------------------------------------------------------------------------------------------

This plugin supports the following configuration:

```
# Keptn API Gateway URL.url: "http://api-gateway-nginx.keptn.svc.cluster.local/api"# Keptn API Token to access events through API Gateway.token: ""# Optional Keptn project.project: ""# Optional Keptn Service name under the project.service: ""# Logging configurationlog:  # Log level  level: info
```

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/plugins/keptn.md)

[Previous ArgoCD](https://docs.botkube.io/plugins/argocd)[Next Plugin development](https://docs.botkube.io/plugins/development/)

*   [Get started](https://docs.botkube.io/plugins/keptn#get-started)
    *   [Enable the plugin](https://docs.botkube.io/plugins/keptn#enable-the-plugin)
*   [Usage](https://docs.botkube.io/plugins/keptn#usage)
*   [Configuration](https://docs.botkube.io/plugins/keptn#configuration)

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
