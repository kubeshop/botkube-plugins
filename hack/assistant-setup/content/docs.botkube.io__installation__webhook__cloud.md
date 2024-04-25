Title: Outgoing webhook for Botkube Cloud | Botkube

URL Source: https://docs.botkube.io/installation/webhook/cloud

Markdown Content:
*   [](https://docs.botkube.io/) * [Installation](https://docs.botkube.io/)
*   [Outgoing webhook](https://docs.botkube.io/installation/webhook/)
*   Outgoing webhook for Botkube Cloud

Version: 1.10

Prerequisites[​](#prerequisites"DirectlinktoPrerequisites")
---------------------------------------------------------------

*   Botkube Cloud account which you can create [here](https://app.botkube.io/) for free.

Create a Botkube Cloud Instance with Webhook[​](#create-a-botkube-cloud-instance-with-webhook"DirectlinktoCreateaBotkubeCloudInstancewithWebhook")
------------------------------------------------------------------------------------------------------------------------------------------------------------

1.  Go to Botkube Cloud [Web App](https://app.botkube.io/) and create a new instance.

You can do it by clicking "Create an Instance" button on Home Page or under this link [Create an Instance](https://app.botkube.io/instances/add)

2.  Fill in the `Instance Display Name` and click `Next` button.

![Image 1: Instance Display Name](https://docs.botkube.io/assets/images/webhook_instance_display_name-b35605d19eef1ecc93de54d6eefacae5.png)

3.  Click `Add platform` dropdown, and select `Webhook` option.

![Image 2: Select Platform](https://docs.botkube.io/assets/images/webhook_platform_select-aac36ca4e34549bef88cc00b3603f4ac.png)

4.  Fill in all required data in the form

![Image 3: Form](https://docs.botkube.io/assets/images/webhook_form-ccd69a8ec75ac03c98154db9bcf32b13.png)

5.  Add plugins you want to enable in your Botkube instance and click `Next` button.

![Image 4: Plugins](https://docs.botkube.io/assets/images/webhook_add_plugins-0bafa9371a2d3bccfc36c138bd442456.png)

6.  Include optional `default aliases` and `default actions` and click `Create` button to create Botkube Cloud instance.

![Image 5: Create](https://docs.botkube.io/assets/images/webhook_create-27a060686f00b0fb21f01839fd959e04.png)

note

If you don't include other platforms which use `Executor` plugins we recommend `default aliases` and `default actions` options unchecked

7.  Follow the instructions in the summary page to deploy Botkube into your environment.

![Image 6: Summary](https://docs.botkube.io/assets/images/webhook_summary-bfdc3ff0af6735b41a17d7219fd6b6f0.png)


Clean up[​](#clean-up"DirectlinktoCleanup")
------------------------------------------------

### Remove Botkube from Kubernetes cluster[​](#remove-botkube-from-kubernetes-cluster"DirectlinktoRemoveBotkubefromKubernetescluster")

1.  Go to Botkube Cloud instances page and click `Manage` button of the instance you want to remove.

2.  Click `Delete instance` button, type instance name in the popup and click `Delete instance`.

caution

Remember to execute the displayed command to completely remove Botkube and related resources from your cluster.

![Image 7: Delete](https://docs.botkube.io/assets/images/webhook_instance_delete-27fe3622760a4cbbd7c92d13d7ddcd41.png)


[Previous Outgoing webhook for self-hosted Botkube](https://docs.botkube.io/installation/webhook/self-hosted)[Next Overview](https://docs.botkube.io/examples-and-tutorials/)
