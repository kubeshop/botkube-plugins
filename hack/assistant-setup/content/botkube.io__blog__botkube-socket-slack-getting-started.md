Title: Getting Started with the New Botkube App for Slack

URL Source: https://botkube.io/blog/botkube-socket-slack-getting-started

Published Time: Dec 05, 2022

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/636df3edbf5389368f6bef9c_cYbM1beBC5tQnSPVfaXCg_W9tkHugByZV2TOleN6pTw.jpeg)

Blair Rampling

Product Leader

Botkube

The new Botkube Slack app provides more great interactive features and better security when compared to the legacy Botkube Slack app.

### Table of Contents

*   [Installing the Slack App in your workspace](https://botkube.io/blog/botkube-socket-slack-getting-started#installing-the-slack-app-in-your-workspace)
*   [Botkube Slack App Tokens](https://botkube.io/blog/botkube-socket-slack-getting-started#botkube-slack-app-tokens)
*   [Invite Botkube to a channel](https://botkube.io/blog/botkube-socket-slack-getting-started#invite-botkube-to-a-channel)
*   [Installing Botkube](https://botkube.io/blog/botkube-socket-slack-getting-started#installing-botkube)
*   [Explore the New Slack App](https://botkube.io/blog/botkube-socket-slack-getting-started#explore-the-new-slack-app)
*   [What Could Go Wrong?](https://botkube.io/blog/botkube-socket-slack-getting-started#what-could-go-wrong-)
*   [Feedback - We’d Love to Hear From You](https://botkube.io/blog/botkube-socket-slack-getting-started#feedback-we-d-love-to-hear-from-you)

#### Start Using Botkube AI-Powered Assistant Today

The [new Botkube Slack app](https://docs.botkube.io/installation/slack/) provides more great interactive features and better security when compared to the [legacy Botkube Slack app](https://docs.botkube.io/0.15/installation/slack/). We announced the new socket mode Slack app in the [Botkube v0.14.0 release notes](https://botkube.io/blog/botkube-v014-release-notes#new-botkube-slack-app). The new Slack app has some specific requirements and a new installation process, so let's have a look at how to get started with the most modern [ChatOps tool for Kubernetes!](http://botkube.io/) You can also use the Botkube [installation documentation](https://docs.botkube.io/installation/slack/) to get started, but this post is to give you more context about the changes to the new app and some caveats to watch out for.

This blog post assumes that we're starting a completely new Botkube installation.

Requirements:

*   A Slack workspace where you have permission to install apps and create channels
    
*   A Kubernetes cluster where you have access to install Botkube
    
*   Working kubectl and helm installation
    

Multi-cluster caveat: The architecture of socket-based Slack apps is different from the older Slack app. If you are using Botkube executors (e.g. kubectl commands) and have multiple Kubernetes clusters, you need to install a Botkube Slack app for each cluster. This is required so that the Slack to Botkube connections go to the right place. We recommend you set the name of each app to reflect the cluster it will connect to in the next steps.

Installing the Slack App in your workspace
------------------------------------------

Let's go through the initial steps for installing the socket-based Slack app itself in your Slack workspace.

1.  Open the [Slack App Console](https://api.slack.com/apps) and click Create an App or Create New App.
    
2.  Click "From an app manifest" as we will use the Botkube-provided manifest for the new app.
    

![Image 2: Kubernetes Deployment Setup wizard](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd5bd91dbba0faccc21_9xHsAYF1sDUzf92VzItiVmpg6pPfW8mLBma5StF_-j8ZfzyqSgBnPbqpPnTLGLzn2nCGy_wEPBw9E5dpwbZBNXbKm0jSWixj-ufmFk-7JtqPw5TXhXO4a4LpwXLVJgh2pRe2U-bx0KhPJoUCB5EkdhT20QMTg4VVQvty5IqPCG3ITKmfL_f1nNdZQ_hlcQ.png)

3.  Select the workspace into which you want to install the app and click Next.
    
4.  You can find the latest app manifests [in the Botkube documentation](https://docs.botkube.io/installation/slack/#create-slack-app). There are three different YAML manifests available. The manifests vary based on whether you want to use Botkube in a public Slack channel, a private channel, or both. This is one of the new improved security features that limits the scope of the token that Botkube uses. This is an example of the manifest for both public and private channels:
    

```
display_information:
 name: Botkube
 description: Botkube
 background_color: "#a653a6"
features:
 bot_user:
   display_name: Botkube
   always_online: false
oauth_config:
 scopes:
   bot:
     - channels:read
     - groups:read
     - app_mentions:read
     - chat:write
     - files:write
settings:
 event_subscriptions:
   bot_events:
     - app_mention
 interactivity:
   is_enabled: true
 org_deploy_enabled: false
 socket_mode_enabled: true
 token_rotation_enabled: false
```

If you are using Botkube with multiple Kubernetes clusters as mentioned in the earlier caveat, set the name and description fields to something specific to the cluster this app will manage. This will make it clear in Slack itself which bot maps to which cluster. Paste this YAML into the app manifest dialog and click Next.

![Image 3: Kubernetes Manifest file for Botkube](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd59be2c7f0b4231ab2_rT70Rd59URrbNTS0SF4iTHlKyh9YkqZylD9a_vi_umpJJghbKTkrlq4dc1SUb1JmLa-TvD-Zv0fgQZiAdNgv--_ptj2Jf7vwNDxX20KMin3O80Xrjyl1BzDCM8Sfe2MXizEODQqKb1Cuz-Rlz_69CzhYKRhFr12eYAXezbr5YnRkjw_k6VUft--i1led9Q.png)

5.  You now have a chance to review the app settings. Click the Create button.
    
6.  Click the Install to Workspace button to add this newly created app to the Slack Workspace.
    

![Image 4: Slack Webhook API for building Slack Bots](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd53a3620a880c1429d_2lBQzqtvc7o-h5cHB7Hp-qyFgGwlqldNSxf1lHVyYUDUwGBXgpmBEPHOH5xj_bEQnudTJx6d1agQgl6EZoptfPbrgxwWU3hwn-C-0-23ImODUiph_cvEFxMcMPEjewJEeL8Gvj-NNj_ysyTstwz8fFB5jtyGIy6-dmpTM1v8D097Tx6VQ12Q-mumwzFGhA.png)

7.  Review the permissions and click Allow. The app is now installed in your Slack workspace.

![Image 5: Allow Botkube permissions into Group Channels](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd5d46fe3cc1ec655d1_kR2N0u3ZQlfHarVS2aC5lfyerRXjRmE25noMdLH2yp3l1Rd8gNeplU-JbQ4aak4LlhHk-H7-G3nxO8CXJyD5eH0mhy5-69gJsjwSUJaWeHijGc94iFse5AYCKYEG9-fxZ5Q57C9JUcPTfRsqiiPpBzU4Bc_7MlJo2lzKrnBhtudCygDS0M35okbc39reEg.png)

Botkube Slack App Tokens
------------------------

Now that the Slack app is installed in your workspace, you can proceed to installing Botkube in your Kubernetes cluster. First, you need two tokens, a bot token and an app-level token. This is another additional security measure in the new Slack app. You can see details on these token types in the [Slack documentation](https://api.slack.com/authentication/token-types). You will use these tokens in the Botkube installation process.

1.  From the Slack app page where you left off after installing the app in your workspace, select OAuth & Permissions in the left navigation bar.
    
2.  Copy the Bot User OAuth Token shown here.
    

![Image 6: Oauth Tokens for Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd57f6d2658565c2802_EbTpqUY7Kxtxi7pDnIVQREtoVR7lLKJaHYlFSo5Ald7PPzcwWuoNxhTIfjSu5b8on04-YKAtxwuwBPSAlO-MBAGNkM9RsjYLb-D9prje8FoOORshIpi3AUvnyPP9aFj08weq-EjK0bMLtb_II93-zftUqceLGV3SropKhaNaRckmI7UHx3BMPSpyU0JhSQ.png)

3.  Click Basic Information in the left navigation bar.
    
4.  Scroll down to the App-Level Tokens section and click the Generate Token and Scopes button.
    
5.  Give the token an appropriate name and click Add Scope.
    
6.  Select the connections:write scope, then click the Generate button.
    

![Image 7: Generating app token for custom Slack Bot](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd57787389c7b4d6eac_Wl7cvk72vCB_r0giRxWLTQlmew8rmPFkFTZUaj7urPf2C4LM_YaVuO5XrDXOl9Ql58cOrnQT3LiDl4dcY0PjyIUWHivxBlIsTOsWdYX4un0JJaI8DIQfIxOfTHjVxD50Yk9QaLTQ_zlbcmhdZzjZN_1-4FSFZCYzO8k8N1uLxDQRcIXRZZIAMLFQJuG5iw.png)

7.  Copy the created token and click Done.

Invite Botkube to a channel
---------------------------

Now that everything is set up for the Slack app, you need to invite the Botkube bot to a Slack channel. This is the channel that you will use for notifications and command execution and is required for the Botkube configuration before it is installed in Kubernetes.

You should see a new Slack bot in your workspace under "Apps" with the name that you specified in the manifest. You can create a new Slack channel or use an existing one. Whether the channel is public or private, make sure that you used the correct manifest earlier when creating the app. If the wrong manifest was used you will see error messages in the Botkube logs about the scope.

Invite the bot to the channel. The easiest way to do this is to open the Slack channel and send a message to _@botkube bot name_. You will be prompted to add the bot to the channel, click the Add to Channel button.

![Image 8: Adding Botube to your cluster chat channel](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd5bd91db676aaccc2a_5FVGjg5gJZRQdEsdM2zokGiw_bH5m2tUJm100sy1unbpCAwupyN2hYkXLc21-emgPzST00zRg5srK_eQXKsJ_X8xnBRe4tHcyCnvH97uePNhnRt4TA2B6y0sqr-Um2-V1V9U-mJcpFvuPxo5LzH3kLBMnSbHWnycZ6BTSNb7e9issQ2kcAwLv6oue8ePUg.png)

Installing Botkube
------------------

We now have everything we need to configure and install Botkube. We will use helm to install it in a Kubernetes cluster. First, configure some environment variables that we will pass to the `helm install` command.

The Slack app bot token from earlier:

`export SLACK_API_BOT_TOKEN="xoxb-{botToken}"`

The Slack app-level token from earlier:

`export SLACK_API_APP_TOKEN="xapp-{appToken}"`

The cluster name that will be used when interacting with this instance of Botkube. Notifications and other interactive commands will use this name to reference this cluster:

`export CLUSTER_NAME={cluster_name}`

Set to true to allow execution of kubectl commands, false to make this bot instance for event notifications only:

`export ALLOW_KUBECTL={allow_kubectl}`

Specify the Slack channel name where you invited the bot earlier. Make sure you do not include the # in the channel name:

`export SLACK_CHANNEL_NAME={channel_name}`

Now that those five environment variables are set, make sure you have the Botkube helm repository configured:

`helm repo add botkube https://charts.botkube.io`

`helm repo update`

Now you can simply run the helm install command shown here, or copy from the [documentation](https://docs.botkube.io/installation/slack/#install-botkube-backend-in-kubernetes-cluster). Note that this installs version v0.16.0 of Botkube which is the latest version as of this article:

```
helm install --version v0.16.0 botkube --namespace botkube --create-namespace \

--set communications.default-group.socketSlack.enabled=true \

--set communications.default-group.socketSlack.channels.default.name=${SLACK_CHANNEL_NAME} \

--set communications.default-group.socketSlack.appToken=${SLACK_API_APP_TOKEN} \

--set communications.default-group.socketSlack.botToken=${SLACK_API_BOT_TOKEN} \

--set settings.clusterName=${CLUSTER_NAME} \

--set executors.kubectl-read-only.kubectl.enabled=${ALLOW_KUBECTL} \

botkube/botkube
```

If everything goes according to plans, you will see Botkube start up in the Slack channel and print the initial interactive help menu.

![Image 9: Kubernetes Alerts now actively alerting Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638defd53a3620ccfbc1429e_T_WAEBM055UezMigtRJmb49_MhQocCrlbgee7nW_ekbsIKvexG3KzQrjn9gUsJkbOYCGPhanm9mYriXQu0JaDrB4Pr7AsQdg3xQfLqLB6Qi1gLRay8uSRfJGke15DOmfBFWoCL9xGQTjBdi50CRU4z_thYxXVS3mASf-cbxluEV7EezHvV4sill1apGpEw.png)

Botkube is now ready to go!

Explore the New Slack App
-------------------------

The new Botkube Slack app adds some great features that weren't previously available, particularly around interactive event response. You saw the interactive help at the end of the installation process in the previous section. You can click the buttons in the interactive help to run the commands as shown. You can also make configuration changes interactively, try clicking the Adjust notifications button and you can set notification settings using a dialog rather than editing YAML.

Another interactive feature is the Run command option shown with events. Any time an event is sent to the channel, you can select a contextual command from the Run command drop-down box. The list of commands will be only those relevant to the resource type in the event. Selecting a command will use the kubectl executor to run the command against the specific resource in the event. This saves typing out the kubectl commands by hand if you want to quickly get some information about a resource, like `describe` a new resource or get `logs` on a pod with errors.

![Image 10: Run Kubernetes troubleshooting commands](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638dfeee173eea4859c13681_aeB3oRSq_LvBSkdLujcv2ls4WiN8M1W-Mo6KDzL2MmC7bQ487s0DSqXd5X4iejHKHgtD1abNQn6LEREE-J4U-TR5VEre1aATDi51rzTn2WQVogIE6qV6bFYRZe_2O7o7ZkH_5iUa5j5DcVMbzzo7mXKi3N2n04WcE7cqj5MEly6On-gg-fRNpxsmDAxn9A.png)

Another interactive feature can be found by running `@botkube kubectl` with no options. This starts the interactive command generator. This presents you with a list of verbs that you can run and as you make selections, automatically offers the contextually-relevant options in drop-down boxes until you have built a complete command. You can then click the Run command button and see the resulting output. You can also filter the output by specifying a match string in the Filter output box.

![Image 11: Pulling K8s Logs](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/638dfeee5c48734b1151fe69_VBgNCtn62GTB7Ee-xc2z5xMXs92IaE5blN1CZPVNeV18fpAF4dSTi_giYZyoA80T4BBTnajpiBnsl7kX-xlWPKTdjTb1-6zKxNSJCQwjsqA9HyYinugvHNE-B0oQKj0MkV-FsRNbuIvmJe77SdtuFJv6tR5WxKwmvLkMzxDMFlLCA6lCItB8uUtFHS7Lmg.png)

What Could Go Wrong?
--------------------

There are several steps in the process to look out for. These are where things could go wrong and you could end up with a Botkube installation that doesn't work. Always check the Botkube pod logs for details on why something might not be working, but here are a few things to look out for:

*   Slack App Manifest Scope: When using the YAML app manifest to create the Slack app, make sure you use the [appropriate version](https://docs.botkube.io/installation/slack/#create-slack-app) for public, private, or both Slack channel types.
    
*   Slack app-level token scope: The app-level token needs to have the `connections:write` scope assigned.
    
*   Multiple clusters: When using Botkube to execute commands in multiple Kubernetes clusters, you need to install a Slack app for each cluster to ensure the commands are routed to the correct one. If you are using Botkube for **notifications only** and have command execution disabled, you can get away with a single app for multiple clusters. If you ever do want to enable command executions though, you will need to revert to the Slack app per cluster model.
    
*   Incorrect tokens: Make sure your bot and app-level tokens are correct and match those in the Slack App Console.
    
*   Slack channel configuration: Make sure you omit the # when specifying the Slack channel that Botkube will use.
    

Feedback - We’d Love to Hear From You
-------------------------------------

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented, and get bugs fixed quickly.

You can talk to us in the Botkube [GitHub issues](https://github.com/kubeshop/botkube/issues), Botkube [Slack community](https://join.botkube.io/), or email our Product Leader at [blair@kubeshop.io](mailto:blair@kubeshop.io).
