Title: Botkube.io

URL Source: https://botkube.io/integration/prometheus

Markdown Content:
### Easy Prometheus Plugin Deployment

There has never been an easier way to get started with alert manager by Prometheus. If it is not already installed into your Kubernetes cluster, just simply go over to Botkube’s setup wizard and select the different K8s plugins you would like to add. No more CLI deployment needed!

![Image 1: Adding Alerts to your Kubernetes Cluster with Prometheus Alert manager ](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6549321f3ab68ce0bef6a1a0_POEEPnLdeD8pzH-pcTbUwkuHvE46ilC2nSw5zPvp4f1DuxdIBV3hbRC-hBnuEAYpCS7UzSU2bpsNdmWXFGIot6nhEtZskDiPOP7K7er6bOy4-1p2AKkjcJYNaYaGHClmW5rrmFsd8rzXLO6MIFJabZg.png)

How does Prometheus work?
-------------------------

Key benefits that Promehteus brought to the Cloud Native landscape include:

*   A multi-dimensional data model with time series data identified by metric name and key/value pairs
*   PromQL, a flexible query language to leverage this dimensionality
*   No reliance on distributed storage; single server nodes are autonomous
*   Time series collection happens via a pull model over HTTP
*   Pushing time series is supported via an intermediary gateway
*   Targets are discovered via service discovery or static configuration
*   Multiple modes of graphing and dashboarding support

All of these features are great among themselves, but the most impressive feature is that Promotheus syncs these to Alertmanager. This creates a trail of collected metrics that can be followed to draw helpful insights about how your K8s cluster is running.

Botkube can be used to follow this trail and offer helpful suggestions to fix common Kubernetes errors or alerts. If your team already uses alertmanager to create prometheus alerts.

Where did Promotheus come from?
-------------------------------

Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud. Since its inception in 2012, many companies and organizations have adopted Prometheus, and the project has a very active developer and user community. It is now a standalone open source project and maintained independently of any company. To emphasize this, and to clarify the project's governance structure, Prometheus joined the Cloud Native Computing Foundation in 2016 as the second hosted project, after Kubernetes.

Prometheus is 100% open source and community-driven. All components are available under the Apache 2 License on GitHub.
