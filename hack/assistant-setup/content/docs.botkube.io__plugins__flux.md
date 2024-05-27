Title: Flux | Botkube

URL Source: https://docs.botkube.io/plugins/flux

Markdown Content:
info

**This plugin is hosted by the [Botkube Cloud](https://app.botkube.io/) plugin repository and requires active Botkube Cloud account.**

The Botkube Flux executor plugin allows you to run the [`flux`](https://fluxcd.io/) CLI commands directly within the chat window of your chosen communication platform.

Get started[​](#get-started "Direct link to Get started")
---------------------------------------------------------

### 1\. Prepare elevated RBAC permissions[​](#1-prepare-elevated-rbac-permissions "Direct link to 1. Prepare elevated RBAC permissions")

One of the plugin capabilities is the `flux diff` command. To use it, you need to update the Flux plugin RBAC configuration. This is necessary because the command performs a server-side dry run that requires patch permissions, as specified in the [Kubernetes documentation](https://kubernetes.io/docs/reference/using-api/api-concepts/#dry-run-authorization).

First, create RBAC resources on your cluster:

    cat > /tmp/flux-rbac.yaml << ENDOFFILE---apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRolemetadata:  name: fluxrules:  - apiGroups: ["*"]    resources: ["*"]    verbs: ["get", "watch", "list", "patch"]---apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRoleBindingmetadata:  name: fluxroleRef:  apiGroup: rbac.authorization.k8s.io  kind: ClusterRole  name: fluxsubjects:- kind: Group  name: flux  apiGroup: rbac.authorization.k8s.ioENDOFFILEkubectl apply -f /tmp/flux-rbac.yaml

Next, use the `flux` group in the plugin RBAC configuration:

![Image 1: Flux RBAC](https://docs.botkube.io/assets/images/flux-rbac-bfe6d7c972bcfd611669afd75a3bab20.png)

### 2\. Enable the plugin[​](#2-enable-the-plugin "Direct link to 2. Enable the plugin")

You can enable the plugin as a part of Botkube instance configuration.

1.  If you don't have an existing Botkube instance, create a new one, according to the [Installation](https://docs.botkube.io/) docs.
    
2.  From the [Botkube Cloud homepage](https://app.botkube.io/), click on a card of a given Botkube instance.
    
3.  Navigate to the platform tab which you want to configure.
    
4.  Click **Add plugin** button.
    
5.  Select the Flux plugin.
    
6.  Click **Configure** button and then **Configuration as Code** tab.
    
7.  Configure optional GitHub access token.
    
    The Flux plugin comes with integrated GitHub support. To enable it, you'll need a valid [GitHub token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/#creating-a-token). Set the token with the following configuration:
    
        github:  auth:    accessToken: "" # your GitHub access token
    
8.  Click **Save** button.
    

By default, the Flux plugin has read-only access. To perform actions like creating or deleting Flux-related sources, you'll need to customize the [RBAC](https://docs.botkube.io/features/rbac).

Usage[​](#usage "Direct link to Usage")
---------------------------------------

To execute the `flux` CLI commands, send a message in the channel where Botkube is present. For example:

### Interactive Usage[​](#interactive-usage "Direct link to Interactive Usage")

The Flux plugin supports interactivity (tables, etc.) to simplify running Flux CLI commands e.g. from mobile devices.

![Image 2: flux-interactivity](https://docs.botkube.io/assets/images/flux-interactivity-36eaec2696dd56fe8924ef36f42a7ac1.gif)

### Simplified Kustomization Diffing Flow[​](#simplified-kustomization-diffing-flow "Direct link to Simplified Kustomization Diffing Flow")

With the Botkube Flux executor, you can execute a single command to perform a diff between a specific pull request and the cluster state. For instance:

    @Botkube flux diff kustomization podinfo --path ./kustomize --github-ref [PR Number| URL | Branch]

![Image 3: flux-diff](https://docs.botkube.io/assets/images/flux-diff-abdd97d5a1b5dd3b64ecf2c1712fa14d.gif)

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

To enhance your workflow's efficiency, you can use the [GitHub Events](https://docs.botkube.io/plugins/github-events) source for automatic notification of pull request events, complete with an integrated `flux diff` button.

    github:  auth:    accessToken: "ghp_" # GitHub PATrepositories:  - name: { owner }/{name}    on:      pullRequests:          - types: [ "open" ]            paths:              # Patterns for included file changes in pull requests.              include: [ 'kustomize/.*' ]            notificationTemplate:              extraButtons:                - displayName: "Flux Diff"                  commandTpl: "flux diff ks podinfo --path ./kustomize --github-ref {{ .HTMLURL }} "

Don't forget to bind the plugin to one of the channels.

Configuration[​](#configuration "Direct link to Configuration")
---------------------------------------------------------------

The plugin supports the following configuration:

    github:  auth:    # GitHub access token.    # Instructions for token creation: https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/#creating-a-token.    # Lack of token may limit functionality, e.g., adding comments to pull requests or approving them.    accessToken: ""log:  level: "info"
