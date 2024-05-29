Title: Tutorial: Streamlining GitOps with the Botkube Flux Plugin

URL Source: https://botkube.io/blog/streamlining-gitops-with-the-botkube-flux-plugin

Published Time: Sep 01, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Learn how to set up Botkube with Flux and GitHub-events plugins in this step-by-step tutorial. Configure permissions, insert tokens, and install seamlessly.

### Table of Contents

*   [Prerequisites](#prerequisites)
*   [Using the Flux Plugin in Action](#using-the-flux-plugin-in-action)
*   [Conclusion](#conclusion)
*   [Sign up now](#sign-up-now)
*   [Feedback](#feedback)

#### Start Using Botkube AI Assistant Today

#### Get started with Botkube Cloud

In today's fast-paced world of Kubernetes management, embracing GitOps is a crucial step towards efficient collaboration and automation. However, it comes with its own set of challenges. Enter [Botkube](https://botkube.io/), a Kubernetes collaborative troubleshooting tool for Kubernetes that seamlessly integrates with popular collaboration platforms like [Slack, Microsoft Teams, Discord, and Mattermost](https://botkube.io/integrations). Botkube not only simplifies alert management but also optimizes GitOps workflows by enhancing automation, real-time collaboration and centralizing knowledge. Learn more about Botkube’s move towards [GitOps](https://botkube.io/blog/enhancing-gitops-workflows-with-botkube) and the new Flux Plugin in our [release announcemen](https://botkube.io/blog/introducing-botkubes-integration-with-flux)t. In this tutorial, we will explore the Botkube Flux plugin, a powerful tool that automates notifications, enhances GitOps practices, and simplifies the interaction between Kubernetes clusters and GitHub repositories. By the end of this tutorial, you'll have a solid understanding of how to create, configure, and leverage the Botkube Flux plugin to streamline your GitOps workflow.

Prerequisites
-------------

Before we dive into the tutorial, make sure you have the following prerequisites ready:

*   A GitHub account with access to tokens
*   Basic familiarity with Kubernetes and its components.
*   Access to a Kubernetes cluster with helm installed
*   Access to Slack workspace
*   A Botkube Cloud Account

### Creating Botkube Cloud Account

1.  On the Botkube [homepage](https://botkube.io/), locate the “Get Started” button and click on it. This will take you to the account registration page.
2.  Fill in the required information in the registration form. You can sign up with your email address or Github account.

*   Click [here](https://botkube.io/blog/step-by-step-tutorial-leveraging-botkubes-cloud-slack-feature-for-kubernetes-collaborative-troubleshooting) for a more indepth Botkube installation tutorial

![Image 2: Easy one click Kubernetes deployment for cloud](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a709dfd8e90bc79339fcd0_cQKP0DfzGkbQO4R8kCAnqr54pgSa_IKaPa756N-FFua5n9N1omSH9fg9nGI1JYNjRS6ZmkbNUYrZLK1Z2BmTjPVHBDP0U9jNpidqq7RIqKWJScUJ32pOPryOAp49HR6OoerKN7yJSu6yHr2DU1GDaoo.png)

### Connecting Kubernetes Cluster to Slack

Select either the Wizard or the Block builder method to initiate the setup process for Botkube's Cloud Slack integration.

1.  Next, enter your instance display name.

![Image 3: K8s cluster building GUI](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a709efd8e90bc7933a1393_nAeC7-04jk70WellyEP2GM4m75jP4jrLhnmbjAkZr3rLlNi7zaD2bMLx8rvebpfqFIrvB8OSIxIqKezCZngk7ooCH6WAOT_1PBSQKz-sAAl9WRSq-GqtR1gmHmwC87Oq443Bzdu_sMKsHw-_g8Jwrfo.png)

2.  Select the Official Botkube Slack option. (Note this option requires you to sign up for a 30 day free trial)

![Image 4: Official Slack Kubernetes bot](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a709ffd8e90bc7933a2249_3GYyjQn-Uklnp1Bn8T7YmSOdKEaFnl3idDQcYJiD1mx7xeBbr6yvoRgbLI3Fir7TaW4a1N8l4tgB_Zbt6b3XryqzyYff4z1I_nffpWkoS6Hx7yPmmTrk2Z9tnAlXWUoM_VrAm0iBje2a8oaIiaGxRx0.png)

3.  Access the app settings within Slack to configure the integration with Botkube. Click the "Add Slack" button to initiate the process.
4.  Select the workspace within Slack where you want to utilize Botkube. You will now have access to public and private channels, enabling seamless communication and collaboration for Kubernetes troubleshooting tasks.

![Image 5: Multicluster setup screen for Slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a0bae43806c67551203_v-0W_ZDNIBT2Z7lbvKemYUyidm6L4ftXfEXxY9t0i5d6NB3_A_wrkVIluEVKfh8ZCCSYan2mS8PfS0YXm8DmViUyII5FXmmaLUPy6deAhqmYypJr0mZCg8aOo1FckVZaX3LOlTK6Epso_FqKUAde3Qw.png)

5.  To include Botkube in private channels, utilize the "Add to Channel" feature. This ensures that Botkube is present in relevant channels where Kubernetes-related discussions and incident responses take place.

![Image 6: One click slack sign in](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a1d00209556840fc2aa_qCmpnXKLE-S-5GKx1PijNsYeJOqKsWffvD0NIp708myAL6SynM44bx0khhpKpQCX-LnUoIQ2t5JAbqjdOfxrQSxIJPZWLRKYrX0O1lKJJQQj0hmkIM_5ADswoPXLaRPrMmLwAtVCSAsGEbySEsGW0WY.png)

6.  Select the Plugins you would like to use.

### Enabling the Plugin

![Image 7](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f219c6901138bd3fb3ecbb__JgXw9_EzpdpHItfymUJr5owqit0mMnoORy7I1PApaua66YV3GKd00wMc4TYYSpNpW-uhYDgGc0FHywymRF2Zf4tGny-TIORkO8LdDfmwTFyuOX5gDR4-najpP4gOHJs7XZ0TgmsOm_kCGFu08ZuRWw.png)

1.  Select the Flux and GitHub-events Plugin
    
    *   Begin by selecting the Flux and GitHub-events plugins within your Botkube setup.
2.  Insert Your GitHub Access Token
    
    *   To do this, navigate to your GitHub developer settings.
        
    *   Locate your GitHub access token and copy it.
        

*   Insert the copied access token into the appropriate field within the Flux plugin settings.

![Image 8](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f218c651c3be1367d23eb4_RxmxZInIIFLvZUTkgh63JlGBKnJlN1oGd5RW1NGQ6dM0ZPgnXaTXJ48wRFBDJWzUvNlIo5H9PuOEBMbB_rCh15bW_VK2f3uXPwOl_y2P7QLtzLNs2aIZBukjuu2WQGfkJD3LznRWAQy-CropKTSXxbc.png)

3.  Configure Permissions

*   In the Flux plugin settings, select the "Permissions" tab.
*   Choose the "Custom" option.
*   Under the custom permissions, select the "Group" option
*   Next, update the permissions for RBAC (Role-Based Access Control).
*   Add "flux-write" and "flux-read-patch" permissions since they are required for using the "flux install" command.

![Image 9](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f218e2755212a14074243d_bvJGmOwBNt4TLjh_yde1mlTZ_X-b0Mk-pXbKoJP3iLSNfw3ilvM_21hkaj0M5qAP8W04glEk1aWetaklzfuvZGGNBmmQefuFjt-4LWGz-WF2XTIsFnIcNliGTAL_izF_RQfj-6OjcUihJYdk66t2jkM.png)

4.  Configure GitHub-events Plugin
    
    *   Now, select the GitHub-events plugin
5.  Insert GitHub Access Token
    
    *   Insert the same GitHub access token that you used for the Flux plugin into the GitHub-events plugin settings.

![Image 10](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f21921f9104d28d7620904_a-5ZlAVnumBSl5orUqAAU2Ma0aWx3xrJ7QjVRWZU-IKvraswRku8RkYh4Ax6QRo_W0lvDHnhGTCRKQm0dF3U-kyVq_1c6THoimLnlRy0uw8oZ0mahQbCxZKINNbdCKTVbNfoSr0KgqF-w9A8YG5fGe0.png)

‍

6.  Repository Configuration

*   Scroll down to the "Repository Configuration" section.
*   Insert the name of your repository into the appropriate field.
*   Scroll further down to the "Pull Request Matcher" section.
*   Select the types of pull requests you want Botkube to check for.
*   For this demo, choose the "open" option.

![Image 11](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f2193d26a35a2059f5a9c2_zV2W8j7eAJbfFdR_eFsbWZQdpJmZY0uXib1XCLCk9HR3n8XuQuCnnmy8T0Lk_eSWv6ZVPy2iTd_KHrhLUvGlSI0kN6OVtwI4K9dLHYxUW_wR3piuJYV0KDB4Hsh9hFOBW-SrYC0bRe2GPg_0bxGLekU.png)

*   In the same section, locate the "File Pattern" field.
*   Type in the file path that you want pull requests to match.
*   For this demo, use your "kustomize" folder.

![Image 12](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f2194b9e5c70cd869dbfb4_X_x4cE-nGwjogL8wBWDcCQOeuxeHW-JzGMp6YmgOJEU1rK9z-79prBw2HNnmkXnmJCBkraY_qsRPwBGuxXfzWFqY1PtocW51l3qfly2Gsu0ToaPZb62pcIR_O-hlNBKu5T_tN7cwrgixqPNxAkDObQc.png)

7\. Click the "Submit" button to save your configurations.

    * If desired, configure any additional plugins you want to use at this stage.
    

![Image 13](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f2195fe9284fa6e17e51d1_9T2t8y3gIhVQPV4qHiXUcKwdYB_vYQ2wV06O1P7wM1zxib3rJlv8h2oTNW25bNehjGHDtQ2qYWwRy_96GYdI-bvl-0yEIKUjmfTqrARjZU4GcgVySZCldUnKHrlA5JjVwEXkxAFnyHu3SFQP4NMbok8.png)

8.  Click "Next" to proceed to the installation step.
    
9.  On the next page, you will have the option to also enable [Command Alias](https://botkube.io/blog/command-line-magic-simplify-your-life-with-custom-kubernetes-kubectrl-aliases-on-botkube) and [Botkube Actions](https://docs.botkube.io/usage/automated-actions/).
    
    *   Make your selection and click the create button.
10.  You are now ready to start playing with your Botkube plugin.
    

Using the Flux Plugin in Action
-------------------------------

### Running a Flux install

![Image 14](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64ecb79001bc7e01e2c88804_flux-interactivity.gif)

With the Botkube Flux plugin, you have the power to streamline your Kubernetes deployment workflow even further. Beyond the ability to effortlessly manage GitHub pull requests, this plugin offers you the convenience of checking prerequisites on your cluster and performing a complete Flux installation – all from your preferred communication platform. This means that not only can you seamlessly collaborate with your team members by sharing and discussing pull request changes within your Slack channel, but you can also ensure that your cluster is properly configured before initiating the installation of Flux. This capability to validate prerequisites directly from your communication platform adds a layer of control and convenience to your Kubernetes operations.

### Automating a Flux Diff Workflow

![Image 15](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64e694bca6bd600e8a7e88dd_flux-diff-1.gif)

This command streamlines several essential tasks in your workflow. It starts by automatically identifying the relevant GitHub repository linked to the provided kustomization. After that, it proceeds to clone this repository and check out a specified pull request. The command then performs a thorough comparison between the changes in the pull request and the current state of your cluster. What sets it apart is its ability to conveniently share the resulting diff report on your Slack channel, facilitating seamless collaboration among team members for reviewing and discussing the alterations. Moreover, the command enhances efficiency by offering additional contextual actions, including posting the diff report as a GitHub comment on the corresponding pull request, approving the pull request, and providing quick access to view the pull request itself. Combining Botkube's GitHub Slack integration for repo control and Flux CD integration for more tools allows for better GitOps workflow management.

Conclusion
----------

In conclusion, the Botkube Flux plugin presents an invaluable asset for simplifying GitOps workflows. From automating notifications to enhancing collaboration, this tool streamlines the process of managing Kubernetes clusters and GitHub repositories. We encourage you to implement the plugin in your workflow and explore further automation possibilities.With seamless integration into major communication channels, Botkube empowers you to take swift action on identified errors, ensuring your Kubernetes environment runs smoothly wherever you are.

Sign up now
-----------

Get started with [Botkube](https://app.botkube.io/). Whether you’re a seasoned Kubernetes pro or just getting started, Botkube has something to offer. Sign up now for free and join the [community of users](https://join.botkube.io/) who are already benefiting from the power of Botkube.

Feedback
--------

We welcome developers and Kubernetes enthusiasts to explore the platform and share their valuable feedback. We want to know what you think of Botkube and how we can make it even better. We're doing quick 15-minute interviews to get your feedback, and as a thank you, we'll give you some cool Botkube plushies and t-shirts and enter you into a raffle for a chance to win a $50 Amazon gift card! Just email our Developer Advocate, Maria or use this calendly [link](https://calendly.com/maria-botkube) to sign up.

You can also talk to us in the Botkube GitHub issues, connect with others and get help in the Botkube [Slack community](https://join.botkube.io/), or email our Product Leader at blair@kubeshop.io.
