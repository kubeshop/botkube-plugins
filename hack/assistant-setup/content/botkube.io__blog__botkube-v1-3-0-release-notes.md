Title: Botkube v1.3.0 Release Notes

URL Source: https://botkube.io/blog/botkube-v1-3-0-release-notes

Published Time: Aug 15, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

Our most recent release of Botkube Cloud & Open Source brings important features that help with Kubernetes troubleshooting. This update introduces several new Botkube plugins.

### Table of Contents

*   [What is Botkube?](#what-is-botkube--2)
*   [Keptn Source Plugin](#keptn-source-plugin-2)
*   [Flux CD Executor Plugin](#flux-cd-executor-plugin-2)
*   [GitHub Events Source Plugin](#github-events-source-plugin-2)
*   [Slack message reactions](#slack-message-reactions-2)
*   [Bug Fixes](#bug-fixes-2)
*   [Feedback - We’d Love to Hear From You!](#feedback-we-d-love-to-hear-from-you--2)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

#### Get started with Botkube Cloud

The Botkube plugin ecosystem is expanding! We are excited to announce the latest release of Botkube, packed with new features and improvements to enhance your Kubernetes collaborative troubleshooting experience. This release introduces:

\* GitHub source plugin

\* Keptn source plugin

\* Flux CD executor plugin

\* Slack message reactions

As usual, all of these plugins can be easily configured in \[Botkube Cloud\](https://app.botkube.io) or via \[manual YAML configuration\](https://docs.botkube.io/configuration/) for the open source Botkube engine.

\[Botkube Actions\](https://docs.botkube.io/configuration/action) also allow you to automate any of the Botkube plugins. That means that any source plugin can trigger a command in any executor plugin. Botkube is the automation glue that can tie together all of the tools you use both inside and outside of your Kubernetes cluster. We have more details coming up in future blog posts on how to use Botkube to manage and monitor your GitOps workflows with diverse tools.

\## What is Botkube?

Botkube is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred messaging and collaboration platforms like Slack, Microsoft Teams, Discord, and Mattermost. This means you'll gain instant visibility and control over your cluster resources, all without ever having to leave your messaging platform. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more. In this blog post, we'll give you an overview of all of Botkube’s newest features from the v1.2.0 release.

\*If you have any feature requests or bug reports, reach out to us on \[Slack\](http://join.botkube.io) or submit a \[Github issue\](https://github.com/kubeshop/botkube/issues)\*

\## Keptn Source Plugin

The Keptn source plugin is the latest edition to Botkube’s source plugin CNCF tool suite. By incorporating Keptn, you can effortlessly orchestrate load tests and validate the results. To simplify the understanding of Keptn's synergy with Botkube, we'll take a hands-on approach. We'll deploy a straightforward task using Keptn, and then seamlessly capture Keptn events through Botkube. The result? Notifications delivered straight to your chosen communication platform, such as Slack. This integration not only enhances your insight into the deployment process but also streamlines your workflow.With this exciting update, we're merging the capabilities of Keptn and Botkube to streamline your workflow, enhance collaboration, and ensure that your deployments are more efficient than ever before.

Find out how to use the Keptn source plugin in the \[documentation\](https://docs.botkube.io/configuration/source/keptn).

You can read all about how we developed the Keptn source plugin (and how to create your own source plugins) in this \[Botkube blog post\](https://botkube.io/blog/implementing-your-own-botkube-plugin-a-real-life-use-case).

\## Flux CD Executor Plugin

Introducing the New Flux Executor Plugin  
The Flux executor plugin marks a new approach to Botkube executor plugins. It doesn't just offer the convenience of running Flux CLI commands directly from communications platforms like Slack or Discord; it goes further by incorporating interactivity, such as tables and more.

But the true game-changer is the integration of a practical, real-world use-case right from the start. With the Botkube Flux executor, executing a single command can initiate a thorough comparison between a specific pull request and the current state of your cluster. For example:  
Sample code  
This command streamlines tasks like:  
Automatically identifying the related GitHub repository linked to the provided kustomization.  
Effortlessly cloning the repository.  
Checking out the specified pull request.  
Seamlessly contrasting the changes in the pull request against the cluster's present state.  
Gone are the days of laborious manual processes.The Botkube Flux executor helps to eliminate these complexities out of the GitOps workflow.

Find out how to use the Flux CD executor plugin in the \[documentation\](https://docs.botkube.io/configuration/executor/flux).

\## GitHub Events Source Plugin

The GitHub Events source plugin is the first Botkube plugin that we have developed that works outside of the Kubernetes cluster. This plugin allows you to stream events from the GitHub repositories you select via Botkube to your collaboration platform. Botkube Actions make this plugin especially exciting. You can use the GitHub plugin to watch for pull requests that make application or configuration changes to your environment and trigger other tools, like GitOps reviews or security scans.

Find out how to use the GitHub Events source plugin in the \[documentation\](https://docs.botkube.io/configuration/source/github-events).

\## Slack message reactions

Botkube has introduced a lot of new executor plugins to help you troubleshoot and manage your Kubernetes clusters. Some commands take some time to process in the background and until now, you couldn't see if the command had completed, you just had to wait for a response. We have introduced emoji reactions to executor requests that show the state of the command. If you see 👀, the command is still running. Once the command is complete, the reaction automatically changes to ✅.

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64db6cd0c67e3e749437884e_Screenshot%202023-08-15%20at%2014.17.12.png)

Note that if you are using Botkube already, you need to add the \`reactions:write\` scope to the Bot Token. If you are using the Botkube \[Cloud Slack\](https://docs.botkube.io/installation/slack/cloud-slack) app, you can update directly from the Botkube Cloud web app. If you are using the Botkube \[Socket Slack\](https://docs.botkube.io/installation/slack/socket-slack/self-hosted) app, you need to go to the \[Slack App console\](https://api.slack.com/apps/), select your Botkube app, navigate to OAuth & Permissions, and add the \`reactions:write\` scope to the Bot Token configuration. You then need to reinstall the Slack app from the console.

\## Bug Fixes

We have also fixed a couple of small bugs in Botkube that were reported to us by users and found by us. You can see the full list of changes in the \[changelog\](https://github.com/kubeshop/botkube/releases/tag/v1.3.0).

We appreciate your bug reports and pull requests! If you notice a bug in Botkube, please report it as a \[GitHub issue\](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

\## Feedback - We’d Love to Hear From You!

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the Botkube \[GitHub issues\](https://github.com/kubeshop/botkube/issues), Botkube \[Slack community\](http://join.botkube.io/), or email our Product Leader at \[blair@kubeshop.io\](mailto:blair@kubeshop.io).
