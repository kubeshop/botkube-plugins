Title: Flux | Botkube

URL Source: https://docs.botkube.io/usage/executor/flux

Markdown Content:
Version: 1.10

info

**This plugin is available as a part of the Botkube Cloud offering.**

Botkube is introducing new plugins with advanced functionality that will be part of the Botkube Team and Enterprise packages. These advanced plugins require cloud services provided by Botkube and are not part of the Botkube open source software.

As part of this change, some of the existing Botkube plugins are being moved to a new repository. This repository requires authentication with a Botkube account. To continue using these Botkube plugins, create an account at [https://app.botkube.io/](https://app.botkube.io/) and configure a Botkube instance, or [migrate an existing installation with the Botkube CLI](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud).

Botkube offers seamless execution of Flux CLI commands within your Kubernetes cluster. By default, Flux command execution is disabled. To enable it, refer to the [**Enabling plugin**](https://docs.botkube.io/configuration/executor/flux#enabling-plugin) section.

To execute the `flux` CLI commands, send a message in the channel where Botkube is present. For example:

Interactive Usage[​](#interactive-usage "Direct link to Interactive Usage")
---------------------------------------------------------------------------

We have also incorporated interactivity (tables, etc.) to simplify running Flux CLI commands e.g. from mobile devices.

![Image 1: flux-interactivity](https://docs.botkube.io/assets/images/flux-interactivity-36eaec2696dd56fe8924ef36f42a7ac1.gif)

Simplified Kustomization Diffing Flow[​](#simplified-kustomization-diffing-flow "Direct link to Simplified Kustomization Diffing Flow")
---------------------------------------------------------------------------------------------------------------------------------------

With the Botkube Flux executor, you can execute a single command to perform a diff between a specific pull request and the cluster state. For instance:

    @Botkube flux diff kustomization podinfo --path ./kustomize --github-ref [PR Number| URL | Branch]

![Image 2: flux-diff](https://docs.botkube.io/assets/images/flux-diff-abdd97d5a1b5dd3b64ecf2c1712fa14d.gif)

This command automates several tasks:

*   Automatically discovering the associated GitHub repository for the given kustomization.
*   Cloning the repository.
*   Checking out a given pull request.
*   Comparing pull request changes with the current cluster state.
*   Sharing the diff report.

The diff results are posted on the Slack channel, making it easy for team members to review and discuss the changes. Additionally, the returned message provides additional contextual actions:

*   Posting the diff report as a GitHub comment on the corresponding pull request.
*   Approving the pull request.
*   Viewing the pull request.

### GitHub automation[​](#github-automation "Direct link to GitHub automation")

To enhance your workflow's efficiency, you can use the [GitHub Events](https://docs.botkube.io/configuration/source/github-events) source for automatic notification of pull request events, complete with an integrated `flux diff` button.

    github:  auth:    accessToken: "ghp_" # GitHub PATrepositories:  - name: { owner }/{name}    on:      pullRequests:          - types: [ "open" ]            paths:              # Patterns for included file changes in pull requests.              include: [ 'kustomize/.*' ]            notificationTemplate:              extraButtons:                - displayName: "Flux Diff"                  commandTpl: "flux diff ks podinfo --path ./kustomize --github-ref {{ .HTMLURL }} "

Don't forget to bind the plugin to one of the channels.
