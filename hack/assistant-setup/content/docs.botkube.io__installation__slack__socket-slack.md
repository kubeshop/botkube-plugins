Title: Socket Slack App | Botkube

URL Source: https://docs.botkube.io/installation/slack/socket-slack

Markdown Content:
The Socket-mode app works with the open-source Botkube Agent and does not require an account or subscription.

Prerequisites[​](#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------

*   Botkube CLI installed according to the [Getting Started guide](https://docs.botkube.io/cli/getting-started#installation)
*   Access to Kubernetes cluster
*   Slack Workspace admin access

Install Socket Slack App in Your Slack workspace[​](#install-socket-slack-app-in-your-slack-workspace "Direct link to Install Socket Slack App in Your Slack workspace")
------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Botkube uses interactive messaging to provide better experience. Interactive messaging needs a Slack App with Socket Mode enabled and currently this is not suitable for Slack App Directory listing. For this reason, you need to create a Slack App in your own Slack workspace and use it for Botkube deployment.

danger

**Multi-cluster caveat:** The architecture of socket-based Slack apps has a limitation on the routing of executor commands. If you would like to use [Botkube executors](https://docs.botkube.io/configuration/executor/) (e.g. kubectl commands) and have multiple Kubernetes clusters, you need to create and install a Botkube Slack app for each cluster. This is required so that the Slack to Botkube connections go to the right place. We recommend you set the name of each app to reflect the cluster it will connect to in the next steps.

To learn more about the Slack Socket API limitation, see the [comment](https://github.com/slackapi/bolt-js/issues/1263#issuecomment-1006372826) in the official Slack bot framework repository.

The [Botkube Cloud App for Slack](#botkube-cloud-slack-app) does not have this limitation.

Follow the steps below to create and install Botkube Slack app to your Slack workspace.

### Create Slack app[​](#create-slack-app "Direct link to Create Slack app")

1.  Go to [Slack App console](https://api.slack.com/apps) to create an application.
    
2.  Click **Create New App** and select **From an app manifest** in the popup to create application from manifest.
    
    ![Image 1: Create App from Manifest](https://docs.botkube.io/assets/images/slack_add_app-716017a6575f393b49e2cd157c67fe48.png)
    
3.  Select a workspace where you want to create application and click **Next**.
    
    ![Image 2: Select Workspace](https://docs.botkube.io/assets/images/slack_select_workspace-082c3680c0653a819d556756493134d8.png)
    
4.  Select **YAML** tab, copy & paste one of the following manifests, and click **Next**, and then **Create**.
    

*   Public channels only
*   Private channels only
*   Public and private channels

    display_information:  name: Botkube  description: Botkube  background_color: "#a653a6"features:  bot_user:    display_name: Botkube    always_online: falseoauth_config:  scopes:    bot:      - channels:read      - app_mentions:read      - reactions:write      - chat:write      - files:write      - users:read # Remote configuration only: Used to get Real Name for audit reportingsettings:  event_subscriptions:    bot_events:      - app_mention  interactivity:    is_enabled: true  org_deploy_enabled: false  socket_mode_enabled: true  token_rotation_enabled: false

### Install Botkube to the Slack workspace[​](#install-botkube-to-the-slack-workspace "Direct link to Install Botkube to the Slack workspace")

Once the application is created, you will be redirected to application details page. Press the **Install your app** button, select the workspace and click **Allow to finish installation**.

![Image 3: Install Slack App](https://docs.botkube.io/assets/images/slack_install_app-0c2fea83804d9a29ffe593d491d699c5.png)

### Obtain Bot Token[​](#obtain-bot-token "Direct link to Obtain Bot Token")

Follow the steps to obtain the Bot Token:

1.  Select **OAuth & Permissions** section on the left sidebar. On this page you can copy the bot token which starts with `xoxb...`.
    
    ![Image 4: Retrieve Slack Bot Token](https://docs.botkube.io/assets/images/slack_retrieve_bot_token-98639453c7d18970dca8a4727a1c149e.png)
    
2.  Export Slack Bot Token as follows:
    
        export SLACK_API_BOT_TOKEN="{botToken}"
    

### Generate and obtain App-Level Token[​](#generate-and-obtain-app-level-token "Direct link to Generate and obtain App-Level Token")

Slack App with Socket Mode requires an App-Level Token for the websocket connection.

Follow the steps to generate an App-Level Token:

1.  Select **Basic Information** link from the left sidebar and scroll down to section **App-Level Token**. Click on the **Generate Token and Scopes** button.
    
2.  Enter a **Name**, select `connections:write` scope, and click **Generate**.
    
    ![Image 5: Generate App-Level Token](https://docs.botkube.io/assets/images/slack_generate_app_token-685ab59995edd5495a5f5cca626ae895.png)
    
    ![Image 6: Retrieve App-Level Token](https://docs.botkube.io/assets/images/slack_retrieve_app_token-512945b00d08d0fdcb7a25a09ec5a950.png)
    
3.  Copy **App-Level Token** and export it as follows:
    
        export SLACK_API_APP_TOKEN="${appToken}"
    

### Add Botkube user to a Slack channel[​](#add-botkube-user-to-a-slack-channel "Direct link to Add Botkube user to a Slack channel")

After installing Botkube app to your Slack workspace, you could see a new bot user with the name "Botkube" added in your workspace. Add that bot to a Slack channel you want to receive notification in. You can add it by inviting `@Botkube` in a channel.

Install Botkube in Kubernetes cluster[​](#install-botkube-in-kubernetes-cluster "Direct link to Install Botkube in Kubernetes cluster")
---------------------------------------------------------------------------------------------------------------------------------------

To deploy Botkube agent in your cluster, run:

    export CLUSTER_NAME={cluster_name}export ALLOW_KUBECTL={allow_kubectl}export SLACK_CHANNEL_NAME={channel_name}botkube install --version v1.10.0 \--set communications.default-group.socketSlack.enabled=true \--set communications.default-group.socketSlack.channels.default.name=${SLACK_CHANNEL_NAME} \--set communications.default-group.socketSlack.appToken=${SLACK_API_APP_TOKEN} \--set communications.default-group.socketSlack.botToken=${SLACK_API_BOT_TOKEN} \--set settings.clusterName=${CLUSTER_NAME} \--set 'executors.k8s-default-tools.botkube/kubectl.enabled'=${ALLOW_KUBECTL}

where:

*   **SLACK\_CHANNEL\_NAME** is the channel name where @Botkube is added
*   **SLACK\_API\_BOT\_TOKEN** is the Token you received after installing Botkube app to your Slack workspace
*   **SLACK\_API\_APP\_TOKEN** is the Token you received after installing Botkube app to your Slack workspace and generate in App-Level Token section
*   **CLUSTER\_NAME** is the cluster name set in the incoming messages
*   **ALLOW\_KUBECTL** set true to allow `kubectl` command execution by Botkube on the cluster.

Configuration syntax is explained [here](https://docs.botkube.io/configuration). All possible installation parameters are documented [here](https://docs.botkube.io/configuration/helm-chart-parameters).

Send `@Botkube ping` in the channel to see if Botkube is running and responding.

### Delete Botkube from Slack workspace[​](#delete-botkube-from-slack-workspace "Direct link to Delete Botkube from Slack workspace")

*   Go to the [Slack apps](https://api.slack.com/apps) page,
*   Click on "Botkube", scroll to bottom, and click on "Delete App" button.

Remove Botkube from Kubernetes cluster[​](#remove-botkube-from-kubernetes-cluster "Direct link to Remove Botkube from Kubernetes cluster")
------------------------------------------------------------------------------------------------------------------------------------------

Execute the following command to completely remove Botkube and related resources from your cluster:
