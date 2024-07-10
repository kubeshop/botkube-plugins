Title: Installation | Botkube

URL Source: https://docs.botkube.io/

Markdown Content:
What is Botkube?[​](https://docs.botkube.io/#what-is-botkube "Direct link to What is Botkube?")
-----------------------------------------------------------------------------------------------

Botkube is a troubleshooting and monitoring solution that empowers DevOps teams to work more efficiently, enables developers to troubleshoot their applications without special Kubernetes access or knowledge, and improves reliability by delivering timely, context-enhanced notifications about events in your Kubernetes environments.

It integrates with multiple platforms, such as Slack, Microsoft Teams, Discord and Mattermost. Botkube monitors events from various sources, allows you to securely run commands, and run automated actions triggered by any of the plugins Botkube supports or your own custom plugins.

Feature map[​](https://docs.botkube.io/#feature-map "Direct link to Feature map")
---------------------------------------------------------------------------------

Learn about Botkube features and their availability in different integrations.

### Bots[​](https://docs.botkube.io/#bots "Direct link to Bots")

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

### Sinks[​](https://docs.botkube.io/#sinks "Direct link to Sinks")

Compare our unidirectional integrations:

| Feature | Elasticsearch | Webhook | PagerDuty |
| --- | --- | --- | --- |
| Source plugins support (e.g. `kubernetes`, `prometheus`, etc.) | ✔️ | ✔️ | ✔️ |
| Multi-cluster support | ✔️ | ✔️ | ✔️ |

Integrations[​](https://docs.botkube.io/#integrations "Direct link to Integrations")
------------------------------------------------------------------------------------

tip

You can enable multiple platform integrations using a single Botkube agent.
