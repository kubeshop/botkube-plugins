Title: Migrating installation to Botkube Cloud | Botkube

URL Source: https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud

Markdown Content:
Migrating installation to Botkube Cloud[​](#migrating-installation-to-botkube-cloud "Direct link to Migrating installation to Botkube Cloud")
---------------------------------------------------------------------------------------------------------------------------------------------

If you have started using Botkube with the Open Source installation, you have the option to migrate this instance to be managed using [Botkube Cloud](https://app.botkube.io/).

To make the migration process easier, we provide a dedicated `botkube cloud migrate` command that seamlessly transfers your Botkube installation to Botkube Cloud.

Supported Botkube platforms:

*   Socket Slack
*   Discord
*   Mattermost

Steps[​](#steps "Direct link to Steps")
---------------------------------------

1.  [Install Botkube CLI](https://docs.botkube.io/cli/getting-started#installation)
    
2.  [Login into Botkube Cloud](https://docs.botkube.io/cli/getting-started#first-use)
    
3.  Run Botkube migrate:
    

Limitations[​](#limitations "Direct link to Limitations")
---------------------------------------------------------

The following list contains current limitations that we will address in the near future:

*   `extraObjects` in Botkube [helm configurations](https://github.com/kubeshop/botkube/blob/593746a70d9eb23469c28e5c0274c9a40a7b651d/helm/botkube/values.yaml#L1040) are ignored. If you have any extra resources under `extraObjects` section, you need to migrate them on your own.
*   Custom `rbac.groups` are ignored.
*   All 3rd-party plugins are ignored.
*   Minimal supported Botkube version is v1.0.0.

See more[​](#see-more "Direct link to See more")
------------------------------------------------

To learn more about `botkube cloud migrate` and all supported settings, visit the [Botkube migrate](https://docs.botkube.io/cli/commands/botkube_migrate) document.
