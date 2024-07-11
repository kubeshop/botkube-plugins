Title: Botkube v0.14 Release Notes

URL Source: https://botkube.io/blog/botkube-v014-release-notes

Published Time: Oct 06, 2022

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

BotKube v0.14.0 addresses a number of community-suggested issues and adds some great features to help you control the notification levels you receive from your Kubernetes clusters.

### Table of Contents

*   [Dynamic notification setting changes](https://botkube.io/blog/botkube-v014-release-notes#dynamic-notification-setting-changes)
*   [New Botkube Slack App](https://botkube.io/blog/botkube-v014-release-notes#new-botkube-slack-app)
*   [Improved Help](https://botkube.io/blog/botkube-v014-release-notes#improved-help)
*   [Persistent Settings](https://botkube.io/blog/botkube-v014-release-notes#persistent-settings)
*   [Mattermost v7 Support](https://botkube.io/blog/botkube-v014-release-notes#mattermost-v7-support)
*   [Bug Fixes](https://botkube.io/blog/botkube-v014-release-notes#bug-fixes)
*   [Feedback - We’d Love to Hear From You](https://botkube.io/blog/botkube-v014-release-notes#feedback-we-d-love-to-hear-from-you)

#### Start Using Botkube AI-Powered Assistant Today

In the previous release of Botkube we promised to accelerate the pace of development and release. Well, here is another new release of Botkube, the most modern [ChatOps tool for Kubernetes](http://botkube.io/)!

BotKube v0.14.0 addresses a number of community-suggested issues and adds some great features to help you control the notification levels you receive from your Kubernetes clusters.

_If you have any feature requests or bug reports, reach out to us on [Slack](http://join.botkube.io/) or submit a [Github issue](https://github.com/kubeshop/botkube/issues)_

In this release, we focused on improvements to the Botkube onboarding experience and notification tuning, persistent settings, and bug fixes. Some of the enhancements are currently Slack-specific using the additional interactive capabilities of the new Slack app. We plan to enhance other platforms with interactive features where technically feasible.

Dynamic notification setting changes
------------------------------------

We've received frequent feedback from users about the notification settings in Botkube. By default, Botkube monitors and shows all Kubernetes events in a configured communication channel. Previously, this could only be configured in the YAML configuration or during the Helm installation/upgrade.

You can now also configure notification settings right from the Botkube interface in your communication platform. Using the command @Botkube edit SourceBindings you can select the notifications you want to receive. After choosing the notifications, the configuration is dynamically updated with an automatic brief restart of Botkube.

![Image 2: Monitoring Kubernetes in Slack Channel](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6360eb5b0c5188f8691d6b94_633ef84e656938a0ae571fc0_iRbu4PGydNxlrRVX0s4UpTz9QFyoUzE-l5Onxt_TW2YVakCXmxLkWfqXpgtJ7ecj136UrCZdgJGubxY_hdd2IEsiq8tyanu5ITCqu-fgUB0mP_tPZLGYtghrVHYx5uU3bhAi0H-lW6PKNIE24trMca9NWzPrlqHj-CYGLZqoA2CiKsNrr3i2TeFGfQ.png)

New Botkube Slack App
---------------------

We have created a brand new Botkube Slack app. The new app uses the Slack Socket Mode which provides a much more robust platform for interactivity and uses the new Slack permissions model for enhanced security.

You will see many enhancements to Botkube in Slack in this and upcoming releases. For example, you can use the interactive selector to change the notification settings mentioned earlier. Just run @Botkube, edit SourceBindings without specifying the notification level, and you can select the sources from an interactive dialog.

![Image 3: Adding Kubernetes Alerts to Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6360eb5b4e8002657684fc0e_633ef84e908df6431e4dd34d_SelrLZepl_wTwzOgctqA0qtliy6LH3SnfV8l8I2r6iMivCxRNdzhtdo9T5TvjNw9boEGHGD1jqnOcld_4B5MxTLxO01xwGq41cZ_SKhyFJLacFJFK7HMXOJ7lrP93TrM9M6CmhTpauoLdoG1D7bWLXYK-rryjw0SCVi5c-xXTh_eKe9JAB73QHI0Dg.png)

The installation process has changed for the new Slack app. See the [installation guide](https://botkube.io/docs/installation/socketslack) to get started with the enhanced interactive experience.

The [legacy Slack app](https://botkube.io/docs/installation/slack/) is still available, but will be deprecated in future versions.

Improved Help
-------------

The @Botkube help command has been enhanced when using the new Botkube Slack app. Run @Botkube help and you will see an interactive help message. You can click any of the buttons in the interactive help to run the commands, change settings, and more.

This new, interactive help will enable new Botkube users to get up and running even faster. You can also send us feedback and access documentation and additional external help resources directly from the interactive help.

![Image 4: Botkube help suggestions for cluster management ](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6360eb5b82ec26bf25ce56ed_633ef850e5ef9f0d7d116776_EO0U7RPVKtnFJKGI9P-P-ra7xxTD04eyiGXF_khHkvFL8Age_YGzOa7YM3iTZZ6g9OfTmH_HgJYyJNEN75LDiP5jij3hsQjAF9z0nWw2xJ42aEJlskc_lNSh6XIiwwybgIl0TGxlJRgOSKkUmGbJv9d5NLGOW-x_iVWUj5GyPevaKTWieW5B1QmpCw.png)

Persistent Settings
-------------------

Changes to settings made from the Botkube interface, such as notifier start/stop, notification level settings, and filter enable/disable now persist across Botkube restarts. In earlier versions, these settings would reset to the settings in the Botkube configuration file upon restarting the Botkube pod.

Mattermost v7 Support
---------------------

Botkube now supports Mattermost version 7.x! See the [installation guide for Mattermost](https://botkube.io/docs/installation/mattermost) to get started.

Bug Fixes
---------

We have fixed several bugs in BotKube that were reported to us by users. We also spent some time refactoring code and increasing test coverage to improve the quality of BotKube. You can see the list of bug fixes in the [changelog](https://github.com/kubeshop/botkube/releases/tag/v0.14.0).

We appreciate your bug reports and pull requests! If you notice a bug in BotKube, please report it as a [GitHub issue](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

Feedback - We’d Love to Hear From You
-------------------------------------

As always, we want to hear your feedback and ideas about BotKube. Help us plan the BotKube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the BotKube [GitHub issues](https://github.com/kubeshop/botkube/issues), BotKube [Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
