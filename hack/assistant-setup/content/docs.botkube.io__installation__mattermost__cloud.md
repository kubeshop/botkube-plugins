Title: Mattermost for Botkube Cloud | Botkube

URL Source: https://docs.botkube.io/installation/mattermost/cloud

Markdown Content:
*   [](https://docs.botkube.io/)
*   [Installation](https://docs.botkube.io/)
*   [Mattermost](https://docs.botkube.io/installation/mattermost/)
*   Mattermost for Botkube Cloud

Version: 1.12

Prerequisites[​](https://docs.botkube.io/installation/mattermost/cloud/#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------------------------------------------------------------

*   Botkube Cloud account which you can create [here](https://app.botkube.io/) for free.

Create a Botkube Cloud Instance with Mattermost[​](https://docs.botkube.io/installation/mattermost/cloud/#create-a-botkube-cloud-instance-with-mattermost "Direct link to Create a Botkube Cloud Instance with Mattermost")
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Follow the steps below to install Botkube in your Mattermost instance.

1.  Go to Botkube Cloud [Web App](https://app.botkube.io/) and create a new instance.
    
    You can do it by clicking "Create an Instance" button on Home Page or under this link [Create an Instance](https://app.botkube.io/instances/add)
    
2.  Fill in the `Instance Display Name` and click `Next` button.
    
    ![Image 1: Instance Display Name](https://docs.botkube.io/assets/images/mattermost_instance_display_name-b35605d19eef1ecc93de54d6eefacae5.png)
    
3.  Click `Add platform` dropdown, and select `Mattermost` option. ![Image 2: Mattermost Platform Select](https://docs.botkube.io/assets/images/mm_platform_select-aac36ca4e34549bef88cc00b3603f4ac.png)
    
4.  Follow the [Mattermost instructions](https://developers.mattermost.com/integrate/reference/bot-accounts/) for creating a bot account. When creating the bot account, use the following details:
    
    *   Username — `Botkube`
        
        note
        
        You can also use a custom username for your bot. Just remember that you'll need to provide this username during a later step of the Botkube installation.
        
    *   Description — `Botkube helps you monitor your Kubernetes cluster, debug critical deployments and gives recommendations for standard practices by running checks on the Kubernetes resources.`.
        
    *   Icon — You can download the Botkube icon from [this link](https://github.com/kubeshop/botkube/blob/main/branding/logos/botkube-black-192x192.png).
        
5.  Paste the bot name in the form
    
    ![Image 3: Bot Name in the form](https://docs.botkube.io/assets/images/mm_form_bot_name-987d8023e3f95e038072afa00b124eef.png)
    
6.  Past the token in the form
    
    ![Image 4: Personal Token in the form](https://docs.botkube.io/assets/images/mm_personal_token_form-1629ebe4b91aef0765809c70b619cba7.png)
    
7.  Add Botkube to a channel
    
    Make sure that the newly created bot account is added to your Mattermost team by following [these instructions](https://developers.mattermost.com/integrate/reference/bot-accounts/#bot-account-creation).
    
    ![Image 5: Invite Bot Account](https://docs.botkube.io/assets/images/invite-93908b3daf15ba3c0b87ab8522107fe6.png)
    
    Add Botkube user created to the channel you want to receive notifications in.
    
    ![Image 6: Channels in the form](https://docs.botkube.io/assets/images/mm_channels_form-2bcedf3a15e879c1c0fb01b22e95edc7.png)
    
8.  Add plugins you want to enable in your Botkube instance and click `Next` button.
    
    ![Image 7: Plugins](https://docs.botkube.io/assets/images/mm_add_plugins-a9d627e9575fd90aa56676a8d809c700.png)
    
9.  Include optional `default aliases` and `default actions` and click `Create` button to create Botkube Cloud instance.
    
    ![Image 8: Create](https://docs.botkube.io/assets/images/mm_create-069ec4b9176f0f58d424e665fa4b2472.png)
    
10.  Follow the instructions in the summary page to deploy Botkube into your environment.
    

![Image 9: Summary](https://docs.botkube.io/assets/images/mm_summary-bfdc3ff0af6735b41a17d7219fd6b6f0.png)

Clean up[​](https://docs.botkube.io/installation/mattermost/cloud/#clean-up "Direct link to Clean up")
------------------------------------------------------------------------------------------------------

### Remove Botkube from Mattermost Team[​](https://docs.botkube.io/installation/mattermost/cloud/#remove-botkube-from-mattermost-team "Direct link to Remove Botkube from Mattermost Team")

*   Deactivate or remove Botkube user from Mattermost Team. Login as System Admin, in the Menu proceed to System console -> Users -> botkube -> Deactivate.
*   Archive Channel created for Botkube communication if required.

### Remove Botkube from Kubernetes cluster[​](https://docs.botkube.io/installation/mattermost/cloud/#remove-botkube-from-kubernetes-cluster "Direct link to Remove Botkube from Kubernetes cluster")

1.  Go to Botkube Cloud instances page and click `Manage` button of the instance you want to remove.
    
2.  Click `Delete instance` button, type instance name in the popup and click `Delete instance`.
    
    caution
    
    Remember to execute the displayed command to completely remove Botkube and related resources from your cluster.
    
    ![Image 10: Delete](https://docs.botkube.io/assets/images/mm_instance_delete-27fe3622760a4cbbd7c92d13d7ddcd41.png)
    

[Previous Mattermost for self-hosted Botkube](https://docs.botkube.io/installation/mattermost/self-hosted)[Next Discord](https://docs.botkube.io/installation/discord/)
