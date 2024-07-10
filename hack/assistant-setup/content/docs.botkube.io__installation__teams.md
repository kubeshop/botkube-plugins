Title: Microsoft Teams | Botkube

URL Source: https://docs.botkube.io/installation/teams/

Markdown Content:
Botkube app for Microsoft Teams[​](https://docs.botkube.io/installation/teams/#botkube-app-for-microsoft-teams "Direct link to Botkube app for Microsoft Teams")
----------------------------------------------------------------------------------------------------------------------------------------------------------------

The Botkube app for Microsoft Teams offers several advanced features:

*   Simplified installation into your Microsoft Teams workspace
*   Multi-cluster executor support with a single Microsoft Teams app
*   Manage Teams channels directly from the Botkube web app and ensure the Botkube bot is invited to the correct channels

The Botkube app for Microsoft Teams uses Botkube's cloud services to manage channels and route source events and executor commands. Currently, it requires a manual side-loading of the app, but we are working on getting it listed in Microsoft AppSource.

You can directly try Botkube app for Microsoft Teams for free by creating an account in the [Botkube Web App](https://app.botkube.io/). Follow the steps below to install the app.

Prerequisites[​](https://docs.botkube.io/installation/teams/#prerequisites "Direct link to Prerequisites")
----------------------------------------------------------------------------------------------------------

*   A Botkube Cloud account.
    
    You can try out the Botkube app for Microsoft Teams for free by creating an account in the [Botkube Cloud app](https://app.botkube.io/).
    

Create a Botkube Cloud Instance with Microsoft Teams[​](https://docs.botkube.io/installation/teams/#create-a-botkube-cloud-instance-with-microsoft-teams "Direct link to Create a Botkube Cloud Instance with Microsoft Teams")
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

### Connect Botkube Cloud to your Kubernetes cluster[​](https://docs.botkube.io/installation/teams/#connect-botkube-cloud-to-your-kubernetes-cluster "Direct link to Connect Botkube Cloud to your Kubernetes cluster")

1.  Go to Botkube Cloud [Web App](https://app.botkube.io/) and click on `New Instance` button.
    
    ![Image 1: New Instance](https://docs.botkube.io/assets/images/cloud_new_instance-65f57c9b6c2e30c7b6250b1ebebf306c.png)
    
2.  Install Botkube Agent on your Kubernetes cluster by following the instructions on the page.
    
    ![Image 2: Install Agent](https://docs.botkube.io/assets/images/cloud_install-c457bdc2758930b79f90849e72f6ebf2.png)
    
3.  Click `Add platform` dropdown, and select `Teams` option.
    
    ![Image 3: Teams Platform Select](https://docs.botkube.io/assets/images/cloud_teams_select_platform-682cdcc197682e6bdb909bdf4e6a5f80.png)
    

Proceed with the next section.

### Install Botkube app to your team[​](https://docs.botkube.io/installation/teams/#install-botkube-app-to-your-team "Direct link to Install Botkube app to your team")

1.  Install Botkube app for Microsoft Teams from the official app catalog by clicking the `Install Botkube App for Teams` button.
    
    ![Image 4: Install Botkube App for Microsoft Teams](https://docs.botkube.io/assets/images/cloud_teams_download_app-ea7e70df0d9de4e294ab0aa8c85f6a4a.png)
    
    Alternatively, you can search for the "Botkube" app in the Microsoft Teams app catalog and click the **Add** button.
    
    ![Image 5: Add Botkube App](https://docs.botkube.io/assets/images/cloud_teams_add_app-64ae4b65439c1e95757ff7e23c9456e6.png)
    
2.  Add the Botkube app to your team.
    
    1.  Click the **Add to a team** button.
        
        ![Image 6: Add app to team](https://docs.botkube.io/assets/images/cloud_teams_botkube_app_add-ab8ab0dc3113336a67d31d5c5bcb7d4a.png)
        
    2.  Select the team and default channel, where you'll get the welcome message.
        
        ![Image 7: Select a team](https://docs.botkube.io/assets/images/cloud_teams_select_team-3df2e3153f0d02101671bf04dc1376b2.png)
        
    3.  Click the **Set up a bot** button.
        

Once the Botkube app is added to your team, you'll receive a welcome message.

![Image 8: Botkube Cloud Microsoft Teams Welcome Message](https://docs.botkube.io/assets/images/cloud_teams_welcome_msg-8adcff99243ea8fd517a4588ba85695b.png)

Proceed with the next section.

### Grant permissions for Botkube app[​](https://docs.botkube.io/installation/teams/#grant-permissions-for-botkube-app "Direct link to Grant permissions for Botkube app")

info

This step requires administrator permissions on your Microsoft Teams organization. Granting permissions is needed only once for the whole organization.

1.  Click on the **Grant access** button.
    
2.  Select your Microsoft account.
    
    ![Image 9: Select account](https://docs.botkube.io/assets/images/cloud_teams_permissions_select_account-9d17c4cfaa2eb163cd114030675293c5.png)
    
3.  Click the **Accept** button.
    
    ![Image 10: Grant access](https://docs.botkube.io/assets/images/cloud_teams_permissions_accept-ff8f1bcdada86ac6a58ed9ffcad78880.png)
    
4.  You will be redirected to the confirmation page.
    
    ![Image 11: Confirmation page](https://docs.botkube.io/assets/images/cloud_teams_permissions_success-74cacd791e1d052fd96ac1dd1ae9dae1.png)
    

Close the page and proceed with the next section.

### Connect your team to Botkube Cloud[​](https://docs.botkube.io/installation/teams/#connect-your-team-to-botkube-cloud "Direct link to Connect your team to Botkube Cloud")

Go back to the Microsoft Teams app channel, where you received the welcome message.

1.  Click the **Connect to Botkube Cloud** button in the welcome message.
    
2.  Log in to Botkube Cloud, if you haven't already. Ensure that you selected the correct organization, where you want to connect your team.
    
3.  Click the **Connect** button.
    
    ![Image 12: Connect to Botkube Cloud](https://docs.botkube.io/assets/images/cloud_teams_team_connect-26c492023cb739b0734ba7b9b9d9dba9.png)
    
4.  You will see a confirmation page.
    
    ![Image 13: Confirmation page](https://docs.botkube.io/assets/images/cloud_teams_team_connect_success-255cd1a9e0f6ccd8cd869d76295ae713.png)
    

Close the page and proceed with the next section.

### Finalize Botkube Cloud instance configuration[​](https://docs.botkube.io/installation/teams/#finalize-botkube-cloud-instance-configuration "Direct link to Finalize Botkube Cloud instance configuration")

Go back to the Botkube Cloud instance creation.

1.  In step 2, select your connected team and provide other details.
    
    *   **Connected Microsoft Teams team:** Teams team that you connected in the previous section.
    *   **Display Name:** Display name of the Microsoft Teams team configuration.
    *   **Channels:** Teams channels where you can execute Botkube commands and receive notification.
    
    ![Image 14: Botkube Cloud Instance Configuration](https://docs.botkube.io/assets/images/cloud_teams_config-c996765f7ed399d6ddd263bfd463a140.png)
    
2.  Add plugins you want to enable in your Botkube instance and click `Next` button.
    
    ![Image 15: Microsoft Teams Plugins](https://docs.botkube.io/assets/images/cloud_teams_add_plugins-0236355cf8424353758934a96ca81112.png)
    
3.  Include optional default command aliases and actions and click `Apply Changes` button to update Botkube Cloud instance.
    
    ![Image 16: Microsoft Teams Create](https://docs.botkube.io/assets/images/cloud_teams_create-c3ff681023bb64d1c92ecf2c85f112a3.png)
    

Using Botkube app for Microsoft Teams[​](https://docs.botkube.io/installation/teams/#using-botkube-app-for-microsoft-teams "Direct link to Using Botkube app for Microsoft Teams")
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

You can start using Botkube Cloud Microsoft Teams by typing `@Botkube cloud help` in one of the channels in the team you configured in one of the previous steps.

![Image 17: Botkube Cloud Microsoft Teams Command Help](https://docs.botkube.io/assets/images/cloud_teams_command_help-7b1a978b74f7bb08c7062b6a80bbea07.png)

### Listing Cloud Instances[​](https://docs.botkube.io/installation/teams/#listing-cloud-instances "Direct link to Listing Cloud Instances")

You can list all your Botkube Cloud instances by typing `@Botkube cloud list` in the Microsoft Teams channel, or click the button `List connected instances` in the help command response. Besides the instance `name`, `ID`, and `status` in the list response, you can also click the `Get details` button to go to instance details on Botkube Cloud Dashboard.

![Image 18: Botkube Cloud Microsoft Teams List Instances](https://docs.botkube.io/assets/images/cloud_teams_list_instances-0079ca8c5f306a230342b447ef8f31cb.png)

### Setting Default Instance[​](https://docs.botkube.io/installation/teams/#setting-default-instance "Direct link to Setting Default Instance")

Once a Botkube command is executed, it will be handled on target Kubernetes cluster specified with `--cluster-name` flag. This is an optional flag, where if you have not specified it, Botkube Cloud will select the first instance. However, you can also achieve setting default instance with command `@Botkube cloud set default-instance`.

![Image 19: Cloud Microsoft Teams Set Default Instances](https://docs.botkube.io/assets/images/cloud_teams_command_set_default-3b37b0496c660bc70a102c910234456e.png)

After this point, all of your commands will be executed on the default instance. Moreover, if you want to execute a command on all the target clusters, you can use `--all-clusters` flag.

![Image 20: Cloud Microsoft Teams All Clusters](https://docs.botkube.io/assets/images/cloud_teams_command_all_clusters-a5c8c474816eb96d277a3dec873d8351.png)

Cleanup[​](https://docs.botkube.io/installation/teams/#cleanup "Direct link to Cleanup")
----------------------------------------------------------------------------------------

1.  Go to Botkube Cloud instances page and click `Manage` button of the instance you want to remove.
    
    ![Image 21: Cloud Teams Instance Manage](https://docs.botkube.io/assets/images/cloud_teams_instance_list_manage-e7ccd6d8aaabb01576a4ba21cd2f722d.png)
    
2.  Click `Delete instance` button, type instance name in the popup and click `Delete instance`.
    
    caution
    
    Remember to execute the displayed command to completely remove Botkube and related resources from your cluster.
    
    ![Image 22: Cloud Teams Instance Delete](https://docs.botkube.io/assets/images/cloud_instance_delete-27fe3622760a4cbbd7c92d13d7ddcd41.png)
    

Caveats[​](https://docs.botkube.io/installation/teams/#caveats "Direct link to Caveats")
----------------------------------------------------------------------------------------

### RBAC `ChannelName` mapping[​](https://docs.botkube.io/installation/teams/#rbac-channelname-mapping "Direct link to rbac-channelname-mapping")

Same as other communication platforms, Botkube app for Microsoft Teams supports RBAC along with [all mappings](https://docs.botkube.io/features/rbac#mapping-types). However, because of the Microsoft Teams API limitation, for the default team channel the `ChannelName` is always `General`, regardless of the actual localized channel name.

Limitations[​](https://docs.botkube.io/installation/teams/#limitations "Direct link to Limitations")
----------------------------------------------------------------------------------------------------

Botkube app for Microsoft Teams currently doesn't support the following features yet:

*   Processing states from selected dropdowns, e.g., used for the `kubectl` command builder. In a result, the command builder is not available.
    
*   Adding 👀 and ✅ reactions to messages indicating process status.
    
    This seems to be a limitation of the Microsoft Teams platform, however we'll investigate if there is a workaround.
    
*   Sending messages visible only to specific users.
    
*   Replacing messages with new content, e.g., used for pagination. Currently, a new card is sent as a new message.
    
*   User mentions in messages. Instead, Botkube app uses plaintext mentions with first and last name.
