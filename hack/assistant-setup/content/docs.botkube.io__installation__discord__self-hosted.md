Title: Discord for self-hosted Botkube | Botkube

URL Source: https://docs.botkube.io/installation/discord/self-hosted

Markdown Content:
Version: 1.12

Prerequisites[​](#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------

*   Botkube CLI installed according to the [Getting Started guide](https://docs.botkube.io/cli/getting-started#installation)
*   Access to Kubernetes cluster
*   Discord server account

Install Botkube to the Discord Server[​](#install-botkube-to-the-discord-server "Direct link to Install Botkube to the Discord Server")
---------------------------------------------------------------------------------------------------------------------------------------

Follow the steps below to install Botkube Discord app to your Discord server.

### Create Botkube app at your Discord Server[​](#create-botkube-app-at-your-discord-server "Direct link to Create Botkube app at your Discord Server")

1.  Reach [https://discordapp.com/developers/applications](https://discordapp.com/developers/applications).
    
    ![Image 1: discord_applications_portal](https://docs.botkube.io/assets/images/discord_applications_portal-a4e1b45cb3df4a271cbd599ec9f3b7ab.png)
    
2.  Create a "New Application" named Botkube and add a bot named **Botkube** into the Application.
    
    ![Image 2: discord_create_new](https://docs.botkube.io/assets/images/discord_create_new-ba9152ffe6f7be4f64af374d836c7062.png)
    
3.  Copy the Application **APPLICATION ID** and export it as the `DISCORD_BOT_ID` environment variable.
    
    ```
    export DISCORD_BOT_ID={APPLICATION_ID}
    ```
    
    ![Image 3: discord_copy_client_id](https://docs.botkube.io/assets/images/discord_copy_application_id-bf48ff3b0d9dc613c35d92dc287bd305.png)
    
4.  Add a description - `Botkube helps you monitor your Kubernetes cluster, debug critical deployments and gives recommendations for standard practices by running checks on the Kubernetes resources.`.
    
    Set the Botkube icon (Botkube icon can be downloaded from [this link](https://github.com/kubeshop/botkube/blob/main/branding/logos/botkube-color-192x192.png)).
    
    Click on Save Changes to update the Bot.
    
5.  Now, reach the **Bot** page and Click **Add Bot** to add a Discord Bot to your application.
    
    ![Image 4: discord_add_bot](https://docs.botkube.io/assets/images/discord_add_bot-867c43f73a079d08996072d3261d2fbc.png)
    
6.  After Bot creation, now you can see a bot is added to your application. Click on the **Reset Token** button.
    
    ![Image 5: discord_bot_created](https://docs.botkube.io/assets/images/discord_bot_created-845172424d2066002bff223d9a3afd36.png)
    
7.  Copy the Token and export it as the `DISCORD_TOKEN` environment variable.
    
    ```
    export DISCORD_TOKEN={TOKEN}
    ```
    
8.  Go to the **OAuth2** page. Generate the URL with suitable permissions using the **OAuth2 URL Generator** available under the OAuth2 section to add bot to your Discord server.
    
    ![Image 6: discord_bot_scope](https://docs.botkube.io/assets/images/discord_bot_scope-4024ad9c61ca1bab846e9181580bcd70.png)
    
    the generated URL contains **YOUR\_CLIENT\_ID**, Scope and permission details.
    
    ```
    https://discord.com/api/oauth2/authorize?client_id={YOUR_CLIENT_ID}&permissions={SET_OF_PERMISSIONS}&scope=bot
    ```
    
9.  Copy and Paste the generated URL in a new tab, select the discord server to which you want to add the bot, click Continue and Authorize Bot addition.
    
    ![Image 7: discord_bot_auth](https://docs.botkube.io/assets/images/discord_bot_auth-54b4a2d05fe3c3a6125c0f2e77f0bc78.png)
    
    ![Image 8: discord_bot_auth_2](https://docs.botkube.io/assets/images/discord_bot_auth_2-3fd072cd239d9cf517cc7c4a6c11e313.png)
    
10.  Switch to the Discord app. Navigate to **User settings** and select **Advanced** tab.
    
    Enable the **Developer Mode**.
    
    ![Image 9: discord_developer_mode](https://docs.botkube.io/assets/images/discord_developer_mode-b11b51acab3db94669c26cd97cde6c50.png)
    
11.  Create a new channel or select an existing one and copy the **CHANNEL ID**.
    
    To get the channel ID, right-click on a channel you want to receive notification in and click on **Copy ID**.
    
    ![Image 10: discord_copy_channel_id.png](https://docs.botkube.io/assets/images/discord_copy_channel_id-53f4972f2c412426a9d8a27ffaa2737a.png)
    
    Copy the channel ID and export it as the `DISCORD_CHANNEL_ID` environment variable.
    
    ```
    export DISCORD_CHANNEL_ID={ID}
    ```
    
12.  Now, go ahead and install the Botkube backend on your Kubernetes cluster.
    

note

Follow the first 4 mins of this [Video Tutorial](https://youtu.be/8o25pRbXdFw) to understand the process visually.

### Install Botkube in Kubernetes cluster[​](#install-botkube-in-kubernetes-cluster "Direct link to Install Botkube in Kubernetes cluster")

To deploy Botkube agent in your cluster, run:

```
export CLUSTER_NAME={cluster_name}export ALLOW_KUBECTL={allow_kubectl}botkube install --version v1.12.0 \--set communications.default-group.discord.enabled=true \--set communications.default-group.discord.channels.default.id=${DISCORD_CHANNEL_ID} \--set communications.default-group.discord.botID=${DISCORD_BOT_ID} \--set communications.default-group.discord.token=${DISCORD_TOKEN} \--set settings.clusterName=${CLUSTER_NAME} \--set 'executors.k8s-default-tools.botkube/kubectl.enabled'=${ALLOW_KUBECTL}
```

where:

*   **DISCORD\_CHANNEL\_ID** is the channel name where @Botkube needs to send notifications,
*   **DISCORD\_BOT\_ID** is the Botkube Application Client ID,
*   **DISCORD\_TOKEN** is the Token you received after adding Botkube bot to your Discord Application,
*   **CLUSTER\_NAME** is the cluster name set in the incoming messages,
*   **ALLOW\_KUBECTL** set true to allow `kubectl` command execution by Botkube on the cluster.

Configuration syntax is explained [here](https://docs.botkube.io/self-hosted-configuration). All possible installation parameters are documented [here](https://docs.botkube.io/self-hosted-configuration/helm-chart-parameters).

Send `@Botkube ping` in the channel to see if Botkube is running and responding.

### Remove Botkube from Discord Server[​](#remove-botkube-from-discord-server "Direct link to Remove Botkube from Discord Server")

*   Go to Discord Developers Portal [Applications](https://discord.com/developers/applications) page,
*   Click on "Botkube" and click on "Delete App" button.

Remove Botkube from Kubernetes cluster[​](#remove-botkube-from-kubernetes-cluster "Direct link to Remove Botkube from Kubernetes cluster")
------------------------------------------------------------------------------------------------------------------------------------------

Execute the following command to completely remove Botkube and related resources from your cluster:
