Title: Elasticsearch notifier development | Botkube

URL Source: https://docs.botkube.io/community/contribute/elasticsearch-develop

Markdown Content:
*   [](https://docs.botkube.io/)
*   [Contribute](https://docs.botkube.io/community/contribute/)
*   Elasticsearch notifier development

Elasticsearch notifier development
----------------------------------

Botkube has stable support for Elasticsearch `v8` and it is backward compatible to Elasticsearch `v7`. If you are using Elasticsearch `v7`, please follow the [Elasticsearch v7 setup](#elasticsearch-v7-setup) section.

Elasticsearch v8 setup[​](#elasticsearch-v8-setup "Direct link to Elasticsearch v8 setup")
------------------------------------------------------------------------------------------

The easiest way to develop Botkube with Elasticsearch notifier enabled is to install Elasticsearch on your local Kubernetes cluster via [ECK](https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-deploy-eck.html).

### Steps[​](#steps "Direct link to Steps")

1.  Install Elasticsearch:
    
    1.  Install ECK custom resource definitions:
        
        ```
        kubectl create -f https://download.elastic.co/downloads/eck/2.9.0/crds.yaml
        ```
        
    2.  Install ECK operator:
        
        ```
        kubectl apply -f https://download.elastic.co/downloads/eck/2.9.0/operator.yaml
        ```
        
    3.  Deploy Elasticsearch:
        
        ```
        cat <<EOF | kubectl apply -f -apiVersion: elasticsearch.k8s.elastic.co/v1kind: Elasticsearchmetadata:  name: elasticsearchspec:  version: 8.9.1  nodeSets:  - name: default    count: 1    config:      node.store.allow_mmap: falseEOF
        ```
        
2.  Retrieve password
    
    ```
    PASSWORD=$(kubectl get secret elasticsearch-es-elastic-user -o go-template='{{.data.elastic | base64decode}}')
    ```
    
3.  Install Botkube with Elasticsearch, according to the [Elasticsearch installation](https://docs.botkube.io/installation/elasticsearch) instruction, where:
    
    *   `ELASTICSEARCH_USERNAME` is `elastic`,
    *   `ELASTICSEARCH_PASSWORD` is `$PASSWORD`,
    *   `ELASTICSEARCH_ADDRESS` is `https://elasticsearch-es-internal-http.default:9200`.
    
    You don't need to set index name, type, shards and replicas. Also, during Botkube installation, you need to use `--set communications.default-group.elasticsearch.skipTLSVerify=true` flag to skip TLS verification.
    

To review if the events are properly saved in Elasticsearch, follow these steps:

1.  Do port forward:
    
    ```
    kubectl port-forward svc/elasticsearch-es-internal-http 9200
    ```
    
2.  Fetch Elasticsearch indices:
    
    ```
    curl -u "elastic:$PASSWORD" -k https://localhost:9200/_cat/indices
    ```
    
3.  Copy the index name with the `botkube-` prefix and export it as environment variable. For example:
    
    ```
    export INDEX_NAME="botkube-2022-06-06"
    ```
    
4.  See Elasticsearch index details with logged events:
    
    ```
    curl -u "elastic:$PASSWORD" -k https://localhost:9200/$INDEX_NAME/_search\?pretty
    ```
    

Elasticsearch v7 setup[​](#elasticsearch-v7-setup "Direct link to Elasticsearch v7 setup")
------------------------------------------------------------------------------------------

The easiest way to develop Botkube with Elasticsearch notifier enabled is to install Elasticsearch on your local Kubernetes cluster.

### Steps[​](#steps-1 "Direct link to Steps")

1.  Install Elasticsearch:
    
    ```
    helm repo add elastic https://helm.elastic.cohelm install elasticsearch elastic/elasticsearch --version 7.17.3  --set replicas=1 --set resources.requests.cpu="100m" --set resources.requests.memory="512M" --wait
    ```
    
2.  Install Botkube with Elasticsearch, according to the [Elasticsearch installation](https://docs.botkube.io/installation/elasticsearch) instruction, where:
    
    *   `ELASTICSEARCH_USERNAME` is `elastic`,
    *   `ELASTICSEARCH_PASSWORD` is `changeme`,
    *   `ELASTICSEARCH_ADDRESS` is `http://elasticsearch-master.default:9200`.
    
    You don't need to set index name, type, shards and replicas.
    

To review if the events are properly saved in Elasticsearch, follow these steps:

1.  Do port forward:
    
    ```
    kubectl port-forward svc/elasticsearch-master 9200
    ```
    
2.  Fetch Elasticsearch indices:
    
    ```
    curl http://localhost:9200/_cat/indices
    ```
    
3.  Copy the index name with the `botkube-` prefix and export it as environment variable. For example:
    
    ```
    export INDEX_NAME="botkube-2022-06-06"
    ```
    
4.  See Elasticsearch index details with logged events:
    
    ```
    curl http://localhost:9200/$INDEX_NAME/_search\?pretty
    ```
    

[](https://docs.botkube.io/community/contribute/)[Outgoing Webhook development](https://docs.botkube.io/community/contribute/webhook-develop)

*   [Elasticsearch v8 setup](#elasticsearch-v8-setup)
    *   [Steps](#steps)
*   [Elasticsearch v7 setup](#elasticsearch-v7-setup)
    *   [Steps](#steps-1)
