Title: Troubleshoot K8s with ChatGPT + Botkube's Doctor Plugin

URL Source: https://botkube.io/blog/doctor-plug-in-tutorial

Published Time: Jul 31, 2023

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

This Tutorial is about the new ChatGPT-powered Doctor plugin which allows for easy diagnosis and solution recommendations for Kubernetes troubleshooting.

### Table of Contents

*   [Requirements](https://botkube.io/blog/doctor-plug-in-tutorial#requirements)
*   [Method 1: Install the Doctor Plug into a new instance](https://botkube.io/blog/doctor-plug-in-tutorial#method-1-install-the-doctor-plug-into-a-new-instance)
*   [Method 2: Install the Doctor plug-in into an already created Botkubeinstance.](https://botkube.io/blog/doctor-plug-in-tutorial#method-2-install-the-doctor-plug-in-into-an-already-created-botkubeinstance-)
*   [Interactive Troubleshooting](https://botkube.io/blog/doctor-plug-in-tutorial#interactive-troubleshooting)
*   [Support for Generic Questions](https://botkube.io/blog/doctor-plug-in-tutorial#support-for-generic-questions)
*   [Scalable Plugin Pipelining](https://botkube.io/blog/doctor-plug-in-tutorial#scalable-plugin-pipelining)
*   [Conclusion](https://botkube.io/blog/doctor-plug-in-tutorial#conclusion)
*   [We'd Love Your Feedback](https://botkube.io/blog/doctor-plug-in-tutorial#we-d-love-your-feedback)

#### Start Using Botkube AI-Powered Assistant Today

Discover the future of efficient Kubernetes troubleshooting with Botkube’s new [Doctor plug-in](https://botkube.io/integration/chatgpt-botkube-kubernetes-integration)! By leveraging ChatGPT's powerful capabilities, you can now troubleshoot your Kubernetes cluster with unparalleled ease and save valuable time in the process. Say goodbye to tedious and time-consuming tasks, as ChatGPT helps to automate much of the troubleshooting process, allowing you to concentrate on more critical aspects of your workflow. Moreover, ChatGPT's comprehensive analysis of your cluster configuration and resource utilization offers invaluable insights to optimize performance and cut down on costs. Real-time errors and recommendations enhance your troubleshooting experience, ensuring you're always up-to-date with the health of your cluster. Seamlessly integrated with major communication channels, Botkube empowers you to act swiftly on identified errors, keeping your Kubernetes environment running smoothly at all times.

Learn more about Botkube's ChatGPT-powered Doctor plugin in our [release announcement](https://botkube.io/blog/use-chatgpt-to-troubleshoot-kubernetes-errors-with-botkubes-doctor).

In this tutorial, we will walk you through the step-by-step process of setting up and utilizing Botkube's new [Doctor plug-in](https://docs.botkube.io/usage/executor/doctor), optimizing your team’s performance, and maximizing your Kubernetes productivity.

Requirements
------------

*   A Slack workspace
*   A [Botkube Cloud](http://app.botkube.io/) account
*   \[OpenAI credits \[(https://openai.com/)

Method 1: Install the Doctor Plug into a new instance
-----------------------------------------------------

### Creating a Botkube Cloud Account

1.  On the Botkube [homepage](https://botkube.io/), locate the “Get Started” button and click on it. This will take you to the account registration page.
2.  Fill in the required information in the registration form. You can sign up with your email address or Github account.

*   Click [here](https://botkube.io/blog/step-by-step-tutorial-leveraging-botkubes-cloud-slack-feature-for-kubernetes-collaborative-troubleshooting) for a more indepth Botkube installation tutorial

### Connecting Kubernetes Cluster to Slack

1.  Select either the Wizard or the Block builder method to initiate the setup process for Botkube's Cloud Slack integration.

![Image 2: Easy one click Kubernetes deployment for cloud](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64a709dfd8e90bc79339fcd0_cQKP0DfzGkbQO4R8kCAnqr54pgSa_IKaPa756N-FFua5n9N1omSH9fg9nGI1JYNjRS6ZmkbNUYrZLK1Z2BmTjPVHBDP0U9jNpidqq7RIqKWJScUJ32pOPryOAp49HR6OoerKN7yJSu6yHr2DU1GDaoo.png)

‍

2.  Next, enter your instance display name.

![Image 3: K8s cluster building GUI](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64a709efd8e90bc7933a1393_nAeC7-04jk70WellyEP2GM4m75jP4jrLhnmbjAkZr3rLlNi7zaD2bMLx8rvebpfqFIrvB8OSIxIqKezCZngk7ooCH6WAOT_1PBSQKz-sAAl9WRSq-GqtR1gmHmwC87Oq443Bzdu_sMKsHw-_g8Jwrfo.png)

‍

3.  Select the _Official Botkube Slack_ option. (Note this option requires you to sign up for a 30 day free trial)

![Image 4: Official Slack Kubernetes bot](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64a709ffd8e90bc7933a2249_3GYyjQn-Uklnp1Bn8T7YmSOdKEaFnl3idDQcYJiD1mx7xeBbr6yvoRgbLI3Fir7TaW4a1N8l4tgB_Zbt6b3XryqzyYff4z1I_nffpWkoS6Hx7yPmmTrk2Z9tnAlXWUoM_VrAm0iBje2a8oaIiaGxRx0.png)

‍

4.  Access the app settings within Slack to configure the integration with Botkube. Click the "Add Slack" button to initiate the process.
    
5.  Select the workspace within Slack where you want to utilize Botkube. Botkube will now have access to public and private channels, enabling seamless communication and collaboration for Kubernetes troubleshooting tasks.
    

![Image 5: Multicluster setup screen for Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64a70a0bae43806c67551203_v-0W_ZDNIBT2Z7lbvKemYUyidm6L4ftXfEXxY9t0i5d6NB3_A_wrkVIluEVKfh8ZCCSYan2mS8PfS0YXm8DmViUyII5FXmmaLUPy6deAhqmYypJr0mZCg8aOo1FckVZaX3LOlTK6Epso_FqKUAde3Qw.png)

‍

6.  To include Botkube in private channels, utilize the "Add to Channel" feature. This ensures that Botkube is present in relevant channels where Kubernetes-related discussions and incident responses take place.

![Image 6: One click slack sign in](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64a70a1d00209556840fc2aa_qCmpnXKLE-S-5GKx1PijNsYeJOqKsWffvD0NIp708myAL6SynM44bx0khhpKpQCX-LnUoIQ2t5JAbqjdOfxrQSxIJPZWLRKYrX0O1lKJJQQj0hmkIM_5ADswoPXLaRPrMmLwAtVCSAsGEbySEsGW0WY.png)

‍

7.  Select the Plugins you would like to use. [Helm](https://docs.botkube.io/usage/executor/helm), [kubectl](https://docs.botkube.io/usage/executor/kubectl), and Kubernetes source are the Botkube default plug-ins. Select the [Doctor plug-in](https://docs.botkube.io/usage/executor/doctor).

![Image 7](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64c821b8724cd267a8f2079e_S_D9N3SAYRhJit2YvTCT1mfdMWr6Ba5CAX33ioUVeqdNvcycYIkd7yrIZ5Av6kG-hVuQXtTODdwDYtzoCWuV2NMH5F0waTITlkQXvBxglTo8aU_feKRZqpp9pZcssPIvxotbXdoKE49GVO9r9CreqyY.png)

‍

8.  Navigate to the OpenAI website [here](https://beta.openai.com/account/api-keys) and create your API key
9.  Paste your API key into the empty box and press submit.

![Image 8](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64c821bb26bb53b2ce43074d_I-Jz17f6Ag1ojKmWY0MT8l-IgRvU82Q2Gfwqk9OQVLXHF4Wx_0yQM-iy3GPnmIgQrDm1_Nck_-M2bpccmV2VY9svcoSjN7yylToOVl21vWoeJb6pp_bldS9zT677JvGJ0022oJd5dWlKN06Nph8UAmw.png)

10.  On the next page, you will have the option to also enable [Command Alias](https://botkube.io/blog/command-line-magic-simplify-your-life-with-custom-kubernetes-kubectrl-aliases-on-botkube)and [Botkube Actions](https://docs.botkube.io/usage/automated-actions/). Make your selection and click the Create button.

![Image 9](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64c821b95d087e15b6d1ec7d_zoOK3PqJeNz8R8Mgngk-Qssdni9rpMmx6-CMe1cGxvODqzVz-vb4cJ62ZbWeGfeoAudmnhkwKCzY8bb1UwAIoeefzAAFgwHfLKt4VNkU9kX_0Q8UJFhbWcqfLdJoLcmeNHAcluZPXeEA5I0FkkfQ4kA.png)

‍

11.  You are now ready to start utilizing the Botkube Doctor plugin.

Method 2: Install the Doctor plug-in into an already created Botkubeinstance.
-----------------------------------------------------------------------------

### Creating your New Plug-In

1.  Select the `Add Plug-in` button

![Image 10](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64c821b87c3808d8b8fb048d_fyMnDpgGrvfDOKs6zbiAj1SXGp5deepJ1_uWTLiYuXgcnv7tonGyf-NolV9U2d-IeH0nFMdnSpwX12xtV4355EoQ6kuWd2FmBEpjtpOrIc7UQQ5Mnc6YrB8kxajMO3yfbORKjh6Y6JYunOB5OLTfZAc.png)

‍

2.  Select the option for the Doctor plug-in

![Image 11](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64c821ba302266c000e029ec_FbpyK0RLgMWySuuwKBVfVBtsX-eDAoIuGqze913li_eZyDDjWoOpj7JmmeXKci6M7xzbGddBcqLz17gYRuDN0ISndRLHtc2gvkApV9GJIXUNfn2o7B-J8Zg-lpiNe9UUMKsffwaQ1A_-qImmfhAw1UY.png)

‍

3.  Navigate to the OpenAI website [here](https://beta.openai.com/account/api-keys) and create your API key
4.  Paste your API key into the empty field and press Submit.

![Image 12](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64c821bb26bb53b2ce43074d_I-Jz17f6Ag1ojKmWY0MT8l-IgRvU82Q2Gfwqk9OQVLXHF4Wx_0yQM-iy3GPnmIgQrDm1_Nck_-M2bpccmV2VY9svcoSjN7yylToOVl21vWoeJb6pp_bldS9zT677JvGJ0022oJd5dWlKN06Nph8UAmw.png)

‍

5.  Select the Save button and refresh your browser.
6.  Start interacting with your new Doctor Plug-in!

Interactive Troubleshooting
---------------------------

When the Doctor AI detects an error or anomaly in the Kubernetes environment, it takes proactive steps to assist users in resolving the issue. Upon detecting an error event, the Doctor AI generates a Slack message with a `Get Help` button. Clicking on this button triggers the Doctor AI to provide users with actionable solutions and next steps to tackle the problem. This interactive troubleshooting approach empowers users with the knowledge and confidence to resolve Kubernetes issues swiftly.

Support for Generic Questions
-----------------------------

Beyond its ability to tackle specific errors, the Doctor AI is also well-versed in answering general questions related to Kubernetes. Users can seek guidance on various Kubernetes concepts, best practices, or configurations directly within the Slack channel. This means no more interruptions to access a separate dashboard or documentation, making it easier to gain quick insights and expand Kubernetes expertise.

![Image 13](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/64c821bb21a0701673853991__UHn0zhtq5mu4YOnbuJN0vHABCKyUNCqyhsohVmShvTAPqmc8RJo4DQNsiN-QDU2vZ03D-W9N8_x-G9jSQpxGfeErfnxuzMt4YCE7BavU_uuXf0nTRhDDzp3NaIedpaXj0_8ehtKPQdYiynaCK6sunk.png)

Scalable Plugin Pipelining
--------------------------

The Doctor AI plugin's versatility extends beyond standalone usage. It can be easily combined with other [Botkube plug-ins](https://botkube.io/integrations), allowing users to create custom workflows tailored to their specific use cases. By pipelining the Doctor AI with other plugins, teams can orchestrate complex actions and integrations, further streamlining Kubernetes management processes.

Conclusion
----------

Botkube's new Doctor plug-in brings the future of efficient Kubernetes troubleshooting to your fingertips. With the power of ChatGPT, you can now effortlessly troubleshoot your Kubernetes cluster, saving precious time and streamlining your workflow. Bid farewell to tedious tasks, as ChatGPT automates much of the troubleshooting process, allowing you to focus on critical aspects of your work.

With seamless integration into [major communication channels](http://botkube.io/integrations), Botkube empowers you to take swift action on identified errors, ensuring your Kubernetes environment runs smoothly wherever you are.

### Sign up now

Whether you're a seasoned Kubernetes pro or just getting started, Botkube helps make troubleshooting Kubernetes easier and faster. [Sign up now for free](http://app.botkube.io/) and [join the community](https://join.botkube.io/) of users who are already benefiting from the power of Botkube.

We'd Love Your Feedback
-----------------------

We welcome developers and Kubernetes enthusiasts to explore the platform and share their valuable feedback. We want to know what you think of Botkube and how we can make it even better. We're doing quick 15-minute interviews to get your feedback, and as a thank you, we'll give away cool Botkube plushies and t-shirts. Just email [maria@kubeshop.io](mailto:maria@kubeshop.io) or [set up a meeting](https://calendly.com/maria-botkube/30min).

You can also talk to us in the Botkube GitHub [issues](https://github.com/kubeshop/botkube/issues), connect with others and get help in the [Botkube Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
