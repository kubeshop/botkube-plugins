Title: Botkube v0.18.0 Release Notes

URL Source: https://botkube.io/blog/botkube-v0-18-0-release-notes

Published Time: Feb 06, 2023

Markdown Content:
![Image 1: a man with a beard wearing a black shirt](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

Botkube v0.18.0 adds a Prometheus source plugin and command aliases.

### Table of Contents

*   [Prometheus Plugin](#prometheus-plugin)
*   [Aliases](#aliases)
*   [Bug Fixes](#bug-fixes)
*   [Feedback - We’d Love to Hear From You!](#feedback-we-d-love-to-hear-from-you-)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

Botkube v0.18.0 has arrived. We've added our first new source plugin for Prometheus, and customizable command aliases. Botkube is the most modern [ChatOps tool for Kubernetes](http://botkube.io/)!

This release builds on the [plugin system](https://botkube.io/blog/build-a-github-issues-reporter-for-failing-kubernetes-apps-with-botkube-plugins) we released in [Botkube v0.17.0](https://botkube.io/blog/botkube-v017-release-notes). In that version we also released the first Botkube-developed executor [plugin for Helm](https://docs.botkube.io/configuration/executor/helm). We have made some other small enhancements to the plugin system in this version.

_If you have any feature requests or bug reports, reach out to us on [Slack](http://join.botkube.io/) or submit a [Github issue](https://github.com/kubeshop/botkube/issues)_

Prometheus Plugin
-----------------

While Botkube has supported Kubernetes events for a long time, K8s clusters tend to have a lot more state events and metrics going on. Typically tools like kube-state-metrics and custom applications are configured to use Prometheus as a store for metrics. Prometheus' Alertmanager can then find metrics that are in anomalous states and trigger alerts. We've built the Botkube source plugin for [Prometheus](https://docs.botkube.io/configuration/source/prometheus) to give Botkube access to those alerts. Once enabled, the Prometheus plugin can be configured to consume any alerts in the specified states and display them in any channels you have configured in your source binding configuration.

![Image 2: Prometheus alerts in Slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63e104c5fb11822e79cc09c8_ijd6WpT3jxksxNoy-fni7NrXRKM3p1iZAHutyVNDPEMjUTsgBkpAZ_pDEHEHgDKIHjjtpl_8V1Eu60mHmFUt2RIcRLv-TKdnmQIbF9op1r-VGZY6d6XTsB6zGkQHWB3r7wuk2ekw0o8Hl6adAZaoJmQ.png)

See the [documentation](https://docs.botkube.io/configuration/source/prometheus) to enable and configure the Prometheus plugin, it only requires a few settings.

Aliases
-------

As a way to save typing, Botkube had implemented the k and kc aliases for the kubectl command, much like DevOps admins like to do on the command line. Now you can [configure your own custom aliases](https://docs.botkube.io/configuration/alias) for any command in Botkube. Not only can you alias a single command like kubectl, you can create an alias for the full command including options and flags. In the example shown here, `kgp` is an alias for the full `kubectl get pods` command.

![Image 3: Kubectl aliases setup](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63e104c5e04d8205b4b47aa9_yHASb87npDigom97MSabXqtK812CEfQ08iOuqf9UU9KgcSY0kRwbcCdapzCZDonNyKi-HuxxAWgAFcKaMt0MYLJKZDTW9LY75QCNi4JA_vS7V0n8A0XmMBh-WzJLxMMoyBms6HUjDEiGcEs_VSZbTms.png)

Maria wrote another article on [Kubectl aliases](https://botkube.io/blog/command-line-magic-simplify-your-life-with-custom-kubernetes-kubectrl-aliases-on-botkube) and their use cases. For more details on how to configure your own Kubernetes aliases read this article and the documentation above.

Bug Fixes
---------

We have also fixed a couple of small bugs in Botkube that were reported to us by users and found by us. You can see the full list of changes in the [changelog](https://github.com/kubeshop/botkube/releases/tag/v0.18.0).

We appreciate your bug reports and pull requests! If you notice a bug in Botkube, please report it as a [GitHub issue](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

Feedback - We’d Love to Hear From You!
--------------------------------------

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the Botkube [GitHub issues](https://github.com/kubeshop/botkube/issues), Botkube [Slack community](http://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
