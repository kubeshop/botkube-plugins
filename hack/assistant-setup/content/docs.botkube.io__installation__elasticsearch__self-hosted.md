Title: Elasticsearch for self-hosted Botkube | Botkube

URL Source: https://docs.botkube.io/installation/elasticsearch/self-hosted

Markdown Content:
```
export CLUSTER_NAME={cluster_name}export ELASTICSEARCH_ADDRESS={elasticsearch_address}export ELASTICSEARCH_USERNAME={elasticsearch_username}export ELASTICSEARCH_PASSWORD={elasticsearch_password}export ELASTICSEARCH_INDEX_NAME={elasticsearch_index_name}botkube install --version v1.12.0 \--set communications.default-group.elasticsearch.enabled=true \--set communications.default-group.elasticsearch.server=${ELASTICSEARCH_ADDRESS} \--set communications.default-group.elasticsearch.username=${ELASTICSEARCH_USERNAME} \--set communications.default-group.elasticsearch.password=${ELASTICSEARCH_PASSWORD} \--set communications.default-group.elasticsearch.indices.default.name=${ELASTICSEARCH_INDEX_NAME} \--set settings.clusterName=${CLUSTER_NAME}
```

Configuration syntax is explained [here](https://docs.botkube.io/self-hosted-configuration). All possible installation parameters are documented [here](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters).

Execute the following command to completely remove Botkube and related resources from your cluster:
