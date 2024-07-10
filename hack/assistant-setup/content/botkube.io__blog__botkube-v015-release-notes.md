Title: Botkube v0.15.0 Release notes

URL Source: https://botkube.io/blog/botkube-v015-release-notes

Published Time: Oct 24, 2022

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

Botkube v0.15.0 adds some great interactivity features aimed at application developers who may not be Kubernetes experts.

### Table of Contents

*   [Interactive kubectl](https://botkube.io/blog/botkube-v015-release-notes#interactive-kubectl)
*   [Actionable Events](https://botkube.io/blog/botkube-v015-release-notes#actionable-events)
*   [Output Filtering](https://botkube.io/blog/botkube-v015-release-notes#output-filtering)
*   [Interactive Output Filtering](https://botkube.io/blog/botkube-v015-release-notes#interactive-output-filtering)
*   [Bug Fixes](https://botkube.io/blog/botkube-v015-release-notes#bug-fixes)
*   [Feedback - We’d Love to Hear From You](https://botkube.io/blog/botkube-v015-release-notes#feedback-we-d-love-to-hear-from-you)

#### Start Using Botkube AI-Powered Assistant Today

We have an exciting early release of Botkube, just in time for KubeCon! We've been working as fast as we can to get some great new features ready to release. Here's v0.15.0 of Botkube, the most modern [ChatOps tool for Kubernetes](http://botkube.io/)!

v0.15.0 adds some great interactivity features aimed at application developers who may not be Kubernetes experts.

_If you have any feature requests or bug reports, reach out to us on [Slack](http://join.botkube.io/) or submit a [Github issue](https://github.com/kubeshop/botkube/issues)_

In this release, we focused on interactivity with Botkube and the kubectl executor. Some of the enhancements are currently Slack-specific using the additional interactive capabilities of the [new Slack app](https://botkube.io/docs/installation/socketslack/). We plan to enhance other platforms with interactive features where technically feasible.

Interactive kubectl
-------------------

If you're using the new Slack app, you can now use kubectl interactively. Simply run @botkube kubectl (or use the alias k or kc) and you will see an interactive kubectl command builder. Select the verb (get, in this example) and you can select the object type to get. Once you have selected those, based on context, you can select a namespace and, optionally, a specific object to act on. Leaving out the object will generally show the list of available objects. Once you've made your selections, you can see a preview of the constructed command. You can also specify a filter string to apply to the output (more about that later!) and click Run command. The command runs and you can see the output in the channel as usual.

This interactive kubectl tool is especially useful for developers and other users who want to see information about their applications but might not be Kubernetes and kubectl experts.

![Image 2: Getting Log reports of Kubernetes reports into Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6360eb442fdfd602024a9acd_6356bdf725fcea68c282c460_Y3wu_jOs-Ujqsrm-K_p5k9ixbT54IMlPnlOAyBDjuVsmkqOsLDkvcN2x3ji2mRn_CU0M-1ZAnp_lwDvLRhJ_eOTYodpOZAVhLHWxpjVfrTildXKVksWotV07CJHe9H7XRhRuxiVqJ2oTLEr7zr_vT6mRQQJKV_hD1EV65oMSXkUu2YYNFTXFdY6thg.png)

Actionable Events
-----------------

If you're receiving events in the new Slack app, you can now run contextual actions right from the event message. When a pod error comes in, for example, you can run kubectl commands that are relevant to the pod object right from the "Run command" drop-down box. The command runs automatically without any typing and returns the output in the channel. This makes it quick an easy to start investigating your alerts.

![Image 3: Run Kubectl commands directly from Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6360eb44adbd3b434f7dcbd4_6356bdf7fd97a85f48b6f4eb_tGR4IkDUpaGrpMiF7u7VXfDpliSzKiF6H8nCPk-xGSddI9heXZx6sKKPVzpy6PZVsMZ4znb37qqyJRsZZ5dQsOmxcVYNu_sXDLGcoFg88Htm7xjt_8pSsXxM3EQd-yigz8RAkChykcJEHGz8gY2zMUJca4nhxcaJoMfGLI2ZmD9YkWHn_4WTgsNPMQ.png)

Output Filtering
----------------

You can now filter the output of any command you run with Botkube, on any platform. Simply run the command with the `--filter "string"` flag and you will only see the matching output. The filter flag does simple linewise string matching. Any line in the output with a matching string will be shown, lines without a match are discarded.

![Image 4: Filtering Kubernetes logs for important alerts](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6360eb44701e5f37a18b49db_6356bdf7f7a6792610d5e3ca_kjx3ejlR6Aaq2y9mENfRr81flF6H3HMljZWwFis4-1rkG6jyjEiPmYwfxeMfU-_NoEE0OLtGHkPUxM90PhuFauG-7iPdakZtFEdZcqkbxRXEcnfkQnErMqivmwplsPmem-JQL-2X71Kpmh5-NS8jZJMSkWtfd2hqinQhht64CWwIJxgV4SBI1uwKzQ.png)

Interactive Output Filtering
----------------------------

Output filtering is even easier if you're using the [new Slack app](https://botkube.io/docs/installation/socketslack/). When you run a command with a lot of output (>15 lines) you will see a Filter output box. If you type a string in the box and press enter, you will then see the output of the previous command filtered to only show lines matching the filter string. This is great for quickly sifting through logs of a problematic pod, for example.

![Image 5: Searching through K8s logs](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6360eb4426a735c2c230b1cb_6356bdf7b95716455530d559_eIDU6Illetw8nvSmjY2RF0DUfo480e8Cg-z9oo9TrZoUMc82X5QeBP_LDPjPAByGdJGZVjqAZAeFF9MphVVECeN-D1KbDxttppO-YtjdKo6dUvOhLo9b1TypQbvEX9vTnvcTJVKaC1kDxQHOdTWFqHFty4VUYZKoIgXBMmch6sNCQp2-7Mde9Zm78g.png)

Bug Fixes
---------

We have also fixed a couple of small bugs in Botkube that were reported to us by users and found by us. You can see the full list of changes in the [changelog](https://github.com/kubeshop/botkube/releases/tag/v0.15.0).

We appreciate your bug reports and pull requests! If you notice a bug in BotKube, please report it as a [GitHub issue](https://github.com/kubeshop/botkube/issues) and we are happy to collaborate with you on getting it resolved.

Feedback - We’d Love to Hear From You
-------------------------------------

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the Botkube [GitHub issues](https://github.com/kubeshop/botkube/issues), Botkube [Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
