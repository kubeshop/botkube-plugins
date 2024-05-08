Title: Streamlining Kubernetes Management: A Closer Look at Botkube ChatOps

URL Source: https://botkube.io/blog/streamlining-kubernetes-management-a-closer-look-at-botkube-chatops

Published Time: Jan 19, 2024

Markdown Content:
![Image 1: a man in a black shirt is standing in front of a circle with the words experts google cloud](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9dc3218bb2465f041cea7_1693115200398.jpeg)

Rohit Ghumare

Google Developer Expert – Google Cloud, CNCF Ambassador

Botkube

Learn about Botkube from Rohit Ghumare, Google Developer Expert – Google Cloud and CNCF Ambassador.

### Table of Contents

*   [Step 1: Signing Up for Botkube](#step-1-signing-up-for-botkube)
*   [Step 2: Setting Your Instance](#step-2-setting-your-instance-)
*   [Debugging Botkube](#debugging-botkube-)
*   [Botkube Meets Helm](#botkube-meets-helm)
*   [Conclusion](#conclusion)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

#### Get started with Botkube Cloud

Hey there, Kubernetes enthusiasts! Let's dive into the world of Botkube and its game-changing. Developed by the wizards at KubeShop, Botkube is not just another tool—it's a revolution in Kubernetes management.

If you're in DevOps, you know the headache of debugging production issues, especially when juggling massive Kubernetes clusters. The complexity can be overwhelming, right? This post will give you the lowdown on Botkube’s Helm magic and its impact on your Kubernetes adventures.

🔎 **Facing the Debugging Challenge:** Imagine dealing with a maze of applications relying on various system services. This scale makes it challenging to see what's going wrong when major issues crop up. Plus, there are always many tools to manage, and collaborating across different teams is only sometimes a walk in the park.

🛠️ **Botkube to the Rescue:** Enter Botkube! This nifty tool is like having a Swiss Army knife for Kubernetes troubleshooting. It's an app that makes monitoring, debugging, and managing Kubernetes clusters a breeze. It hooks up with your favorite chat tools, like Slack or Microsoft Teams, turning them into a command center. You get real-time alerts, can mess around with deployments, check logs, and even use kubectl and Helm from your chat window!

Ready to dive into the hands-on demo where we use Botkube for Kubernetes troubleshooting? Here's how to get started:

**Step 1: Signing Up for Botkube**
----------------------------------

*   **New Users:** If you're new to Botkube, welcome aboard! You can [sign up for a free](https://app.botkube.io/#first-login) account using your email or GitHub account. It's quick and easy.
*   **Existing Users:** Already part of the Botkube family? Just [log in](https://app.botkube.io/#first-login) with your existing credentials, and you're good to go.
*   You can expect your dashboard to be like this:

![Image 2: a screen shot of a website with a blue background](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d4f685f0a8da71eabce5_Screenshot%202024-01-13%20at%2018.33.18.png)

> **Linking Your Kubernetes Cluster to Slack:** Now that you're in, let's connect your Kubernetes cluster to Slack for real-time monitoring magic.

**Step 2: Setting Your Instance**
---------------------------------

*   Click “Create an Instance” and run the below commands in your terminal.
*   Homebrew is the easiest way to install the CLI tool. Alternatively, you can download the binary directly for your system or use helm commands for installations.

### **1\. Install the Botkube CLI**

Use [Homebrew](https://brew.sh/) to install the latest Botkube CLI:

`brew install kubeshop/botkube/botkube`

Alternatively, download the Botkube CLI binary and move it to a directory under your `$PATH:`

curl -Lo botkube https://github.com/kubeshop/botkube/releases/download/v1.7.0/botkube-darwin-arm64chmod +x botkube &amp;&amp; mv botkube /usr/local/bin/botkube


### **2\. Install Botkube on your cluster**

Install or upgrade Botkube and connect it to Botkube Cloud:

botkube install --version=v1.7.0 \ 
--set config.provider.endpoint=https://api.botkube.io/graphql \  
--set config.provider.identifier=b2b56b7d-392a-4614-b1ba-93eb0e92f424 \  
--set config.provider.apiKey=key:1d4ffcee-9df8-489c-bcd2-f288eb781dde


Now, Choose a display name for your Botkube instance. This is like giving a cool nickname to your new digital helper.

![Image 3: a screen showing the creation of a house in adobe adobe adobe ado](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d575d9deb7aaf9866988_Screenshot%202024-01-13%20at%2020.15.55.png)

### **Step 3: Choosing Your Integration Method**

*   Time to pick your integration method. As we know, Botkube provides multiple integrations, but I’m using Slack Integration here because I use it a lot for personal and community use.
*   Go for the "**Official Botkube Slack**" option. Heads up, though – this requires a 30-day free trial sign-up. It's like a test drive for the whole experience!
*   It also provides a free version if you want to set it up yourself.
*   You can also check the [official documentation](https://docs.botkube.io/) for Botkube Integration with other chat platforms.

### Comparison for Botkube bidirectional communication integrations:

![Image 4: a table showing the differences between azure and youtube integrations](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65aed164c12e1b6851dff536_Screenshot%202024-01-22%20at%203.34.16%20PM.png)

### **Step 4: Syncing with Slack**

*   Jump into your Slack app settings. Look for the "Add Slack" button and give it a click. This is where the real magic begins – integrating Botkube with your Slack world.

![Image 5: a screen shot of a page with a button on it](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d5b8de7e130a53622129_Screenshot%202024-01-13%20at%2020.22.38.png)

### **Step 5: Selecting Your Slack Workspace**

*   Head on to the Slack console and Create an app from an app manifest.

![Image 6: a screen shot of a web page showing the add apps button](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d5d3de7e130a5362396c_Screenshot%202024-01-13%20at%2021.00.59.png)

Pick a workspace of your choice where you want to install the Botkube

![Image 7: a screenshot of the pick workspace to develop your app](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d5f1035f985e43bc7225_Screenshot%202024-01-13%20at%2021.02.50.png)

You can use the manifest YAML file below for app creation.

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="botkube-app-creation.yaml"><tbody><tr><td id="file-botkube-app-creation-yaml-L1" data-line-number="1"></td><td id="file-botkube-app-creation-yaml-LC1"><span>display_information</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L2" data-line-number="2"></td><td id="file-botkube-app-creation-yaml-LC2"><span>name</span>: <span>devreler </span><span><span>#</span> name your app as you want it to be.</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L3" data-line-number="3"></td><td id="file-botkube-app-creation-yaml-LC3"><span>description</span>: <span>devrel as service </span><span><span>#</span> description for your app.</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L4" data-line-number="4"></td><td id="file-botkube-app-creation-yaml-LC4"><span>background_color</span>: <span><span>"</span>#a653a6<span>"</span></span></td></tr><tr><td id="file-botkube-app-creation-yaml-L5" data-line-number="5"></td><td id="file-botkube-app-creation-yaml-LC5"><span>features</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L6" data-line-number="6"></td><td id="file-botkube-app-creation-yaml-LC6"><span>bot_user</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L7" data-line-number="7"></td><td id="file-botkube-app-creation-yaml-LC7"><span>display_name</span>: <span>devreler </span><span><span>#</span> this name will be displayed when you call it.</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L8" data-line-number="8"></td><td id="file-botkube-app-creation-yaml-LC8"><span>always_online</span>: <span>false</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L9" data-line-number="9"></td><td id="file-botkube-app-creation-yaml-LC9"><span>oauth_config</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L10" data-line-number="10"></td><td id="file-botkube-app-creation-yaml-LC10"><span>scopes</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L11" data-line-number="11"></td><td id="file-botkube-app-creation-yaml-LC11"><span>bot</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L12" data-line-number="12"></td><td id="file-botkube-app-creation-yaml-LC12">- <span>channels:read</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L13" data-line-number="13"></td><td id="file-botkube-app-creation-yaml-LC13">- <span>groups:read</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L14" data-line-number="14"></td><td id="file-botkube-app-creation-yaml-LC14">- <span>app_mentions:read</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L15" data-line-number="15"></td><td id="file-botkube-app-creation-yaml-LC15">- <span>reactions:write</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L16" data-line-number="16"></td><td id="file-botkube-app-creation-yaml-LC16">- <span>chat:write</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L17" data-line-number="17"></td><td id="file-botkube-app-creation-yaml-LC17">- <span>files:write</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L18" data-line-number="18"></td><td id="file-botkube-app-creation-yaml-LC18">- <span>users:read</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L19" data-line-number="19"></td><td id="file-botkube-app-creation-yaml-LC19"><span>settings</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L20" data-line-number="20"></td><td id="file-botkube-app-creation-yaml-LC20"><span>event_subscriptions</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L21" data-line-number="21"></td><td id="file-botkube-app-creation-yaml-LC21"><span>bot_events</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L22" data-line-number="22"></td><td id="file-botkube-app-creation-yaml-LC22">- <span>app_mention</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L23" data-line-number="23"></td><td id="file-botkube-app-creation-yaml-LC23"><span>interactivity</span>:</td></tr><tr><td id="file-botkube-app-creation-yaml-L24" data-line-number="24"></td><td id="file-botkube-app-creation-yaml-LC24"><span>is_enabled</span>: <span>true</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L25" data-line-number="25"></td><td id="file-botkube-app-creation-yaml-LC25"><span>org_deploy_enabled</span>: <span>false</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L26" data-line-number="26"></td><td id="file-botkube-app-creation-yaml-LC26"><span>socket_mode_enabled</span>: <span>true</span></td></tr><tr><td id="file-botkube-app-creation-yaml-L27" data-line-number="27"></td><td id="file-botkube-app-creation-yaml-LC27"><span>token_rotation_enabled</span>: <span>false</span></td></tr></tbody></table>

Choose the Slack workspace where you want Botkube to be your sidekick. Now you can bridge the gap between Slack and Kubernetes, bringing them together in harmony.

![Image 8: a screenshot of the adobe adobe adobe adobe adobe](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d62aeb09b9191aa90344_Screenshot%202024-01-13%20at%2021.28.43.png)

You need to invite Botkube to your Slack workspace first, so click on “invite bot to channel” and add the name of the channel you want to invite Botkube from.

![Image 9: a screen shot of a screen with a message box on it](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d674bcf59b0f274bfccb_Screenshot%202024-01-13%20at%2021.30.26.png)

*   Refer this [documentation](https://docs.botkube.io/installation/slack/socket-slack) for more details about integrating Slack with Botkube.
*   Install an App in your workspace

![Image 10: a screenshot of a screen showing a user's permission to access the theatrical stack workspace](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d6a944c186ec25c1d101_Screenshot%202024-01-13%20at%2021.09.44.png)

*   Now, We need to copy the **_Bot token and App level token_** from the Slack Console that is required afterward.
*   Bot Token is directly available on Slack console - [Oauth token](https://www.notion.so/Streamlining-Kubernetes-Management-A-Closer-Look-at-Botkube-ChatOps-d5d0c9761ca449ea84d449aaf8031fa8?pvs=21)
*   App Token can be created as follows:

1.  Select the **Basic Information** link from the left sidebar and scroll down to the **App-Level Token section**. Click on the **Generate Token and Scopes** button.

![Image 11: a screen showing the app level tokens page](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d6c60f818c7e0fa3d23d_Screenshot%202024-01-13%20at%2021.36.17.png)

2.  Enter a **Name**, select `connections:write` scope, and click **Generate**.

![Image 12: generating an app level token](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d7029fa7ebdd68ebe769_Screenshot%202024-01-13%20at%2021.38.49.png)

Export slack tokens from your terminal as follows:

export SLACK_API_BOT_TOKEN="botToken"
export SLACK_API_APP_TOKEN="appToken"


*   You need to make sure you’re adding Botkube to your channel where you’ll be using Botkube, as shown below.
*   After installing the Botkube app to your Slack workspace, you could see a new bot user with the name "Botkube" added to your workspace. Add that bot to a Slack channel you want to receive notifications in. You can add it by inviting @Botkube to a channel.

![Image 13: how to add a person to a group on tumblr](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d76a9e3dce9e6797a30d_Screenshot%202024-01-13%20at%2022.53.41.png)

We’re going to install Botkube on our Kubernetes cluster now. Run below commands:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="" data-tagsearch-path="install-bk"><tbody><tr><td id="file-install-bk-L1" data-line-number="1"></td><td id="file-install-bk-LC1">export CLUSTER_NAME={cluster_name} #your cluster name here</td></tr><tr><td id="file-install-bk-L2" data-line-number="2"></td><td id="file-install-bk-LC2">export ALLOW_KUBECTL=true</td></tr><tr><td id="file-install-bk-L3" data-line-number="3"></td><td id="file-install-bk-LC3">export ALLOW_HELM=true</td></tr><tr><td id="file-install-bk-L4" data-line-number="4"></td><td id="file-install-bk-LC4">export SLACK_CHANNEL_NAME={channel_name} #CHANNEL_NAME is the channel name where @Botkube is added</td></tr><tr><td id="file-install-bk-L5" data-line-number="5"></td><td id="file-install-bk-LC5"></td></tr><tr><td id="file-install-bk-L6" data-line-number="6"></td><td id="file-install-bk-LC6">botkube install --version v1.7.0 \</td></tr><tr><td id="file-install-bk-L7" data-line-number="7"></td><td id="file-install-bk-LC7">--set communications.default-group.socketSlack.enabled=true \</td></tr><tr><td id="file-install-bk-L8" data-line-number="8"></td><td id="file-install-bk-LC8">--set communications.default-group.socketSlack.channels.default.name=${SLACK_CHANNEL_NAME} \</td></tr><tr><td id="file-install-bk-L9" data-line-number="9"></td><td id="file-install-bk-LC9">--set communications.default-group.socketSlack.appToken=${SLACK_API_APP_TOKEN} \</td></tr><tr><td id="file-install-bk-L10" data-line-number="10"></td><td id="file-install-bk-LC10">--set communications.default-group.socketSlack.botToken=${SLACK_API_BOT_TOKEN} \</td></tr><tr><td id="file-install-bk-L11" data-line-number="11"></td><td id="file-install-bk-LC11">--set settings.clusterName=${CLUSTER_NAME} \</td></tr><tr><td id="file-install-bk-L12" data-line-number="12"></td><td id="file-install-bk-LC12">--set 'executors.k8s-default-tools.botkube/kubectl.enabled'=${ALLOW_KUBECTL} \</td></tr><tr><td id="file-install-bk-L13" data-line-number="13"></td><td id="file-install-bk-LC13">--set 'executors.k8s-default-tools.botkube/helm.enabled'=${ALLOW_HELM}</td></tr></tbody></table>

You should expect the output as given below:

![Image 14: a screen shot of a black screen with text on it](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d79c9e3dce9e6797ca20_Screenshot%202024-01-14%20at%2001.13.47.png)

![Image 15: a screen shot of a pc with a black screen](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d7a4e97f05231a340be0_Screenshot%202024-01-14%20at%2001.14.35.png)

### **Step 6: Bringing Botkube to Your Channels**

*   Want Botkube in your private Slack channels? No problem! Just use the "Add to Channel" feature. This way, Botkube becomes part of the conversation, ready to assist in Kubernetes-related discussions and incident management.
*   It’s time to head on to our web platform for Botkube, where we can now easily select the channels where we’ve already installed Botkube.

![Image 16: a screenshot of the adwords settings page](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d7eb0f818c7e0fa4e848_Screenshot%202024-01-13%20at%2023.12.43.png)

### **Step 7: Customizing Your Plugins**

*   The final touch is selecting the plugins you want to use. This is like picking the right tools for your toolbox. Choose the ones that best fit your Kubernetes monitoring and troubleshooting needs. It's all about making Botkube work for you and your team.
*   We had already configured Helm and Kubernetes in the above commands. If you have other use-cases, you can explore argocd, flux, keptn, etc. plugins. All the best 🚀

![Image 17: a screen shot of a google docs page](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d865ae8cb06de78b6563_Screenshot%202024-01-13%20at%2023.13.41.png)

Create Instance by making changes:

![Image 18: a screen shot of the google analytics dashboard](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d88aae0351187ae818a8_Screenshot%202024-01-13%20at%2023.14.04.png)

**Congratulations!** You're all set up with Botkube Cloud and Slack. Now you've got a powerful ally in managing your Kubernetes clusters right within your Slack workspace. Go ahead, give it a spin, and see how it transforms your workflow!

![Image 19: a screen shot of a website with a white background](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d8ae0f818c7e0fa582e7_Screenshot%202024-01-13%20at%2023.15.22.png)

Go to your Slack workspace, you can expect below updates automatically!

![Image 20: a screen shot of the adobe video editor in adobe adobe adobe video](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d8d6e97f05231a350620_Screenshot%202024-01-13%20at%2023.16.31.png)

### Verify Botkube Installation

In Slack, try running kubectl commands under the _kubernetes_ channel like this `@SLACK_APP_NAME kubectl get nodes`

![Image 21: a screenshot of the chat screen on a computer](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65a9d8fcb5f4e48c29b03978_Screenshot%202024-01-14%20at%2000.28.44.png)

Debugging Botkube
-----------------

‍

> Most likely, you’ll come across some common errors around slack channel name or botkube not recognizing your slack workspace channels. To resolve this, you can manually edit the botkube YAML configuration file.

*   Configure Botkube by editing the values.yaml file.

vi botkube/values.yaml


Enable **_socketSlack_**_:_ under ”_communications:”_ and set the name of the default channel to match the name of your actual Slack channel. Set the **botToken** and **appToken** values to the tokens you retrieved from above.

communications:
'default-group':
socketSlack:
enabled: true
channels:
'default':
name: 'kubernetes'  ...
...
botToken: 'BOT_TOKEN'
appToken: 'APP_TOKEN'


Enable the Helm and Kubectl executors

executors:
k8s-default-tools:
botkube/helm:
enabled: true
...
...
...
botkube/kubectl:
enabled: true
...
...


Alternatively, You can also use **helm** for the Installation process for efficient Kubernetes application management.

**Botkube Meets Helm**
----------------------

A standout feature of Botkube? It's a [Helm executor plugin](https://botkube.io/integration/helm). This nifty integration lets you wield Helm commands right from your chat platform. Need to roll back a release or check an app’s status? Botkube’s got your back.

**Botkube’s Toolkit ‍**Botkube isn’t just about deploying and monitoring. It’s a powerhouse for executing a range of Helm commands like install, upgrade, and rollback. Just remember to tweak those RBAC settings for read-write commands.

**Why Botkube and Helm Are a Dream Team**

*   **Efficient Management:** Streamlines Kubernetes application management.
*   **Instant Updates:** Get real-time feedback right in your chat window.
*   **Accessibility:** Manage Kubernetes on the go, even from your phone.

**Getting Hands-On: Botkube and Helm in Action**

**‍**Let’s add practical magic to this post with a quick demo and code snippets.

**Setting Up Botkube with Helm: A Step-by-Step Guide**

# Adding Botkube to helm
helm repo add botkube <a href="https://charts.botkube.io">https://charts.botkube.io</a>.
helm search repo botkube

# Reveals all available charts
helm search repo botkube 

# Botkube Installation
helm install botkube --namespace botkube --create-namespace botkube/botkube

# Verifying Installation
kubectl get pods -n botkube

# If you're using Slack, try this
@Botkube helm list


Conclusion
----------

‍**‍**Botkube emphasizes the challenges DevOps faces with Kubernetes, such as debugging in complex environments and coordinating across teams.This ingenious tool is all about making life easier. It offers streamlined monitoring, efficient debugging, and simplified management of Kubernetes clusters. Think of it as your go-to for instant alerts and the ability to fire off Kubernetes commands right from your chat window. From signing up to integrating your Kubernetes clusters with your favorite chat tools, it’s all about boosting your operational game.

For more details, join the [Botkube community on Slack](https://join.botkube.io/) or reach out to the Botkube team via the [Botkube Dashboard](http://app.botkube.io/).
