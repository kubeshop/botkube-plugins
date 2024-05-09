Title: Exec | Botkube

URL Source: https://docs.botkube.io/configuration/executor/exec

Markdown Content:
info

**This plugin is available as a part of the Botkube Cloud offering.**

Botkube is introducing new plugins with advanced functionality that will be part of the Botkube Team and Enterprise packages. These advanced plugins require cloud services provided by Botkube and are not part of the Botkube open source software.

As part of this change, some of the existing Botkube plugins are being moved to a new repository. This repository requires authentication with a Botkube account. To continue using these Botkube plugins, create an account at [https://app.botkube.io/](https://app.botkube.io/) and configure a Botkube instance, or [migrate an existing installation with the Botkube CLI](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud).

The `exec` executor plugin allows you to install and run CLI applications directly from chat (e.g., Slack, Discord, or Mattermost) without any hassle.

The `exec` plugin is hosted by the Botkube Cloud plugin repository and requires active Botkube Cloud account.

By default, the read-only `KUBECONFIG` is assigned. For enabling commands that require create, update or delete rules, you need to create specific (Cluster)Role and (Cluster)RoleBinding and reference it from plugin's `context` configuration. To learn more refer to the [RBAC section](https://docs.botkube.io/configuration/rbac).

You can enable the plugin as a part of Botkube instance configuration.

    # An array of templates that define how to convert the command output into an interactive message.templates:  # Link to templates source  # It uses the go-getter library, which supports multiple URL formats (such as HTTP, Git repositories, or S3) and is able to unpack archives.  # For more details, see the documentation at https://github.com/hashicorp/go-getter.  - ref: github.com/kubeshop/botkube//cmd/executor/exec/templates?ref=release-1.8
