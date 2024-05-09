Title: Botkube Cloud App for Slack | Botkube

URL Source: https://docs.botkube.io/installation/slack/cloud-slack

Markdown Content:
The Botkube Cloud App for Slack uses Botkube Cloud services to manage channels and route executor commands. This allows multi-cluster support without a need to create a dedicated Slack application for each cluster. Events and alerts are sent directly from your cluster to your Slack workspace for reliable, fast notifications.

Prerequisites[​](#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------

*   A Botkube Cloud account.
    
    You can try out the Botkube Cloud App for Slack for free by creating an account in the [Botkube Cloud app](https://app.botkube.io/).
    

Create a Botkube Cloud Instance with Cloud Slack[​](#create-a-botkube-cloud-instance-with-cloud-slack "Direct link to Create a Botkube Cloud Instance with Cloud Slack")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------

1.  Go to Botkube Cloud [Web App](https://app.botkube.io/) and click on `New Instance` button.
    
    ![Image 1: New Instance](https://docs.botkube.io/assets/images/cloud_slack_new_instance-65f57c9b6c2e30c7b6250b1ebebf306c.png)
    
2.  Install Botkube Agent on your Kubernetes cluster by following the instructions on the page.
    
    ![Image 2: Install Agent](https://docs.botkube.io/assets/images/cloud_slack_install-c457bdc2758930b79f90849e72f6ebf2.png)
    
3.  Click `Add platform` dropdown, and select `Slack` option.
    
    ![Image 3: Slack Platform Select](https://docs.botkube.io/assets/images/cloud_slack_select_slack-c32779447fc66ed5091a858d0c7e2e46.png)
    
4.  Click `Add to Slack` button to add Cloud Slack integration to your Slack workspace.
    
    ![Image 4: Add to Slack](https://docs.botkube.io/assets/images/cloud_slack_add_to_slack-29428c5a907db8aeb0bd91a1488b35a6.png)
    
5.  Click `Allow` to grant permission for Botkube app to access your Slack workspace.
    
    ![Image 5: Cloud Slack Grant](https://docs.botkube.io/assets/images/cloud_slack_grant-891b3b884c149cc8285622770dbdb140.png)
    
6.  Provide the Slack app details as described follows and click `Next` button.
    
    *   **Connected Slack Workspace:** Slack workspace that you granted permission in the previous step.
    *   **Display Name:** Display name of the Cloud Slack configuration.
    *   **Channels:** Slack channes where you can execute Botkube commands and receive notification.
    
    ![Image 6: Cloud Slack Workspace](https://docs.botkube.io/assets/images/cloud_slack_workspace_details-769c2e33f6d5b18a9c20071d671af71e.png)
    
7.  Add plugins you want to enable in your Botkube instance and click `Next` button.
    
    ![Image 7: Cloud Slack Plugins](https://docs.botkube.io/assets/images/cloud_slack_add_plugins-889f253b85edf139ee9c89e70400e28a.png)
    
8.  Include optional default command aliases and actions and click `Apply Changes` button to update Botkube Cloud instance.
    
    ![Image 8: Cloud Slack Create](https://docs.botkube.io/assets/images/cloud_slack_create-e87f2bc4d8da6a31a5c3b18c095735de.png)
    

Using Botkube Cloud App for Slack[​](#using-botkube-cloud-app-for-slack "Direct link to Using Botkube Cloud App for Slack")
---------------------------------------------------------------------------------------------------------------------------

You can start using Botkube Cloud App for Slack by typing `@Botkube cloud help` in the Slack channel you configured in one of the previous steps.

![Image 9: Cloud Slack Command Help](https://docs.botkube.io/assets/images/cloud_slack_command_help-806b8a1c8238a93638259241e41741e3.png)

### Listing Cloud Instances[​](#listing-cloud-instances "Direct link to Listing Cloud Instances")

You can list all the Botkube Cloud instances by typing `@Botkube cloud list instances` in the Slack channel, or click the button `List connected instances` in the help command response. Besides the instance `name`, `ID`, and `status` in the list response, you can also click the `Get details` button to go to instance details on Botkube Cloud Dashboard.

![Image 10: Cloud Slack List Instances](https://docs.botkube.io/assets/images/cloud_slack_command_list_instances-b9d9661ea10f591eb72f92a79430d9cf.png)

### Setting Default Instances[​](#setting-default-instances "Direct link to Setting Default Instances")

Once a Botkube command is executed, it will be handled on target Kubernetes cluster specified with `--cluster-name` flag. This is an optional flag, where if you have not specified it, Botkube Cloud will select the first instance. However, you can also achieve setting default instance with command `@Botkube cloud set default-instance {instance-id}`.

![Image 11: Cloud Slack Set Default Instances](https://docs.botkube.io/assets/images/cloud_slack_command_set_default-8bdb16f71ec8f0cb97d41967707477e4.png)

After this point, all of your commands will be executed on the default instance. Moreover, if you want to execute a command on all the target clusters, you can use `--all-clusters` flag.

![Image 12: Cloud Slack All Clusters](https://docs.botkube.io/assets/images/cloud_slack_command_all_clusters-ff8c8984bb8097d34b419a6cfddb7cd0.png)

### Setting Public Alias for Private Channels[​](#setting-public-alias-for-private-channels "Direct link to Setting Public Alias for Private Channels")

In order to maintain your confidentiality while using Botkube's plugins, you need to create a public alias for your private Slack channels. This alias will only be visible to Botkube Cloud administrators.

#### During the Botkube Invitation to Private Channels[​](#during-the-botkube-invitation-to-private-channels "Direct link to During the Botkube Invitation to Private Channels")

When you invite Botkube to a private channel, a prompt will appear to guide you through the process of creating a public alias.

#### For Existing Private Channels[​](#for-existing-private-channels "Direct link to For Existing Private Channels")

To update an existing alias for a private channel, or if Botkube is already integrated, simply use this command:

    @Botkube cloud set channel-alias <your-public-alias-name>

Cleanup[​](#cleanup "Direct link to Cleanup")
---------------------------------------------

1.  Go to Botkube Cloud instances page and click `Manage` button of the instance you want to remove.
    
    ![Image 13: Cloud Slack Instance Manage](https://docs.botkube.io/assets/images/cloud_slack_instance_list_manage-49d417014a51479f9513b83a7ca2c9a2.png)
    
2.  Click `Delete instance` button, type instance name in the popup and click `Delete instance`.
    
    caution
    
    Remember to execute the displayed command to completely remove Botkube and related resources from your cluster.
    
    ![Image 14: Cloud Slack Instance Delete](https://docs.botkube.io/assets/images/cloud_slack_instance_delete-27fe3622760a4cbbd7c92d13d7ddcd41.png)
