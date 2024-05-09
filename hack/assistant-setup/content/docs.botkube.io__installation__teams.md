Title: Microsoft Teams | Botkube

URL Source: https://docs.botkube.io/installation/teams/

Markdown Content:
info

The previous Microsoft Teams integration has been deprecated. If you need to use the legacy Microsoft Teams integration, see the [Botkube 1.5 Documentation](https://docs.botkube.io/1.5/installation/teams/). It is recommended to migrate to the new Microsoft Teams app per the below instructions.

Botkube Cloud Microsoft Teams App[​](#botkube-cloud-microsoft-teams-app "Direct link to Botkube Cloud Microsoft Teams App")
---------------------------------------------------------------------------------------------------------------------------

The Botkube Cloud Microsoft Teams app offers several advanced features:

*   Simplified installation into your Microsoft Teams workspace
*   Multi-cluster executor support with a single Microsoft Teams app
*   Manage Teams channels directly from the Botkube web app and ensure the Botkube bot is invited to the correct channels

The Botkube Cloud Microsoft Teams app uses Botkube's cloud services to manage channels and route source events and executor commands. Currently, it requires a manual side-loading of the app, but we are working on getting it listed in Microsoft AppSource.

You can directly try Botkube Cloud Microsoft Teams app for free by creating an account in the [Botkube Web App](https://app.botkube.io/). Follow the steps below to install the app.

Prerequisites[​](#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------

*   A Botkube Cloud account.
    
    You can try out the Botkube Cloud Microsoft Teams app for free by creating an account in the [Botkube Cloud app](https://app.botkube.io/).
    

Create a Botkube Cloud Instance with Microsoft Teams[​](#create-a-botkube-cloud-instance-with-microsoft-teams "Direct link to Create a Botkube Cloud Instance with Microsoft Teams")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

### Connect Botkube Cloud to your Kubernetes cluster[​](#connect-botkube-cloud-to-your-kubernetes-cluster "Direct link to Connect Botkube Cloud to your Kubernetes cluster")

1.  Go to Botkube Cloud [Web App](https://app.botkube.io/) and click on `New Instance` button.
    
    ![Image 1: New Instance](https://docs.botkube.io/assets/images/cloud_teams_new_instance-65f57c9b6c2e30c7b6250b1ebebf306c.png)
    
2.  Install Botkube Agent on your Kubernetes cluster by following the instructions on the page.
    
    ![Image 2: Install Agent](https://docs.botkube.io/assets/images/cloud_teams_install-c457bdc2758930b79f90849e72f6ebf2.png)
    
3.  Click `Add platform` dropdown, and select `Teams` option.
    
    ![Image 3: Teams Platform Select](https://docs.botkube.io/assets/images/cloud_teams_select_platform-682cdcc197682e6bdb909bdf4e6a5f80.png)
    

Proceed with the next section.

### Install Botkube app to your Microsoft Teams organization and add it to your team[​](#install-botkube-app-to-your-microsoft-teams-organization-and-add-it-to-your-team "Direct link to Install Botkube app to your Microsoft Teams organization and add it to your team")

1.  Download the Botkube Cloud Microsoft Teams app by clicking the `Download Botkube App for Teams` button.
    
    ![Image 4: Download Teams App](https://docs.botkube.io/assets/images/cloud_teams_download_app-d7196c1a1d06c798599e6d798558823b.png)
    
2.  Sideload the Botkube app to your Microsoft Teams organization via Teams Admin Center, following the [official documentation](https://learn.microsoft.com/en-us/microsoftteams/teams-custom-app-policies-and-settings#upload-a-custom-app-using-teams-admin-center).
    
    info
    
    This step requires administrator permissions on your Microsoft Teams organization. Sideloading app is needed only once for the whole organization.
    
    *   Ensure the Botkube app is allowed for the organization in the [Teams Admin Center](https://admin.teams.microsoft.com/policies/manage-apps)
    
    ![Image 5: Teams Admin Center](https://docs.botkube.io/assets/images/cloud_teams_admin_center-8b857abd2f5aef5416b820878bbaa2a7.png)
    
3.  Add the Botkube app to your team.
    
    1.  In your Microsoft Teams application, navigate to the **Apps** tab.
        
    2.  Select the **Built for your organization** tab.
        
    3.  On the Botkube app card, click on the **Add** button.
        
        ![Image 6: Add Botkube App](https://docs.botkube.io/assets/images/cloud_teams_add_app-559531c09a0c06d3c0af4f70dc6f741b.png)
        
    4.  Click the **Add to a team** button.
        
        ![Image 7: Add app to team](https://docs.botkube.io/assets/images/cloud_teams_botkube_app_add-472d56ad47f55d7f7782a09eeacf2552.png)
        
    5.  Select the team and default channel, where you'll get the welcome message.
        
        ![Image 8: Select a team](https://docs.botkube.io/assets/images/cloud_teams_select_team-3df2e3153f0d02101671bf04dc1376b2.png)
        
    6.  Click the **Set up a bot** button.
        

Once the Botkube app is added to your team, you'll receive a welcome message.

![Image 9: Botkube Cloud Microsoft Teams Welcome Message](https://docs.botkube.io/assets/images/cloud_teams_welcome_msg-ff190292b2cd49e74dc49c4ba286c6bb.png)

Proceed with the next section.

### Grant permissions for Botkube app[​](#grant-permissions-for-botkube-app "Direct link to Grant permissions for Botkube app")

info

This step requires administrator permissions on your Microsoft Teams organization. Granting permissions is needed only once for the whole organization.

1.  Click on the **Grant access** button.
    
2.  Select your Microsoft account.
    
    ![Image 10: Select account](https://docs.botkube.io/assets/images/cloud_teams_permissions_select_account-9d17c4cfaa2eb163cd114030675293c5.png)
    
3.  Click the **Accept** button.
    
    ![Image 11: Grant access](https://docs.botkube.io/assets/images/cloud_teams_permissions_accept-9579403a6317274d931ab5fb5b37e8a0.png)
    
4.  You will be redirected to the confirmation page.
    
    ![Image 12: Confirmation page](https://docs.botkube.io/assets/images/cloud_teams_permissions_success-05aa285e1c8e448c475932925a9ffb02.png)
    

Close the page and proceed with the next section.

### Connect your team to Botkube Cloud[​](#connect-your-team-to-botkube-cloud "Direct link to Connect your team to Botkube Cloud")

Go back to the Microsoft Teams app channel, where you received the welcome message.

1.  Click the **Connect to Botkube Cloud** button in the welcome message.
    
2.  Log in to Botkube Cloud, if you haven't already. Ensure that you selected the correct organization, where you want to connect your team.
    
3.  Click the **Connect** button.
    
    ![Image 13: Connect to Botkube Cloud](https://docs.botkube.io/assets/images/cloud_teams_team_connect-46d7192e61f8dcd220f7dd333e8cbaa3.png)
    
4.  You will see a confirmation page.
    
    ![Image 14: Confirmation page](https://docs.botkube.io/assets/images/cloud_teams_team_connect_success-2121d79ce97a8f66bd660bf852a32cff.png)
    

Close the page and proceed with the next section.

### Finalize Botkube Cloud instance configuration[​](#finalize-botkube-cloud-instance-configuration "Direct link to Finalize Botkube Cloud instance configuration")

Go back to the Botkube Cloud instance creation.

1.  In step 2, select your connected team and provide other details.
    
    *   **Connected Microsoft Teams team:** Teams team that you connected in the previous section.
    *   **Display Name:** Display name of the Microsoft Teams team configuration.
    *   **Channels:** Teams channels where you can execute Botkube commands and receive notification.
    
    ![Image 15: Botkube Cloud Instance Configuration](https://docs.botkube.io/assets/images/cloud_teams_config-c996765f7ed399d6ddd263bfd463a140.png)
    
2.  Add plugins you want to enable in your Botkube instance and click `Next` button.
    
    ![Image 16: Microsoft Teams Plugins](https://docs.botkube.io/assets/images/cloud_teams_add_plugins-0236355cf8424353758934a96ca81112.png)
    
3.  Include optional default command aliases and actions and click `Apply Changes` button to update Botkube Cloud instance.
    
    ![Image 17: Microsoft Teams Create](https://docs.botkube.io/assets/images/cloud_teams_create-c3ff681023bb64d1c92ecf2c85f112a3.png)
    

Using Botkube Cloud Microsoft Teams App[​](#using-botkube-cloud-microsoft-teams-app "Direct link to Using Botkube Cloud Microsoft Teams App")
---------------------------------------------------------------------------------------------------------------------------------------------

You can start using Botkube Cloud Microsoft Teams by typing `@Botkube cloud help` in one of the channels in the team you configured in one of the previous steps.

![Image 18: Botkube Cloud Microsoft Teams Command Help](https://docs.botkube.io/assets/images/cloud_teams_command_help-7b1a978b74f7bb08c7062b6a80bbea07.png)

### Listing Cloud Instances[​](#listing-cloud-instances "Direct link to Listing Cloud Instances")

You can list all your Botkube Cloud instances by typing `@Botkube cloud list` in the Microsoft Teams channel, or click the button `List connected instances` in the help command response. Besides the instance `name`, `ID`, and `status` in the list response, you can also click the `Get details` button to go to instance details on Botkube Cloud Dashboard.

![Image 19: Botkube Cloud Microsoft Teams List Instances](https://docs.botkube.io/assets/images/cloud_teams_list_instances-0079ca8c5f306a230342b447ef8f31cb.png)

### Setting Default Instance[​](#setting-default-instance "Direct link to Setting Default Instance")

Once a Botkube command is executed, it will be handled on target Kubernetes cluster specified with `--cluster-name` flag. This is an optional flag, where if you have not specified it, Botkube Cloud will select the first instance. However, you can also achieve setting default instance with command `@Botkube cloud set default-instance`.

![Image 20: Cloud Microsoft Teams Set Default Instances](https://docs.botkube.io/assets/images/cloud_teams_command_set_default-3b37b0496c660bc70a102c910234456e.png)

After this point, all of your commands will be executed on the default instance. Moreover, if you want to execute a command on all the target clusters, you can use `--all-clusters` flag.

![Image 21: Cloud Microsoft Teams All Clusters](https://docs.botkube.io/assets/images/cloud_teams_command_all_clusters-a5c8c474816eb96d277a3dec873d8351.png)

Cleanup[​](#cleanup "Direct link to Cleanup")
---------------------------------------------

1.  Go to Botkube Cloud instances page and click `Manage` button of the instance you want to remove.
    
    ![Image 22: Cloud Teams Instance Manage](https://docs.botkube.io/assets/images/cloud_teams_instance_list_manage-e7ccd6d8aaabb01576a4ba21cd2f722d.png)
    
2.  Click `Delete instance` button, type instance name in the popup and click `Delete instance`.
    
    caution
    
    Remember to execute the displayed command to completely remove Botkube and related resources from your cluster.
    
    ![Image 23: Cloud Teams Instance Delete](https://docs.botkube.io/assets/images/cloud_teams_instance_delete-27fe3622760a4cbbd7c92d13d7ddcd41.png)
    

Caveats[​](#caveats "Direct link to Caveats")
---------------------------------------------

### RBAC `ChannelName` mapping[​](#rbac-channelname-mapping "Direct link to rbac-channelname-mapping")

Same as other communication platforms, Botkube Cloud Microsoft Teams app supports RBAC along with [all mappings](https://docs.botkube.io/configuration/rbac#mapping-types). However, because of the Microsoft Teams API limitation, for the default team channel the `ChannelName` is always `General`, regardless of the actual localized channel name.

Limitations[​](#limitations "Direct link to Limitations")
---------------------------------------------------------

Botkube Cloud Microsoft Teams App currently doesn't support the following features yet:

*   Processing states from selected dropdowns, e.g., used for the `kubectl` command builder. In a result, the command builder is not available.
    
*   Adding 👀 and ✅ reactions to messages indicating process status.
    
    This seems to be a limitation of the Microsoft Teams platform, however we'll investigate if there is a workaround.
    
*   Sending messages visible only to specific users.
    
*   Replacing messages with new content, e.g., used for pagination. Currently, a new card is sent as a new message.
    
*   User mentions in messages. Instead, Botkube app uses plaintext mentions with first and last name.
