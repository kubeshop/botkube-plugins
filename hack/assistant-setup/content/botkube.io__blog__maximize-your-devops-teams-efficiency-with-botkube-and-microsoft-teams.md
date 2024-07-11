Title: Maximize Your DevOps Teams Efficiency with Botkube + Microsoft Teams

URL Source: https://botkube.io/blog/maximize-your-devops-teams-efficiency-with-botkube-and-microsoft-teams

Published Time: May 20, 2024

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Get Botkube up and running with Microsoft Teams in less than 5 mins with this walk-through tutorial

### Table of Contents

*   [Introduction](https://botkube.io/blog/maximize-your-devops-teams-efficiency-with-botkube-and-microsoft-teams#introduction)
*   [Benefits of Using Botkube and Microsoft Teams](https://botkube.io/blog/maximize-your-devops-teams-efficiency-with-botkube-and-microsoft-teams#benefits-of-using-botkube-and-microsoft-teams)
*   [Tutorial Guide](https://botkube.io/blog/maximize-your-devops-teams-efficiency-with-botkube-and-microsoft-teams#tutorial-guide-)
*   [Conclusion](https://botkube.io/blog/maximize-your-devops-teams-efficiency-with-botkube-and-microsoft-teams#conclusion-)
*   [Get Started with Botkube](https://botkube.io/blog/maximize-your-devops-teams-efficiency-with-botkube-and-microsoft-teams#get-started-with-botkube)

#### Start Using Botkube AI-Powered Assistant Today

Introduction
------------

Microsoft Teams is now a leading choice for collaboration in the business world, particularly in large enterprises. The platform is pivotal in the evolving landscape of hybrid work environments. Its growth aligns with the rapid adoption of Azure Cloud services and Azure Kubernetes Service (AKS), which are steering organizations towards more innovative, scalable, and efficient deployment of cloud-based applications. Yet, this advancement also introduces a key challenge: the complex task of integrating Kubernetes into DevOps workflows.

Enter Botkube, a collaborative Kubernetes troubleshooting and monitoring tool. It empowers DevOps teams to work more efficiently by enabling developers to troubleshoot applications without requiring Kubernetes expertise. Botkube significantly improves team reliability by delivering timely, context-enhanced notifications about events in Kubernetes environments. This integration facilitates a seamless troubleshooting environment, perfectly tailored to the workflows of Microsoft Teams' user base. Beyond alerting, Botkube allows teams to automate actions based on events, executing kubectl and Helm commands, receiving best practices recommendations, and much more.

Benefits of Using Botkube and Microsoft Teams
---------------------------------------------

Botkube’s integration offers useful features for Microsoft users working with Kubernetes. It simplifies interactions with Kubernetes and the broader cloud-native ecosystem. With the Botkube plugin system, users can easily integrate tools like [Prometheus for monitoring](https://botkube.io/learn/how-botkube-makes-monitoring-kubernetes-easy) and [Helm for application package management](https://botkube.io/learn/helm-charts) within the Microsoft ecosystem. With Botkube, you can manage your deployments with Helm and customize your alerting set with Prometheus in the Botkube web app and receive alerts directly in Teams. This integration facilitates smoother deployment processes and more consistent monitoring. Our [engineering team](http://learn%20more%20about%20the%20engineering%20process/) built this new Teams app with enterprise developer in-mind.

Botkube empowers developers with self-service access while ensuring a controlled and secure environment. It offers a controlled environment for developers to access their Kubernetes resources. It enables the whitelist of powerful commands like 'create' and 'delete,' allowing developers to experiment with Kubernetes tools without granting them full control over the clusters. This is particularly useful for enterprise teams because it allows for a balance between developer autonomy and maintaining security standards. In essence, Botkube enhances the Kubernetes experience by streamlining tool integration and offering controlled access for developers.

Tutorial Guide
--------------

In this tutorial, we will guide you through the step-by-step process of configuring and leveraging Botkube for Microsoft Teams.The previous (now deprecated) version of Botkube’s Teams integration included a lot of manual steps that required you to manually expose an ingress point. In this new version, you'll simply download the zip file of the Teams application from the Botkube Cloud dashboard, upload it directly to your Team app directory (with admin privileges), and connect your cluster to Botkube. Get up and running with Botkube in Teams in less than 5 minutes.

### Prerequisites

*   [Botkube Cloud](https://app.botkube.io/) account
*   Access to a Kubernetes cluster
*   Microsoft Teams account with administrator privileges

#### One-Time Setup and Ongoing Flexibility

_\*\* A few steps are one-time operations requiring administrative privileges. Once set up, Botkube can be configured across various channels and clusters without needing further admin access in Teams. This means after the initial setup, users can easily manage alerts and configurations through the_ [_Botkube Web App_](https://app.botkube.io/)_.\*\*_

### Creating Botkube Cloud Account

1.  On the [Botkube homepage](https://botkube.io/), locate the “Get Started” button and click on it. This will take you to the account registration page.
2.  Fill in the required information in the registration form. You can sign up with your email address or Github account.

### Dashboard setup

1.  Select create a new instance

![Image 2: Create a instance for the Bot](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7dd390269cbe2d3ba4_jCUYYPLNDFZDkosCGhXjsC4Cvk9OsKaPJowAXS_Yi3-gdAekdM-YYj_QvgqvMCkAOIDbqXTaJGZuJFAjb5pIwZWo0kFlQwPBcwAzKW6X7ax6gK3rQVjbGKOJg_9Ps9i28sE-f7xg0hdp8hoY5mPwnNI.png)

![Image 3: Install code for Botkube into a Kubernetes Cluster](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7e072ab825022ca51e_w3OobTivy6lb1zbPozEcTOySEmAZPSRU0WVO120nal_egmZ-HVayv2FIuTzLsJ6vBJuZBfrFLiMkzjzpOS2kJash0C8_p3scSVIAFUY5Rb_1YqE2xACl2811ugQ1E-VazSxtzki-AirkeARSEZ5sKq0.png)

2.  Next, use homebrew or curl to install Botkube CLI on your terminal (Step 1)
3.  Copy and paste the `Botkube install`… command (Step 2)
4.  Wait for the cluster to connect and click, **Next**

‍

![Image 4: Selecting Chat Platform to connect to](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7e1938916887fbf34f_WVBv6m9k1C5RWt4FI7QuA8VU-lSDXKQZOJWyqvfe7YZVDnNxquO3DBkznU2LP9TulrVxeDPloV4O7w40n6OVt3NjPPkMynGNKA_6wbc1knG-znVU3N5E8J6H1fqmpWafhnh4eOgb37W1Di2MgHUrrtU.png)

5.  Name your instance and select the **Teams** option

![Image 5: Connecting Kubernetes to Microsoft Teams](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/664f5adaddbd09c3d1dfc64e_Teams%20download%20for%20Botkube.png)

### Installing Botkube to Microsoft Teams

1.  Navigate to the Microsoft Teams application and select **Manage your Apps** at the button of the screen

![Image 6: teams App Store](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7d14da41d53c779535_oErOXojrwoCSfbs3A0830mnbkI1xb2rOD83YopN5B6k7Ti_O5OZdusUD4VrNjdFs_xHSaTxCjdipKQJuaLhpksHgW3Fwd8xx8MSFSq1HiuWVzsoTG90t1Dy4nlcAAkIuNdAsumH4sjQWXljJJ9y_Lbo.png)

2.  Next either **Search to find** Botkube in the Teams App store or simply select **Add.**

![Image 7: Add Botkube App](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/664cdd1749371ee8efa0e9ba_Botkube%20Teams%20APp%20pic.png)

3.  Select the **Add to Team** buttton.

![Image 8: Botkube's Microsoft Teams App Store Listing](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/664cde0777d69d63cd727d52_Botkube%20Teams%20listing.png)

4\. Select your preferred team or channel you would like to add Botkube to

![Image 9: Choose chat channels](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7e778b3a2f89c145c5_2I5Ickc_OXlcgym-Ju-rDd9gVc0kVk-QW0_js_3gURGp0dgqj5jS0lOzoIox8vr07ky4hiHkn3LHiTIJ0JKU4jK6Q9rIsan8_lavBnIF7WETp-F_LEM0bTZW4keiuLe3cM3VA_9leLry29hAZ1_vXfU.png)

‍

```
    5. Next, Navigate to the chat window and a welcome message should appear

    6. Navigate to the chat window and select <strong>Connect to Botkube Cloud</strong>
```

![Image 10: Bot message in Teams](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7db157112561870666_ZfA2XrAwSEqm7yLkbUMoOEL_3sIULsNu2rHAaerMIvhOyKO79bjSFvausKcnNF_rxPc8rq0rFAl2VGRNct3Gb0aJimPt3pYbJInhW0-Z6ffQWq3_gPAbiJbHQHMhefcBbAbROu95icJakohqR-patiY.png)

```
 7. This will take you to the Botkube Cloud dashboard. Your Teams channel will appear. Select the <strong>Connect </strong>Button.
```

![Image 11: One last button to click](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/656eefd26cc09bfeec2845b3_Screen%20Shot%202023-12-05%20at%201.38.28%20AM.png)

![Image 12: k8s connected to Microsoft Teams](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7d165a59345b2b698b__rEh9UH04sSf8XxhQi_3JhGSjynIaVLwD--bRFqQa3v2Rqrahxpnna3yryM1a4omthQ-Fize-gyhNgRAXDTl-DYQXkJ1LUhp1OvRWNwn62jwfra7qa806TPcVm13W3pbeA52XN47_MlPkUVEvvyb6KA.png)

```
  8. All of your Teams information will appear in the dashboard. Select the channels you want to use with Botkube. 
```

![Image 13: Customize alerts coming into Teams from Kubernetes](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7ead44d850b34e7f35_nSgmglbxW-mAFvthFI-1q0c0dcMg2m2wejlA2__CmM-vyDDeyLjUh84vTLufTx77jaJ-ifWWa1bzkFpn7bPK0KbehAARBOg2Zle9HbNUk3SXAP9-jcNbNXFypsfgSPY75R2BLLwmDi9nUFUY88StBNg.png)

```
 10. Connect your preferred plugins 
```

![Image 14: Add more cloud native plugins](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7ddec01b0e38f3eb6d_pVkuqIiFGr5zQ8lNXQcSMQXQhDbPx5rZ6m0OptpyWKHIfpzzTwn1UbTR44-HU_YM2NOBfoOhvapjYfohK0AUjF5dsvV_8JGnujLfhEzupnaCMLJoD4pzbAE6aHZemjv1Rzzi3rsu8HoFbqrSxbu1TVI.png)

```
 11. Review and select your preferred Botkube defaults. 
```

![Image 15: Save the new instance](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65678f7d6f8e75cab928022a_PVwnI51Br_JjDcyKQsNjixKmrtTDW10Ug0gHIUJMFxDo51P-_aRrStZiIrK-Pmqxg3DnFwrtTRZqhHUGhZv2d37pHnWMmUvx9p5o4FnQe8M-YkkM8wVZk7_P3-tUs0umUmnfE37H4QbtR8E-yg36w2A.png)

```
    12. Select <strong>Apply changes</strong>

    13. Congratulations! You have successfully deployed Botkube for Microsoft Teams!
```

Conclusion
----------

The updated Botkube Microsoft Teams integration offers an all-in-one solution for Microsoft Teams users. Botkube’s features like real-time notifications support across multiple channels streamline deployment management and open up a range of potential use cases for incident response management and troubleshooting. This integration empowers developers with self-service access while ensuring control and security, making it a valuable tool for enterprise teams seeking a balance between autonomy and compliance with security standards.

Get Started with Botkube
------------------------

Whether you're a seasoned Kubernetes pro or just getting started, Botkube can help supercharge your troubleshooting process. [Sign up now](https://app.botkube.io/) for free and join the community of users who are already benefiting from the power of Botkube. Once signed in, follow the tutorial above or read the [official Botkube documentation](https://docs.botkube.io/installation/teams/).

We want to know what you think of Botkube and how we can make it even better. [Email our Developer Advocate, Maria](mailto:maria@kubeshop.io) or schedule [a quick 15 meeting](https://calendly.com/maria-botkube/15min) at your preferred time. As a thank you, we’ll send you some great Botkube swag.
