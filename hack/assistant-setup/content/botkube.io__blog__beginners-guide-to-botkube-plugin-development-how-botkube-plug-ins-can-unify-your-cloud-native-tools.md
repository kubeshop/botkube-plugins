Title: Beginner's Guide to Botkube Plugin Development: Part 1

URL Source: https://botkube.io/blog/beginners-guide-to-botkube-plugin-development-how-botkube-plug-ins-can-unify-your-cloud-native-tools

Published Time: Mar 30, 2023

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

This introduction will teach you how Botkube plug-ins can unify your cloud native tools.

### Table of Contents

*   [Why make a Botkube Plug-in: The Power of CNCF Tools Together](https://botkube.io/blog/beginners-guide-to-botkube-plugin-development-how-botkube-plug-ins-can-unify-your-cloud-native-tools#why-make-a-botkube-plug-in-the-power-of-cncf-tools-together-)
*   [How Botkube Plug-Ins Work:](https://botkube.io/blog/beginners-guide-to-botkube-plugin-development-how-botkube-plug-ins-can-unify-your-cloud-native-tools#how-botkube-plug-ins-work-)
*   [Conclusion](https://botkube.io/blog/beginners-guide-to-botkube-plugin-development-how-botkube-plug-ins-can-unify-your-cloud-native-tools#conclusion-)

#### Start Using Botkube AI-Powered Assistant Today

So you want to make a [Botkube Plug-in](https://botkube.io/integrations), but don't know where to begin? Fear not, because this series has got you covered! In short, Botkube plugins are tools that enhance the Botkube user experience to make it more suited to fit their goals. They provide users with extra features and options to customize their experience. However, to fully utilize Botkube, it is necessary to configure it according to your specific needs. Starting from scratch with only documentation to guide you can be intimidating, which is why I am writing this series to help jumpstart your plug-in development journey regardless of your experience level. With this series, you’ll have everything you need to build your own Botkube Plug-in.

![Image 2](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64259de0f04e131b26b323ab_YDm2BILjPQoWfmfEClCKc4da4Acv3ASayZTaEEEdd27wxMLEb-eWsECh1qpEKoqabyu43YFgDmSULoDFTIwBMCF7ndEFnTzG0bLBSjA1xFx0v_cBNjx8zBD1owLk8IfqA8nK2IiIJ_qtrJcU3dGb-s8.png)

But first things first, let's talk about why plug-ins are so important.

**Why make a Botkube Plug-in: The Power of CNCF Tools Together**
----------------------------------------------------------------

As a DevOps or platform engineer, your role is to ensure that your team has access to a suite of tools that enable efficient and quality deployments. However, with over 1,178 open-source tools available in the [CNCF Landscape](https://landscape.cncf.io/?organization=kubeshop&selected=botkube) across multiple categories, choosing which tools to use and integrating them into your organization's workflow can be overwhelming.

![Image 3](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64259def371cfe07bc0f2b8e_s7cvlLKT1mIpSir81gNynMEUtfKsi34yzhstUPE8tTuzP8sRizuzesbF3elxcr1ynfIKa_fWLqAsX2a0jT9Z52xisGCDxGUlKZACDTlJkmQ7lKOttNSomQgDDiYKJ0yt_mZF2KhHRlSbygHiNiod_3g.png)

Choosing independent tools that work well together and unifying them to work alongside other collaborative conversation tools can be a daunting task. Furthermore, the lack of established or certified routes for Kubernetes deployment, monitoring, and management can lead to a situation where your tools don't communicate well with each other, which is a significant problem.

Fortunately, [Botkube](https://botkube.io/) and its plugin system provide an effective solution to this problem.

Botkube makes it easy to unify all the cloud-native tools into a single platform. As a developer or operations team member, you know that automation is the key to simplifying your workflow and increasing productivity. If you're using Botkube for monitoring and troubleshooting Kubernetes clusters, you're already ahead of the game. But with the [introduction of a plugin system](https://botkube.io/blog/botkube-v017-release-notes) in Botkube v0.17.0, you can take Botkube to the next level with your own custom plugins.

The plugin system allows anyone to write new plugins for sources and executors, which can be published for other Botkube users to install and use. This means that you can extend the functionality of Botkube to fit your specific needs, making it an even more valuable tool in your toolkit.

**How Botkube Plug-Ins Work:**
------------------------------

![Image 4](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64259e0b4fdb15ca2aa6e839_iKasmcohH7JW9C8d9Y10Y6AM8n1uSpA-HtyxcqE6jxI9XBN53Nott737B5XTDLMGB1nXUnkGPN-fHeNbq4RIbsuy6Kko3LdT9hrdf-YpXukG2kVCCgbSDNvPltY4coW4PmCmXIdeyO0luWUNcDxVk0I.png)

Botkube's architecture is designed to make it easy for developers and operations teams to interact with APIs and other tools through sources and executors. Think of __sources as the ears of Botkube__ - they listen for events or notifications from any tool that can send them. This means that you can easily integrate any tool that can send events, like [Prometheus](https://botkube.io/integration/prometheus) into Botkube, and have those events sent to communication channels based on the source bindings that you've configured.

Alternatively, __executors are like the hands of Botkube__ - they're used to run commands from other tools within a chat channel. This means that you can execute commands from your favorite tools, like kubectl or Helm, directly within your chat channel and have the output of those commands returned to the same channel for easy visibility.

Botkube already comes with a native source for Kubernetes events, which has been part of Botkube since the beginning. But with the release of the new plugin system, you can now add additional sources and executors to Botkube, including helm and Prometheus plugins that we already support.

**Conclusion**
--------------

Now that you understand the importance of Botkube's plugin system and how it works, it's time to start building your own Botkube plugin. In the next blog post, we'll walk through a step-by-step guide on how to build a Botkube plugin from scratch, using a simple example.

Stay tuned for the next blog post, where we'll dive deep into the world of Botkube plugins and show you how to build one yourself.

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap and get the features you’d like implemented.

There are plenty of options to contact us:

*   [**GitHub issues**](https://github.com/kubeshop/botkube/issues)
*   [**Slack**](https://join.botkube.io/)
*   or email our Product Leader at blair@kubeshop.io.
