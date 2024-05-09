Title: Introducing Botkube's Integration with Flux

URL Source: https://botkube.io/blog/introducing-botkubes-integration-with-flux

Published Time: Aug 28, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

As the demand for seamless scalability and efficient teamwork increases, an innovative solution is needed. Enter the Botkube Flux integration to tackle collaboration, automation, and scaling head-on.

### Table of Contents

*   [Introduction](#introduction)
*   [What is Flux?](#what-is-flux-)
*   [What is Botkube?](#what-is-botkube-)
*   [Why is it important?](#why-is-it-important-)
*   [Optimized Flux Workflow](#optimized-flux-workflow)
*   [Get Started with Botkube’s new Flux Plugin](#get-started-with-botkube-s-new-flux-plugin)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

#### Get started with Botkube Cloud

Introduction
------------

Struggling to manage scaling operations and collaborative efforts with GitOps tools like Flux? When it comes to checking what changes have happened in your repositories and configurations, things can get complex. And as your organization grows, it becomes trickier to keep track of pull requests and issues in a straightforward way. As the demand for seamless scalability and efficient teamwork increases, an innovative solution is needed. Enter the [Botkube Flux integration](https://botkube.io/integration/botkube-flux-kubernetes-integration)—a game-changer in tackling your collaboration, automation, and scaling challenges head-on. It enables real-time collaboration by delivering immediate alerts about pod crashes, resource issues, or deployment failures to chat platforms, facilitating rapid decision-making. With Botkube, monitoring and actions take place in the chat platform, creating a centralized knowledge hub that streamlines information sharing and actions, while its automation capabilities reduce repetitive tasks and increase productivity for tasks like scaling deployments or examining logs.With Botkube, say goodbye to tedious manual processes and say hello to a new world of event-driven possibilities, empowering you with unparalleled cluster interactivity options.

What is Flux?
-------------

[Flux CD](https://fluxcd.io/) is an open-source continuous delivery and GitOps solution tailored to simplify and automate the deployment and lifecycle management of applications and infrastructure on Kubernetes. Designed for developers and operations professionals, Flux CD empowers teams to declaratively define the desired state of their applications and configurations using code stored in a Git repository.

Flux continuously monitors your repository for any changes and seamlessly applies updates to your Kubernetes cluster. This ensures that your actual state matches the desired state.

What is Botkube?
----------------

![Image 2: Showing how to run Kubernetes commands from Slack with Botkube](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64b96a341b5ccb59ffb87637_act-on-events.gif)

Botkube is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred collaboration platforms like [Slack, Microsoft Teams, Discord, and Mattermost.](https://botkube.io/integrations) This means you'll gain instant visibility and control over your cluster resources, all without ever having to leave your messaging platform. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more. Botkube now works with Flux!

Why is it important?
--------------------

Before we dive into the realm of Botkube, let's address the pain points that many of us have encountered with traditional manual GitOps workflows. Consider the series of steps you'd need to navigate without Botkube's assistance

### List of Common GitOps Issues

#### Installing Flux CLI

The first hurdle is configuring the Flux CLI on your local system, a process that often involves multiple configurations and settings.

#### Cluster Connection

Connecting to your Kubernetes cluster manually, and potentially switching contexts between different environments like development, staging, and production, can be time-consuming and error-prone.

#### Repository Cloning/PR Checkout

Locally cloning the repository housing the pull request for testing changes is yet another step in the manual approach.

#### Sharing Insights

Sharing updates and information requires toggling between platforms like Slack and GitHub to post diff reports, which can lead to fragmented communication.

Optimized Flux Workflow
-----------------------

![Image 3: Optimizing Flux CD workflow for Kubernetes](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64e694bca6bd600e8a7e88dd_flux-diff-1.gif)

With Botkube's new Flux executor, you can simplify complex tasks using a single command

    @BotKube flux diff kustomization podinfo --path ./kustomize --github-ref [PR Number| URL | Branch]
    

This command works right in your preferred chat platform like Slack or Teams, making everything easy. Get ready to experience a world where innovation and user-friendly simplicity come together!

### Auto-discovery of GitHub Repository

Seamlessly identifies the associated GitHub repository linked to the provided kustomization.

### Effortless Repository Cloning

The git repository is cloned without manual intervention.

### Precision with Pull Requests

The specified pull request is accurately reviewed by our [AI assistant](https://botkube.io/integration/chatgpt-botkube-kubernetes-integration) for processing.

### Seamless State Comparison

A comprehensive comparison between the pull request changes and the current cluster state is performed.

### Accessible Diff Reports

The outcome of the comparison is shared conveniently via the designated Slack channel.

Get Started with Botkube’s new Flux Plugin
------------------------------------------

![Image 4: Live Kubernetes Cluster connected to slack with Flux CD workflows](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64ecb79001bc7e01e2c88804_flux-interactivity.gif)

Ready to try it out on your own? The easiest way to configure it is through the [Botkube web app](https://app.botkube.io/) if your cluster is connected. Otherwise you can enable it in your [Botkube YAML configuration](https://docs.botkube.io/configuration/executor/flux).

Once enabled, you can ask questions about specific resources or ask free-form questions, directly from any enabled channel. Find out how to use [the Flux plugin](https://docs.botkube.io/usage/executor/flux) in the documentation.

We’d love to hear how you are using GitOps! Share your experiences with us in the Botkube [Slack community](https://join.botkube.io/) or [email our Developer Advocate, Maria](mailto:maria@kubeshop.io) and we’ll send you some fun swag.
