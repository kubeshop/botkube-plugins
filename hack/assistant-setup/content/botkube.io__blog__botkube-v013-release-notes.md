Title: Botkube v0.13 Release Notes

URL Source: https://botkube.io/blog/botkube-v013-release-notes

Published Time: Sep 01, 2022

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

The latest release of Botkube is here! This release from the new Botkube team includes user-requested multi-channel support and much more.

### Table of Contents

*   [Multi-Channel Support](#multi-channel-support)
*   [Bug Fixes](#bug-fixes)
*   [Release Pipeline Enhancements](#release-pipeline-enhancements)
*   [Telemetry](#telemetry)
*   [Feedback - We’d Love to Hear From You](#feedback-we-d-love-to-hear-from-you)

#### Start Using Botkube AI Assistant Today!

BotKube, welcome to Kubeshop! We're happy to have the most modern [ChatOps tool for Kubernetes](http://botkube.io/) join the team.

BotKube v0.13.0 is the first release under Kubeshop and the newly formed dedicated Botkube team. We are picking up the pace of development and you can expect to see more frequent BotKube releases with lots of new, great features and bug fixes.

_If you have any feature requests or bug reports, reach out to us on [Slack](http://join.botkube.io/) or submit a [Github issue](https://github.com/kubeshop/botkube/issues)_

In this release, we focused on multi-channel support, bug fixes, and optimizing the BotKube release pipeline.

Multi-Channel Support
---------------------

![Image 2: Connecting Multiple Kubernetes Clusters to Multiple Slack Channels](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6360eb6a2fdfd662d24a9ca0_6310c1dc595cf2e50475bc23_2dGs46sYiazq09JJsR0RtJZwNSQPpS_FFV3EcATxfr7Om0vQQgafkiLDUdFNfVg31adYxk0VRVgyPxjr4nFPLENhF8XSeLPmBAAWVS6dDYtRkgat9EqCF5ApyMZyVhDBJX0cW0y9knrOPcDdSQIsRg.png)

While BotKube is now part of the Kubeshop team, BotKube development is still community-driven. Multi-channel support was requested by several users in the BotKube community and the reason why we implemented it in our first release. We also worked with the community to review the [configuration proposal](https://github.com/kubeshop/botkube/issues/628) and will continue this kind of community involvement as we go.

BotKube now supports multiple channels and communication tools using a single BotKube installation in your Kubernetes cluster. You can group and send events to different channels and tools. For example, some events are sent as Slack chat messages for an immediate response and others to ElasticSearch for archiving. You can even have multiple instances of these tools and send specific events to different channels in each of them. This previously required multiple BotKube installations in a cluster or complex annotations. You can also toggle filters and recommendations per-rule now, where filters and recommendation settings were global in previous versions.

In future releases, we plan to streamline this source to channel mapping further and make the configuration even easier.

It is also possible to configure the executor configuration on a per-channel basis, so that kubectl permissions can be controlled. You can set different kubectl permissions that restrict the commands, resources, and namespaces users in a specified channel can access. This is great for environments with multi-tenant Kubernetes clusters where teams use separate namespaces.

We have simplified and enhanced the BotKube configuration syntax to enable the new multi-channel support. If you are running BotKube v0.12.4 (or earlier) using a default configuration, you can specify an updated configuration when using the helm upgrade command. If you are using a custom configuration, see the [changelog](https://github.com/kubeshop/botkube/blob/main/CHANGELOG.md#v0130-2022-08-29) for details on the configuration syntax changes. You can modify your config.yaml file with the new structure and pass that to the helm upgrade -f command.

You can also see the configuration details for [sources](https://botkube.io/configuration/source/) and [executors](https://botkube.io/configuration/executor/) in the BotKube documentation.

Bug Fixes
---------

We have fixed several bugs in BotKube that were reported to us by users. We also spent some time refactoring code and increasing test coverage to improve the quality of BotKube. You can see the list of bug fixes in the [changelog](https://github.com/kubeshop/botkube/blob/main/CHANGELOG.md#v0130-2022-08-29).

We appreciate your bug reports and pull requests! If you notice a bug in BotKube, please report it as a [GitHub issue](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

Release Pipeline Enhancements
-----------------------------

We are working on improving the BotKube release pipeline with each release. Our goal is to fully automate the release pipeline, so that we can collect all pull requests in a release branch, perform automated end-to-end testing with the various communication platforms, and build and release the BotKube container image.

In this release we have cleaned up the overall CI pipeline and automated end-to-end testing of BotKube on Slack. As we add more automation, BotKube releases will be more frequent and more reliable.

Telemetry
---------

BotKube has introduced anonymized telemetry collection in v0.13.

Why: We are collecting this telemetry data in order to focus our development efforts on the communication tools and actions that BotKube users are using the most.

What:All data is anonymized. We use UUIDs in order to aggregate telemetry from individual BotKube installations. Other data that could identify users or environments is redacted, such as namespace names, pod names, etc. For transparency, we collect the following data:

*   BotKube version
    
*   Kubernetes version
    
*   Names of enabled integrations (notifiers and bots)
    
*   Handled events in anonymized form, grouped by the integration (communication platform) name
    
*   For each event, we collect its type (e.g. create or delete), resource API Version and resource Kind. Any custom resource API groups or Kinds are excluded from the analytics collection.
    
*   Executed commands in anonymized form
    
*   For kubectl commands, only the command verb is collected. Resource name and namespace are excluded from the analytics collection.
    
*   App errors (crashes, configuration and notification errors)
    
*   For identifying BotKube installations, we use unique identifiers generated in the following way:
    
*   As an anonymous cluster identifier, we use the uid of kube-system Namespace,
    
*   As an anonymous installation identifier, we use UUID generated during Helm chart installation and persisted in a ConfigMap.
    

Opt-out: **You can opt out of telemetry collection by setting the `analytics.disable: true` parameter when installing or upgrading the Helm chart.** This prevents all analytics collection and we do not track the installation and opt-out actions.

Please see the BotKube [Privacy Policy](https://botkube.io/privacy/) for more details.

Feedback - We’d Love to Hear From You
-------------------------------------

As always, we want to hear your feedback and ideas about BotKube. Help us plan the BotKube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the BotKube [GitHub issues](https://github.com/kubeshop/botkube/issues), BotKube [Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
