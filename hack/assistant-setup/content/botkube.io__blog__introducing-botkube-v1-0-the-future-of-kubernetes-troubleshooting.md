Title: Botkube v1.0 - the Future of Kubernetes Troubleshooting

URL Source: https://botkube.io/blog/introducing-botkube-v1-0-the-future-of-kubernetes-troubleshooting

Published Time: May 04, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Something new is coming to the Botkube universe...

### Table of Contents

*   [What is Botkube?](#what-is-botkube-)
*   [RBAC support](#rbac-support-)
*   [Interactive Control Plane](#interactive-control-plane-)
*   [Additional Features](#additional-features)
*   [Sign up now](#sign-up-now)
*   [Feedback](#feedback)

#### Start Using Botkube AI Assistant Today!

We are thrilled to announce the launch of [Botkube v1.0](https://app.botkube.io/) This new version of Botkube introduces a hosted control plane for multi-cluster management and monitoring of all of your Botkube instances.Troubleshooting in the K8s world can be like trying to navigate through a mosh pit - chaotic and unpredictable. With alerts flying in from all directions and collaboration being a somewhat ad hoc process, it's easy to feel like you're drowning in a sea of chaos. Even if developers have access to the K8s cluster, it's often not a smooth process - secret keys can be tough to revoke and retrieving logs is done manually. And let's not forget about security - accessing the cluster to perform operations must be limited to approved networks and devices. Sometimes, remote troubleshooting just isn't an option. It's time for a better way to troubleshoot K8s.That’s where Botkube comes in.

What is Botkube?
----------------

Botkube is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred messaging and collaboration platforms like Slack, Microsoft Teams, Discord, and Mattermost. This means you'll gain instant visibility and control over your cluster resources, all without ever having to leave your messaging platform. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more.. In this blog post, we'll give you an introduction to Botkube and highlight some of its key features and benefits.

Key Features of Botkube v1.0
----------------------------

RBAC support
------------

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6437222690593713ca726589_botkube-read-only-717ed01cf9fa5e6621f2a09c7b29a32d.svg)

Botkube RBAC Architecture

Botkube features role-based access control ([RBAC](https://docs.botkube.io/configuration/rbac/)), enabling plugins or communication channels to map to Kubernetes users or groups. This means that when a plugin or channel is mapped, it takes on the permissions of the user or group, thereby limiting the actions that the plugin or channel can perform in the cluster. With this feature, Botkube's security and access controls are flexible and precise, providing enhanced security and control to developers and DevOps engineers alike.

Interactive Control Plane
-------------------------

Botkube brings an exciting new feature to the table - a [web-based control plane](https://app.botkube.io/) that works across multiple clusters. With this powerful tool, you can create new Botkube instances in a flash, making installation into a K8s cluster a breeze. The control plane also enables you to modify the configuration of communication platforms and plugins, and even synchronize the configuration to Botkube instances automatically.

![Image 3](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/643728ace69f4305d971ce6f_Screen%20Shot%202023-04-12%20at%202.54.18%20PM.png)

Botkube Control Plane

The control plane also collects and displays event logs and audit logs in a single, easy-to-read list. This feature is incredibly useful for monitoring event notifications sent from all sources, as well as commands run by all executors. With this aggregated list, you can quickly identify and resolve any issues that arise in your K8s clusters.

Additional Features
-------------------

Botkube has made some exciting additions to its core features:

*   The Kubernetes source and kubectl executors have been extracted as plugins to take advantage of all the plugin functionalities.
*   Executors and sources now support more advanced message formatting and interactivity to enhance user experience.
*   Prometheus messages are now more readable with new message formatting.
*   Negative regex patterns for event and resource constraints are now supported.
*   Plugin development guides have been updated to cover the new interactive features.

These improvements provide greater flexibility and functionality for Botkube users, making it easier to troubleshoot Kubernetes clusters and streamline communication between teams.

Sign up now
-----------

[Get started with Botkube](https://app.botkube.io/)! Whether you're a seasoned Kubernetes pro or just getting started, Botkube has something to offer. Sign up now for free and join the community of users who are already benefiting from the power of Botkube. Don't miss out on this opportunity to streamline your K8s workflow and take your skills to the next level - try Botkube today!

Feedback
--------

We welcome developers and Kubernetes enthusiasts to explore the platform and share their valuable feedback. We want to know what you think of Botkube and how we can make it even better. We're doing quick 15-minute interviews to get your feedback, and as a thank you, we'll give you some cool Botkube plushies and t-shirts and enter you into a raffle for a chance to win a $50 Amazon gift card.

You can also talk to us in the Botkube [**GitHub issues**](https://github.com/kubeshop/botkube/issues), Botkube [**Slack community**](https://join.botkube.io/), or email our Product Leader at [**blair@kubeshop.io**](mailto:blair@kubeshop.io).
