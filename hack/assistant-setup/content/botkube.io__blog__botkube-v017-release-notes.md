Title: Botkube v0.17.0 Release Notes

URL Source: https://botkube.io/blog/botkube-v017-release-notes

Published Time: Jan 09, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

Botkube v0.17.0 adds a plugin system and the first executor plugin allowing you to run Helm commands right from your chat tools.

### Table of Contents

*   [Plugin System](#plugin-system-2)
*   [Helm Plugin](#helm-plugin-2)
*   [Kubectl Prefix Required](#kubectl-prefix-required-2)
*   [Botkube Command Tidying](#botkube-command-tidying-2)
*   [Bug Fixes](#bug-fixes-2)
*   [Feedback - We’d Love to Hear From You!](#feedback-we-d-love-to-hear-from-you--2)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

Botkube v0.17.0 is here, and it's huge! We've introduced a plugin system for sources and executors along with the first plugin for Helm. Botkube is the most modern \[ChatOps tool for Kubernetes\](http://botkube.io)!

This release is a major upgrade for Botkube. Kubernetes users are embracing all kinds of new tools and technologies for their application delivery workflows. GitOps and other DevOps tooling is changing the way Kubernetes is managed and monitored and automating a lot of the traditional processes. We wanted to go beyond the Kubernetes basics of monitoring the Kubernetes event stream and running kubectl commands. We decided to do that by implementing a plugin system that allows anyone to build a plugin for Botkube to add sources and executors.

With the addition of plugins in v0.17.0 combined with \[actions\](https://docs.botkube.io/next/configuration/action), which were added in \[v0.16.0\](https://botkube.io/blog/botkube-v016-release-notes), Botkube gains the ability to tie almost any tools together in the Kubernetes/CNCF ecosystem. Botkube is growing into an API automation engine that can listen for events from any source and respond with automated actions via any executor, all while reporting in to channels in your communication platform.

\*If you have any feature requests or bug reports, reach out to us on \[Slack\](http://join.botkube.io) or submit a \[Github issue\](https://github.com/kubeshop/botkube/issues)\*

\## Plugin System

We have introduced a plugin system to Botkube in v0.17.0. This allows anyone to write new plugins for sources and executors and those plugins can be published for other Botkube users to install and use. We'll be leading the way with our own plugins, including the first executor plugin included in this release.

Botkube has two ways of interacting with APIs and other tools, sources and executors. Sources are event sources, so any tool that sends events or notifications can be integrated into Botkube and those events and notifications can be sent to communication channels based on the configured source bindings. Kubernetes native events are an example of a source and have been included in Botkube since the beginning. Executors are used to run commands from other tools from within a chat channel. The output of the command is then returned back to the channel where the command was run and can be filtered if needed. The original executor in Botkube is kubectl, and we have added a helm plugin with this release.

We will be publishing documentation and tutorials on the Botkube plugin system soon so you can design and create your own plugins. We are also working on additional plugins including the first new source plugin for Prometheus.

\## Helm Plugin

Along with the plugin system, we have introduced Helm as the first executor plugin for Botkube. The Helm plugin allows you to run Helm commands right from your communication platform. The Botkube Helm plugin provides a subset of the most useful Helm commands, you can list installed releases, rollback upgrades, and even install charts and uninstall releases.

![Image 2: Helm list in Slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63bc34953fa404bdf9e9b66c_lJAFGnZKp4HY98JtcnteqhObKFnrZ1RoxmMHj8jPZ3XcCk3TWjekXD_hQJNrucD7JANNaYNyHh4E5uIHDl1SS3RWRisuAd8boq7fXK388pca9Tae_CN2k0ZdMy1_QTC9ZGVNmLPwECEzvNKSRai2p3d4CfAMoQjNrV0VJzsfnjkjipqdmCkwT0FkrOr3EA.png)

Imagine seeing an alert in your communications platform that a deployment was updated but something went wrong and the new release is failing. You can check the application pod logs, the helm release status, then run \`@botkube helm rollback <release>\` and fix the problem without ever leaving the chat!

See the \[documentation\](https://docs.botkube.io/next/configuration/executor/helm) for how to enable and use the Helm plugin. Much like the kubectl executor, Helm is disabled by default and when enabled, only allows read-only operations by default. You can enable write actions if you want, and you can even segregate those to specific, more restricted channels in your communication platform.

\## Kubectl Prefix Required

The kubectl executor was the only executor available in older versions of Botkube. With the addition of the Botkube plugin system and additional executor plugins, kubectl has company. You now need to specify the executor when running commands with Botkube. For kubectl this would be something like \`@botkube kubectl get pods\`, and we've even added the aliases \`@botkube k\` and \`@botkube kc\` to save you some typing. For Helm it would look like \`@botkube helm list\`. We added a notification about this to the help message in Botkube v0.16.0, but as of v0.17.0 the implied kubectl (i.e. \`@botkube get pods\`) will no longer work.

\## Botkube Command Tidying

There were some inconsistencies in how Botkube commands were structured in previous versions. We put some effort into unifying the command structure. Now all Botkube commands follow the same \`@botkube <verb> <object>\` structure. For example: \`@botkube notifier start\` is now changed to \`@botkube enable notifications\`. You can see a list of all of the changed commands in the \[changelog\](https://github.com/kubeshop/botkube/releases/tag/v0.17.0).

\## Bug Fixes

We have also fixed a couple of small bugs in Botkube that were reported to us by users and found by us. You can see the full list of changes in the \[changelog\](https://github.com/kubeshop/botkube/releases/tag/v0.17.0). 

We appreciate your bug reports and pull requests! If you notice a bug in Botkube, please report it as a \[GitHub issue\](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

\## Feedback - We’d Love to Hear From You!

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly. 

You can talk to us in the Botkube \[GitHub issues\](https://github.com/kubeshop/botkube/issues), Botkube \[Slack community\](http://join.botkube.io/), or email our Product Leader at \[blair@kubeshop.io\](mailto:blair@kubeshop.io).
