Title: Elasticsearch for self-hosted Botkube | Botkube

URL Source: https://docs.botkube.io/installation/elasticsearch/self-hosted

Markdown Content:
Elasticsearch for self-hosted Botkube | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/installation/elasticsearch/self-hosted#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/installation/elasticsearch/self-hosted)
*   [1.12](https://docs.botkube.io/installation/elasticsearch/self-hosted)
*   [1.11](https://docs.botkube.io/1.11/installation/elasticsearch/self-hosted)
*   [1.10](https://docs.botkube.io/1.10/installation/elasticsearch/self-hosted)
*   [1.9](https://docs.botkube.io/1.9/installation/elasticsearch/self-hosted)
*   [1.8](https://docs.botkube.io/1.8/installation/elasticsearch/self-hosted)
*   [1.7](https://docs.botkube.io/1.7/installation/elasticsearch/self-hosted)
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
        
        *   [Elasticsearch for self-hosted Botkube](https://docs.botkube.io/installation/elasticsearch/self-hosted)
        *   [Elasticsearch for Botkube Cloud](https://docs.botkube.io/installation/elasticsearch/cloud)
    *   [PagerDuty](https://docs.botkube.io/installation/pagerduty/)
    *   [Outgoing webhook](https://docs.botkube.io/installation/webhook/)
        
*   [Tutorials and examples](https://docs.botkube.io/examples-and-tutorials/)
    
*   [Features](https://docs.botkube.io/features/event-notifications)
    
*   [Plugins](https://docs.botkube.io/plugins/)
    
*   [CLI](https://docs.botkube.io/cli/getting-started)
    
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
    
*   [Troubleshooting](https://docs.botkube.io/troubleshooting/common-problems)
    
*   [Architecture](https://docs.botkube.io/architecture/)
    

*   [](https://docs.botkube.io/)
*   [Installation](https://docs.botkube.io/)
*   [Elasticsearch](https://docs.botkube.io/installation/elasticsearch/)
*   Elasticsearch for self-hosted Botkube

Version: 1.12

On this page

Elasticsearch for self-hosted Botkube
=====================================

Prerequisites[​](https://docs.botkube.io/installation/elasticsearch/self-hosted#prerequisites "Direct link to Prerequisites")
-----------------------------------------------------------------------------------------------------------------------------

*   Botkube CLI installed according to the [Getting Started guide](https://docs.botkube.io/cli/getting-started#installation)
*   Access to Kubernetes cluster
*   Elasticsearch server

Install Botkube in Kubernetes cluster[​](https://docs.botkube.io/installation/elasticsearch/self-hosted#install-botkube-in-kubernetes-cluster "Direct link to Install Botkube in Kubernetes cluster")
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

To deploy Botkube agent in your cluster, run:

```
export CLUSTER_NAME={cluster_name}export ELASTICSEARCH_ADDRESS={elasticsearch_address}export ELASTICSEARCH_USERNAME={elasticsearch_username}export ELASTICSEARCH_PASSWORD={elasticsearch_password}export ELASTICSEARCH_INDEX_NAME={elasticsearch_index_name}botkube install --version v1.12.0 \--set communications.default-group.elasticsearch.enabled=true \--set communications.default-group.elasticsearch.server=${ELASTICSEARCH_ADDRESS} \--set communications.default-group.elasticsearch.username=${ELASTICSEARCH_USERNAME} \--set communications.default-group.elasticsearch.password=${ELASTICSEARCH_PASSWORD} \--set communications.default-group.elasticsearch.indices.default.name=${ELASTICSEARCH_INDEX_NAME} \--set settings.clusterName=${CLUSTER_NAME}
```

where:

*   **ELASTICSEARCH\_ADDRESS** is an address on which Elasticsearch server is reachable e.g [https://example.com:9243](https://example.com:9243/),
*   **ELASTICSEARCH\_USERNAME** is the username for authentication to Els server,
*   **ELASTICSEARCH\_PASSWORD** is a password for the username to authenticate with Els server,
*   **ELASTICSEARCH\_INDEX\_NAME** _(optional)_ is an index name on which Botkube events will be stored _(default: botkube)_.

Configuration syntax is explained [here](https://docs.botkube.io/self-hosted-configuration). All possible installation parameters are documented [here](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters).

Remove Botkube from Kubernetes cluster[​](https://docs.botkube.io/installation/elasticsearch/self-hosted#remove-botkube-from-kubernetes-cluster "Direct link to Remove Botkube from Kubernetes cluster")
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Execute the following command to completely remove Botkube and related resources from your cluster:

```
botkube uninstall
```

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/installation/elasticsearch/self-hosted.md)

[Previous Elasticsearch](https://docs.botkube.io/installation/elasticsearch/)[Next Elasticsearch for Botkube Cloud](https://docs.botkube.io/installation/elasticsearch/cloud)

*   [Prerequisites](https://docs.botkube.io/installation/elasticsearch/self-hosted#prerequisites)
*   [Install Botkube in Kubernetes cluster](https://docs.botkube.io/installation/elasticsearch/self-hosted#install-botkube-in-kubernetes-cluster)
*   [Remove Botkube from Kubernetes cluster](https://docs.botkube.io/installation/elasticsearch/self-hosted#remove-botkube-from-kubernetes-cluster)

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
