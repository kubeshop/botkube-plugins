Title: Outgoing webhook for self-hosted Botkube | Botkube

URL Source: https://docs.botkube.io/installation/webhook/self-hosted

Markdown Content:
Botkube can be integrated with external apps via Webhooks. A webhook is essentially a POST request sent to a callback URL. So you can configure Botkube to send events on specified URL.

    export CLUSTER_NAME={cluster_name}export WEBHOOK_URL={url}botkube install --version v1.12.0 \--set communications.default-group.webhook.enabled=true \--set communications.default-group.webhook.url=${WEBHOOK_URL} \--set settings.clusterName=${CLUSTER_NAME}

Configuration syntax is explained [here](https://docs.botkube.io/self-hosted-configuration). All possible installation parameters are documented [here](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters).

Execute the following command to completely remove Botkube and related resources from your cluster:
