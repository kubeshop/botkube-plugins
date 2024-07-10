Title: Outgoing webhook for self-hosted Botkube | Botkube

URL Source: https://docs.botkube.io/installation/webhook/self-hosted

Markdown Content:
Outgoing webhook for self-hosted Botkube | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/installation/webhook/self-hosted#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/installation/webhook/self-hosted)
*   [1.12](https://docs.botkube.io/installation/webhook/self-hosted)
*   [1.11](https://docs.botkube.io/1.11/installation/webhook/self-hosted)
*   [1.10](https://docs.botkube.io/1.10/installation/webhook/self-hosted)
*   [1.9](https://docs.botkube.io/1.9/installation/webhook/self-hosted)
*   [1.8](https://docs.botkube.io/1.8/installation/webhook/self-hosted)
*   [1.7](https://docs.botkube.io/1.7/installation/webhook/self-hosted)
*   * * *
    
*   [All versions](https://docs.botkube.io/versions)

[GitHub](https://github.com/kubeshop/botkube)[Slack](https://join.botkube.io/)

Search

*   [Installation](https://docs.botkube.io/)
    
    *   [Slack](https://docs.botkube.io/installation/socketslack)
    *   [Slack](https://docs.botkube.io/installation/slack/)
        
    *   [Mattermost](https://docs.botkube.io/installation/mattermost/)
        
    *   [Discord](https://docs.botkube.io/installation/discord/)
        
    *   [Microsoft Teams](https://docs.botkube.io/installation/teams/)
    *   [Elasticsearch](https://docs.botkube.io/installation/elasticsearch/)
        
    *   [PagerDuty](https://docs.botkube.io/installation/pagerduty/)
    *   [Outgoing webhook](https://docs.botkube.io/installation/webhook/)
        
        *   [Outgoing webhook for self-hosted Botkube](https://docs.botkube.io/installation/webhook/self-hosted)
        *   [Outgoing webhook for Botkube Cloud](https://docs.botkube.io/installation/webhook/cloud)
*   [Tutorials and examples](https://docs.botkube.io/examples-and-tutorials/)
    
*   [Features](https://docs.botkube.io/features/event-notifications)
    
*   [Plugins](https://docs.botkube.io/plugins/)
    
*   [CLI](https://docs.botkube.io/cli/getting-started)
    
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
    
*   [Troubleshooting](https://docs.botkube.io/troubleshooting/common-problems)
    
*   [Architecture](https://docs.botkube.io/architecture/)
    

*   [](https://docs.botkube.io/)
*   [Installation](https://docs.botkube.io/)
*   [Outgoing webhook](https://docs.botkube.io/installation/webhook/)
*   Outgoing webhook for self-hosted Botkube

Version: 1.12

On this page

Outgoing webhook for self-hosted Botkube
========================================

Install Botkube in Kubernetes cluster[​](https://docs.botkube.io/installation/webhook/self-hosted#install-botkube-in-kubernetes-cluster "Direct link to Install Botkube in Kubernetes cluster")
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Botkube can be integrated with external apps via Webhooks. A webhook is essentially a POST request sent to a callback URL. So you can configure Botkube to send events on specified URL.

To deploy Botkube agent in your cluster, run:

```
export CLUSTER_NAME={cluster_name}export WEBHOOK_URL={url}botkube install --version v1.12.0 \--set communications.default-group.webhook.enabled=true \--set communications.default-group.webhook.url=${WEBHOOK_URL} \--set settings.clusterName=${CLUSTER_NAME}
```

where:

*   **WEBHOOK\_URL** is an outgoing webhook URL to which Botkube will POST the events,
*   **CLUSTER\_NAME** is the cluster name set in the incoming messages.

Configuration syntax is explained [here](https://docs.botkube.io/self-hosted-configuration). All possible installation parameters are documented [here](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters).

Remove Botkube from Kubernetes cluster[​](https://docs.botkube.io/installation/webhook/self-hosted#remove-botkube-from-kubernetes-cluster "Direct link to Remove Botkube from Kubernetes cluster")
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Execute the following command to completely remove Botkube and related resources from your cluster:

```
botkube uninstall
```

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/installation/webhook/self-hosted.md)

[Previous Outgoing webhook](https://docs.botkube.io/installation/webhook/)[Next Outgoing webhook for Botkube Cloud](https://docs.botkube.io/installation/webhook/cloud)

*   [Install Botkube in Kubernetes cluster](https://docs.botkube.io/installation/webhook/self-hosted#install-botkube-in-kubernetes-cluster)
*   [Remove Botkube from Kubernetes cluster](https://docs.botkube.io/installation/webhook/self-hosted#remove-botkube-from-kubernetes-cluster)

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
