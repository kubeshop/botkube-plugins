Title: The Botkube Cloud App for Slack

URL Source: https://botkube.io/slack-app

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/64e8e512edbae7755379b2fe_bk-slack-logo-hero-bg.webp)

Cloud App for Slack
-------------------

![Image 2](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500bcab958c3be57a5523a5_slack-logo-color.svg)

#### Cloud App for Slack

The Botkube Cloud App for Slack uses Botkube Cloud services to manage channels and route executor commands. This allows multi-cluster support without a need to create a dedicated application for Slack for each cluster. Events and alerts are sent directly from your cluster to your Slack workspace for reliable, fast notifications.

Installing Botkube in Slack
---------------------------

1

[![Image 3](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c008d898dab566f70d85_install-1.svg) ![Image 4](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)

Go to Botkube Cloud  [Web App](https://app.botkube.io/)

and click on New Instance button.

2

[![Image 5](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c1702080a8b5a26efb6d_install-2.svg) ![Image 6](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Fill in the Instance Display Name and click Next button.

3

[![Image 7](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c17b5527fae1a4c5bc57_install-3.svg) ![Image 8](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Click Add platform dropdown, and select Slack option.

4

[![Image 9](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c191387a00b2b474e1b3_install-5.svg) ![Image 10](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Click Add to Slack button to add Cloud integration for Slack to your Slack workspace.

5

[![Image 11](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c4fab232b58125653577_install-6.svg) ![Image 12](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Click Allow to grant permission for Botkube Cloud App for Slack to access your Slack workspace.

6

[![Image 13](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c505324fd1b24eeb1ab6_install-7.svg) ![Image 14](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Provide the details as described follows and click Next button.

7

[![Image 15](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c510d21725cf9c03a0c9_install-8.svg) ![Image 16](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Add plugins you want  to enable in your Botkube instance and click Next button.

8

[![Image 17](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c51cd898dab566fcb743_install-9.svg) ![Image 18](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Include optional default aliases and default actions and click Create button to create Botkube Cloud instance.

9

[![Image 19](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/6500c528a593ecc0b3dcfa98_install-10.svg) ![Image 20](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)Follow the instructions  in the summary page to deploy Botkube into your environment.

#### Prerequisites

A Botkube Cloud account with active subscription is required.

You can try out the Botkube Cloud App for Slack for free by creating an account in the [Botkube Cloud app](https://app.botkube.io/) and starting a free trial.

NOTE: If you downgrade a paid subscription to the Free plan, you will not be able to execute Botkube commands from Slack. However, you will still receive notifications from your clusters.

Steps to get started with Botkube
---------------------------------

1

You can start using Botkube Cloud App for Slack by typing @Botkube cloud help in the Slack channel you configured in one of the previous steps.

[![Image 21](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650b312e6823f3bf4930026a_step1-cloud-help_pr.svg) ![Image 22](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)

2

You can list all the Botkube Cloud instances by typing @Botkube cloud list instances in the Slack channel or click the button List connected instances in the help command response. Besides the instance name, ID, and status in the list response, you can also click the Get details button to go to instance details on Botkube Cloud Dashboard.

[![Image 23](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650b321b09f2a16899841a41_step2-cloud-instance_pr.svg) ![Image 24](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)

3

Once a Botkube command is executed, it will be handled on target Kubernetes cluster specified with \--cluster-name flag. However, this is an optional flag, where if you have not specified it, Botkube Cloud will select the first instance. However, you can also achieve setting default instance with command @Botkube cloud set default-instance {instance-id}.

After this point, all of your commands will be executed on the default instance. Moreover, if you want to execute a command on all the target clusters, you can use \--all-clusters flag.

4

1\. Go to Botkube Cloud instances page and click Manage button of the instance you want to remove.

[![Image 25](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650b3581864c2eaa967f7417_step4-cloud_list_manage_pr.svg) ![Image 26](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)

5

2\. Click Delete instance button, type instance name in the popup and click Delete instance.

[![Image 27](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650b35df6390bea044f2eb3b_step5-cloud_delete_pr.svg) ![Image 28](https://cdn.prod.website-files.com/633705de6adaa38599d8e258/650214f73dc8a48ae4075f8f_magnifier.svg)](#)
