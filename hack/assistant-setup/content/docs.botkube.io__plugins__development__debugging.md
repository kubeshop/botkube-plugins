Title: Debugging | Botkube

URL Source: https://docs.botkube.io/plugins/development/debugging

Markdown Content:
Debugging | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/plugins/development/debugging#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/plugins/development/debugging)
*   [1.12](https://docs.botkube.io/plugins/development/debugging)
*   [1.11](https://docs.botkube.io/1.11/plugins/development/debugging)
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
        
        *   [Quick start](https://docs.botkube.io/plugins/development/quick-start)
        *   [Custom executor](https://docs.botkube.io/plugins/development/custom-executor)
        *   [Custom source](https://docs.botkube.io/plugins/development/custom-source)
        *   [Using kubeconfig](https://docs.botkube.io/plugins/development/using-kubeconfig)
        *   [Interactive messages](https://docs.botkube.io/plugins/development/interactive-messages)
        *   [Dependencies](https://docs.botkube.io/plugins/development/dependencies)
        *   [Local testing](https://docs.botkube.io/plugins/development/local-testing)
        *   [Repository](https://docs.botkube.io/plugins/development/repo)
        *   [Troubleshooting](https://docs.botkube.io/plugins/development/troubleshooting)
        *   [Debugging](https://docs.botkube.io/plugins/development/debugging)
*   [CLI](https://docs.botkube.io/cli/getting-started)
    
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
    
*   [Troubleshooting](https://docs.botkube.io/troubleshooting/common-problems)
    
*   [Architecture](https://docs.botkube.io/architecture/)
    

*   [](https://docs.botkube.io/)
*   [Plugins](https://docs.botkube.io/plugins/)
*   [Plugin development](https://docs.botkube.io/plugins/development/)
*   Debugging

Version: 1.12

Debugging
=========

Embedded logging allows you to access more information about the runtime operations of Botkube plugins.

By default:

*   the gRPC client log level is set to `info`,
*   the standard error ([`stderr`](https://en.wikipedia.org/wiki/Standard_streams#Standard_error_\(stderr\))) of a plugin binary is logged at `error` level,
*   the standard output ([`stdout`](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_\(stdout\))) of a plugin binary is ignored.

To change the default log level, export a dedicated environment variable following this pattern `LOG_LEVEL_{pluginType}_{pluginRepo}_{pluginName}`, e.g., `LOG_LEVEL_EXECUTOR_BOTKUBE_KUBECTL`. The possible log level values are:

*   `trace`
*   `debug`
*   `info`
*   `warning`
*   `error`
*   `fatal`
*   `panic`

The plugin standard output is logged only if `debug` level is set.

info

The plugin name is normalized and all characters different from letters, digits, and the underscore (`_`) are replaced with underscore (`_`).

To change the log level for a given plugin directly in the Botkube deployment, specify `extraEnv` in the [values.yaml](https://github.com/kubeshop/botkube/blob/main/helm/botkube/values.yaml) file. For example:

```
extraEnv:  - name: LOG_LEVEL_EXECUTOR_BOTKUBE_KUBECTL    value: "debug"
```

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/plugins/development/debug.md)

[Previous Troubleshooting](https://docs.botkube.io/plugins/development/troubleshooting)[Next Getting Started](https://docs.botkube.io/cli/getting-started)

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
