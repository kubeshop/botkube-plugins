Title: Slack | Botkube

URL Source: https://docs.botkube.io/installation/slack/

Markdown Content:
Botkube App for Slack Versions[​](https://docs.botkube.io/installation/slack/#botkube-app-for-slack-versions "Direct link to Botkube App for Slack Versions")
-------------------------------------------------------------------------------------------------------------------------------------------------------------

In order to connect with Slack, Botkube requires an application for Slack. There are two versions of the Botkube App for Slack available:

### Botkube Cloud App for Slack[​](https://docs.botkube.io/installation/slack/#botkube-cloud-app-for-slack "Direct link to Botkube Cloud App for Slack")

The Botkube Cloud App for Slack offers several advanced features:

*   One-click installation into your Slack workspace
*   Multi-cluster executor support with a single App for Slack
*   Manage Slack channels directly from the Botkube web app and ensure the Botkube bot is invited to the correct channels

The Botkube Cloud App for Slack uses Botkube's cloud services to manage channels and route executor commands. Events and alerts are sent directly from your cluster to your Slack workspace for reliable, fast notifications.

You can directly try Botkube Cloud App for Slack for free by creating an account in the [Botkube Web App](https://app.botkube.io/). Follow the [Cloud app for Slack tutorial](https://docs.botkube.io/installation/slack/cloud-slack) to learn more.

### Botkube App for Socket Slack[​](https://docs.botkube.io/installation/slack/#botkube-app-for-socket-slack "Direct link to Botkube App for Socket Slack")

The Socket-mode app works with the open-source Botkube Agent. The Botkube App for Socket-mode Slack has the following caveats:

*   Must be installed manually into your Slack workspace using the provided configuration
*   Slack channels must be managed manually, and you need to ensure the Botkube bot is invited to any channel you want to use with Botkube
*   When using executor plugins (e.g. kubectl) in a multi-cluster environment, each cluster needs to have a dedicated Botkube bot for Slack in order to route commands to the correct cluster.

Follow the [instruction](https://docs.botkube.io/installation/slack/socket-slack) for more details.
