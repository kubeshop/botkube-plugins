Title: Automated actions | Botkube

URL Source: https://docs.botkube.io/usage/automated-actions

Markdown Content:
Actions allow you to automate your workflows by executing custom commands based on specific events. To read how to configure actions, see the [Action](https://docs.botkube.io/configuration/action) configuration document.

Manage actions[​](#manage-actions "Direct link to Manage actions")
------------------------------------------------------------------

Botkube allows you to manage actions using `@Botkube` commands.

### List available actions[​](#list-available-actions "Direct link to List available actions")

Run `@Botkube list actions` to get list of configured actions and their running status:

![Image 1: List available actions](https://docs.botkube.io/assets/images/list-actions-e1d1d86e622d7a10077d5347958a3559.png)

### Disable action[​](#disable-action "Direct link to Disable action")

Run `@Botkube disable action {action-name}` to disable an action named `{action-name}`. The action settings are persisted across Botkube app restarts.

![Image 2: Disable action](https://docs.botkube.io/assets/images/disable-action-414dd23e8a7bcb9efc1d52251f68999c.png)

When you disable an action, changes are applied once the Botkube is restarted. It is an automated process which usually takes a few seconds.

### Enable action[​](#enable-action "Direct link to Enable action")

Run `@Botkube enable action {action-name}` to enable an action named `{action-name}`. The action settings are persisted across Botkube app restarts.

![Image 3: Enable action](https://docs.botkube.io/assets/images/enable-action-08c9232d0d21939ec91201abdcb70a50.png)

When you enable an action, changes are applied once the Botkube is restarted. It is an automated process which usually takes a few seconds.
