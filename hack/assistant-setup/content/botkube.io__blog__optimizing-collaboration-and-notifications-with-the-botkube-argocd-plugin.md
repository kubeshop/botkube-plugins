Title: Botkube's ArgoCD Plugin Optimizes Collaboration + Notifications

URL Source: https://botkube.io/blog/optimizing-collaboration-and-notifications-with-the-botkube-argocd-plugin

Published Time: Oct 26, 2023

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/636df3f0810c273feb4b4ad3_VEyGvbbIe6fYlHwidQsOYSS3FvzlHuQsOUjeuZzXWJw.jpeg)

Josef Karásek

Software Engineer

Botkube

A closer look into the Botkube team's process for developing the ArgoCD plugin

### Table of Contents

*   [ArgoCD Plugin for Kubernetes Deployment Management](https://botkube.io/blog/optimizing-collaboration-and-notifications-with-the-botkube-argocd-plugin#argocd-plugin-for-kubernetes-deployment-management)
*   [How Botkube Uses ArgoCD for GitOps Management](https://botkube.io/blog/optimizing-collaboration-and-notifications-with-the-botkube-argocd-plugin#how-botkube-uses-argocd-for-gitops-management)
*   [Conclusion](https://botkube.io/blog/optimizing-collaboration-and-notifications-with-the-botkube-argocd-plugin#conclusion)
*   [Sign up now!](https://botkube.io/blog/optimizing-collaboration-and-notifications-with-the-botkube-argocd-plugin#sign-up-now-)
*   [Feedback](https://botkube.io/blog/optimizing-collaboration-and-notifications-with-the-botkube-argocd-plugin#feedback)

#### Start Using Botkube AI-Powered Assistant Today

IT operations have always been very dynamic and rapidly evolving and even more so in the cloud-native world. Insight into the state of the system is crucial for maintaining operational awareness and ensuring that the system is running as expected. This is especially true in Kubernetes where the system is constantly changing and evolving. At Botkube we believe that notifications from Continuous Delivery tools, such as ArgoCD, are a key part of this.

ArgoCD Plugin for Kubernetes Deployment Management
--------------------------------------------------

![Image 2: Argo CD branches being checked for app health](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/650e09b69191288d41cf2161_rih56gy96kbqx4UzlDDmVadKW9ieXnfmbXLOwzJiDWgHBDzmW0vG867PZM74YdzH5YkNHY-9F2xaVfJTam8eFpvSgzoB4EX-FxDPzLzqMvKJmSNtSBwIRifp2EctcHW3oeh_ruepqkKpwhfFyDzS5Kc.gif)

‍

ArgoCD has become the standard solution for managing Kubernetes applications. It provides a declarative GitOps workflow for continuous delivery and automated updates. Which is why we decided to invest our time and effort into developing a new [Botkube plugin for ArgoCD](https://botkube.io/integration/argo-cd-botkube-kubernetes-integration). We wanted to make it easier for ArgoCD users to get started with Botkube by auto-configuring notifications and providing default message templates and triggers.

‍

The Botkube engineering team developed a [new plugin that automatically configures ArgoCD notifications](https://docs.botkube.io/usage/source/argocd) for all communication platforms supported by Botkube, such as [Slack, Microsoft Teams, and Discord](https://botkube.io/integrations). This plugin greatly reduces the barrier to entry for ArgoCD users and enhances collaboration and notifications in Kubernetes management. In this article, we'll explore the Botkube engineering team's journey from using ArgoCD to developing their own Botkube ArgoCD plugin.

How Botkube Uses ArgoCD for GitOps Management
---------------------------------------------

![Image 3: Diagram of Developer teams use of Argo CD and Git](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/653a912003e50d0ea9eb0c15_Argo%20Sync%20Image%20(1).png)

Diagram trying to show how Botkube uses ArgoCD and Git to manage K8s Clusters

At Botkube, we use two long-lived Kubernetes clusters which both run in Google Kubernetes Engine (GKE). One cluster is used for production and the other one is used for staging. Each cluster runs its own Argo instance, which is created during cluster creation. A majority of development and bug fixing is done in local dev environments, such as Kind or k3d clusters and is only merged into the main branch after passing all tests in CI. This means that the main branch is always in a deployable state and ArgoCD in the staging cluster consumes directly from the main branch with the auto-sync feature on. Releases into production need to be triggered manually and can happen less frequently. Although we incorporate time for changes to be tested during the staging process, we have a short feature development cycle and our team is quite comfortable with releasing to prod on Friday evenings.

‍

As with every new Botkube feature, we were the first users. We wanted to make sure that the plugin was easy to use and that it provided value to our users. We also wanted to make sure that the plugin was stable and that it didn't cause any issues in our own ArgoCD instances. We were also able to get valuable feedback from our own team members and use that to improve the plugin.

‍

[Once the plugin is installed](https://docs.botkube.io/configuration/source/argocd/), it automatically configures ArgoCD notifications for all communication platforms supported by Botkube, such as Slack, Microsoft Teams, and Discord. The plugin also provides real-time updates about ArgoCD events, such as "New Application Version Detected" or "Sync Started", detailed change information, and customizable alerts. This aligns with the GitOps philosophy of declarative configuration and enables users to easily monitor and manage their Kubernetes applications.

The plugin can be installed either via web UI at [app.botkube.io](https://app.botkube.io/) or with the Botkube CLI.

  Botkube install --version v1.5.0 \\
  --set sources.argocd.Botkube/argocd.enabled=true \\
  --set 'rbac.groups.argocd.create'=true \\
  --set 'sources.argocd.Botkube/argocd.config.defaultSubscriptions.applications\[0\].name'=guestbook \\
  --set 'sources.argocd.Botkube/argocd.config.defaultSubscriptions.applications\[0\].namespace'=argocd \\
  --set communications.default-group.socketSlack.enabled=true \\
  --set communications.default-group.socketSlack.channels.default.name=${SLACK\_CHANNEL\_NAME} \\
  --set communications.default-group.socketSlack.appToken=${SLACK\_API\_APP\_TOKEN} \\
  --set communications.default-group.socketSlack.botToken=${SLACK\_API\_BOT\_TOKEN}

‍

‍

Of course, the Botkube engineering team uses Botkube in both staging and prod.

We have a dedicated channel for Botkube notifications in our Slack workspace and we have configured the Botkube ArgoCD plugin to send notifications to this channel as well.

![Image 4: Showing Kubernetes Cluster created and deleted in Slack with Botkube](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/653a68dff56c4c123324282e_L72F7n2Dmu8c1Ua4Zpyw7FLyLF4LTUh7AjQ8cKUg5u8TguoHXxfwbYBJYMtZjMXCm6PXk3xyEj-dwF83OsRdwoA-RXiMUHSnIQppKb6WnZSim6V8x5_1vp94dlRVuFj7L_fFlwG7Ir_VYrORpIZkzmA.png)

‍

### Setting Up ArgoCD Notifications

First lesson after enabling the plugin was that setting up ArgoCD notifications, in this case for Slack, was really easy. We just had to point Botkube at which ArgoCD Applications it should watch and the plugin took care of the rest. We didn't have to worry about creating new message templates or event triggers. Botkube took care of all of that for us. After a while of using the plugin, we were happy that the defaults we'd gone with were working well for us. We didn't need to change anything.

### Real Use Case: Horizontal Pod Autoscaler

‍

The next big lesson was an issue between ArgoCD and the HorizontalPodAutoscaler controller which we found thanks to the incoming notifications. We stored HorizontalPodAutoscaler definition in our git, which then ArgoCD applies to the cluster and the HorizontalPodAutoscaler controller reads it (or rather HPO conversion [webhook](https://github.com/kubernetes/kubernetes/issues/74099)). What we did not know was that the webhook also maintains order of some of the fields and in this case it differed from what we had in git. This caused ArgoCD to think that the resource was out of sync and it kept trying to apply the same definition which the webhook reverted over and over again. Naturally, Botkube kept sending notifications about this: "OutOfSync" and "Synced". We fixed the issue by storing the resource definition in git in the same order as the webhook was expecting it.

‍

Several staging and prod releases have now gone by and the value of the plugin has become clear. Easy path to greater operational awareness. The plugin handled all necessary configuration tasks and we're now able to keep track of changes and ensure consistency across environments. This is important to us especially when we run ArgoCD with auto-sync enabled.

Conclusion
----------

As we continue to improve the capabilities of the plugin, we have identified potential improvements, including the implementation of features like simplified onboarding for a large number of ArgoCD Applications through label selectors or regular expressions.

‍

Your feedback, whether you've already used the plugin or are just considering trying it out, is valuable to us. We encourage you to share your insights and suggestions.The Botkube team welcomes [contributions](https://github.com/kubeshop/botkube) from all its users. Together, we can improve and expand the functionality of this tool for the benefit of the entire community.

Sign up now!
------------

Get started with Botkube! Whether you’re a seasoned Kubernetes pro or just getting started, Botkube has something to offer. Sign up now for free and join the community of users who are already benefiting from the power of Botkube.

Feedback
--------

We welcome developers and Kubernetes enthusiasts to explore the platform and share their valuable feedback. We want to know what you think of Botkube and how we can make it even better. We're doing quick 15-minute interviews to get your feedback, and as a thank you, we'll give you some cool Botkube plushies and t-shirts and enter you into a raffle for a chance to win a $50 Amazon gift card! Just email our Developer Advocate, Maria or use this calendly [link](https://calendly.com/maria-botkube) to sign up.You can also talk to us in the Botkube GitHub issues, connect with others and get help in the Botkube [Slack community](http://join.botkube.io/), or email our Product Leader at blair@kubeshop.io.
