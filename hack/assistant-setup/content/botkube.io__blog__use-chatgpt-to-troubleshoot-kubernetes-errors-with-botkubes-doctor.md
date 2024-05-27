Title: ChatGPT to Troubleshoot K8s Errors with Botkube’s Doctor Plug-in

URL Source: https://botkube.io/blog/use-chatgpt-to-troubleshoot-kubernetes-errors-with-botkubes-doctor

Published Time: Jul 31, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3fb36b4e60920a3b1b2_hPLC9itV8zp-raGDFmvOZMfn2hV8RFcl237qzT8Wa1g.jpeg)

Kelly Revenaugh

Developer Relations Lead

Botkube

People have been using ChatGPT and AI for many uses over the past year. But have you heard of ChatGPT recommendations for Kubernetes troubleshooting? Learn more about our latest feature.

### Table of Contents

*   [Using ChatGPT with Kubernetes](#using-chatgpt-with-kubernetes)
*   [How Botkube Makes ChatGPT for Kubernetes Even Better](#how-botkube-makes-chatgpt-for-kubernetes-even-better)
*   [Common Use Cases with ChatGPT and Botkube](#common-use-cases-with-chatgpt-and-botkube)
*   [Is ChatGPT Safe to Use in Kubernetes?](#is-chatgpt-safe-to-use-in-kubernetes-)
*   [Get Started with Botkube’s Doctor Plugin](#get-started-with-botkube-s-doctor-plugin)

#### Start Using Botkube AI Assistant Today!

#### Get started with Botkube Cloud

Artificial Intelligence (AI) and its counterpart, GPT (generative pre-trained transformer), have been making incredible progress over the past year. Large Language Models power tools like ChatGPT to answer almost any question using the collective knowledge stored over years on the Internet.

And more recently, AI has entered the Kubernetes ecosystem. Throughout the past year, the industry has seen AI help with Kubernetes cost optimization, security management, troubleshooting, and more. So, the Botkube team harnessed the power of ChatGPT to answer questions and give recommendations on troubleshooting Kubernetes clusters with our new [AI Troubleshooting feature called Doctor](https://docs.botkube.io/usage/executor/doctor). ‍

‍

Learn more about how we developed Botkube’s Doctor (ChatGPT) plugin during an internal Hackathon in Pawel’s [blog post](https://botkube.io/blog/building-a-chatgpt-plugin-from-ideation-to-implementation).

Using ChatGPT with Kubernetes
-----------------------------

Using ChatGPT to troubleshoot issues in your Kubernetes cluster can save time and streamline your workflow. GPT (generative pre-trained transformer) automates many of the tedious and time-consuming tasks involved in troubleshooting, allowing you to focus on more important tasks. Additionally, ChatGPT's ability to analyze your cluster configuration and resource usage can help you optimize your cluster for better performance and cost savings.

How Botkube Makes ChatGPT for Kubernetes Even Better
----------------------------------------------------

Now, through the ChatOps-based interface in communication platforms like [Slack, MS Teams, Discord](http://botkube.io/integrations), Botkube users can now interact with ChatGPT and get recommendations while troubleshooting Kubernetes. Even though Kubernetes has been widely adopted over the past few years, there are still many developers and Ops teams who struggle with the complexity of Kubernetes. The ChatGPT plugin will help beginner and intermediate K8s users who are faced with troubleshooting issues or questions without the need for a Google search on Kubernetes documentation. every error they encounter.

Common Use Cases with ChatGPT and Botkube
-----------------------------------------

There are many ways to use ChatGPT when troubleshooting Kubernetes clusters. AI can help analyze logs, metrics, and other data from your Kubernetes cluster to identify issues – then returns recommendations on how to fix and improve errors and issues.

Here's a few use cases as examples:

### Assistance for Beginner Kubernetes Users

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64c7f3a7c42051dab4872cab_8hyWVq9NOkFLCQpe-EqzxKczhU5VIqqG_bm2kP876TdzysK0Z3PGJOXCBF3aPo8wIV9w8bC5n77ksg5I62jg7KlzmZpnmxNRmP2yvLTrxWZaHYv2tZFxGfQAo21ky2infXvJraVs4RbpiM4Jiyl1ulA.png)

After the Doctor plugin is installed, users are able to use the `Get Help` button located under incoming error events, like the one shown above. This helps lower the barrier of entry for new Kubernetes and/or Botkube users.

The Get Help action passes the error data to ChatGPT, which returns a response with actionable buttons to execute commands, like `kubectl describe pod` or `kubectl logs pod`, instead of scouring the Internet for the correct answer.

Users can also ask ChatGPT specific questions directly from the interface. Users can ask questions about Kubernetes, like “Where does Kubernetes store its state?” to tutorial-like questions like “How do I create a Pod?”

__Keep in mind that the ChatGPT plugin will only give recommendations and expects the Botkube user to execute commands.__

### ChatGPT Recommendations

![Image 3](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64c7f3b7be2361ae1bb5afd0_LEr5f9Rr-O1pgKD_dYqFRG8GcwopzDWXYkDiVyEL3as3vJF3r1DQhDHa4SZs0sQD2NmW8sHJ3XUVMxUAp5z8WJx-mIuyXanea4788oniZnR0o4m2UrCZXKe-Uj8RZufiLihfB__BQYzgNo3uG2IrIwY.png)

The new Doctor plugin allows you to integrate [Botkube's existing Kubernetes recommendations](https://www.youtube.com/watch?v=9D2tASyx7eA) derived from cluster logs to receive personalized suggestions based on the specific information extracted from your Kubernetes cluster.

Whether it's optimizing resource allocation, enhancing security measures, or improving overall performance, the Doctor plug-in empowers users with valuable insights to make informed decisions for their Kubernetes environments.

### Troubleshooting Automations

Users can harness ChatGPT capabilities to craft custom Botkube configurations and [automated actions](https://docs.botkube.io/usage/automated-actions) effortlessly, working seamlessly with other plugins. By leveraging Slack interactivity as a command builder, users can efficiently interact with the system, receiving tailored suggestions for automations that directly address the specific challenges encountered within their Kubernetes cluster. The Doctor plug-in empowers users to find precise solutions and execute commands with ease, streamlining the troubleshooting process and enhancing overall cluster management.

Is ChatGPT Safe to Use in Kubernetes?
-------------------------------------

Giving access to your Kubernetes clusters in production to __any__ tool can be nerve-wracking. We take our users’ data and security seriously. Luckily, ChatGPT in Botkube only generates information and recommendations based on what you ask it and will not execute any commands. Take a look at our [Terms of Service](https://app.botkube.io/terms-of-service), [End User License Agreement](https://app.botkube.io/eula) and [Privacy Policy](https://app.botkube.io/privacy-policy).

Get Started with Botkube’s Doctor Plugin
----------------------------------------

Ready to try it out on your own? The easiest way to configure it is through the [Botkube web app](https://app.botkube.io/) if your cluster is connected. Otherwise you can enable it in your [Botkube YAML configuration](https://docs.botkube.io/configuration/executor/doctor).

Once enabled, you can ask questions about specific resources or ask free-form questions, directly from any enabled channel. Find out how to use the Doctor plugin in the [documentation](https://docs.botkube.io/usage/executor/doctor).

We’d love to hear how you are using ChatGPT to troubleshoot your Kubernetes clusters! Share your experiences with us in the Botkube [Slack community](http://join.botkube.io/) or [email our Developer Advocate, Maria](mailto:maria@kubeshop.io) and we’ll send you some fun swag.
