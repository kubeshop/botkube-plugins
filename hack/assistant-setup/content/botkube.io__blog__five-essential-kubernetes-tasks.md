Title: 5 Essential Kubernetes Troubleshooting + Monitoring Tasks to Automate

URL Source: https://botkube.io/blog/five-essential-kubernetes-tasks

Published Time: Jan 11, 2024

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Learn how to simplify DevOps tasks directly within your favorite communication platform like Slack and Microsoft Teams

### Table of Contents

*   [Real-time Kubernetes Monitoring with Prometheus](#real-time-kubernetes-monitoring-with-prometheus)
*   [Receive Real Time ArgoCD Notifications in Slack](#receive-real-time-argocd-notifications-in-slack)
*   [Kubernetes Insights with Botkube’s ChatGPT Assistant](#kubernetes-insights-with-botkube-s-chatgpt-assistant)
*   [Kubernetes Troubleshooting Automations](#kubernetes-troubleshooting-automations)
*   [Executor Plugin](#executor-plugin-)
*   [Conclusion](#conclusion)

#### Start Using Botkube AI Assistant Today

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams

If you're curious about the exciting possibilities of using Kubernetes in your chat platform, you've come to the right place. With Botkube you can complete all your collaborative Kubernetes troubleshooting and monitoring tasks right from your chat window! In this blog post, we'll explore five exciting ways to get the most out of Botkube. Whether you're a beginner or an experienced Kubernetes user, you'll learn how to simplify DevOps tasks, empower your teams, and enhance your Kubernetes experience.

[Botkube](http://app.botkube.io/) is a collaborative Kubernetes troubleshooting and monitoring tool designed for both DevOps experts and developers who may not be Kubernetes experts. Botkube helps teams quickly respond to issues by sending timely alerts about what's happening in their Kubernetes environments. It's not just about alerts though; Botkube also lets teams automate responses, run Kubernetes commands, and follow Kubernetes best practices. Plus, it integrates with popular communication platforms like Slack, Microsoft Teams, Discord, and Mattermost, making it a valuable asset for any team working with Kubernetes.

Real-time Kubernetes Monitoring with Prometheus
-----------------------------------------------

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65baac094d6cd1122a89f719_Screenshot_Template_Botkube%20(3).png)

[Real-time monitoring](https://botkube.io/solutions) of Kubernetes is the foundation of Botkube's features. With this feature, you gain access to a live, dynamic view of your Kubernetes environment. While Botkube has a long-standing history of supporting Kubernetes events, Kubernetes clusters often generate an abundance of state events and metrics. Typically, tools like kube-state-metrics and custom applications are configured to utilize Prometheus as a metrics store. Prometheus' Alertmanager then identifies metrics in anomalous states and triggers corresponding alerts. To seamlessly integrate this functionality into Botkube, we've developed the [Botkube source plugin for Prometheus](https://botkube.io/integration/prometheus).

This feature enables you to instantaneously assess the performance and health of your applications, pods, and services. With real-time insight into your cluster's state, you can swiftly identify and address issues, minimizing downtime and optimizing resource allocation.

Real-time monitoring brings significant advantages that help improve application reliability and system efficiency. DevOps teams, in particular, will find this feature invaluable for maintaining high-performance Kubernetes environments.

The versatility of real-time monitoring extends across various use cases, from microservices health monitoring to resource consumption tracking and performance bottleneck identification. With Botkube's real-time monitoring capabilities, you're well-equipped to tackle a wide range of Kubernetes challenges.

Receive Real Time ArgoCD Notifications in Slack
-----------------------------------------------

Botkube's [integration with Argo CD](https://botkube.io/integration/argo-cd-botkube-kubernetes-integration) simplifies deployments and enables a seamless GitOps workflow. Argo CD is a continuous delivery tool that streamlines the code updates from Git repositories and deployment to Kubernetes resources, managing infrastructure configuration and application updates within a single system.

The Botkube ArgoCD plugin simplifies Argo CD notifications by unifying platform secrets and notifications across projects like Kubernetes, Prometheus, and Argo. It enhances visibility by consolidating GitHub and Argo events, so users can now receive real-time notifications from their Argo CD deployments directly in their communication channels, especially valuable for teams managing Kubernetes workflows within GitHub. Setting up ArgoCD notification natively is a complex process that requires installing triggers and templates, adding email credentials to a secret, registering the email notification service with Gmail credentials, and subscribing to notifications for specific applications or projects. To begin receiving notifications with Botkube, simply provide your application name, Argo URL, and namespace, and ensure ArgoCD is enabled in your RBAC permissions. This plugin also allows you to utilize Argo CD's native notification system in chat platforms like MS Teams and Discord which are not supported natively.

![Image 3](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/650e09c855b42178c42a1d9b_jOhrHB90gwPhqwSU94v3y1Q7Q2Y_1Ltfap5j-mY6XbgieOkVITkVOoOboVTaVHT55onYtmncvcVt_zMrOQehiIOKbM2unJi5NKvWpXhjN222CbEB31JP_oSxT9QowgHWFcKv0YoK2FvZZvJMwGpET4s.png)

Botkube sending ArgoCD notifications to Slack

### Automating Kubernetes Operations

![Image 4](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a0710c644fa0ebb76293d8_DJDInRt7FR5LTwmVqnG4WM9OBv7o9_FmRKnG5sA9F-UU-kqljSWEtByVtVP37PhGh2wq7eezjjCNzzjlYyIOyqlAfEMDA6UdSCs5AUJLKfcy3qqXg8cEOoJTdi4S-5Z_Otd9bgcKLoeY5gEcWNa0D4U.gif)

‍

Botkube excels in fostering collaboration and efficiency through its integration with popular chat platforms. Botkube's [automated action feature](https://docs.botkube.io/configuration/action) transforms Kubernetes management for DevOps teams.It introduces automated commands triggered by events, delivering context to communication channels seamlessly. Some examples of these automations are “get-created-resource”

![Image 5](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a039d1bc5724b39fa1b44d_GXqDS0fufecSqTlSqV-h3ch5pmFTNSdEY7mE5IXYXidytlJBIM8f2EfS2_xpxIhxJqxK23lp5aXMrw6zlD5qK6RaFBh4l9ZZ5PZ5v7ud5JrE65atfkRPW0FdTqzm55LEYKPhvjjDSl8DQYShs_l1GAg.png)

‍

Actions consist of two components: the command (currently kubectl) and the event source binding. The command, defined with Go template syntax, utilizes event metadata for execution. Event source bindings connect actions to specific sources, enhancing integration with event filters. When a source binds to both an action and a communication channel, action results accompany the original event. Actions offer flexibility, from read-only kubectl configurations to customized commands like auto-restarting hanging pods, sending results to channels.

### Botkube includes two preconfigured actions in the Cloud Dashboard

![Image 6](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a0710b68ad02138d5720ff_v4OOZo5A1CCE3VZToEZfBiIuLviwI71gpu5gWv_erbnK1YfbO1a7oYWiqmNlwJZO-9ZuXEWD50ls3ylgorzhg76Lf6PpXpvjleX-sA3vrhuboF-61-bn37dMqsNd5Q6BvC9FIdMYbJ_KBfDJu3LtIkA.gif)

#### describe-created resource:

Automatically runs kubectl describe on resource creation events, delivering results to bound channels.

‍

![Image 7](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a0710c7946a3e903c20f78_gh7vWb3pXGzzEya6P_SNo6ZdoeJHG1rE2whT7wfwf4-kF9hFunVFf2_O2j2zjJk6zUXCaT62Yw4GxcQCoPpmKaWTIzKXnOQu-fO9niWeypy2nsigHIBM1DyEu4Rr2WLf7KnXK0zlyBk_9I3HN9I8GJM.gif)

#### Show-logs-on-error

Executes kubectl logs on error events, delivering logs to relevant channels.

Botkube's automated action feature simplifies Kubernetes management by automating commands in response to events, improving DevOps workflows, and ensuring effective collaboration.

Kubernetes Insights with Botkube’s ChatGPT Assistant
----------------------------------------------------

![Image 8](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64c21a2723cb531baf9f0867_doctor-interactive-help.gif)

‍

Botkube doesn't stop at monitoring and alerting; [Botkube’s Doctor Plugin](https://botkube.io/blog/use-chatgpt-to-troubleshoot-kubernetes-errors-with-botkubes-doctor) allows you to tap into the power of ChatGPT right from your communication platform. There are many ways to use the Doctor plugin when troubleshooting Kubernetes clusters. It can help analyze logs, metrics, and other data from your Kubernetes cluster to identify issues – then returns recommendations on how to fix and improve errors and issues.

### Here's a few use cases as examples:

#### Assistance for Beginner Kubernetes Users

![Image 9](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a039d1b4b8569bc008cf16_d3_gGf91GFWCqjVrQEfoSLImw32RhITxnMPRZVlJg2Ybc5N6JFl5nXM2zJCM4goiFIgIAFe5vEjA5KJd4HIkFM8-ZBtLYvInlk_zYFHYK6ksTvaP1yo32cZf-3TpFnP97sl7WXWZj8SViQXXDA8ZMuo.png)

‍

With the [Doctor plugin](https://botkube.io/integration/chatgpt-botkube-kubernetes-integration) installed, users are able to use the **Get Help** button located under incoming error events, like the one shown above. This helps lower the barrier of entry for new Kubernetes users. The Get Help button passes the error data to ChatGPT, which returns a response with actionable buttons to execute commands, like kubectl describe pod or kubectl logs pod which saves time instead of searching the internet or looking through Kubernetes docs.

Users can also ask ChatGPT specific questions directly from the interface. Users can ask questions about Kubernetes, like “Where does Kubernetes store its state?” to tutorial-like questions like “How do I create a Pod?”

‍

![Image 10](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a039d0a46841c35338b848_SuBeHwTBNG6pmzxQrBlBPNA4tbpFK17w_DWF3D_xXJK4gcjuMnNhFd5uCuKyjUSdCuplS-JAlksGeGcjil9V94LBJn0ZY2sNJ0RhwYC0UYvEIoHJkBK9c0IygnpLse9jwQmVuj8S6t6dwIvUfrmCtmE.png)

#### ChatGPT Recommendations

The Doctor plugin integrates with Botkube's existing Kubernetes recommendations, leveraging insights derived from cluster logs to offer personalized suggestions tailored to your Kubernetes environment. Whether it's optimizing resource allocation, enhancing security, or boosting overall performance, the Doctor plugin equips users with valuable insights to make informed decisions for their Kubernetes clusters.

‍

![Image 11](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a0710b0d200385a983c1f3_I196XuDKfbF_0Tq-brChXFXESPbAuAMK7UVl7oWnbS5TY0DEDg5mK4dOqO5WDIuTdcr9kFtivAoOezeqDYb6AkCWFCwW37jskmlzHgz6FDtz0AQwH9TWUUTCEUOE_yw_2zcBkYqYymnHml4Si8X0Rfw.png)

Kubernetes Troubleshooting Automations
--------------------------------------

Botkube users can use the Doctor plugin to harness ChatGPT's capabilities to effortlessly create custom Botkube configurations and automated actions, seamlessly integrating with other plugins. With Slack interactivity as a command builder, users can interact efficiently with the system, receiving tailored suggestions for automations that address specific challenges within their Kubernetes clusters. The plugin empowers users to pinpoint precise solutions and execute commands with ease, streamlining the troubleshooting process and enhancing overall cluster management.

This integration is particularly useful for developers and Ops teams seeking to navigate the complexities of Kubernetes, offering real-time assistance and guidance without the need for extensive Google searches or sifting through Kubernetes documentation for every encountered error.

Executor Plugin
---------------

![Image 12](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a039d14c51c24d23107d67_be3NRGddT9iSoKYyGzqKfm0vAygUfhC7Tjt1sqpMpSKxIALefMbB_GeBXrxs4-Xy3vlMRz7Dq7vJtts06n1UhgydkWVbJs0WVBoEtZwu_UqM26jLT7u4Dxbh1993ioMwhORJwnbh0F_pcYmpJBm1Mgg.gif)

Last but not least, Botkube's [Executor plugin](https://docs.botkube.io/usage/executor/) plays a pivotal role in allowing users to explore the flexibility within the plugin system. With the Exec plugin, you can seamlessly install and run any command-line interface (CLI) directly from your chosen communication platform through Botkube. It empowers you to create your personalized interactive interface for these CLIs.

While the Executor plugin is not designed to replace our dedicated executor plugins, it serves as an excellent tool for exploring Botkube's [plugin system](https://botkube.io/integrations) by being able to quickly test new CLIs within Botkube, or in controlled lab environments. What sets it apart is its support for interactivity, featuring buttons and dropdowns that enable you to effortlessly utilize your preferred tools without the need for extensive typing, even when using your mobile device.

Conclusion
----------

In conclusion, Botkube offers a variety of powerful [features](https://botkube.io/features) that cater to both beginners and seasoned Kubernetes users and can enhance your Kubernetes troubleshooting experience. From real-time monitoring with Prometheus for immediate insights into your cluster's health to a seamless ArgoCD integration that helps simplify deployments and [GitOps](https://botkube.io/blog/enhancing-gitops-workflows-with-botkube) workflows, Botkube empowers DevOps teams to work more efficiently and enables developers to troubleshoot their applications without Kubernetes expertise. Dive into these cool Botkube features and elevate your Kubernetes troubleshooting experience.

### Getting Started with Botkube Today

[Try it for yourself](https://app.botkube.io/). Follow our step-by-step [tutorial](https://botkube.io/blog/maximize-your-devops-teams-efficiency-with-botkube-and-microsoft-teams) to set up Botkube using our web app. We're excited to hear how you use Botkube. Create a support ticket directly in the dashboard, share your stories with us in the [Botkube Slack community](https://join.botkube.io/). We’ll even send you some cool swag as a thank you.
