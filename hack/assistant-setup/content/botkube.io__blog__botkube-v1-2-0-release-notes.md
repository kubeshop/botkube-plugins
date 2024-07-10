Title: Botkube v1.2.0 Release Notes

URL Source: https://botkube.io/blog/botkube-v1-2-0-release-notes

Published Time: Jul 18, 2023

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

This release brings easy Doctor plugin configuration with an OpenAI key; CLI for easy installs and cloud migration; and Exec plugin to install + run any CLI directly from your communication platform.

### Table of Contents

*   [What is Botkube?](https://botkube.io/blog/botkube-v1-2-0-release-notes#what-is-botkube-)
*   [Doctor Plugin](https://botkube.io/blog/botkube-v1-2-0-release-notes#doctor-plugin)
*   [Botkube CLI for Installation](https://botkube.io/blog/botkube-v1-2-0-release-notes#botkube-cli-for-installation)
*   [Botkube CLI for Easy Cloud Migration](https://botkube.io/blog/botkube-v1-2-0-release-notes#botkube-cli-for-easy-cloud-migration)
*   [Exec Plugin](https://botkube.io/blog/botkube-v1-2-0-release-notes#exec-plugin)
*   [Bug Fixes](https://botkube.io/blog/botkube-v1-2-0-release-notes#bug-fixes)
*   [Feedback - We’d Love to Hear From You](https://botkube.io/blog/botkube-v1-2-0-release-notes#feedback-we-d-love-to-hear-from-you)

#### Start Using Botkube AI-Powered Assistant Today

Botkube is introducing the power of AI to Kubernetes troubleshooting! We are excited to announce the latest release of Botkube, packed with new features and improvements to enhance your Kubernetes collaborative troubleshooting experience. This release introduces

*   Doctor, a ChatGPT plugin for Botkube
    
*   Botkube CLI for easier installation and simple migration to Botkube Cloud
    
*   Exec, a Botkube plugin that allows you to run any CLI command from your collaboration platform
    

What is Botkube?
----------------

Botkube is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred messaging and collaboration platforms like Slack, Microsoft Teams, Discord, and Mattermost. This means you'll gain instant visibility and control over your cluster resources, all without ever having to leave your messaging platform. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more. In this blog post, we'll give you an overview of all of Botkube’s newest features from the v1.2.0 release.

_If you have any feature requests or bug reports, reach out to us on [Slack](http://join.botkube.io/) or submit a [Github issue](https://github.com/kubeshop/botkube/issues)_

Doctor Plugin
-------------

AI has been making incredible progress recently. Large Language Models power tools like ChatGPT to answer almost any question using the collective knowledge stored over years on the Internet. New applications for these tools are being recognized every day and now you can use the power of ChatGPT to troubleshoot your Kubernetes clusters with Botkube.

The Doctor plugin only requires an OpenAI API key to configure. The easiest way to configure it is through the [Botkube web app](https://app.botkube.io/) if your cluster is connected. Otherwise you can enable it in your [Botkube YAML configuration](https://docs.botkube.io/configuration/executor/doctor). Once enabled, you can ask questions about specific resources or ask free-form questions, directly from any enabled channel.

‍

![Image 2](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64b6a35168a1227527eb6501_Screenshot%202023-07-18%20at%2016.13.06.png)

Find out how to use the Doctor plugin in the [documentation](https://docs.botkube.io/usage/executor/doctor).

Botkube CLI for Installation
----------------------------

We've introduced the Botkube CLI in order to streamline the installation process. Previous versions of Botkube were installed using Helm which was quick and easy but if misconfigurations or other errors occurred, they weren't so easy to find. Helm is great for unattended installations, but not very interactive.

The new Botkube CLI still uses Helm in the background. The CLI ensures all of the necessary resources are created and streams the installation log to you in real time. This lets you respond to any errors or misconfigurations much quicker.

‍

![Image 3](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64b6a3374b2b9c3d5500b9c2_Screenshot%202023-07-18%20at%2014.51.10.png)

Botkube can still be installed using the Helm chart alone if you want to perform automated installations with CI or GitOps tools.

It's easy to get started with the Botkube CLI, check out the [documentation](https://docs.botkube.io/cli/getting-started).

Botkube CLI for Easy Cloud Migration
------------------------------------

[Botkube Cloud](https://app.botkube.io/) was released with Botkube v1.0 and enhanced with advanced Slack support in [v1.1.0](https://botkube.io/blog/botkube-v1-1-0-release-notes). Botkube Cloud provides quicker configuration and installation, multi-cluster configuration management, aggregated audit and event logs for all of your clusters, and advanced Slack multi-cluster executor routing.

If you are using Botkube in your cluster with local configuration and you didn't install via Botkube Cloud, you now have a simple on-ramp to the advanced Botkube Cloud services. Here's how to migrate your Botkube installation:

1.  Create a [Botkube Cloud](https://app.botkube.io/) account
    
2.  Install the [Botkube CLI](https://docs.botkube.io/cli/getting-started)
    
3.  Run botkube login to [authenticate](https://docs.botkube.io/cli/getting-started#first-use) your CLI
    
4.  Run botkube migrate to [move your configuration](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud) to the cloud and connect your Botkube installation automatically
    

Botkube Cloud is free to use for a single cluster up to 20 nodes. For multi-cluster management, larger clusters (>20 nodes), and advanced Slack functionality, a Botkube Cloud subscription is required.

Exec Plugin
-----------

The Exec plugin is the result of an experiment to find out how flexible the Botkube plugin system really can be. Exec allows you to install and run any CLI directly from your communication platform with Botkube. It also lets you create your own interactive Slack interface for those CLIs.

The Exec plugin is not meant as a replacement for dedicated executor plugins, but you can use it to get acquainted with the plugin system, to quickly test out a new CLI in Botkube, or in a lab environment.

Check out the [installation](https://docs.botkube.io/configuration/executor/exec) and [usage](https://docs.botkube.io/usage/executor/exec) documentation for details on how to get started.

Bug Fixes
---------

We have also fixed a couple of small bugs in Botkube that were reported to us by users and found by us. You can see the full list of changes in the [changelog](https://github.com/kubeshop/botkube/releases/tag/v1.2.0).

We appreciate your bug reports and pull requests! If you notice a bug in Botkube, please report it as a [GitHub issue](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

Feedback - We’d Love to Hear From You
-------------------------------------

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the Botkube [GitHub issues](https://github.com/kubeshop/botkube/issues), Botkube [Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
