Title: Leveraging Our Cloud Slack for K8s Collaborative Troubleshooting

URL Source: https://botkube.io/blog/step-by-step-tutorial-leveraging-botkubes-cloud-slack-feature-for-kubernetes-collaborative-troubleshooting

Published Time: Jul 06, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

This step-by-step tutorial gets you started with Botkube Cloud Slack App, which enables collaborative troubleshooting by seamlessly integrating with your organization's Slack workspace.

### Table of Contents

*   [Creating Botkube Cloud Account](#creating-botkube-cloud-account-2)
*   [Connecting Kubernetes Cluster to Slack](#connecting-kubernetes-cluster-to-slack-2)
*   [Conclusion](#conclusion-2)
*   [Sign up now](#sign-up-now-2)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

![Image 2: Diagram of how Botkube connects K8s Clusters to Slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a709cc5761791c793a83cb_FXqJSS9KtaIOKugq8s7dLxGlnmtLjrwb6L7gurT9lYJdDRz12RFZoLngWUxSDtvrMKpEnhFCEKOroy2rvA9MJiSEZV4DUJwBa58Vl2JiUXfXJ6b3RrHK-sXsLbaqGbijRlnbXXLmuqKB6ckrNR36yFE.png)

‍  
Botkube is a powerful tool designed to streamline incident and event response in Kubernetes environments. With its \[Cloud\](https://app.botkube.io/) \[Slack integration\](https://botkube.io/integration/slack), Botkube enables collaborative troubleshooting by seamlessly integrating with your organization's Slack workspace. In this tutorial, we will walk you through the step-by-step process of setting up and utilizing Botkube's new Cloud Slack feature, empowering your team to effectively create a kubernetes connection to slack and manage multiple Kubernetes clusters to improve incident resolution time

This blog post assumes that we're starting a completely new Botkube installation, but will show the quickest way to get Kubernetes alerts into Slack.

Requirements:

\* A Slack workspace where you have permission to install apps and create channels  
\* A Kubernetes cluster where you have access to install Botkube  
\* Working kubectl and helm installation  
\* A \[Botkube Cloud\](https://app.botkube.io/) account

Here's a video walkthrough of the installation:  
‍  
<iframe width="800" height="450" src="https://www.youtube.com/embed/UptGoHnYh14" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

\## Creating Botkube Cloud Account

1\. On the Botkube \[homepage\](https://botkube.io/), locate the “Get Started” button and click on it. This will take you to the account registration page.  
2\. Fill in the required information in the registration form. You can sign up with your email address or Github account.  
\## Connecting Kubernetes Cluster to Slack  
3\. Select either the Wizard or the Block builder method to initiate the setup process for Botkube's Cloud Slack integration.

![Image 3: Easy one click Kubernetes deployment for cloud](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a709dfd8e90bc79339fcd0_cQKP0DfzGkbQO4R8kCAnqr54pgSa_IKaPa756N-FFua5n9N1omSH9fg9nGI1JYNjRS6ZmkbNUYrZLK1Z2BmTjPVHBDP0U9jNpidqq7RIqKWJScUJ32pOPryOAp49HR6OoerKN7yJSu6yHr2DU1GDaoo.png)

‍  
4\. Next, enter your instance display name.

![Image 4: K8s cluster building GUI](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a709efd8e90bc7933a1393_nAeC7-04jk70WellyEP2GM4m75jP4jrLhnmbjAkZr3rLlNi7zaD2bMLx8rvebpfqFIrvB8OSIxIqKezCZngk7ooCH6WAOT_1PBSQKz-sAAl9WRSq-GqtR1gmHmwC87Oq443Bzdu_sMKsHw-_g8Jwrfo.png)

5\. Select the Official Botkube Slack option. (Note this option requires you to sign up for a 30 day free trial)

![Image 5: Official Slack Kubernetes bot](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a709ffd8e90bc7933a2249_3GYyjQn-Uklnp1Bn8T7YmSOdKEaFnl3idDQcYJiD1mx7xeBbr6yvoRgbLI3Fir7TaW4a1N8l4tgB_Zbt6b3XryqzyYff4z1I_nffpWkoS6Hx7yPmmTrk2Z9tnAlXWUoM_VrAm0iBje2a8oaIiaGxRx0.png)

‍  
6\. Access the app settings within Slack to configure the integration with Botkube. Click the "Add Slack" button to initiate the process.  
7\. Select the workspace within Slack where you want to utilize Botkube. You will now have access to public and private channels, enabling seamless communication and collaboration for Kubernetes troubleshooting tasks.

![Image 6: Multicluster setup screen for Slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a0bae43806c67551203_v-0W_ZDNIBT2Z7lbvKemYUyidm6L4ftXfEXxY9t0i5d6NB3_A_wrkVIluEVKfh8ZCCSYan2mS8PfS0YXm8DmViUyII5FXmmaLUPy6deAhqmYypJr0mZCg8aOo1FckVZaX3LOlTK6Epso_FqKUAde3Qw.png)

‍  
8\. To include Botkube in private channels, utilize the "Add to Channel" feature. This ensures that Botkube is present in relevant channels where Kubernetes-related discussions and incident responses take place.

![Image 7: One click slack sign in](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a1d00209556840fc2aa_qCmpnXKLE-S-5GKx1PijNsYeJOqKsWffvD0NIp708myAL6SynM44bx0khhpKpQCX-LnUoIQ2t5JAbqjdOfxrQSxIJPZWLRKYrX0O1lKJJQQj0hmkIM_5ADswoPXLaRPrMmLwAtVCSAsGEbySEsGW0WY.png)

‍

9\. Select the Plugins you would like to use. Helm, kubectl, and kubernetes source are the botkube default plug-ins.

![Image 8: Add plugins to your Kubernetes Cluster quickly with Botkube](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a29a107604e69f10a44_EqV_jhVu5WsrFY2awVlpBZ5UGrulD-EtQrKoYnYoyZP-7TuapKozeQFXiLnQB3g0sUT8YdFZX_yYBgeaJUHhuXpYq3fUuaV9SyJgI0MAwYeJM3to-VfmwRuNyLOkBupW9r32e61df73T4HIa50KMVlc.png)

‍

10\. Click the create button  
11\. Copy and paste the helm commands into your terminal and then click next

![Image 9: Helm command to copy into command prompt ](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a33b3210007e280a92a_kqmlF1iCEr1KDQYsMNCGC83a_ZPx0agVAUb6crdZHQOONeg4BlQbwXSWferYj26OZkxyG2cRZ7jtLoDQtbUdEQ9eriQ-KmQmeBcLGxc7QQTtraL3VOAUQW0rNCWGNjJj5HBdzIv8lbk6HgjLwIJwTNM.png)

‍  
12\. You should see your Botkube instance being created.

![Image 10: Slack Kubernetes integration setup complete](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a3e978de29171f3a547_pjoT_YHQoMpAqVlVcPpsEl0oupPk1cJLMMLXHbehvipxrb0tni3hVtXkLE52YZMypptKk1Uozszf0pPUCN_SpzzP4W49mZy7NwJfLYWEGMBpwjHwIvIvD--mO22yCj9kV3wE4T8jIA532dDf2oUzVY0.png)

‍

\### Let’s see it in action

![Image 11: ChatOps for Kubernetes in Slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64a70a5bb3210007e280d3f8_Untitled%20design.gif)

The new Botkube Cloud Slack app adds some great features that weren't previously available, particularly around multi-cluster management. With new \*\*@botkube cloud\*\* commands you can easily navigate between instances.  By using \*\*"@botkube cloud list instances\*\*," you can conveniently view your connected deployments. You can also establish a default instance for a specific workspace and easily switch between them using our slack interactivity feature. Increase your team’s efficiency with the ability to run commands across all your connected clusters by simply adding " \*\*--all-clusters.\*\*" This feature saves you valuable time and lets you manage all your resources with a single command.  
\## Conclusion

Botkube's \[Cloud Slack integration\](https://botkube.io/integration/slack) offers a range of key features and functionalities. From seamless\[ incident response\](https://docs.botkube.io/usage/automated-actions) to enhanced \[resource monitoring\](https://docs.botkube.io/usage/source/prometheus/), Botkube empowers teams to troubleshoot and manage Kubernetes clusters effectively. Utilizing Botkube's Cloud Slack feature is crucial for teams seeking to optimize Kubernetes workflows and streamline troubleshooting processes. By leveraging Botkube, teams can improve incident resolution time, collaborate efficiently, and enhance overall operational efficiency in Kubernetes environments.

‍  
Integrate Botkube's Cloud with Slack to enhance your Kubernetes operations and improve incident response and collaboration. By utilizing Botkube's powerful features, you can streamline troubleshooting processes and achieve quicker resolutions for any Kubernetes-related issues.  
\## Sign up now

Get started with \[Botkube\](https://app.botkube.io/)! Whether you're a seasoned Kubernetes pro or just getting started, Botkube has something to offer. Sign up now for free and join the community of users who are already benefiting from the power of Botkube.  
\### Feedback

We welcome developers and Kubernetes enthusiasts to explore the platform and share their valuable feedback. We want to know what you think of Botkube and how we can make it even better. We're doing quick 15-minute interviews to get your feedback, and as a thank you, we'll give you some cool Botkube plushies and t-shirts and enter you into a raffle for a chance to win a $50 Amazon gift card! Just email maria@kubeshop.io or use this calendly link to sign up.

You can also talk to us in the Botkube GitHub \[issues\](https://github.com/kubeshop/botkube/issues), connect with others and get help in the Botkube Slack community, or email our Product Leader at \[blair@kubeshop.io\](mailto:blair@kubeshop.io).
