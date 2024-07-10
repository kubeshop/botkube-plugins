Title: Integrating Microsoft Teams with Azure for K8s Deployments

URL Source: https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments

Published Time: Oct 30, 2023

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Botkube is entering the MS teams world

### Table of Contents

*   [Benefits of Using Botkube and MS Teams](https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments#benefits-of-using-botkube-and-ms-teams)
*   [Tutorial Guide](https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments#tutorial-guide-)
*   [Prerequisites](https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments#prerequisites)
*   [Creating Botkube Cloud Account](https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments#creating-botkube-cloud-account)
*   [Dashboard setup](https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments#dashboard-setup)
*   [Conclusion](https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments#conclusion)
*   [Get Started with Botkube](https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments#get-started-with-botkube)

#### Start Using Botkube AI-Powered Assistant Today

As Microsoft Teams continues its steady rise as the go to platform for enterprise-level workplace collaboration, it is parallel to the rapid adoption of Azure and Azure Kubernetes Service (AKS). These collective forces drive organizations toward innovation, scalability, and streamlined cloud-native application deployment. However, this highlights a major challenge: as teams begin to incorporate Kubernetes and all of its complexities into their workflows, the need for a collaborative troubleshooting tool becomes essential.

Enter Botkube - a Kubernetes troubleshooting and monitoring tool designed to empower DevOps teams to work more efficiently. Botkube has now expanded its capabilities to seamlessly integrate with Microsoft Teams, fostering collaborative troubleshooting tailored to your organization's Teams workspace.

**_Botkube has a new version of MS Teams. It is easy to install and get up in running in five minutes or less. Click_** [**_here_**](https://botkube.io/blog/revolutionize-your-kubernetes-troubleshooting-workflow-with-microsoft-teams-and-botkube) **_to find out more!_**

Benefits of Using Botkube and MS Teams
--------------------------------------

Botkube’s integration offers useful features for Microsoft users working with Azure Kubernetes Service (AKS). It simplifies interactions with Kubernetes and the broader cloud-native ecosystem. With the Botkube plugin system, users can easily integrate tools like [Prometheus for monitoring](https://botkube.io/learn/how-botkube-makes-monitoring-kubernetes-easy) and [Helm for application package management](https://botkube.io/learn/helm-charts) within the Microsoft ecosystem. With Botkube, you can manage your deployments with Helm and customize your alerting set with Prometheus directly in Teams! This integration facilitates smoother deployment processes and more consistent monitoring.

Botkube empowers developers with self-service access while ensuring a controlled and secure environment. It offers a controlled environment for developers to access their Kubernetes resources. It enables the whitelist of potent commands like 'create' and 'delete,' allowing developers to experiment with Kubernetes tools without granting them full control over the clusters. This is particularly useful for enterprise teams because it allows for a balance between developer autonomy and maintaining security standards. In essence, Botkube enhances the AKS experience by streamlining tool integration and offering controlled access for developers.

Tutorial Guide
--------------

In this tutorial, we will guide you through the step-by-step process of configuring and leveraging [Botkube for Microsoft Teams and AKS](https://docs.botkube.io/installation/teams/). This enhancement empowers your team to efficiently establish a connection between Kubernetes and Microsoft Teams, facilitating the streamlined management of multiple Kubernetes clusters and significantly improving incident resolution times.

Prerequisites
-------------

*   [Botkube Cloud](https://app.botkube.io/) account
*   Access to a Kubernetes cluster
*   MS Teams account

#### One-Time Setup and Ongoing Flexibility

_\*\* Steps 1-4 and 7-10 are one-time operations requiring administrative privileges. Once set up, Botkube can be configured across various channels and clusters without needing further admin access in Teams. This means after the initial setup, users can easily manage alerts and configurations through the Botkube Web App.\*\*_

Creating Botkube Cloud Account
------------------------------

1.  On the [Botkube homepage](https://botkube.io/), locate the “Get Started” button and click on it. This will take you to the account registration page.
2.  Fill in the required information in the registration form. You can sign up with your email address or Github account.

Dashboard setup
---------------

1.  Select create a new instance

![Image 2: Web Dashboard to connect Botkube to your Cluster](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7dd390269cbe2d3ba4_jCUYYPLNDFZDkosCGhXjsC4Cvk9OsKaPJowAXS_Yi3-gdAekdM-YYj_QvgqvMCkAOIDbqXTaJGZuJFAjb5pIwZWo0kFlQwPBcwAzKW6X7ax6gK3rQVjbGKOJg_9Ps9i28sE-f7xg0hdp8hoY5mPwnNI.png)

![Image 3: Commands to run to add Botkube to Kubernetes Cluster for Teams](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7e072ab825022ca51e_w3OobTivy6lb1zbPozEcTOySEmAZPSRU0WVO120nal_egmZ-HVayv2FIuTzLsJ6vBJuZBfrFLiMkzjzpOS2kJash0C8_p3scSVIAFUY5Rb_1YqE2xACl2811ugQ1E-VazSxtzki-AirkeARSEZ5sKq0.png)

2.  Next, use homebrew or curl to install Botkube CLI on your terminal (Step 1)
3.  Copy and paste the `Botkube install`… command (Step 2)
4.  Wait for the cluster to connect and click, **Next**

‍

![Image 4: Select preferred communication platform connection](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7e1938916887fbf34f_WVBv6m9k1C5RWt4FI7QuA8VU-lSDXKQZOJWyqvfe7YZVDnNxquO3DBkznU2LP9TulrVxeDPloV4O7w40n6OVt3NjPPkMynGNKA_6wbc1knG-znVU3N5E8J6H1fqmpWafhnh4eOgb37W1Di2MgHUrrtU.png)

5.  Name your instance and select the **Teams** option

![Image 5: Setting up Teams to receive Kubernetes alerts](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7e6a852b2620a9b5fb_7fpjvjcPJcqBZKRJ2oNOisODdOt9vN9J-sYcsuFLYMzlKF8CzLOXI3vM83HZ0YvdShckXSZcDBTclvCv9IRPGrNqeIIMBA6KdnKhilwSek9nqKBLZOrwUUFMZYe5zIBJmv2nh3jwJuJK2kjgO5Lh1ZM.png)

### Installing Botkube to Microsoft Teams

If you set this connection up previously, it should still be connected. However Botkube recently got listed on the [Microsoft Teams App Store](https://botkube.io/blog/botkube-is-now-verified-in-microsoft-appsource), so the connection just got way easier with less steps. So if you want to set the connection up again you may notice it went from 18 steps to now only 11 with Botkube being an approved Teams and Azure application.

1.  Navigate to the Microsoft Teams application and select **Manage your Apps** at the button of the screen

![Image 6: Adding Teams Apps from Store](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7d14da41d53c779535_oErOXojrwoCSfbs3A0830mnbkI1xb2rOD83YopN5B6k7Ti_O5OZdusUD4VrNjdFs_xHSaTxCjdipKQJuaLhpksHgW3Fwd8xx8MSFSq1HiuWVzsoTG90t1Dy4nlcAAkIuNdAsumH4sjQWXljJJ9y_Lbo.png)

2.  Next select **Search for Botkube** in the search bar at the top of the screen.

![Image 7: Botkube's listing on Microsoft teams store](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/664cdd1749371ee8efa0e9ba_Botkube%20Teams%20APp%20pic.png)

3.  And then select the **Add to team** button

![Image 8: Further info on the Microsoft Teams Kubernetes connection](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/664cde0777d69d63cd727d52_Botkube%20Teams%20listing.png)

4\. Select your preferred team or channel you would like to add Botkube to

![Image 9: Adding bot to chat channel](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7e778b3a2f89c145c5_2I5Ickc_OXlcgym-Ju-rDd9gVc0kVk-QW0_js_3gURGp0dgqj5jS0lOzoIox8vr07ky4hiHkn3LHiTIJ0JKU4jK6Q9rIsan8_lavBnIF7WETp-F_LEM0bTZW4keiuLe3cM3VA_9leLry29hAZ1_vXfU.png)

5.  Next, Navigate to the chat window and a welcome message should appear

![Image 10: Welcome message](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7db157112561870666_ZfA2XrAwSEqm7yLkbUMoOEL_3sIULsNu2rHAaerMIvhOyKO79bjSFvausKcnNF_rxPc8rq0rFAl2VGRNct3Gb0aJimPt3pYbJInhW0-Z6ffQWq3_gPAbiJbHQHMhefcBbAbROu95icJakohqR-patiY.png)

6.  Select **Connect to Botkube Cloud**

![Image 11: Further adding the bot](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7db157112561870666_ZfA2XrAwSEqm7yLkbUMoOEL_3sIULsNu2rHAaerMIvhOyKO79bjSFvausKcnNF_rxPc8rq0rFAl2VGRNct3Gb0aJimPt3pYbJInhW0-Z6ffQWq3_gPAbiJbHQHMhefcBbAbROu95icJakohqR-patiY.png)

7.  This will take you to the Botkube Cloud dashboard. Your Teams channel will appear. Select the **Connect** Button.

![Image 12: Simple connect button](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/656eefd26cc09bfeec2845b3_Screen%20Shot%202023-12-05%20at%201.38.28%20AM.png)

![Image 13: Success message for alerts to show up in channel](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7d165a59345b2b698b__rEh9UH04sSf8XxhQi_3JhGSjynIaVLwD--bRFqQa3v2Rqrahxpnna3yryM1a4omthQ-Fize-gyhNgRAXDTl-DYQXkJ1LUhp1OvRWNwn62jwfra7qa806TPcVm13W3pbeA52XN47_MlPkUVEvvyb6KA.png)

8.  All of your Teams information will appear in the dashboard. Select the channels you want to use with Botkube.

![Image 14: Control all channels that the Azure Kubernetes cluster is connected to in a dashboard](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7ead44d850b34e7f35_nSgmglbxW-mAFvthFI-1q0c0dcMg2m2wejlA2__CmM-vyDDeyLjUh84vTLufTx77jaJ-ifWWa1bzkFpn7bPK0KbehAARBOg2Zle9HbNUk3SXAP9-jcNbNXFypsfgSPY75R2BLLwmDi9nUFUY88StBNg.png)

9.  Connect your preferred plugins

![Image 15: Add extra DevOps related plugins to the cluster](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7ddec01b0e38f3eb6d_pVkuqIiFGr5zQ8lNXQcSMQXQhDbPx5rZ6m0OptpyWKHIfpzzTwn1UbTR44-HU_YM2NOBfoOhvapjYfohK0AUjF5dsvV_8JGnujLfhEzupnaCMLJoD4pzbAE6aHZemjv1Rzzi3rsu8HoFbqrSxbu1TVI.png)

10.  Review and select your preferred Botkube defaults.

![Image 16: Adding kubectl aliases to Teams](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7d6f8e75cab928022a_PVwnI51Br_JjDcyKQsNjixKmrtTDW10Ug0gHIUJMFxDo51P-_aRrStZiIrK-Pmqxg3DnFwrtTRZqhHUGhZv2d37pHnWMmUvx9p5o4FnQe8M-YkkM8wVZk7_P3-tUs0umUmnfE37H4QbtR8E-yg36w2A.png)

11.  Select **Apply changes**‍
    
12.  Congratulations! You have successfully deployed Botkube for Microsoft Teams!
    

Conclusion
----------

The new Botkube Microsoft Teams integration offers an all-in-one solution for MS Teams users. This integration enhances the Azure Kubernetes experience by simplifying interactions with Kubernetes and the broader cloud-native ecosystem. Botkube’s features like real-time notifications support across multiple channels streamline deployment management and open up a range of potential use cases for incident response management and troubleshooting. This integration empowers developers with self-service access while ensuring control and security, making it a valuable tool for enterprise teams seeking a balance between autonomy and compliance with security standards.

Check out [Botkube Microsoft Teams integration here.](https://botkube.io/integration/teams)

Get Started with Botkube
------------------------

Whether you're a seasoned Kubernetes pro or just getting started, Botkube can help supercharge your troubleshooting process. [Sign up now](https://app.botkube.io/) for free and join the community of users who are already benefiting from the power of Botkube.

We want to know what you think of Botkube and how we can make it even better. [Email our Developer Advocate, Maria](mailto:maria@kubeshop.io) or schedule [a quick 15 meeting](https://calendly.com/maria-botkube/15min) at your preferred time. As a thank you, we’ll send you some great Botkube swag!
