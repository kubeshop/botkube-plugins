Title: Simplify Kubernetes Security Troubleshooting with AI

URL Source: https://botkube.io/blog/simplify-kubernetes-security-troubleshooting-with-ai

Published Time: May 08, 2024

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/663ad9425dee853429acf422_T01TRG6U4AV-U06CH7HDJTS-34c4e3601ac1-512.jpg)

Vaidas Jablonskis

Software Engineer

Botkube

Learn how Botkube can streamline troubleshooting and help overcome common Kubernetes PSA-related issues

### Table of Contents

*   [What is Kubernetes Pod Security Admission (PSA)?](#what-is-kubernetes-pod-security-admission-psa--2)
*   [Common Challenges with Pod Security Policies](#common-challenges-with-pod-security-policies-2)
*   [Getting Started with Kubernetes PSA Troubleshooting Example](#getting-started-with-kubernetes-psa-troubleshooting-example-2)
*   [Enforce Pod Security Standards with Namespace Labels](#enforce-pod-security-standards-with-namespace-labels-2)
*   [Summary](#summary-2)
*   [Try Out Botkube for Kubernetes Security Best Practices Today](#try-out-botkube-for-kubernetes-security-best-practices-today-2)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

#### Start Using Botkube AI Assistant Today!

Kubernetes is a powerful container orchestration platform, but its complexity can leave even seasoned developers scratching their heads. One particularly tricky area is \[Pod Security Admission (PSA)\](https://kubernetes.io/docs/concepts/security/pod-security-admission/). While PSA plays a crucial role in cluster security, it can also be the source of unexpected deployment failures. That's where the AI-powered Botkube Assistant, specifically designed for Kubernetes, comes to the rescue. Botkube understands the complexities of your cluster and guides you towards solutions in clear, simple language.

In this blog post, we'll demonstrate how Botkube can streamline troubleshooting and help you overcome common PSA-related issues.

\## What is Kubernetes Pod Security Admission (PSA)?

\- The Pod Security Admission Controller (PSA) acts as a gatekeeper within your Kubernetes cluster. It evaluates Pod specifications against predefined security policies to ensure your applications adhere to security best practices. Think of PSA as your cluster's security guard, making sure that potentially risky deployments are blocked before they can cause problems.

\## Common Challenges with Pod Security Policies

\- **\*\*Privileged Containers\*\***: Some applications may demand privileged access, but this can expose your cluster to vulnerabilities. Pod Security Admission Controller can enforce strict policies to prevent privileged containers unless absolutely necessary.

\- \*\*Host Network/Ports\*\*: Allowing pods direct access to your host network or specific privileged ports increases attack vectors. PSA can restrict this access.

\- \*\*Root User:\*\* Running containers as the root user is a major security risk. PSA can block or warn about pods attempting to run as root.

\- \*\*Volume Mounts\*\*: PSA can be used to prevent your containers from accessing sensitive directories on your host system for enhanced protection.

\- \*\*Capabilities\*\*: Linux capabilities give containers a wide range of permissions. PSA can help you enforce a 'least privilege' model by dropping unnecessary capabilities and minimizing risks.

\## Getting Started with Kubernetes PSA Troubleshooting Example

Deploying applications in Kubernetes can sometimes lead to unexpected hiccups when the Pod Security Admission (PSA) controller steps in. Understanding why your Pod fails to launch and how to quickly resolve these security-related errors is crucial for smooth operations. In this tutorial, we'll walk you through a real-world scenario of a deployment blocked by PSA.  We'll demonstrate how Botkube’s AI assistant can help you:

\- Pinpoint the exact reasons behind the deployment failure.

\- Understand the relevant k8s security concepts.

\- Guide you through fixing the issues step-by-step.

Let's get started!

\### Prerequisites

\- Kubernetes Cluster with PSA: Consult your provider's documentation on enabling the Pod Security Admission Controller.

\- \[Botkube Instance\](http://app.botkube.io/): Follow the installation guide in the \[Botkube documentation\](https://docs.botkube.io/examples-and-tutorials/getstarted).

\### A Note on Environments

To fully explore the examples in this tutorial, you'll need a few things in place:

\- \*\*Kubernetes Cluster with PSA:\*\* You'll need a Kubernetes cluster where the Pod Security Admission (PSA) controller is active. If you're using a managed Kubernetes service (like GKE, EKS, or AKS), PSA is likely enabled by default. For self-managed clusters, consult your setup documentation for instructions on enabling PSA.

If you're new to Kubernetes or want to experiment without setting up a full production cluster, you can leverage a local Kubernetes environment. The official Kubernetes documentation provides a great tutorial specifically focused on deploying a local cluster with Pod Security Standards enabled: \[https://kubernetes.io/docs/concepts/security/pod-security-standards/\](https://kubernetes.io/docs/concepts/security/pod-security-standards/). This guide walks you through setting up Kind, a popular local Kubernetes tool.

\- \*\*Botkube Installation\*\*: You'll want to have \[Botkube installed and connected to your Kubernetes cluster\](https://docs.botkube.io/examples-and-tutorials/getstarted). Botkube provides clear installation guides for different environments. Refer to the installation documentation for the process that best suits your setup.

_\*Note: While exploring Botkube with a local cluster is a great way to learn, keep in mind that real-world production deployments might require additional considerations.\*_

\## Enforce Pod Security Standards with Namespace Labels

\- Create a Namespace Manifest: Define your namespace and include the desired Pod Security Admission (PSA) labels. Here's the annotated version:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="NamespaceManifest-PSABotkube.yaml"><tbody><tr><td id="file-namespacemanifest-psabotkube-yaml-L1" data-line-number="1"></td><td id="file-namespacemanifest-psabotkube-yaml-LC1">---</td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L2" data-line-number="2"></td><td id="file-namespacemanifest-psabotkube-yaml-LC2"><span>apiVersion</span>: <span>v1</span></td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L3" data-line-number="3"></td><td id="file-namespacemanifest-psabotkube-yaml-LC3"><span>kind</span>: <span>Namespace</span></td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L4" data-line-number="4"></td><td id="file-namespacemanifest-psabotkube-yaml-LC4"><span>metadata</span>:</td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L5" data-line-number="5"></td><td id="file-namespacemanifest-psabotkube-yaml-LC5"><span>labels</span>:</td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L6" data-line-number="6"></td><td id="file-namespacemanifest-psabotkube-yaml-LC6"><span>pod-security.kubernetes.io/enforce</span>: <span>restricted</span></td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L7" data-line-number="7"></td><td id="file-namespacemanifest-psabotkube-yaml-LC7"><span>pod-security.kubernetes.io/audit</span>: <span>restricted</span></td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L8" data-line-number="8"></td><td id="file-namespacemanifest-psabotkube-yaml-LC8"><span>pod-security.kubernetes.io/warn</span>: <span>restricted</span></td></tr><tr><td id="file-namespacemanifest-psabotkube-yaml-L9" data-line-number="9"></td><td id="file-namespacemanifest-psabotkube-yaml-LC9"><span>name</span>: <span>psa</span></td></tr></tbody></table>

\- Apply the Manifest: Use kubectl to create the namespace:

`kubectl apply -f namespace.yaml`

Key points:

\- Enforcement: Pods deployed in this namespace must meet the 'restricted' requirements. Violations will block pod creation.

\- Audit and Warning: The 'restricted' profile is used for logging purposes, giving you insights into potential issues.

Now, you’ll see that we have a very simple deployment below. It’s a shell one-liner which runs in a loop and prints data every 10 seconds. We'll deploy our simple 'sleeper' service, which is intentionally designed with potential security flaws. This will allow us to see how the Pod Security Admission controller reacts and how Botkube helps us fix it.

`kubectl -n psa apply -f deployment.yaml`

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="deployment.yaml"><tbody><tr><td id="file-deployment-yaml-L1" data-line-number="1"></td><td id="file-deployment-yaml-LC1">---</td></tr><tr><td id="file-deployment-yaml-L2" data-line-number="2"></td><td id="file-deployment-yaml-LC2"><span><span>#</span> https://kubernetes.io/docs/concepts/workloads/controllers/deployment/</span></td></tr><tr><td id="file-deployment-yaml-L3" data-line-number="3"></td><td id="file-deployment-yaml-LC3"><span>apiVersion</span>: <span>apps/v1</span></td></tr><tr><td id="file-deployment-yaml-L4" data-line-number="4"></td><td id="file-deployment-yaml-LC4"><span>kind</span>: <span>Deployment</span></td></tr><tr><td id="file-deployment-yaml-L5" data-line-number="5"></td><td id="file-deployment-yaml-LC5"><span>metadata</span>:</td></tr><tr><td id="file-deployment-yaml-L6" data-line-number="6"></td><td id="file-deployment-yaml-LC6"><span>name</span>: <span>sleeper</span></td></tr><tr><td id="file-deployment-yaml-L7" data-line-number="7"></td><td id="file-deployment-yaml-LC7"><span>namespace</span>: <span>psa</span></td></tr><tr><td id="file-deployment-yaml-L8" data-line-number="8"></td><td id="file-deployment-yaml-LC8"><span>labels</span>:</td></tr><tr><td id="file-deployment-yaml-L9" data-line-number="9"></td><td id="file-deployment-yaml-LC9"><span>app</span>: <span>sleeper</span></td></tr><tr><td id="file-deployment-yaml-L10" data-line-number="10"></td><td id="file-deployment-yaml-LC10"><span>spec</span>:</td></tr><tr><td id="file-deployment-yaml-L11" data-line-number="11"></td><td id="file-deployment-yaml-LC11"><span>selector</span>:</td></tr><tr><td id="file-deployment-yaml-L12" data-line-number="12"></td><td id="file-deployment-yaml-LC12"><span>matchLabels</span>:</td></tr><tr><td id="file-deployment-yaml-L13" data-line-number="13"></td><td id="file-deployment-yaml-LC13"><span>app</span>: <span>sleeper</span></td></tr><tr><td id="file-deployment-yaml-L14" data-line-number="14"></td><td id="file-deployment-yaml-LC14"><span>replicas</span>: <span>1</span></td></tr><tr><td id="file-deployment-yaml-L15" data-line-number="15"></td><td id="file-deployment-yaml-LC15"><span>template</span>:</td></tr><tr><td id="file-deployment-yaml-L16" data-line-number="16"></td><td id="file-deployment-yaml-LC16"><span>metadata</span>:</td></tr><tr><td id="file-deployment-yaml-L17" data-line-number="17"></td><td id="file-deployment-yaml-LC17"><span>labels</span>:</td></tr><tr><td id="file-deployment-yaml-L18" data-line-number="18"></td><td id="file-deployment-yaml-LC18"><span>app</span>: <span>sleeper</span></td></tr><tr><td id="file-deployment-yaml-L19" data-line-number="19"></td><td id="file-deployment-yaml-LC19"><span>spec</span>:</td></tr><tr><td id="file-deployment-yaml-L20" data-line-number="20"></td><td id="file-deployment-yaml-LC20"><span>containers</span>:</td></tr><tr><td id="file-deployment-yaml-L21" data-line-number="21"></td><td id="file-deployment-yaml-LC21">- <span>name</span>: <span>sleeper</span></td></tr><tr><td id="file-deployment-yaml-L22" data-line-number="22"></td><td id="file-deployment-yaml-LC22"><span>image</span>: <span>alpine:latest</span></td></tr><tr><td id="file-deployment-yaml-L23" data-line-number="23"></td><td id="file-deployment-yaml-LC23"><span>imagePullPolicy</span>: <span>IfNotPresent</span></td></tr><tr><td id="file-deployment-yaml-L24" data-line-number="24"></td><td id="file-deployment-yaml-LC24"><span>command</span>:</td></tr><tr><td id="file-deployment-yaml-L25" data-line-number="25"></td><td id="file-deployment-yaml-LC25">- <span><span>"</span>/bin/sh<span>"</span></span></td></tr><tr><td id="file-deployment-yaml-L26" data-line-number="26"></td><td id="file-deployment-yaml-LC26">- <span><span>"</span>-c<span>"</span></span></td></tr><tr><td id="file-deployment-yaml-L27" data-line-number="27"></td><td id="file-deployment-yaml-LC27">- <span>|</span></td></tr><tr><td id="file-deployment-yaml-L28" data-line-number="28"></td><td id="file-deployment-yaml-LC28"><span>while true; do date; sleep 10; done</span></td></tr><tr><td id="file-deployment-yaml-L29" data-line-number="29"></td><td id="file-deployment-yaml-LC29"><span></span><span>volumeMounts</span>:</td></tr><tr><td id="file-deployment-yaml-L30" data-line-number="30"></td><td id="file-deployment-yaml-LC30">- <span>name</span>: <span>localtime</span></td></tr><tr><td id="file-deployment-yaml-L31" data-line-number="31"></td><td id="file-deployment-yaml-LC31"><span>mountPath</span>: <span>/etc/localtime</span></td></tr><tr><td id="file-deployment-yaml-L32" data-line-number="32"></td><td id="file-deployment-yaml-LC32"><span>volumes</span>:</td></tr><tr><td id="file-deployment-yaml-L33" data-line-number="33"></td><td id="file-deployment-yaml-LC33">- <span>name</span>: <span>localtime</span></td></tr><tr><td id="file-deployment-yaml-L34" data-line-number="34"></td><td id="file-deployment-yaml-LC34"><span>hostPath</span>:</td></tr><tr><td id="file-deployment-yaml-L35" data-line-number="35"></td><td id="file-deployment-yaml-LC35"><span>path</span>: <span>/usr/share/zoneinfo/Asia/Tokyo</span></td></tr><tr><td id="file-deployment-yaml-L36" data-line-number="36"></td><td id="file-deployment-yaml-LC36"><span>restartPolicy</span>: <span>Always</span></td></tr><tr><td id="file-deployment-yaml-L37" data-line-number="37"></td><td id="file-deployment-yaml-LC37"></td></tr></tbody></table>

‍

1\. Our deployment failed! The error messages may not be immediately clear, especially if we're not very familiar with Pod Security Admission policies.

Let's call in our expert assistant, Botkube. We'll ask Botkube to analyze the error output and provide guidance on resolving the issues.

2\. Now that the assistant suggested how we could fix the issue, let’s apply the changes.

3\. Argh, we've fixed one thing, but another issue has cropped up! This is common in troubleshooting scenarios, especially with complex systems like Kubernetes.

Not to worry, Botkube is still here to help! We'll provide the new error messages to Botkube and get further recommendations.

4\. We'll carefully follow Botkube's instructions to modify the sleeper deployment configuration. This will include adding the missing runAsUser property.

5\. We've made the necessary changes, and our deployment succeeds! Our 'sleeper' service is now running in compliance with our cluster's security policies.

6\. Troubleshooting can be tedious, but Botkube made the process much smoother. Its ability to explain the underlying problems and guide us towards solutions saved us valuable time.

\## Summary

The beauty of using an AI assistant like Botkube lies in its ability to streamline troubleshooting through an iterative, conversational process. Instead of poring over complex error messages or Kubernetes documentation, you can engage with Botkube as if you were explaining the problem to a knowledgeable colleague. With each interaction, it pinpoints potential issues, suggests solutions, and helps you verify those fixes. This not only saves you valuable time but also presents a fantastic opportunity to learn about Kubernetes security principles along the way. By seeing how each change addresses specific Pod Security Admission violations, you gain a deeper understanding of how to build secure and compliant deployments from the ground up.

\## Try Out Botkube for Kubernetes Security Best Practices Today

Ready to streamline your Kubernetes security practices? \[Sign up for Botkube today\](http://app.botkube.io/)—it's free to get started and can quickly empower your team to proactively monitor, manage, and secure your cluster all within your chat platform like Slack or Microsoft Teams. Join the growing community of users who simplify Kubernetes management easily with just a few clicks.

‍
