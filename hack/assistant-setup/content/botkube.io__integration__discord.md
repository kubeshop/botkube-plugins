Title: Discord & Botkube Kubernetes Integration

URL Source: https://botkube.io/integration/discord

Markdown Content:
**Discord** is a popular communication platform that was initially developed for gamers, but it has since evolved to become a hub for all kinds of communities. The platform offers text, voice, and video communication features, enabling users to connect and interact with others in real-time. Discord allows users to join servers or create their own communities, making it easy to connect with like-minded people from all over the world. It also offers a range of customization options, including the ability to add bots, customize notifications, and more. With its user-friendly interface and extensive feature set, Discord has become a go-to platform for online communication and community building.

Key Value Propositions:

*   Real-time communication: Discord allows users to communicate with others in real-time using text, voice, and video chat features.
*   Community building: Users can join servers or create their own communities, making it easy to connect with people who share similar interests.
*   Customization: Discord offers a range of customization options, allowing users to tailor their experience to their preferences.
*   Accessibility: Discord is available on multiple platforms, including desktop, mobile, and web, making it accessible to users across different devices.
*   Security: Discord offers a range of security features, including two-factor authentication, IP location tracking, and more, to help ensure the safety of its users.

Kubernetes in Discord Before Botkube
------------------------------------

In the past, developers had to create their own custom webhooks to receive Kubernetes alerts in their Discord channel, which could be a time-consuming and complex process, especially when dealing with multiple alerts. Building custom webhooks was a time-intensive process that required developers to create and maintain multiple webhooks for each alert they wanted to receive, making it similar to building mini applications just to test their own applications. As a result, Botkube's integration with Discord has become a time-saving and efficient solution that simplifies Kubernetes alerting and allows developers to focus on building applications without the added burden of webhook maintenance.

Discord + Botkube K8s Integration
---------------------------------

With the Botkube integration into Discord, developers can now receive Kubernetes and Cluster alerts directly in their collaboration tool of choice, including Discord. This integration allows for real-time notifications of Kubernetes cluster events, such as pod failures and autoscaling, to be sent directly to Discord channels. Even Helm deployment alerts can now show up in the channel, providing developers with the option to troubleshoot directly from the Chat. With Botkube's integration into Discord, developers can focus on their work and stay on top of Kubernetes alerts without having to worry about setting up and managing custom webhooks.

‍

### New Kubernetes AI Bot for Discord

After release 1.10 of Botkube, every install now comes with Botkube AI by default. What does that mean? It means that you now can troubleshoot your Kubernetes cluster directly from Discord with the power of AI. If you keep getting a repetitive error, such as the [OOMKILLED](https://botkube.io/learn/what-is-oomkilled) error, you can ask Botkube's AI what pod is causing that error. Once the AI finds the failing pod, it will suggest the fix which can be applied in with click thanks to Botkube's ability to run kubectl commands.
