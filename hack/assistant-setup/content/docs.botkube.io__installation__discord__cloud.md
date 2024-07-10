Title: Discord for Botkube Cloud | Botkube

URL Source: https://docs.botkube.io/installation/discord/cloud

Markdown Content:
Discord for Botkube Cloud | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/installation/discord/cloud#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/installation/discord/cloud)
*   [1.12](https://docs.botkube.io/installation/discord/cloud)
*   [1.11](https://docs.botkube.io/1.11/installation/discord/cloud)
*   [1.10](https://docs.botkube.io/1.10/installation/discord/cloud)
*   [1.9](https://docs.botkube.io/1.9/installation/discord/cloud)
*   [1.8](https://docs.botkube.io/1.8/installation/discord/cloud)
*   [1.7](https://docs.botkube.io/1.7/installation/discord/cloud)
*   * * *
    
*   [All versions](https://docs.botkube.io/versions)

[GitHub](https://github.com/kubeshop/botkube)[Slack](https://join.botkube.io/)

Search

*   [Installation](https://docs.botkube.io/)
    
    *   [Slack](https://docs.botkube.io/installation/socketslack)
    *   [Slack](https://docs.botkube.io/installation/slack/)
        
    *   [Mattermost](https://docs.botkube.io/installation/mattermost/)
        
    *   [Discord](https://docs.botkube.io/installation/discord/)
        
        *   [Discord for self-hosted Botkube](https://docs.botkube.io/installation/discord/self-hosted)
        *   [Discord for Botkube Cloud](https://docs.botkube.io/installation/discord/cloud)
    *   [Microsoft Teams](https://docs.botkube.io/installation/teams/)
    *   [Elasticsearch](https://docs.botkube.io/installation/elasticsearch/)
        
    *   [PagerDuty](https://docs.botkube.io/installation/pagerduty/)
    *   [Outgoing webhook](https://docs.botkube.io/installation/webhook/)
        
*   [Tutorials and examples](https://docs.botkube.io/examples-and-tutorials/)
    
*   [Features](https://docs.botkube.io/features/event-notifications)
    
*   [Plugins](https://docs.botkube.io/plugins/)
    
*   [CLI](https://docs.botkube.io/cli/getting-started)
    
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
    
*   [Troubleshooting](https://docs.botkube.io/troubleshooting/common-problems)
    
*   [Architecture](https://docs.botkube.io/architecture/)
    

*   [](https://docs.botkube.io/)
*   [Installation](https://docs.botkube.io/)
*   [Discord](https://docs.botkube.io/installation/discord/)
*   Discord for Botkube Cloud

Version: 1.12

On this page

Discord for Botkube Cloud
=========================

Prerequisites[​](https://docs.botkube.io/installation/discord/cloud#prerequisites "Direct link to Prerequisites")
-----------------------------------------------------------------------------------------------------------------

*   Botkube Cloud account which you can create [here](https://app.botkube.io/) for free.

Create a Botkube Cloud Instance with Discord[​](https://docs.botkube.io/installation/discord/cloud#create-a-botkube-cloud-instance-with-discord "Direct link to Create a Botkube Cloud Instance with Discord")
--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

1.  Go to Botkube Cloud [Web App](https://app.botkube.io/) and create a new instance.
    
    You can do it by clicking "Create an Instance" button on Home Page or under this link [Create an Instance](https://app.botkube.io/instances/add)
    
2.  Fill in the `Instance Display Name` and click `Next` button.
    
    ![Image 3: Instance Display Name](https://docs.botkube.io/assets/images/discord_instance_display_name-b35605d19eef1ecc93de54d6eefacae5.png)
    
3.  Click `Add platform` dropdown, and select `Discord` option. ![Image 4: Select Platform](https://docs.botkube.io/assets/images/discord_platform_select-aac36ca4e34549bef88cc00b3603f4ac.png)
    
4.  Create Botkube app at your Discord Server
    
    Reach [https://discordapp.com/developers/applications](https://discordapp.com/developers/applications).
    
    ![Image 5: discord_applications_portal](https://docs.botkube.io/assets/images/discord_applications_portal-a4e1b45cb3df4a271cbd599ec9f3b7ab.png)
    
5.  Create a "New Application" named Botkube and add a bot named **Botkube** into the Application.
    
    ![Image 6: discord_create_new](https://docs.botkube.io/assets/images/discord_create_new-ba9152ffe6f7be4f64af374d836c7062.png)
    
6.  Copy the Application **APPLICATION ID**
    
    ![Image 7: discord_copy_client_id](https://docs.botkube.io/assets/images/discord_copy_application_id-bf48ff3b0d9dc613c35d92dc287bd305.png)
    
    and paste it in the `BotID` field in the form.
    
    ![Image 8: bot_id_form](https://docs.botkube.io/assets/images/discord_bot_id_form-a9a0d728ad26361d5454b3eac4af8838.png)
    
7.  Add a description - `Botkube helps you monitor your Kubernetes cluster, debug critical deployments and gives recommendations for standard practices by running checks on the Kubernetes resources.`.
    
    Set the Botkube icon (Botkube icon can be downloaded from [this link](https://github.com/kubeshop/botkube/blob/main/branding/logos/botkube-color-192x192.png)).
    
    Click on Save Changes to update the Bot.
    
8.  Now, reach the **Bot** page and Click **Add Bot** to add a Discord Bot to your application.
    
    ![Image 9: discord_add_bot](https://docs.botkube.io/assets/images/discord_add_bot-867c43f73a079d08996072d3261d2fbc.png)
    
9.  After Bot creation, now you can see a bot is added to your application. Click on the **Reset Token** button.
    
    ![Image 10: discord_bot_created](https://docs.botkube.io/assets/images/discord_bot_created-845172424d2066002bff223d9a3afd36.png)
    
10.  Copy the Token and paste it in `Token` field the form.
    
    ![Image 11: discord_token_form](https://docs.botkube.io/assets/images/discord_token_form-0885aae7fa9636d7f896aae91ee10cdf.png)
    
11.  Go to the **OAuth2** page. Generate the URL with suitable permissions using the **OAuth2 URL Generator** available under the OAuth2 section to add bot to your Discord server.
    

![Image 12: discord_bot_scope](https://docs.botkube.io/assets/images/discord_bot_scope-4024ad9c61ca1bab846e9181580bcd70.png)

the generated URL contains **YOUR\_CLIENT\_ID**, Scope and permission details.

```
https://discord.com/api/oauth2/authorize?client_id={YOUR_CLIENT_ID}&permissions={SET_OF_PERMISSIONS}&scope=bot
```

12.  Copy and Paste the generated URL in a new tab, select the discord server to which you want to add the bot, click Continue and Authorize Bot addition.

![Image 13: discord_bot_auth](https://docs.botkube.io/assets/images/discord_bot_auth-54b4a2d05fe3c3a6125c0f2e77f0bc78.png)

![Image 14: discord_bot_auth_2](https://docs.botkube.io/assets/images/discord_bot_auth_2-3fd072cd239d9cf517cc7c4a6c11e313.png)

13.  Switch to the Discord app. Navigate to **User settings** and select **Advanced** tab.
    
    Enable the **Developer Mode**.
    
    ![Image 15: discord_developer_mode](https://docs.botkube.io/assets/images/discord_developer_mode-b11b51acab3db94669c26cd97cde6c50.png)
    
14.  Create a new channel or select an existing one and copy the **CHANNEL ID**.
    
    To get the channel ID, right-click on a channel you want to receive notification in and click on **Copy ID**.
    
    ![Image 16: discord_copy_channel_id.png](https://docs.botkube.io/assets/images/discord_copy_channel_id-53f4972f2c412426a9d8a27ffaa2737a.png)
    
    Copy the channel ID and create it in `Channel ID` field in the form.
    
    ![Image 17: discord_channel_id_form](https://docs.botkube.io/assets/images/discord_channel_id_form-2606624f82107523e93fc964593f0196.png)
    
15.  Add plugins you want to enable in your Botkube instance and click `Next` button.
    
    ![Image 18: Plugins](https://docs.botkube.io/assets/images/discord_add_plugins-e4f6481639f30536b32e059681064548.png)
    
16.  Include optional `default aliases` and `default actions` and click `Create` button to create Botkube Cloud instance.
    
    ![Image 19: Create](https://docs.botkube.io/assets/images/discord_create-c00e7b9eab8f8fb4b4c117acc930522a.png)
    
17.  Follow the instructions in the summary page to deploy Botkube into your environment.
    
    ![Image 20: Summary](https://docs.botkube.io/assets/images/discord_summary-bfdc3ff0af6735b41a17d7219fd6b6f0.png)
    

Clean up[​](https://docs.botkube.io/installation/discord/cloud#clean-up "Direct link to Clean up")
--------------------------------------------------------------------------------------------------

### Remove Botkube from Discord Server[​](https://docs.botkube.io/installation/discord/cloud#remove-botkube-from-discord-server "Direct link to Remove Botkube from Discord Server")

*   Go to Discord Developers Portal [Applications](https://discord.com/developers/applications) page,
*   Click on "Botkube" and click on "Delete App" button.

### Remove Botkube from Kubernetes cluster[​](https://docs.botkube.io/installation/discord/cloud#remove-botkube-from-kubernetes-cluster "Direct link to Remove Botkube from Kubernetes cluster")

1.  Go to Botkube Cloud instances page and click `Manage` button of the instance you want to remove.
    
2.  Click `Delete instance` button, type instance name in the popup and click `Delete instance`.
    
    caution
    
    Remember to execute the displayed command to completely remove Botkube and related resources from your cluster.
    
    ![Image 21: Delete](https://docs.botkube.io/assets/images/discord_instance_delete-27fe3622760a4cbbd7c92d13d7ddcd41.png)
    

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/installation/discord/cloud.md)

[Previous Discord for self-hosted Botkube](https://docs.botkube.io/installation/discord/self-hosted)[Next Microsoft Teams](https://docs.botkube.io/installation/teams/)

*   [Prerequisites](https://docs.botkube.io/installation/discord/cloud#prerequisites)
*   [Create a Botkube Cloud Instance with Discord](https://docs.botkube.io/installation/discord/cloud#create-a-botkube-cloud-instance-with-discord)
*   [Clean up](https://docs.botkube.io/installation/discord/cloud#clean-up)
    *   [Remove Botkube from Discord Server](https://docs.botkube.io/installation/discord/cloud#remove-botkube-from-discord-server)
    *   [Remove Botkube from Kubernetes cluster](https://docs.botkube.io/installation/discord/cloud#remove-botkube-from-kubernetes-cluster)

Community

*   [Contribute](https://docs.botkube.io/community/contribute)
*   [GitHub](https://github.com/kubeshop/botkube)
*   [Community Slack](https://join.botkube.io/)
*   [Support](https://docs.botkube.io/support)

Legal

*   [License](https://docs.botkube.io/license)
*   [Privacy & Legal](https://botkube.io/privacy-policy)
*   [Telemetry](https://docs.botkube.io/telemetry)

Learn

*   [Installation](https://docs.botkube.io/)

Social

*   [Twitter](https://twitter.com/Botkube_io)

Copyright © 2024 Kubeshop, Inc. Built with Docusaurus.
