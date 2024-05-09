Title: Integrating Microsoft Teams with Azure for K8s Deployments

URL Source: https://botkube.io/blog/integrating-microsoft-teams-with-azure-for-kubernetes-deployments

Published Time: Oct 30, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

Botkube is entering the MS teams world

### Table of Contents

*   [Benefits of Using Botkube and MS Teams](#benefits-of-using-botkube-and-ms-teams)
*   [Tutorial Guide](#tutorial-guide-)
*   [Prerequisites](#prerequisites)
*   [Creating Botkube Cloud Account](#creating-botkube-cloud-account)
*   [Dashboard setup](#dashboard-setup)
*   [Conclusion](#conclusion)
*   [Get Started with Botkube](#get-started-with-botkube)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

#### Start Receiving Kubernetes Notifications Now with Botkube Cloud

As Microsoft Teams continues its steady rise as the go to platform for enterprise-level workplace collaboration, it is parallel to the rapid adoption of Azure and Azure Kubernetes Service (AKS). These collective forces drive organizations toward innovation, scalability, and streamlined cloud-native application deployment. However, this highlights a major challenge: as teams begin to incorporate Kubernetes and all of its complexities into their workflows, the need for a collaborative troubleshooting tool becomes essential.

Enter Botkube - a Kubernetes troubleshooting and monitoring tool designed to empower DevOps teams to work more efficiently. Botkube has now expanded its capabilities to seamlessly integrate with Microsoft Teams, fostering collaborative troubleshooting tailored to your organization's Teams workspace.

**_Botkube has a new version of MS Teams. It is easy to install and get up in running in five minutes or less. Click_** [**_here_**](http://botkube.io/blog/revolutionize-your-kubernetes-troubleshooting-workflow-with-microsoft-teams-and-botkube) **_to find out more!_**

Benefits of Using Botkube and MS Teams
--------------------------------------

Botkube’s integration offers useful features for Microsoft users working with Azure Kubernetes Service (AKS). It simplifies interactions with Kubernetes and the broader cloud-native ecosystem. With the Botkube plugin system, users can easily integrate tools like [Prometheus for monitoring](https://botkube.io/learn/how-botkube-makes-monitoring-kubernetes-easy) and [Helm for application package management](https://botkube.io/learn/helm-charts) within the Microsoft ecosystem. With Botkube, you can manage your deployments with Helm and customize your alerting set with Prometheus directly in Teams! This integration facilitates smoother deployment processes and more consistent monitoring.

Botkube empowers developers with self-service access while ensuring a controlled and secure environment. It offers a controlled environment for developers to access their Kubernetes resources. It enables the whitelist of potent commands like 'create' and 'delete,' allowing developers to experiment with Kubernetes tools without granting them full control over the clusters. This is particularly useful for enterprise teams because it allows for a balance between developer autonomy and maintaining security standards. In essence, Botkube enhances the AKS experience by streamlining tool integration and offering controlled access for developers.

Tutorial Guide
--------------

In this tutorial, we will guide you through the step-by-step process of configuring and leveraging [Botkube for Microsoft Teams and AKS](https://docs.botkube.io/installation/teams/). This enhancement empowers your team to efficiently establish a connection between Kubernetes and Microsoft Teams, facilitating the streamlined management of multiple Kubernetes clusters and significantly improving incident resolution times.

Prerequisites
-------------

*   [Botkube Cloud](http://app.botkube.io/) account
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

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7dd390269cbe2d3ba4_jCUYYPLNDFZDkosCGhXjsC4Cvk9OsKaPJowAXS_Yi3-gdAekdM-YYj_QvgqvMCkAOIDbqXTaJGZuJFAjb5pIwZWo0kFlQwPBcwAzKW6X7ax6gK3rQVjbGKOJg_9Ps9i28sE-f7xg0hdp8hoY5mPwnNI.png)

![Image 3](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7e072ab825022ca51e_w3OobTivy6lb1zbPozEcTOySEmAZPSRU0WVO120nal_egmZ-HVayv2FIuTzLsJ6vBJuZBfrFLiMkzjzpOS2kJash0C8_p3scSVIAFUY5Rb_1YqE2xACl2811ugQ1E-VazSxtzki-AirkeARSEZ5sKq0.png)

2.  Next, use homebrew or curl to install Botkube CLI on your terminal (Step 1)
3.  Copy and paste the `Botkube install`… command (Step 2)
4.  Wait for the cluster to connect and click, **Next**

‍

![Image 4](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7e1938916887fbf34f_WVBv6m9k1C5RWt4FI7QuA8VU-lSDXKQZOJWyqvfe7YZVDnNxquO3DBkznU2LP9TulrVxeDPloV4O7w40n6OVt3NjPPkMynGNKA_6wbc1knG-znVU3N5E8J6H1fqmpWafhnh4eOgb37W1Di2MgHUrrtU.png)

5.  Name your instance and select the **Teams** option

![Image 5](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7e6a852b2620a9b5fb_7fpjvjcPJcqBZKRJ2oNOisODdOt9vN9J-sYcsuFLYMzlKF8CzLOXI3vM83HZ0YvdShckXSZcDBTclvCv9IRPGrNqeIIMBA6KdnKhilwSek9nqKBLZOrwUUFMZYe5zIBJmv2nh3jwJuJK2kjgO5Lh1ZM.png)

6.  Select the Download button to download the Botkube App. It will show up as **Botkube.zip**

‍

![Image 6](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7db9251f0d1db86ab7_w9yhNzeHzqqc4JawGTersD2qogEalGSUFs9xlGLhX_8OoCXcxSEEaJNnXaAdB3KyA25t6XnaNAzAM_P1cBjLNN68lZZGZ6GjAwWL8iAlXn4hyrwJ5FM3p8MLBCyUtwoSvw1ZRWBs7ds6jYHeYaJJJAU.png)

### Installing Botkube to Microsoft Teams

‍

1.  Navigate to the Microsoft Teams application and select **Manage your Apps** at the button of the screen

![Image 7](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7d14da41d53c779535_oErOXojrwoCSfbs3A0830mnbkI1xb2rOD83YopN5B6k7Ti_O5OZdusUD4VrNjdFs_xHSaTxCjdipKQJuaLhpksHgW3Fwd8xx8MSFSq1HiuWVzsoTG90t1Dy4nlcAAkIuNdAsumH4sjQWXljJJ9y_Lbo.png)

2.  Next select **Upload an app** and select **Upload an app to your org’s app catalog**

![Image 8](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/656eef4cf0cfa6050f79acc5_Screen%20Shot%202023-12-05%20at%201.36.02%20AM.png)

3.  Select the **Botkube.zip** file from the previous section
4.  Navigate to the **Apps** section on the left hand side and select the **Built for your Org section.** The Botkube application will be there.

‍

5.  Select the **Add** button

![Image 9](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7d640c39e96661f69e_zpXrkocqwakQB8jrA8pwYH2CSrmRim4pp13-Pt8lZRgdgw33jCDfqnrxpYK3jt2fhSctK5cIyMt9taA-VAAUSD3tiicgAdTafZ7gAqBs08Bkljt7BJlUsxUh3OUdg1fhjCiCO5lTztUHXd8CrccWxXs.png)

![Image 10](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7d053e39713de9aa28_FNgbV-UuRUprEaIcHxqlypAX2PnrAGJjbfHZu8I2uJ4zRqP3ZXt3L4ez8rPqWhZZv_-ruLTTXoWLZFO5Nv4vKTfoJ8hjqFjrHX-M_RwhlaAPrF5IDEcZjDg523tVuwYOMranVLyV3IOVrbDzRI4OY5M.png)

6.  And then select the **Add to team** button

![Image 11](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7e778b3a2f89c145c5_2I5Ickc_OXlcgym-Ju-rDd9gVc0kVk-QW0_js_3gURGp0dgqj5jS0lOzoIox8vr07ky4hiHkn3LHiTIJ0JKU4jK6Q9rIsan8_lavBnIF7WETp-F_LEM0bTZW4keiuLe3cM3VA_9leLry29hAZ1_vXfU.png)

7.  Select your preferred team or channel you would like to add Botkube to

![Image 12](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7db157112561870666_ZfA2XrAwSEqm7yLkbUMoOEL_3sIULsNu2rHAaerMIvhOyKO79bjSFvausKcnNF_rxPc8rq0rFAl2VGRNct3Gb0aJimPt3pYbJInhW0-Z6ffQWq3_gPAbiJbHQHMhefcBbAbROu95icJakohqR-patiY.png)

8.  Next, Navigate to the chat window and a welcome message should appear
9.  Select the **Grant Access** button

![Image 13](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7d3cb1d0c7707a8a18_OHBLhhF5yDy9WTlI0AGKR5CnwnFjre9_Vz2BKdaUsnpAnxwk4hdqExTzBxp3yX0MH9wKGbktN0E6RG9YPeLlMShX0ybFtJ2eNbkt2WmHrLzTkBm7bZPnYcDLmy9YbQEcsbhNh5ZWj1HhibNRxe223IQ.png)

10.  A new window will pop up asking you to grant access. Click **Accept**

‍

11.  Navigate to the chat window and select **Connect to Botkube Cloud**

![Image 14](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7db157112561870666_ZfA2XrAwSEqm7yLkbUMoOEL_3sIULsNu2rHAaerMIvhOyKO79bjSFvausKcnNF_rxPc8rq0rFAl2VGRNct3Gb0aJimPt3pYbJInhW0-Z6ffQWq3_gPAbiJbHQHMhefcBbAbROu95icJakohqR-patiY.png)

12.  This will take you to the Botkube Cloud dashboard. Your Teams channel will appear. Select the **Connect** Button.

![Image 15](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/656eefd26cc09bfeec2845b3_Screen%20Shot%202023-12-05%20at%201.38.28%20AM.png)

![Image 16](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7d165a59345b2b698b__rEh9UH04sSf8XxhQi_3JhGSjynIaVLwD--bRFqQa3v2Rqrahxpnna3yryM1a4omthQ-Fize-gyhNgRAXDTl-DYQXkJ1LUhp1OvRWNwn62jwfra7qa806TPcVm13W3pbeA52XN47_MlPkUVEvvyb6KA.png)

13.  All of your Teams information will appear in the dashboard. Select the channels you want to use with Botkube.

![Image 17](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7ead44d850b34e7f35_nSgmglbxW-mAFvthFI-1q0c0dcMg2m2wejlA2__CmM-vyDDeyLjUh84vTLufTx77jaJ-ifWWa1bzkFpn7bPK0KbehAARBOg2Zle9HbNUk3SXAP9-jcNbNXFypsfgSPY75R2BLLwmDi9nUFUY88StBNg.png)

14.  Connect your preferred plugins

![Image 18](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7ddec01b0e38f3eb6d_pVkuqIiFGr5zQ8lNXQcSMQXQhDbPx5rZ6m0OptpyWKHIfpzzTwn1UbTR44-HU_YM2NOBfoOhvapjYfohK0AUjF5dsvV_8JGnujLfhEzupnaCMLJoD4pzbAE6aHZemjv1Rzzi3rsu8HoFbqrSxbu1TVI.png)

15.  Review and select your preferred Botkube defaults.

![Image 19](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65678f7d6f8e75cab928022a_PVwnI51Br_JjDcyKQsNjixKmrtTDW10Ug0gHIUJMFxDo51P-_aRrStZiIrK-Pmqxg3DnFwrtTRZqhHUGhZv2d37pHnWMmUvx9p5o4FnQe8M-YkkM8wVZk7_P3-tUs0umUmnfE37H4QbtR8E-yg36w2A.png)

16.  Select **Apply changes**
17.  Congratulations! You have successfully deployed Botkube for Microsoft Teams!

Conclusion
----------

The new Botkube Microsoft Teams integration offers an all-in-one solution for MS Teams users. This integration enhances the Azure Kubernetes experience by simplifying interactions with Kubernetes and the broader cloud-native ecosystem. Botkube’s features like real-time notifications support across multiple channels streamline deployment management and open up a range of potential use cases for incident response management and troubleshooting. This integration empowers developers with self-service access while ensuring control and security, making it a valuable tool for enterprise teams seeking a balance between autonomy and compliance with security standards.

Check out [Botkube Microsoft Teams integration here.](https://botkube.io/integration/teams)

Get Started with Botkube
------------------------

Whether you're a seasoned Kubernetes pro or just getting started, Botkube can help supercharge your troubleshooting process. [Sign up now](https://app.botkube.io/) for free and join the community of users who are already benefiting from the power of Botkube.

We want to know what you think of Botkube and how we can make it even better. [Email our Developer Advocate, Maria](mailto:maria@kubeshop.io) or schedule [a quick 15 meeting](https://calendly.com/maria-botkube/15min) at your preferred time. As a thank you, we’ll send you some great Botkube swag!
