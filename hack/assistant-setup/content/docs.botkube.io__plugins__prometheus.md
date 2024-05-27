Title: Prometheus | Botkube

URL Source: https://docs.botkube.io/plugins/prometheus

Markdown Content:
The Botkube Prometheus source plugin allows you to fetch alerts from AlertManager of Prometheus deployment and notify in configured platforms.

You can enable the plugin as a part of Botkube instance configuration.

Once it is enabled, Botkube Prometheus plugin will consume Prometheus alerts and send them to configured platforms as shown below.

    # Prometheus endpoint without api version and resource.url: "http://localhost:9090"# If set as true, Prometheus source plugin will not send alerts that is created before plugin start time.ignoreOldAlerts: true# Only the alerts that have state provided in this config will be sent as notification. https://pkg.go.dev/github.com/prometheus/prometheus/rules#AlertStatealertStates: ["firing", "pending", "inactive"]# Logging configurationlog:  # Log level  level: info

For all collected `prometheus` sources bindings, configuration properties are overridden based on the order of the binding list for a given channel. The priority is given to the last binding specified on the list. Empty properties are omitted.
