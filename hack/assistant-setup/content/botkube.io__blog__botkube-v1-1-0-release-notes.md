Title: Botkube v1.1.0 Release Notes

URL Source: https://botkube.io/blog/botkube-v1-1-0-release-notes

Published Time: Jun 21, 2023

Markdown Content:
![Image 1: a woman wearing headphones on an airplane](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Our most recent realese of Botkube Cloud & Open Source brings important features that help with Kubernetes troubleshooting. This update brings quick Slack connection to your cluster!

### Table of Contents

*   [What is Botkube?](#what-is-botkube-)
*   [Botkube Cloud with Cloud Slack Integration](#botkube-cloud-with-cloud-slack-integration)
*   [Here are the key features of Botkube Cloud:](#here-are-the-key-features-of-botkube-cloud-)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

We are excited to announce the latest release of Botkube, packed with new features and improvements to enhance your Kubernetes collaborative troubleshooting experience. In this release, we focused on two major areas: the introduction of Botkube Cloud with Cloud Slack integration and several bug fixes and enhancements for our open-source version.

What is Botkube?
----------------

Botkube is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred messaging and collaboration platforms like Slack, Microsoft Teams, Discord, and Mattermost. This means you'll gain instant visibility and control over your cluster resources, all without ever having to leave your messaging platform. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more. In this blog post, we'll give you an overview of all of Botkube’s newest features from the v1.1.0 release.

Botkube Cloud with Cloud Slack Integration
------------------------------------------

With the growing popularity of cloud-based collaboration tools, we recognized the need to provide seamless integration with popular platforms like Slack. That's why we are thrilled to introduce the Botkube Cloud Slack app that integrates with the Botkube web app. This makes it easier than ever to receive Kubernetes alerts and notifications directly in your Slack workspace, and manage all of your Kubernetes clusters using a single Slack app.

Here are the key features of Botkube Cloud:
-------------------------------------------

### Easy Installation:

Setting up Botkube Cloud is a breeze. Simply connect your Slack workspace to the Botkube app with the click of a button.

![Image 2: connect cluster in azure stack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64940fc9213e4ff10a1b4672_jaaGtHC2ScR0Nac7Bts5ZqRHFfLVeeZYeY4oQk6BN3mYKpUVH09FW11MtSqAA7UBZAc-YOWj38D9F54ahYH5xRE0oDlxc9K17hYum4BIp30W64-cZxHWJJZwKrv7mNYV9-yJ46QlBcSWkiiLgIzIt20.png)

‍

Invite the Botkube bot to the channels you want to use, and continue your configuration. Install Botkube into your Kubernetes clusters, and you're ready to manage your clusters from Slack.

![Image 3: a screenshot of the adwords dashboard](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64940fe3950f20abf76fa3ce_StacAl4VR0BYqGhlJz4HOSii505G8uK8S4WNAptRmeQ0zBU3u_tzqj6Cnwmi5fjYC1aR0ssDiaKbxrrwFatfr0nolxyrFGtBbHKnlNXHpqE2oLA4GwMNU_qc6ufRj-4c71QTo5fDSK0GkHhH31kQjlY.png)

‍

### Multi-Cluster Support:

Botkube Cloud allows you to manage multiple Kubernetes clusters using a single Slack app. The new Botkube Slack app uses our cloud services to route executor commands to the correct cluster. There's no longer a need for a separate Slack app per cluster. This centralized approach saves time and effort when managing alerts and notifications across different environments.

### Open Source Enhancements:

While the majority of our efforts went into Botkube Cloud, we also dedicated some time to improving the open-source version. In this release, we focused on bug fixes and enhancing security and usability. Here are some of the notable changes:

### Early Configuration Error Detection:

We understand the frustration of misconfigurations, so we added a new feature to provide users with immediate feedback when there is an error in their configuration. This helps users identify and rectify any issues quickly, ensuring smooth operation.

### Binary SHA Validation:

To improve security, we now validate the SHA of downloaded binaries. This ensures the integrity of the Botkube installation and protects against potential tampering or unauthorized modifications.

### Ignored Namespace Optimization:

We addressed an issue where Botkube was spamming notifications about events in ignored namespaces. With the fix, you can now confidently ignore specific namespaces without being bombarded with unnecessary alerts.

### Mattermost Bug Fix:

Thanks to an external contribution, we resolved a bug related to Mattermost integration. The fix ensures smooth and reliable communication with Mattermost, enhancing the overall user experience.

We would like to express our gratitude to the community for their ongoing support and feedback. Your contributions have helped us refine and enhance Botkube, making it a more robust and reliable tool for Kubernetes troubleshooting.

We encourage you to update your Botkube installation to benefit from these enhancements. Your feedback and suggestions are invaluable in shaping the future of Botkube.

You can talk to us in the Botkube [GitHub issues](https://github.com/kubeshop/botkube/issues), Botkube [Slack community](http://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).

‍
