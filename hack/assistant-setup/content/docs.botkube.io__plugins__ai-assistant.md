Title: AI Assistant | Botkube

URL Source: https://docs.botkube.io/plugins/ai-assistant

Markdown Content:
info

**This plugin is hosted by the [Botkube Cloud](https://app.botkube.io/) plugin repository and requires active Botkube Cloud account.**

The AI Assistant plugin makes it easier than ever for teams to identify, diagnose, and troubleshoot issues in any Kubernetes cluster. It is powered by the OpenAI's `GPT-4o` large language model with additional Kubernetes-related capabilities.

Features[​](#features "Direct link to Features")
------------------------------------------------

The plugin provides the following features:

*   "AI Help" button for Kubernetes event analysis.
*   Kubernetes cluster troubleshooting based on:
    *   `kubectl` commands,
    *   container logs,
    *   Kubernetes events,
    *   resource usage (CPU/memory).
*   Kubernetes cluster scan.
*   Answering arbitrary questions or instructions.
*   Answering questions based on user-supplied documentation.
*   Answering Botkube questions or instructions based on:
    *   Agent status and configuration,
    *   the latest Botkube documentation and other Botkube content.
*   Keeping conversation context in the message threads.

Get started[​](#get-started "Direct link to Get started")
---------------------------------------------------------

Enable the plugin[​](#enable-the-plugin "Direct link to Enable the plugin")
---------------------------------------------------------------------------

You can enable the plugin as a part of Botkube instance configuration.

1.  If you don't have an existing Botkube instance, create a new one, according to the [Installation](https://docs.botkube.io/) docs.
2.  From the [Botkube Cloud homepage](https://app.botkube.io/), click on a card of a given Botkube instance.
3.  Navigate to the platform tab which you want to configure.
4.  Click **Add plugin** button.
5.  Select the AI plugin.
6.  Click **Save** button.

Usage[​](#usage "Direct link to Usage")
---------------------------------------

### Highlights[​](#highlights "Direct link to Highlights")

To see the AI assistant in action, watch the video:

### AI Help for Kubernetes events[​](#ai-help-for-kubernetes-events "Direct link to AI Help for Kubernetes events")

Click on the **AI Help** button to get an analysis of the event.

![Image 1: AI Help](https://docs.botkube.io/assets/images/ai-help-usage-f2d200f6e8233c8528366a07763f9aac.png)

### AI questions and instructions[​](#ai-questions-and-instructions "Direct link to AI questions and instructions")

To ask the Botkube AI Assistant any question, use the `@Botkube ai` command prefix.

Here are a few examples to get you started:

*   `@Botkube ai are my pods in default namespace up and running?`
*   `@Botkube ai how many nodes are there?`
*   `@Botkube ai output a yaml for Kubernetes pod with nginx image`
*   `@Botkube what's the current Botkube status?`
*   `@Botkube ai what's the difference between Botkube sources and executors?`
*   `@Botkube ai what Botkube plugins are available to install? List all of them with one sentence description and link to the docs.`
*   `@Botkube ai who's the Botkube maintainer?`

### AI cluster scan[​](#ai-cluster-scan "Direct link to AI cluster scan")

The assistant can help you troubleshoot the Kubernetes cluster by performing a cluster-wide scan for any potential issues, such as failing pods, high resource usage, or any other anomalies. To start the scan, use the `@Botkube ai scan` command.

### User-supplied documentation[​](#user-supplied-documentation "Direct link to User-supplied documentation")

info

This feature is only available for Botkube Cloud paid plans.

Botkube Assistant can use supplied documentation to give better answers, specific to your architecture and processes. For example: architectural documents, onboarding material, runbooks, post-mortems, etc.

The user-supplied documentation is scoped to a given Botkube Cloud organization.

#### Manage uploaded files[​](#manage-uploaded-files "Direct link to Manage uploaded files")

To manage uploaded files, navigate to the ["AI Assistant"](https://app.botkube.io/ai-assistant) tab in Botkube Cloud. From the tab, you can upload files with the **Upload File(s)** button. Supported file formats include: `.doc`, `.docx`, `.md`, `.pdf`, `.txt`. See the [full list](https://platform.openai.com/docs/assistants/tools/file-search/supported-files).

To delete file(s), select the file(s) you want to delete and click **Delete selected file(s)**.

Once the files are uploaded, you can ask Botkube AI Assistant questions related to the uploaded content. To learn more, see the [AI questions and instructions](#ai-questions-and-instructions) section.

Configuration[​](#configuration "Direct link to Configuration")
---------------------------------------------------------------

Currently, the AI plugin doesn't have any configuration options.

Limitations[​](#limitations "Direct link to Limitations")
---------------------------------------------------------

*   The conversation context is stored in-memory and will be cleared on Botkube Agent restart.
