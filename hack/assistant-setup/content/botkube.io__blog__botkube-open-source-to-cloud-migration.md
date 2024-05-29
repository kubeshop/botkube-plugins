Title: Seamlessly Migrate from Botkube Open Source to Cloud

URL Source: https://botkube.io/blog/botkube-open-source-to-cloud-migration

Published Time: Aug 09, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3fb36b4e60920a3b1b2_hPLC9itV8zp-raGDFmvOZMfn2hV8RFcl237qzT8Wa1g.jpeg)

Kelly Revenaugh

Developer Relations Lead

Botkube

With our new CLI migration tool, current open-source users are able to migrate their data and configurations over to Botkube Cloud in minutes.

### Table of Contents

*   [Differences Between Botkube Open Source and Botkube Cloud](#differences-between-botkube-open-source-and-botkube-cloud)
*   [What is the Botkube CLI?](#what-is-the-botkube-cli-)
*   [How to Migrate Your Existing Botkube Account to Cloud](#how-to-migrate-your-existing-botkube-account-to-cloud)
*   [Give it a try!](#give-it-a-try-)

#### Start Using Botkube AI Assistant Today

#### Get started with Botkube Cloud

With our latest release, we’ve made transitioning from the open source version of Botkube to our Cloud version even easier. [Botkube Cloud](http://botkube.io/) is built on top of our open source core and gives users even more benefits – including audit and event logs to track changes in your Kubernetes clusters, an easy-to-use visual dashboard to configure your settings, multi-cluster management, and much more.

With our new [CLI migration too](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud)l, current open-source users are able to migrate their data and configurations over to the Cloud version of Botkube with a few simple steps.

Differences Between Botkube Open Source and Botkube Cloud
---------------------------------------------------------

We believe that both the open source version and cloud version of Botkube help supercharge Kubernetes teams’ technical workflows with seamless alert consolidation, monitoring, and troubleshooting – all from the same communication platforms that teams use every day.

Botkube was developed in 2019 and has helped hundreds of teams collaborate and take on challenging Kubernetes-related tasks each day. We’ve create Botkube Cloud for teams that have growing multi-cluster environments and/or want to take of advanced features, not available in our open source version.

Here are some of the highlights of Botkube Cloud:**‍**

**Easy configuration and installation**

Botkube Cloud makes it easy to configure and install Botkube in a new Kubernetes cluster. Before Botkube Cloud, users had to manually install Botkube with a YAML file and update their Helm charts manually for each configuration change. To install with Cloud, just create an instance in the web app, then configure your communication platforms and Botkube plugins. You'll get an installation command that you can copy and paste to install Botkube. Once installed, Botkube will synchronize its configuration with the cloud and any changes you make in the web management interface will automatically be pushed to the cluster.**‍**

**Multi-cluster management**

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64d5072717f5862f1c0ba90e_64385d9ea0e4ca059b5a7a1d_Botkube.png)

Botkube Cloud makes it easy to manage Botkube in multiple Kubernetes clusters. Previously, if users wanted to use Botkube executors (e.g. kubectl commands) and have multiple Kubernetes clusters, users needed to create and install a Botkube Slack app for each cluster.

Today, with Botkube Cloud, you can see the status and manage configuration for all of your clusters in one place in the web hosted app. When adding new clusters, you can copy an existing configuration to get going even faster.**‍**

**Advanced Multi-Cluster Slack bot**

The Botkube Cloud Slack bot can be installed in your Slack workspace with a single click. It allows you to see your Slack channels in your instance configuration, ensure the Botkube bot is invited to the correct channels, and select which channels to use with Botkube plugins. It uses Botkube cloud services to route executor commands to the correct Kubernetes cluster so you can control multiple clusters from one Slack channel. This solves the multi-cluster problem that many open source users encountered previously.**‍**

**Audit and Event logs**

![Image 3](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64d507420a0b1ce9ecb2d0dd_64385e6998cfac2dfc8d887f_Event%20and%20Audit%20Logs.png)

Botkube Cloud collects and aggregates audit and event logs from all of your Botkube-enabled Kubernetes clusters. All Botkube notifications are stored as events and any Botkube executor commands are stored as audit log entries. Logs can be filtered and sorted and you can see what activity is happening in your clusters, correlate events, and find the root cause of problems that arise.**‍**

What is the Botkube CLI?
------------------------

[Botkube CLI](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud) is a tool for installing and migrating Botkube installations from an open source instance to a Botkube Cloud instance.

Botkube CLI is the best way to do an interactive installation of Botkube. The botkube install command runs the installation process and reports the progress and any errors back to you immediately. For automated installations using CI or GitOps tools, you can still use Helm to install the Botkube Helm chart unattended.**‍**

How to Migrate Your Existing Botkube Account to Cloud
-----------------------------------------------------

If you are using Botkube in your Kubernetes cluster with local configuration, the CLI provides a simple on-ramp to the advanced Botkube Cloud services. Migrating to the Cloud version of Botkube is extremely easy and only takes a few commands!

Here's how to [migrate your Botkube instance](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud):

1.  Create a [Botkube Cloud](https://app.botkube.io/) account
    
2.  Install the [Botkube CLI](https://docs.botkube.io/cli/getting-started)
    
3.  Run `botkube login` to [authenticate](https://docs.botkube.io/cli/getting-started#first-use) your CLI
    
4.  Run `botkube migrate` to [move your configuration](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud) to the cloud and connect your Botkube installation automatically
    

Give it a try!
--------------

**‍**We’re here to help! If you need some help about migrating over, please reach out to us on the Botkube Slack workspace. Our team is happy to lend a hand as you make the transition over and answer any questions you may have. We’d love for you to try out the advanced features that we’ve developed over the past few months to make managing and troubleshooting your Kubernetes clusters as painless as possible.

Existing Botkube users can take advantage of new features like the interactive control plane and configuration dashboard immediately with the Free tier after migrating. New Cloud users also receive the Pro tier for free with a 30-day trial period (no credit card required!). If your team is looking to upgrade to the Pro or Enterprise tiers, we are giving a special discounted pricing to existing Botkube users. Please reach out to our [project leader, Blair](mailto:blair@kubeshop.io), to inquire about the discount.

Hope you enjoy the new features.
