Title: Installation | Botkube

URL Source: https://docs.botkube.io/

Markdown Content:
*   [](https://docs.botkube.io/)
*   Installation

Version: 1.10

Botkube has two components that need to be installed.

1.  Botkube App Integration in your Slack/Mattermost/Microsoft Teams/Discord
2.  Botkube agent in your Kubernetes cluster

Feature map[​](#feature-map"DirectlinktoFeaturemap")
---------------------------------------------------------

Learn about Botkube features and their availability in different integrations.

### Bots[​](#bots"DirectlinktoBots")

Compare our bidirectional integrations:

| Feature | Cloud Slack | Slack | Microsoft Teams | Discord | Mattermost |
| --- | --- | --- | --- | --- | --- |
| Source plugins support (e.g. `kubernetes`) | ✔️ | ✔️ | ✔️ | ✔️ | ✔️ |
| Executor plugins support (e.g. `kubectl`) | ✔️ | ✔️ | ✔️ | ✔️ | ✔️ |
| Multi-cluster support | ✔️ | ❌ | ✔️ | ✔️ | ✔️ |
| Enhanced per-channel plugin configuration including RBAC policy. | ✔️ | ✔️ | ✔️ | ✔️ | ✔️ |
| Interactive messages | ✔️ | ✔️ | ✔️ | ❌ | ❌ |
| Actionable notifications | ✔️ | ✔️ | ✔️ | ❌ | ❌ |
| Emoji reactions | ✔️ | ❌ | ❌ | ❌ | ❌ |

### Sinks[​](#sinks"DirectlinktoSinks")

Compare our unidirectional integrations:

| Feature | Elasticsearch | Webhook |
| --- | --- | --- |
| Source plugins support (e.g. `kubernetes`, `prometheus`, etc.) | ✔️ | ✔️ |
| Multi-cluster support | ✔️ | ✔️ |

Integrations[​](#integrations"DirectlinktoIntegrations")
------------------------------------------------------------

tip

You can use a single Botkube agent to serve all the interfaces - Slack, Mattermost, Microsoft Teams, Elasticsearch and Webhook.  
You just need to enable required mediums through the settings and add a necessary configuration.  
_see the [configuration](https://docs.botkube.io/configuration) section for more information_

[Next Slack](https://docs.botkube.io/installation/socketslack)
