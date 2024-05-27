Title: Keptn | Botkube

URL Source: https://docs.botkube.io/plugins/keptn

Markdown Content:
The Botkube Keptn source plugin allows you to consume events from Keptn deployment and notify in configured platforms.

You can enable the plugin as a part of Botkube instance configuration.

Once it is enabled, Botkube Keptn plugin will consume Keptn events and send them to configured platforms as shown below.

    # Keptn API Gateway URL.url: "http://api-gateway-nginx.keptn.svc.cluster.local/api"# Keptn API Token to access events through API Gateway.token: ""# Optional Keptn project.project: ""# Optional Keptn Service name under the project.service: ""# Logging configurationlog:  # Log level  level: info
