Title: Tutorial: Infuse collaboration into your Argo workflows with Botkube

URL Source: https://botkube.io/blog/getting-started-with-botkube-and-argocd

Published Time: Sep 29, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Learn how to set up Botkube with ArgoCD plugins in this step-by-step tutorial.

### Table of Contents

*   [Prerequisites](#prerequisites)
*   [Install the ArgoCD Plugin into a new instance](#install-the-argocd-plugin-into-a-new-instance-)
*   [Using the ArgoCD Plugin in Action](#using-the-argocd-plugin-in-action)
*   [Conclusion](#conclusion)
*   [Sign up now!](#sign-up-now-)
*   [Feedback](#feedback)

#### Start Using Botkube AI Assistant Today!

#### Start Receiving Kubernetes Notifications Now with Botkube Cloud

Navigating through the complexity of scaling operations and collaborative workflows with GitOps tools like ArgoCD can pose significant challenges. As teams scale, real-time updates and troubleshooting efficiency become paramount.

Enter [Botkube](http://botkube.io/), a Kubernetes collaborative troubleshooting solution designed to seamlessly integrate with widely-used collaboration platforms such as [Slack, Microsoft Teams, Discord, and Mattermost](http://botkube.io/integrations). Botkube doesn't just simplify Kubernetes monitoring; it also optimizes GitOps workflows through enhanced automation, real-time collaboration, and centralized knowledge management.

Learn more about Botkube’s move towards [GitOps](https://botkube.io/blog/enhancing-gitops-workflows-with-botkube) and the new ArgoCD Plugin– the second installment in the Botkube GitOps plugin series in our [release announcement](https://botkube.io/blog/argo-cd-botkube-integration).

In this tutorial, we'll delve into the capabilities of the [Botkube ArgoCD plugin](https://docs.botkube.io/configuration/source/argocd/). This powerful tool automates notifications, and enables bi-directional action. By the end of this tutorial, you will know how to create, configure, and effectively leverage the Botkube ArgoCD plugin to enhance and simplify your GitOps workflow.

Prerequisites
-------------

Before you begin, make sure you have the following prerequisites in place:

*   Basic familiarity with Kubernetes and its components
*   Access to a Kubernetes cluster with Helm installed
*   Access to Slack workspace
*   A [Botkube Cloud Account](http://app.botkube.io/)
*   [Argo CD](https://argoproj.github.io/cd/) must be installed on your Kubernetes cluster
*   [Install the Argo CD CLI](https://argo-cd.readthedocs.io/en/stable/?_gl=1*10c1kh8*_ga*NDc0Nzg3NTU3LjE2OTU2NTg1MzI.*_ga_5Z1VTPDL73*MTY5NTkxNTMyMC4yLjEuMTY5NTkxNTM0NC4wLjAuMA..#getting-started) on your local machine

*   Set up port forwarding to securely access your Argo CD instance, avoiding exposure over the internet

*   Ensure that you have at least one application configured within ArgoCD

Install the ArgoCD Plugin into a new instance
---------------------------------------------

### Creating Botkube Cloud Account

1.  Sign in to Botkube Cloud, either as a new user with a free account with an email or Github account or as an existing user.

\* Click [here](https://botkube.io/blog/step-by-step-tutorial-leveraging-botkubes-cloud-slack-feature-for-kubernetes-collaborative-troubleshooting) for a more indepth Botkube installation tutorial

### Connecting Kubernetes Cluster to Slack

1.  Next, enter your instance display name.

2.  Select the Official Botkube Slack option. (Note this option requires you to sign up for a 30 day free trial)

3.  Access the app settings within Slack to configure the integration with Botkube. Click the "Add Slack" button to initiate the process.

4.  Select the workspace within Slack where you want to utilize Botkube. You will now have access to public and private channels, enabling seamless communication and collaboration for Kubernetes troubleshooting tasks.

5.  To include Botkube in private channels, utilize the "Add to Channel" feature. This ensures that Botkube is present in relevant channels where Kubernetes-related discussions and incident responses take place.

6.  Select the Plugins you would like to use.

### Enabling the Argo CD Source Plugin

![Image 2: Botkube Setup Wizard for easy K8s tool deployment](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6516ed34d32d63db102a63d7_OKVRz-x5BxiFYrNR8MaOf7fpyY77Rel8xxgioY5wrTA1HSrT_B9mUrZXxY_BpJ2p71X-Ovr6eN1tHhoVryzsABM3sj8GmHJkY84sQu72IMwDrZieUtJDMvLcYKoLml5oggeDqgtsie5TboIxEDntW2M.png)

1.  Select the ArgoCD and kubectl Plugin

2\. Begin by selecting the ArgoCD and kubectl plugins within your Botkube setup

![Image 3: Adding Argo CD API Keys](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6516ed347d480ce1978008ad_ISJShwJ0wGBlQUqSnpW1Zt-9vzEKnLzwFnLMlxGIX0WI25KZ7tWnyapg0LSixCOslONWZfHErb-qmr_MvlqlWTDK3VxGnPKsZfDCnNKGJNAy90orbvT3HHXdkgm-D3JeArzT4pea8mUOExvfS7QY0rY.png)

      3. Configure the plugin with your Botkube instance
    
      4. Insert your resources from your ArgoCD UI
    

*   Fill in your “Name” and “Namespace”
*   Make sure your BaseURl matches the one found on your ArgoCD UI

![Image 4: Setting Argo CD permissions on deployment](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6516ed337ad79fb0c8bb0436_zRBc4WDwcmJW7sKZjJaItEuxSAKmzxrqxw3C-QhoAPTf7Br_i67Eyk5XN6jYPowsCQ836d4ogBZ3Lh6rC42cbw1Ato5chhbOP9UOxTy6hQy_F0prcvSRmD7IBZtfCFMoKqcjlnUYMYCC9SVqJYl6NJ4.png)

4.  Configure Permissions

*   In the Argo plugin settings, select the "Permissions" tab.
*   Choose the "Custom" option.
*   Under the custom permissions, select the "Group" option
*   Next, update the permissions for RBAC (Role-Based Access Control).
*   Add "argocd”

![Image 5: RBAC controls for Argo](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6516ed3352cdb1553efc2f21_7bqFl-gQuFLiZqajb4AXF0r0BOJ-_D0SseHOHPGjUQ6DKi6M_YWpc2qNljNslsn7UMMDmAULu_cbURDPd6ilRAbbtKE3sQHURZPpcGMMwzgEIuXq1dm0m0R1LTVkQirBVpOsBWF-ooL1EBd1bb2hogs.png)

5.  Click "Next" to proceed to the installation step.
6.  On the next page, you will have the option to also enable [Command Alias](https://botkube.io/blog/command-line-magic-simplify-your-life-with-custom-kubernetes-kubectrl-aliases-on-botkube) and [Botkube Actions](https://docs.botkube.io/usage/automated-actions/).

7\. Make your selection and click the create button.

      8. You are now ready to start playing with your Botkube plugin.
    

Using the ArgoCD Plugin in Action
---------------------------------

### Checking the Health Status of Apps

![Image 6: ArgoCD events](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6509a59c63441b36226ea80d_argocd-events-e6eabb1f581e9822020d55461539bfcd.png)

‍

Use Botkube’s ArgoCD plugin for seamless health status checks on your ArgoCD apps. Dive into the world of troubleshooting with just a couple of commands. Utilize "kubectl describe" and "kubectl get" within Botkube to verify the optimal functioning of the Argo CD source plugin. Leverage slack interactivity and automation to make health status checks a breeze.

Conclusion
----------

In summary, the [ArgoCD Plugin by Botkube](https://botkube.io/integration/argo-cd-botkube-kubernetes-integration) stands as an indispensable resource for streamlining GitOps workflows. It excels in automating notifications, fostering collaboration, and bi-directional control. This makes Gitops workflows more efficient and easier to scale.

Botkube significantly simplifies the ArgoCD troubleshooting process by offering a streamlined and user-friendly alternative to the complex manual process. Once enabled, Botkube takes charge of configuring notifications, utilizing its incoming webhook to effortlessly receive and forward events. This eliminates the need for intricate procedures like generating ArgoCD webhook secrets and manually setting up triggers for events. Botkube goes beyond simple alerting by allowing users to directly engage with ArgoCD events, performing actions like running commands on applications, viewing applications in the ArgoCD UI, or opening the source repository in a web browser—all from your communication platform of choice. This streamlined and interactive approach to troubleshooting ArgoCD takes your team’s efficiency to the next level.

Sign up now!
------------

Get started with Botkube! Whether you're a seasoned Kubernetes pro or just getting started, Botkube can help supercharge your troubleshooting process. [Sign up now for free](http://app.botkube.io/) and [join the community of users on Slack](https://join.botkube.io/) who are already benefiting from the power of Botkube.

Feedback
--------

We welcome developers and Kubernetes enthusiasts to explore the platform and share their valuable feedback. We want to know what you think of Botkube and how we can make it even better. [Reach out to our Developer Advocate, Maria](mailto:maria@kubeshop.io) or schedule a quick 15min meeting at your preferred time. As a thank you, we’ll send you some great Botkube swag.
