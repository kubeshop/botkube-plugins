Title: Argo & Botkube: A Match Made in Heaven?

URL Source: https://botkube.io/blog/argo-cd-botkube-integration

Published Time: Sep 22, 2023

Markdown Content:
![Image 1: a woman wearing headphones on an airplane](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Argo CD now integrates with Botkube, allowing users to receive and interact with notifications in Slack, Discord, and more.

### Table of Contents

*   [What is Argo CD?](#what-is-argo-cd-)
*   [Why was a Botkube & Argo Connection Necessary?](#why-was-a-botkube-argo-connection-necessary-)
*   [What is Botkube?](#what-is-botkube-)
*   [Using Argo CD with Botkube](#using-argo-cd-with-botkube)
*   [Get Started with Botkube’s new Argo CD Plugin](#get-started-with-botkube-s-new-argo-cd-plugin)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

#### Get started with Botkube Cloud

Are you dealing with the complexities of scaling operations and collaborative workflows using GitOps tools like ArgoCD? As your organization expands, staying on top of repository and configuration changes, as well as managing pull requests and issues, can become increasingly difficult.

‍ Today, we are happy to announce the new [Botkube Argo CD integration](https://docs.botkube.io/configuration/source/argocd/)! This integration enables you and your team to leverage automation to scale and optimize your Argo CD troubleshooting workflow.

What is Argo CD?
----------------

![Image 2: a screen shot of a computer screen showing a number of different devices](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/650e09b69191288d41cf2161_rih56gy96kbqx4UzlDDmVadKW9ieXnfmbXLOwzJiDWgHBDzmW0vG867PZM74YdzH5YkNHY-9F2xaVfJTam8eFpvSgzoB4EX-FxDPzLzqMvKJmSNtSBwIRifp2EctcHW3oeh_ruepqkKpwhfFyDzS5Kc.gif)

‍

[Argo CD](https://botkube.io/integration/argo-cd-botkube-kubernetes-integration) is a Kubernetes controller that oversees the status of all active applications. It can pull updated code from Git repositories and deploy it directly to Kubernetes resources. It enables developers to manage both infrastructure configuration and application updates in one system.

Why was a Botkube & Argo Connection Necessary?
----------------------------------------------

With the new Botkube ArgoCD plugin, users can significantly enhance their Argo CD notifications experience. This plugin offers a streamlined approach to managing platform secrets and notifications, eliminating the need for project-specific configurations, whether it's Kubernetes, Prometheus, Argo, or others. Enabling Argo notifications becomes more straightforward, providing users with an efficient and centralized solution.

Furthermore, the plugin enhances visibility into the Argo deployment process by consolidating GitHub and Argo events and delivering them directly to the user's configured communication platform, ensuring a more informed and responsive workflow.

What is Botkube?
----------------

Botkube is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred collaboration platforms like [Slack, Microsoft Teams, Discord, and Mattermost.](https://botkube.io/integrations) This means you'll gain instant visibility and control over your cluster resources, all without ever having to leave your messaging platform. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more.

Using Argo CD with Botkube
--------------------------

![Image 3: a screenshot of a screen showing the status of an application](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/650e09c855b42178c42a1d9b_jOhrHB90gwPhqwSU94v3y1Q7Q2Y_1Ltfap5j-mY6XbgieOkVITkVOoOboVTaVHT55onYtmncvcVt_zMrOQehiIOKbM2unJi5NKvWpXhjN222CbEB31JP_oSxT9QowgHWFcKv0YoK2FvZZvJMwGpET4s.png)

‍

The new [Botkube ArgoCD](https://docs.botkube.io/configuration/source/argocd/) plugin offers a compelling use case for users seeking seamless integration between Argo CD and their preferred communication platforms. With this plugin, users can effortlessly receive real-time notifications from their Argo CD deployments directly in their communication channels. This feature proves invaluable for teams relying on Argo CD for managing Kubernetes workflows within GitHub. Botkube's ArgoCD plugin harnesses Argo CD's native notification system, expanding its functionality by provinteractive notifications for popular cloud-based services like Slack, as well as support for platforms like Discord, which are not covered by Argo CD's default notifications.

One of the standout benefits of this plugin is its automation of the entire notification system setup, simplifying a traditionally complex process. By configuring the plugin in Botkube, users can save time and effort previously required to define webhooks, triggers, templates, and annotations within Argo CD.

Additionally, the plugin enhances notifications, allowing users to run commands and access application details conveniently in a slack channel, ultimately aiding in debugging and troubleshooting. Overall, the Botkube ArgoCD plugin enhances the user experience by bridging the gap between Argo CD and their communication platforms, offering a streamlined, efficient, and interactive approach to managing Kubernetes deployments.

Get Started with Botkube’s new Argo CD Plugin
---------------------------------------------

Ready to try it out on your own? The easiest way to configure it is through the [Botkube web app](https://app.botkube.io/) if your cluster is connected. Otherwise you can enable it in your [Botkube YAML configuration](https://docs.botkube.io/configuration/source/argocd). Check our [tutorial](https://botkube.io/blog/getting-started-with-botkube-and-argocd) for step by step instruction on how to get started.

Once enabled, you can ask questions about specific resources or ask free-form questions, directly from any enabled channel. Find out how to use [the Argo plugin](https://docs.botkube.io/usage/source/argocd) in the documentation.

We’d love to hear how you are using the new Argo CD plugin! Share your experiences with us in the Botkube [Slack community](https://join.botkube.io/) or [email our Developer Advocate, Maria](mailto:maria@kubeshop.io) and we’ll send you some Botkube swag.

‍
