Title: Botkube v1.4.0 Release Notes - Now with Native Argo CD Support

URL Source: https://botkube.io/blog/botkube-v1-4-0-release-notes

Published Time: Sep 19, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

Our most recent release of Botkube brings important features that help with Kubernetes troubleshooting. This update introduces Botkube support for Argo CD sources and more.

### Table of Contents

*   [What is Botkube?](#what-is-botkube-)
*   [Argo CD Source Plugin](#argo-cd-source-plugin)
*   [Mattermost Bot Account Support](#mattermost-bot-account-support)
*   [Custom OpenAI API Parameters for the Doctor Plugin](#custom-openai-api-parameters-for-the-doctor-plugin)
*   [Elasticsearch V8 Support](#elasticsearch-v8-support)
*   [Restart Crashed Plugins](#restart-crashed-plugins)
*   [Bug Fixes](#bug-fixes)
*   [Feedback - We’d Love to Hear From You](#feedback-we-d-love-to-hear-from-you)

#### Start Using Botkube AI Assistant Today!

#### Get started with Botkube Cloud

Botkube continues to enhance your GitOps experience by adding support for Argo CD via a new source plugin! We are excited to announce the latest release of Botkube, packed with new features and improvements to enhance your Kubernetes collaborative troubleshooting experience. This release introduces:

*   [Argo CD source plugin](https://docs.botkube.io/configuration/source/argocd)
    
*   [Mattermost Bot Account support](https://docs.botkube.io/installation/mattermost/self-hosted)
    
*   [Custom OpenAI API parameters for the Doctor plugin](https://docs.botkube.io/configuration/executor/doctor)
    
*   [Elasticsearch V8 support](https://docs.botkube.io/installation/elasticsearch/self-hosted))
    
*   Restart crashed plugins
    

As usual, plugins can be easily configured in [Botkube Cloud](https://app.botkube.io/) or via [manual YAML configuration](https://docs.botkube.io/configuration/) for the open source Botkube engine.

[Botkube Actions](https://docs.botkube.io/configuration/action) also allow you to automate any of the Botkube plugins. That means that any source plugin can trigger a command in any executor plugin. Botkube is the automation glue that can tie together all of the tools you use both inside and outside of your Kubernetes cluster. We have more details coming up in future blog posts on how to use Botkube to manage and monitor your GitOps workflows with diverse tools.

What is Botkube?
----------------

[Botkube](http://app.botkube.io/) is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred messaging and collaboration platforms like Slack, Microsoft Teams, Discord, and Mattermost. This means you'll gain instant visibility and control over your cluster resources, all without ever having to leave your messaging platform. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more.

_If you have any feature requests or bug reports, reach out to us on [Slack](http://join.botkube.io/) or submit a [Github issue](https://github.com/kubeshop/botkube/issues)_

Argo CD Source Plugin
---------------------

The [Argo CD source plugin](https://docs.botkube.io/configuration/source/argocd) is another addition to Botkube's [support for GitOps processes and lifecycles](https://botkube.io/blog/enhancing-gitops-workflows-with-botkube). The Argo CD source plugin allows users to receive and interact with crucial real time notifications when deploying applications in production -- how the deployment is progressing, a health check on the deployment, and more — all directly into communication platforms like Slack, Mattermost, MS Teams, and Discord. While Argo CD does provide a native notification system in some of these platforms, the set-up process can be cumbersome and difficult with many manual steps to get up and running. Botkube’s plug-and-play-style native Argo CD plugin can be configured in minutes directly in [Botkube’s web interface](http://app.botkube.io/).

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6509a59c63441b36226ea80d_argocd-events-e6eabb1f581e9822020d55461539bfcd.png)

You can get started with the Argo CD with the [Documentation](https://docs.botkube.io/configuration/source/argocd).

Mattermost Bot Account Support
------------------------------

Botkube now supports running as a [Mattermost bot account](https://developers.mattermost.com/integrate/reference/bot-accounts/). Mattermost bot accounts have more limited privileges, do not consume a user license, and are not associated with a specific user so there's no concern about Botkube continuing to run if someone leaves your organization.

Get started with Mattermost and Botkube with the [documentation](https://docs.botkube.io/installation/mattermost/self-hosted).

Custom OpenAI API Parameters for the Doctor Plugin
--------------------------------------------------

The Botkube Doctor plugin now supports configurable alternate base URLs for the OpenAI API. This means you can use other OpenAI-compatible APIs with the doctor plugin. You can use Azure OpenAI Service, 3rd party Llama models, OpenAI proxies, and more.

Find out how to configure custom parameters for Doctor in the [documentation](https://docs.botkube.io/configuration/executor/doctor).

Elasticsearch V8 Support
------------------------

Botkube now supports Elasticsearch V8 as a platform. If you're using the latest and greatest version of Elasticsearch, Botkube has you covered and will happily send your events over.

Get set up with Elasticsearch with the [documentation](https://docs.botkube.io/installation/elasticsearch/self-hosted).

Restart Crashed Plugins
-----------------------

We have added advanced logic to the Botkube plugin system in order to detect plugins that might have crashed and automatically restart them. This makes Botkube more reliable as a core part of your observability and troubleshooting platform. Plugin crashes are logged to the Botkube pod log and will only attempt to restart a fixed number of times in order to avoid continuous restart loops.

Bug Fixes
---------

We have also fixed a couple of small bugs in Botkube that were reported to us by users and found by us. You can see the full list of changes in the [changelog](https://github.com/kubeshop/botkube/releases/tag/v1.4.0).

We appreciate your bug reports and pull requests! If you notice a bug in Botkube, please report it as a [GitHub issue](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

Feedback - We’d Love to Hear From You
-------------------------------------

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the Botkube [GitHub issues](https://github.com/kubeshop/botkube/issues), Botkube [Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
