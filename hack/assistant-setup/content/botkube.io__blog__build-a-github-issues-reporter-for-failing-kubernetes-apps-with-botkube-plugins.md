Title: Build a GitHub Issues Reporter for failing K8s Apps

URL Source: https://botkube.io/blog/build-a-github-issues-reporter-for-failing-kubernetes-apps-with-botkube-plugins

Published Time: Feb 01, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3ec19d99029e7cfb5bd_n-vUM0Lz3yXA1WLyfxEMUxXQpUDEiKv9EmEe4NnCbCI.png)

Mateusz Szostok

Software Engineer

Botkube

Botkube 0.17 introduces plugin support. In this guide you will learn how to create your own Botkube executor plugin to build a GitHub issues reporter for Kubernetes apps that are failing.

### Table of Contents

*   [Goal](#goal)
*   [Prerequisites](#prerequisites)
*   [What's under the hood](#what-s-under-the-hood)
*   [Step-By-Step Instructions](#step-by-step-instructions)
*   [Summary](#summary)
*   [How can I get involved?](#how-can-i-get-involved-)

#### Start Using Botkube AI Assistant Today!

Botkube gives you fast and simple access to your clusters right from your communication platform. It does that by sending [actionable notifications](https://docs.botkube.io/configuration/action) (via [_sources_](https://docs.botkube.io/architecture/#source)) and allowing you to run _kubectl_ and _helm_ commands (via [_executors_](https://docs.botkube.io/architecture/#executor)) straight from the platform (Slack, Discord, Microsoft Teams and Mattermost).

In the recent Botkube 0.17 release, we took it to the next level by giving you an easy way to bring your own tools to a chat window!

In this blog post, you will learn how to develop your own executor plugin to fill the gap in your daily workflow.

Goal
----

To make it simple but functional, I will show you how to develop an executor that creates an issue for failing Kubernetes resources such as Job, Deployment, StatefulSet, and Pod.

![Image 2: Github reporter bot for updating Slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63d810824fa1ec68968b8ee2_gh-demo.gif)

GitHub executor demo

### Why it's worth it

With just a few lines of code, we will automate the process of creating a GitHub issue that out-of-the box contains Kubernetes-specific information useful for further debugging. All of that, directly from Slack, Discord, Mattermost, or MS Teams! No need for connecting to your cluster in your terminal, installing and running _kubectl_ commands and copy-pasting fetched details into your browser.

Instead, you will be able to type **_@Botkube gh create issue pod/foo_** from any device that has access to your chat, including mobile apps.

Prerequisites
-------------

*   Access to a Kubernetes cluster**‍**‍

To create a local cluster for testing purposes using [`k3d`](https://k3d.io/v5.0.1/#installation), run:

*   [Botkube installed and configured](https://docs.botkube.io/)[](https://docs.botkube.io/installation/)
    
*   Basic understanding of the Go language[‍](https://go.dev/doc/install)
    
*   [Go](https://go.dev/doc/install) at least 1.18
    

‍

What's under the hood
---------------------

To understand better what we will develop, I want to give you a bigger picture of the Botkube plugin system. The animation below focuses only on the executor part, but it's almost the same for sources.

![Image 3: Custom built plugin system for Kubernetes tools](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63d811d4762a4de9804bc215_arch-executor-plugin.gif)

Botkube Plugin System

The new part is a plugin repository that we introduced in 0.17. Plugin repository is a place where you store your plugin binaries and index files. Any static file server can be used, and in this blog post we will use a GitHub release. It’s similar to what you know from the Helm ecosystem.

The plugin manager consumes user's configuration, and downloads and starts **only** enabled plugins from a given repository. Plugins are running directly on the Kubernetes cluster where the Botkube core was installed.

Such approach allows us to decouple the Botkube core and its extensions. Thanks to that, we can:

*   Avoid having the Botkube core crash if a given plugin malfunctions
    
*   Write extensions in any language as gRPC is used
    

From the end-user's perspective, you can:

*   Specify and use multiple plugin repositories at the same time
    
*   Enable different plugin versions within the same Botkube version
    

To learn more, see [Botkube architecture](https://docs.botkube.io/architecture/).

Step-By-Step Instructions
-------------------------

To quickly onboard you to Botkube plugins, we maintain the [kubeshop/botkube-plugins-template](https://github.com/kubeshop/botkube-plugins-template) repository that has all batteries included. Our first step is to bootstrap your own GitHub repository.

To check out the entire code, visit the [Botkube GitHub repository](https://github.com/kubeshop/botkube/tree/main/cmd/executor/gh/main.go).

### Repository setup

1.  Navigate to [botkube-plugins-template](https://github.com/kubeshop/botkube-plugins-template).
    
2.  Click **Use this template**, and then **Create a new repository**.
    

‍

![Image 4: github pull request automation](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63d812e448eb8e82ad2d44d9_tpl-repo.png)

By doing so, you will create your own plugin repository with a single commit.

3.  Clone your repository locally:_‍_

4.  Create and push a new tag to perform the initial release:

After a few minutes, you should see a new GitHub release.

Voilà! You are already an owner of fully functional Botkube plugins. Now it's time to add your own brick by creating a GitHub executor.

‍

In this blog post, we use only GitHub releases that work out-of-the-box. Releasing plugins on GitHub Pages requires additional setup. To support them too, see the [use template document](https://docs.botkube.io/plugin/quick-start#use-template)

‍

### Develop GitHub executor

To make the code-snippets more readable, I skipped the error handling. However, it will be useful if you will add error handling for the final implementation. You can check the full [gh source-code](https://github.com/kubeshop/botkube/tree/main/cmd/executor/gh/main.go) for the reference.

‍

1.  Create a `cmd/gh/main.go` file with the following template:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="plugin-gh-tpl.go"><tbody><tr><td id="file-plugin-gh-tpl-go-L1" data-line-number="1"></td><td id="file-plugin-gh-tpl-go-LC1"><span>package</span> main</td></tr><tr><td id="file-plugin-gh-tpl-go-L2" data-line-number="2"></td><td id="file-plugin-gh-tpl-go-LC2"></td></tr><tr><td id="file-plugin-gh-tpl-go-L3" data-line-number="3"></td><td id="file-plugin-gh-tpl-go-LC3"><span>import</span> (</td></tr><tr><td id="file-plugin-gh-tpl-go-L4" data-line-number="4"></td><td id="file-plugin-gh-tpl-go-LC4"><span>"context"</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L5" data-line-number="5"></td><td id="file-plugin-gh-tpl-go-LC5"></td></tr><tr><td id="file-plugin-gh-tpl-go-L6" data-line-number="6"></td><td id="file-plugin-gh-tpl-go-LC6"><span>"github.com/hashicorp/go-plugin"</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L7" data-line-number="7"></td><td id="file-plugin-gh-tpl-go-LC7"><span>"github.com/kubeshop/botkube/pkg/api"</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L8" data-line-number="8"></td><td id="file-plugin-gh-tpl-go-LC8"><span>"github.com/kubeshop/botkube/pkg/api/executor"</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L9" data-line-number="9"></td><td id="file-plugin-gh-tpl-go-LC9">)</td></tr><tr><td id="file-plugin-gh-tpl-go-L10" data-line-number="10"></td><td id="file-plugin-gh-tpl-go-LC10"></td></tr><tr><td id="file-plugin-gh-tpl-go-L11" data-line-number="11"></td><td id="file-plugin-gh-tpl-go-LC11"><span>const</span> (</td></tr><tr><td id="file-plugin-gh-tpl-go-L12" data-line-number="12"></td><td id="file-plugin-gh-tpl-go-LC12"><span>pluginName</span> <span>=</span> <span>"gh"</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L13" data-line-number="13"></td><td id="file-plugin-gh-tpl-go-LC13">)</td></tr><tr><td id="file-plugin-gh-tpl-go-L14" data-line-number="14"></td><td id="file-plugin-gh-tpl-go-LC14"></td></tr><tr><td id="file-plugin-gh-tpl-go-L15" data-line-number="15"></td><td id="file-plugin-gh-tpl-go-LC15"><span>// GHExecutor implements the Botkube executor plugin interface.</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L16" data-line-number="16"></td><td id="file-plugin-gh-tpl-go-LC16"><span>type</span> <span>GHExecutor</span> <span>struct</span>{}</td></tr><tr><td id="file-plugin-gh-tpl-go-L17" data-line-number="17"></td><td id="file-plugin-gh-tpl-go-LC17"></td></tr><tr><td id="file-plugin-gh-tpl-go-L18" data-line-number="18"></td><td id="file-plugin-gh-tpl-go-LC18"><span>// Metadata returns details about the GitHub plugin.</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L19" data-line-number="19"></td><td id="file-plugin-gh-tpl-go-LC19"><span>func</span> (<span>*</span><span>GHExecutor</span>) <span>Metadata</span>(context.<span>Context</span>) (api.<span>MetadataOutput</span>, <span>error</span>) {</td></tr><tr><td id="file-plugin-gh-tpl-go-L20" data-line-number="20"></td><td id="file-plugin-gh-tpl-go-LC20"><span>return</span> api.<span>MetadataOutput</span>{</td></tr><tr><td id="file-plugin-gh-tpl-go-L21" data-line-number="21"></td><td id="file-plugin-gh-tpl-go-LC21"><span>Version</span>: <span>"v1.0.0"</span>,</td></tr><tr><td id="file-plugin-gh-tpl-go-L22" data-line-number="22"></td><td id="file-plugin-gh-tpl-go-LC22"><span>Description</span>: <span>"GH creates an issue on GitHub for a related Kubernetes resource."</span>,</td></tr><tr><td id="file-plugin-gh-tpl-go-L23" data-line-number="23"></td><td id="file-plugin-gh-tpl-go-LC23">}, <span>nil</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L24" data-line-number="24"></td><td id="file-plugin-gh-tpl-go-LC24">}</td></tr><tr><td id="file-plugin-gh-tpl-go-L25" data-line-number="25"></td><td id="file-plugin-gh-tpl-go-LC25"></td></tr><tr><td id="file-plugin-gh-tpl-go-L26" data-line-number="26"></td><td id="file-plugin-gh-tpl-go-LC26"><span>// Execute returns a given command as a response.</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L27" data-line-number="27"></td><td id="file-plugin-gh-tpl-go-LC27"><span>func</span> (<span>e</span> <span>*</span><span>GHExecutor</span>) <span>Execute</span>(<span>ctx</span> context.<span>Context</span>, <span>in</span> executor.<span>ExecuteInput</span>) (executor.<span>ExecuteOutput</span>, <span>error</span>) {</td></tr><tr><td id="file-plugin-gh-tpl-go-L28" data-line-number="28"></td><td id="file-plugin-gh-tpl-go-LC28"><span>return</span> executor.<span>ExecuteOutput</span>{</td></tr><tr><td id="file-plugin-gh-tpl-go-L29" data-line-number="29"></td><td id="file-plugin-gh-tpl-go-LC29"><span>Data</span>: <span>"Aloha! 🏄"</span>,</td></tr><tr><td id="file-plugin-gh-tpl-go-L30" data-line-number="30"></td><td id="file-plugin-gh-tpl-go-LC30">}, <span>nil</span></td></tr><tr><td id="file-plugin-gh-tpl-go-L31" data-line-number="31"></td><td id="file-plugin-gh-tpl-go-LC31">}</td></tr><tr><td id="file-plugin-gh-tpl-go-L32" data-line-number="32"></td><td id="file-plugin-gh-tpl-go-LC32"></td></tr><tr><td id="file-plugin-gh-tpl-go-L33" data-line-number="33"></td><td id="file-plugin-gh-tpl-go-LC33"><span>func</span> <span>main</span>() {</td></tr><tr><td id="file-plugin-gh-tpl-go-L34" data-line-number="34"></td><td id="file-plugin-gh-tpl-go-LC34"><span>executor</span>.<span>Serve</span>(<span>map</span>[<span>string</span>]plugin.<span>Plugin</span>{</td></tr><tr><td id="file-plugin-gh-tpl-go-L35" data-line-number="35"></td><td id="file-plugin-gh-tpl-go-LC35"><span>pluginName</span>: <span>&amp;</span>executor.<span>Plugin</span>{</td></tr><tr><td id="file-plugin-gh-tpl-go-L36" data-line-number="36"></td><td id="file-plugin-gh-tpl-go-LC36"><span>Executor</span>: <span>&amp;</span><span>GHExecutor</span>{},</td></tr><tr><td id="file-plugin-gh-tpl-go-L37" data-line-number="37"></td><td id="file-plugin-gh-tpl-go-LC37">},</td></tr><tr><td id="file-plugin-gh-tpl-go-L38" data-line-number="38"></td><td id="file-plugin-gh-tpl-go-LC38">})</td></tr><tr><td id="file-plugin-gh-tpl-go-L39" data-line-number="39"></td><td id="file-plugin-gh-tpl-go-LC39">}</td></tr></tbody></table>

This template code imports required packages and registers `GHExecutor` as the gRPC plugin. Our `GHExecutor` service already implements the required [Protocol Buffers](https://github.com/kubeshop/botkube/blob/main/proto/executor.proto) contract. As you can see, we require **only 2 methods**.

\- The `Metadata` method returns basic information about your plugin. This data is used when the plugin index is [generated in an automated way](https://docs.botkube.io/plugin/repo).

\- The `Execute` method is the heart of your executor plugin. This method runs your business logic and returns the output as plaintext. Next, the Botkube core sends back the response to a given communication platform.

‍

2.  To download the imported dependencies, in your terminal, run:

‍

3.  Great! At this stage, you already have a functional Botkube executor plugin. However, for now, it only responds with a hard-coded "Aloha!" greeting. But it can do that already on all supported communication platforms. ‍

![Image 5: Getting GitHub reports in Slack about repo](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63d814cf12ac7e2d4a40d002_demo-gh-aloha.gif)

Don't worry, in the next steps, we will implement our business logic.

4.  Add support for user configuration:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="config.go"><tbody><tr><td id="file-config-go-L1" data-line-number="1"></td><td id="file-config-go-LC1"><span>package</span> main</td></tr><tr><td id="file-config-go-L2" data-line-number="2"></td><td id="file-config-go-LC2"></td></tr><tr><td id="file-config-go-L3" data-line-number="3"></td><td id="file-config-go-LC3"><span>import</span> (</td></tr><tr><td id="file-config-go-L4" data-line-number="4"></td><td id="file-config-go-LC4"><span>"context"</span></td></tr><tr><td id="file-config-go-L5" data-line-number="5"></td><td id="file-config-go-LC5"><span>"fmt"</span></td></tr><tr><td id="file-config-go-L6" data-line-number="6"></td><td id="file-config-go-LC6"></td></tr><tr><td id="file-config-go-L7" data-line-number="7"></td><td id="file-config-go-LC7"><span>"github.com/kubeshop/botkube/pkg/api/executor"</span></td></tr><tr><td id="file-config-go-L8" data-line-number="8"></td><td id="file-config-go-LC8"><span>"github.com/kubeshop/botkube/pkg/pluginx"</span></td></tr><tr><td id="file-config-go-L9" data-line-number="9"></td><td id="file-config-go-LC9">)</td></tr><tr><td id="file-config-go-L10" data-line-number="10"></td><td id="file-config-go-LC10"></td></tr><tr><td id="file-config-go-L11" data-line-number="11"></td><td id="file-config-go-LC11"><span>// Config holds the GitHub executor configuration.</span></td></tr><tr><td id="file-config-go-L12" data-line-number="12"></td><td id="file-config-go-LC12"><span>type</span> <span>Config</span> <span>struct</span> {</td></tr><tr><td id="file-config-go-L13" data-line-number="13"></td><td id="file-config-go-LC13"><span>GitHub</span> <span>struct</span> {</td></tr><tr><td id="file-config-go-L14" data-line-number="14"></td><td id="file-config-go-LC14"><span>Token</span> <span>string</span> <span>`yaml:"token"`</span></td></tr><tr><td id="file-config-go-L15" data-line-number="15"></td><td id="file-config-go-LC15"><span>Repository</span> <span>string</span> <span>`yaml:"repository"`</span></td></tr><tr><td id="file-config-go-L16" data-line-number="16"></td><td id="file-config-go-LC16"><span>IssueTemplate</span> <span>string</span> <span>`yaml:"issueTemplate"`</span></td></tr><tr><td id="file-config-go-L17" data-line-number="17"></td><td id="file-config-go-LC17">} <span>`yaml:"gitHub"`</span></td></tr><tr><td id="file-config-go-L18" data-line-number="18"></td><td id="file-config-go-LC18">}</td></tr><tr><td id="file-config-go-L19" data-line-number="19"></td><td id="file-config-go-LC19"></td></tr><tr><td id="file-config-go-L20" data-line-number="20"></td><td id="file-config-go-LC20"><span>// Execute returns a given command as a response.</span></td></tr><tr><td id="file-config-go-L21" data-line-number="21"></td><td id="file-config-go-LC21"><span>func</span> (<span>e</span> <span>*</span><span>GHExecutor</span>) <span>Execute</span>(<span>_</span> context.<span>Context</span>, <span>in</span> executor.<span>ExecuteInput</span>) (executor.<span>ExecuteOutput</span>, <span>error</span>) {</td></tr><tr><td id="file-config-go-L22" data-line-number="22"></td><td id="file-config-go-LC22"><span>var</span> <span>cfg</span> <span>Config</span></td></tr><tr><td id="file-config-go-L23" data-line-number="23"></td><td id="file-config-go-LC23"><span>pluginx</span>.<span>MergeExecutorConfigs</span>(<span>in</span>.<span>Configs</span>, <span>&amp;</span><span>cfg</span>)</td></tr><tr><td id="file-config-go-L24" data-line-number="24"></td><td id="file-config-go-LC24">}</td></tr></tbody></table>

For each `Execute` method call, Botkube attaches the list of associated configurations. The input parameters are defined by the user, when enabling a given plugin:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="params.yaml"><tbody><tr><td id="file-params-yaml-L1" data-line-number="1"></td><td id="file-params-yaml-LC1"><span>executors</span>:</td></tr><tr><td id="file-params-yaml-L2" data-line-number="2"></td><td id="file-params-yaml-LC2"><span><span>"</span><span>plugin-based</span><span>"</span></span>:</td></tr><tr><td id="file-params-yaml-L3" data-line-number="3"></td><td id="file-params-yaml-LC3"><span>repo-name/gh</span>:</td></tr><tr><td id="file-params-yaml-L4" data-line-number="4"></td><td id="file-params-yaml-LC4"><span>enabled</span>: <span>true </span><span><span>#</span> If not enabled, plugin is not downloaded and started.</span></td></tr><tr><td id="file-params-yaml-L5" data-line-number="5"></td><td id="file-params-yaml-LC5"><span>config</span>: <span><span>#</span> Plugin's specific configuration.</span></td></tr><tr><td id="file-params-yaml-L6" data-line-number="6"></td><td id="file-params-yaml-LC6"><span>github</span>:</td></tr><tr><td id="file-params-yaml-L7" data-line-number="7"></td><td id="file-params-yaml-LC7"><span>repository</span>: <span><span>"</span>mszostok/repository<span>"</span></span></td></tr><tr><td id="file-params-yaml-L8" data-line-number="8"></td><td id="file-params-yaml-LC8"><span>token</span>: <span><span>"</span>github_pat_foo<span>"</span></span></td></tr><tr><td id="file-params-yaml-L9" data-line-number="9"></td><td id="file-params-yaml-LC9"><span>issueTemplate</span>: <span>|</span></td></tr><tr><td id="file-params-yaml-L10" data-line-number="10"></td><td id="file-params-yaml-LC10"><span>## Description</span></td></tr><tr><td id="file-params-yaml-L11" data-line-number="11"></td><td id="file-params-yaml-LC11"><span></span></td></tr><tr><td id="file-params-yaml-L12" data-line-number="12"></td><td id="file-params-yaml-LC12"><span>This issue refers to the problems connected with {{ .Type | code "bash" }} in namespace {{ .Namespace | code "bash" }}</span></td></tr><tr><td id="file-params-yaml-L13" data-line-number="13"></td><td id="file-params-yaml-LC13"><span></span></td></tr><tr><td id="file-params-yaml-L14" data-line-number="14"></td><td id="file-params-yaml-LC14"><span>{{ .Logs | code "bash"}}</span></td></tr></tbody></table>

In our case, we need to have a GitHub token, a GitHub repository where the issue should be created, and an issue template. The remaining parameters can be hard-coded on the plugin side, however, your plugin will be more flexible if you allow your users to change it without rebuilding your plugins.It's up to the plugin author to merge the passed configurations. You can use our helper function from the plugin extension package (`pluginx`). To learn more, see [Passing configuration to your plugin](https://docs.botkube.io/plugin/custom-executor#passing-configuration-to-your-plugin).

‍**‍**‍

5.  Let's implement command parsing to properly understand command syntax:

Our goal is to parse `gh create issue KIND/NAME [-n, --namespace]`

There are a lot of great libraries supporting command parsing. The most popular is probably [cobra](https://github.com/spf13/cobra), but for our use case, we will just use the helper function from our plugin extension package.

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="command-parsing.go"><tbody><tr><td id="file-command-parsing-go-L1" data-line-number="1"></td><td id="file-command-parsing-go-LC1"><span>const</span> (</td></tr><tr><td id="file-command-parsing-go-L2" data-line-number="2"></td><td id="file-command-parsing-go-LC2"><span>pluginName</span> <span>=</span> <span>"gh"</span></td></tr><tr><td id="file-command-parsing-go-L3" data-line-number="3"></td><td id="file-command-parsing-go-LC3">)</td></tr><tr><td id="file-command-parsing-go-L4" data-line-number="4"></td><td id="file-command-parsing-go-LC4"></td></tr><tr><td id="file-command-parsing-go-L5" data-line-number="5"></td><td id="file-command-parsing-go-LC5"><span>// Commands defines all supported GitHub plugin commands and their flags.</span></td></tr><tr><td id="file-command-parsing-go-L6" data-line-number="6"></td><td id="file-command-parsing-go-LC6"><span>type</span> (</td></tr><tr><td id="file-command-parsing-go-L7" data-line-number="7"></td><td id="file-command-parsing-go-LC7"><span>Commands</span> <span>struct</span> {</td></tr><tr><td id="file-command-parsing-go-L8" data-line-number="8"></td><td id="file-command-parsing-go-LC8"><span>Create</span> <span>*</span><span>CreateCommand</span> <span>`arg:"subcommand:create"`</span></td></tr><tr><td id="file-command-parsing-go-L9" data-line-number="9"></td><td id="file-command-parsing-go-LC9">}</td></tr><tr><td id="file-command-parsing-go-L10" data-line-number="10"></td><td id="file-command-parsing-go-LC10"><span>CreateCommand</span> <span>struct</span> {</td></tr><tr><td id="file-command-parsing-go-L11" data-line-number="11"></td><td id="file-command-parsing-go-LC11"><span>Issue</span> <span>*</span><span>CreateIssueCommand</span> <span>`arg:"subcommand:issue"`</span></td></tr><tr><td id="file-command-parsing-go-L12" data-line-number="12"></td><td id="file-command-parsing-go-LC12">}</td></tr><tr><td id="file-command-parsing-go-L13" data-line-number="13"></td><td id="file-command-parsing-go-LC13"><span>CreateIssueCommand</span> <span>struct</span> {</td></tr><tr><td id="file-command-parsing-go-L14" data-line-number="14"></td><td id="file-command-parsing-go-LC14"><span>Type</span> <span>string</span> <span>`arg:"positional"`</span></td></tr><tr><td id="file-command-parsing-go-L15" data-line-number="15"></td><td id="file-command-parsing-go-LC15"><span>Namespace</span> <span>string</span> <span>`arg:"-n,--namespace"`</span></td></tr><tr><td id="file-command-parsing-go-L16" data-line-number="16"></td><td id="file-command-parsing-go-LC16">}</td></tr><tr><td id="file-command-parsing-go-L17" data-line-number="17"></td><td id="file-command-parsing-go-LC17">)</td></tr><tr><td id="file-command-parsing-go-L18" data-line-number="18"></td><td id="file-command-parsing-go-LC18"></td></tr><tr><td id="file-command-parsing-go-L19" data-line-number="19"></td><td id="file-command-parsing-go-LC19"><span>// Execute returns a given command as a response.</span></td></tr><tr><td id="file-command-parsing-go-L20" data-line-number="20"></td><td id="file-command-parsing-go-LC20"><span>func</span> (<span>e</span> <span>*</span><span>GHExecutor</span>) <span>Execute</span>(<span>ctx</span> context.<span>Context</span>, <span>in</span> executor.<span>ExecuteInput</span>) (executor.<span>ExecuteOutput</span>, <span>error</span>) {</td></tr><tr><td id="file-command-parsing-go-L21" data-line-number="21"></td><td id="file-command-parsing-go-LC21"><span>// ...</span></td></tr><tr><td id="file-command-parsing-go-L22" data-line-number="22"></td><td id="file-command-parsing-go-LC22"><span>var</span> <span>cmd</span> <span>Commands</span></td></tr><tr><td id="file-command-parsing-go-L23" data-line-number="23"></td><td id="file-command-parsing-go-LC23"><span>pluginx</span>.<span>ParseCommand</span>(<span>pluginName</span>, <span>in</span>.<span>Command</span>, <span>&amp;</span><span>cmd</span>)</td></tr><tr><td id="file-command-parsing-go-L24" data-line-number="24"></td><td id="file-command-parsing-go-LC24"></td></tr><tr><td id="file-command-parsing-go-L25" data-line-number="25"></td><td id="file-command-parsing-go-LC25"><span>if</span> <span>cmd</span>.<span>Create</span> <span>==</span> <span>nil</span> <span>||</span> <span>cmd</span>.<span>Create</span>.<span>Issue</span> <span>==</span> <span>nil</span> {</td></tr><tr><td id="file-command-parsing-go-L26" data-line-number="26"></td><td id="file-command-parsing-go-LC26"><span>return</span> executor.<span>ExecuteOutput</span>{</td></tr><tr><td id="file-command-parsing-go-L27" data-line-number="27"></td><td id="file-command-parsing-go-LC27"><span>Data</span>: <span>fmt</span>.<span>Sprintf</span>(<span>"Usage: %s create issue KIND/NAME"</span>, <span>pluginName</span>),</td></tr><tr><td id="file-command-parsing-go-L28" data-line-number="28"></td><td id="file-command-parsing-go-LC28">}, <span>nil</span></td></tr><tr><td id="file-command-parsing-go-L29" data-line-number="29"></td><td id="file-command-parsing-go-LC29">}</td></tr><tr><td id="file-command-parsing-go-L30" data-line-number="30"></td><td id="file-command-parsing-go-LC30"><span>// ...</span></td></tr><tr><td id="file-command-parsing-go-L31" data-line-number="31"></td><td id="file-command-parsing-go-LC31">}</td></tr></tbody></table>

Under the hood, the `pluginx.ParseCommand` method uses [go-arg](https://github.com/alexflint/go-arg).

‍

6.  We are almost there! Now let's fetch the issue details:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="fetch-issue-details.go"><tbody><tr><td id="file-fetch-issue-details-go-L1" data-line-number="1"></td><td id="file-fetch-issue-details-go-LC1"><span>const</span> (</td></tr><tr><td id="file-fetch-issue-details-go-L2" data-line-number="2"></td><td id="file-fetch-issue-details-go-LC2"><span>logsTailLines</span> <span>=</span> <span>150</span></td></tr><tr><td id="file-fetch-issue-details-go-L3" data-line-number="3"></td><td id="file-fetch-issue-details-go-LC3"><span>defaultNamespace</span> <span>=</span> <span>"default"</span></td></tr><tr><td id="file-fetch-issue-details-go-L4" data-line-number="4"></td><td id="file-fetch-issue-details-go-LC4">)</td></tr><tr><td id="file-fetch-issue-details-go-L5" data-line-number="5"></td><td id="file-fetch-issue-details-go-LC5"></td></tr><tr><td id="file-fetch-issue-details-go-L6" data-line-number="6"></td><td id="file-fetch-issue-details-go-LC6"><span>// Execute returns a given command as a response.</span></td></tr><tr><td id="file-fetch-issue-details-go-L7" data-line-number="7"></td><td id="file-fetch-issue-details-go-LC7"><span>func</span> (<span>e</span> <span>*</span><span>GHExecutor</span>) <span>Execute</span>(<span>ctx</span> context.<span>Context</span>, <span>in</span> executor.<span>ExecuteInput</span>) (executor.<span>ExecuteOutput</span>, <span>error</span>) {</td></tr><tr><td id="file-fetch-issue-details-go-L8" data-line-number="8"></td><td id="file-fetch-issue-details-go-LC8"><span>// ...</span></td></tr><tr><td id="file-fetch-issue-details-go-L9" data-line-number="9"></td><td id="file-fetch-issue-details-go-LC9"><span>issueDetails</span>, <span>_</span> <span>:=</span> <span>getIssueDetails</span>(<span>ctx</span>, <span>cmd</span>.<span>Create</span>.<span>Issue</span>.<span>Namespace</span>, <span>cmd</span>.<span>Create</span>.<span>Issue</span>.<span>Type</span>)</td></tr><tr><td id="file-fetch-issue-details-go-L10" data-line-number="10"></td><td id="file-fetch-issue-details-go-LC10"><span>// ...</span></td></tr><tr><td id="file-fetch-issue-details-go-L11" data-line-number="11"></td><td id="file-fetch-issue-details-go-LC11">}</td></tr><tr><td id="file-fetch-issue-details-go-L12" data-line-number="12"></td><td id="file-fetch-issue-details-go-LC12"></td></tr><tr><td id="file-fetch-issue-details-go-L13" data-line-number="13"></td><td id="file-fetch-issue-details-go-LC13"><span>// IssueDetails holds all available information about a given issue.</span></td></tr><tr><td id="file-fetch-issue-details-go-L14" data-line-number="14"></td><td id="file-fetch-issue-details-go-LC14"><span>type</span> <span>IssueDetails</span> <span>struct</span> {</td></tr><tr><td id="file-fetch-issue-details-go-L15" data-line-number="15"></td><td id="file-fetch-issue-details-go-LC15"><span>Type</span> <span>string</span></td></tr><tr><td id="file-fetch-issue-details-go-L16" data-line-number="16"></td><td id="file-fetch-issue-details-go-LC16"><span>Namespace</span> <span>string</span></td></tr><tr><td id="file-fetch-issue-details-go-L17" data-line-number="17"></td><td id="file-fetch-issue-details-go-LC17"><span>Logs</span> <span>string</span></td></tr><tr><td id="file-fetch-issue-details-go-L18" data-line-number="18"></td><td id="file-fetch-issue-details-go-LC18"><span>Version</span> <span>string</span></td></tr><tr><td id="file-fetch-issue-details-go-L19" data-line-number="19"></td><td id="file-fetch-issue-details-go-LC19">}</td></tr><tr><td id="file-fetch-issue-details-go-L20" data-line-number="20"></td><td id="file-fetch-issue-details-go-LC20"></td></tr><tr><td id="file-fetch-issue-details-go-L21" data-line-number="21"></td><td id="file-fetch-issue-details-go-LC21"><span>func</span> <span>getIssueDetails</span>(<span>ctx</span> context.<span>Context</span>, <span>namespace</span>, <span>name</span> <span>string</span>) (<span>IssueDetails</span>, <span>error</span>) {</td></tr><tr><td id="file-fetch-issue-details-go-L22" data-line-number="22"></td><td id="file-fetch-issue-details-go-LC22"><span>if</span> <span>namespace</span> <span>==</span> <span>""</span> {</td></tr><tr><td id="file-fetch-issue-details-go-L23" data-line-number="23"></td><td id="file-fetch-issue-details-go-LC23"><span>namespace</span> <span>=</span> <span>defaultNamespace</span></td></tr><tr><td id="file-fetch-issue-details-go-L24" data-line-number="24"></td><td id="file-fetch-issue-details-go-LC24">}</td></tr><tr><td id="file-fetch-issue-details-go-L25" data-line-number="25"></td><td id="file-fetch-issue-details-go-LC25"><span>logs</span>, <span>_</span> <span>:=</span> <span>pluginx</span>.<span>ExecuteCommand</span>(<span>ctx</span>, <span>fmt</span>.<span>Sprintf</span>(<span>"kubectl logs %s -n %s --tail %d"</span>, <span>name</span>, <span>namespace</span>, <span>logsTailLines</span>))</td></tr><tr><td id="file-fetch-issue-details-go-L26" data-line-number="26"></td><td id="file-fetch-issue-details-go-LC26"><span>ver</span>, <span>_</span> <span>:=</span> <span>pluginx</span>.<span>ExecuteCommand</span>(<span>ctx</span>, <span>"kubectl version -o yaml"</span>)</td></tr><tr><td id="file-fetch-issue-details-go-L27" data-line-number="27"></td><td id="file-fetch-issue-details-go-LC27"></td></tr><tr><td id="file-fetch-issue-details-go-L28" data-line-number="28"></td><td id="file-fetch-issue-details-go-LC28"><span>return</span> <span>IssueDetails</span>{</td></tr><tr><td id="file-fetch-issue-details-go-L29" data-line-number="29"></td><td id="file-fetch-issue-details-go-LC29"><span>Type</span>: <span>name</span>,</td></tr><tr><td id="file-fetch-issue-details-go-L30" data-line-number="30"></td><td id="file-fetch-issue-details-go-LC30"><span>Namespace</span>: <span>namespace</span>,</td></tr><tr><td id="file-fetch-issue-details-go-L31" data-line-number="31"></td><td id="file-fetch-issue-details-go-LC31"><span>Logs</span>: <span>logs</span>,</td></tr><tr><td id="file-fetch-issue-details-go-L32" data-line-number="32"></td><td id="file-fetch-issue-details-go-LC32"><span>Version</span>: <span>ver</span>,</td></tr><tr><td id="file-fetch-issue-details-go-L33" data-line-number="33"></td><td id="file-fetch-issue-details-go-LC33">}, <span>nil</span></td></tr><tr><td id="file-fetch-issue-details-go-L34" data-line-number="34"></td><td id="file-fetch-issue-details-go-LC34">}</td></tr></tbody></table>

Here, we fetch logs and the cluster version, but you can extend it to fetch other details about your cluster. To make it as simple as possible, we use the kubectl CLI and avoid reimplementing a bit more complicated logic for fetching logs from all different Kubernetes kinds.

‍

7.  Render issue description:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="render-issue.go"><tbody><tr><td id="file-render-issue-go-L1" data-line-number="1"></td><td id="file-render-issue-go-LC1"><span>import</span> (</td></tr><tr><td id="file-render-issue-go-L2" data-line-number="2"></td><td id="file-render-issue-go-LC2"><span>"bytes"</span></td></tr><tr><td id="file-render-issue-go-L3" data-line-number="3"></td><td id="file-render-issue-go-LC3"><span>"context"</span></td></tr><tr><td id="file-render-issue-go-L4" data-line-number="4"></td><td id="file-render-issue-go-LC4"><span>"fmt"</span></td></tr><tr><td id="file-render-issue-go-L5" data-line-number="5"></td><td id="file-render-issue-go-LC5"><span>"text/template"</span></td></tr><tr><td id="file-render-issue-go-L6" data-line-number="6"></td><td id="file-render-issue-go-LC6"></td></tr><tr><td id="file-render-issue-go-L7" data-line-number="7"></td><td id="file-render-issue-go-LC7"><span>"github.com/kubeshop/botkube/pkg/api/executor"</span></td></tr><tr><td id="file-render-issue-go-L8" data-line-number="8"></td><td id="file-render-issue-go-LC8">)</td></tr><tr><td id="file-render-issue-go-L9" data-line-number="9"></td><td id="file-render-issue-go-LC9"></td></tr><tr><td id="file-render-issue-go-L10" data-line-number="10"></td><td id="file-render-issue-go-LC10"><span>// Execute returns a given command as a response.</span></td></tr><tr><td id="file-render-issue-go-L11" data-line-number="11"></td><td id="file-render-issue-go-LC11"><span>func</span> (<span>e</span> <span>*</span><span>GHExecutor</span>) <span>Execute</span>(<span>ctx</span> context.<span>Context</span>, <span>in</span> executor.<span>ExecuteInput</span>) (executor.<span>ExecuteOutput</span>, <span>error</span>) {</td></tr><tr><td id="file-render-issue-go-L12" data-line-number="12"></td><td id="file-render-issue-go-LC12"><span>// ...</span></td></tr><tr><td id="file-render-issue-go-L13" data-line-number="13"></td><td id="file-render-issue-go-LC13"><span>mdBody</span>, <span>_</span> <span>:=</span> <span>renderIssueBody</span>(<span>cfg</span>.<span>GitHub</span>.<span>IssueTemplate</span>, <span>issueDetails</span>)</td></tr><tr><td id="file-render-issue-go-L14" data-line-number="14"></td><td id="file-render-issue-go-LC14"><span>// ...</span></td></tr><tr><td id="file-render-issue-go-L15" data-line-number="15"></td><td id="file-render-issue-go-LC15">}</td></tr><tr><td id="file-render-issue-go-L16" data-line-number="16"></td><td id="file-render-issue-go-LC16"></td></tr><tr><td id="file-render-issue-go-L17" data-line-number="17"></td><td id="file-render-issue-go-LC17"><span>const</span> <span>defaultIssueBody</span> <span>=</span> <span>`</span></td></tr><tr><td id="file-render-issue-go-L18" data-line-number="18"></td><td id="file-render-issue-go-LC18"><span>## Description</span></td></tr><tr><td id="file-render-issue-go-L19" data-line-number="19"></td><td id="file-render-issue-go-LC19"><span></span></td></tr><tr><td id="file-render-issue-go-L20" data-line-number="20"></td><td id="file-render-issue-go-LC20"><span>This issue refers to the problems connected with {{ .Type | code "bash" }} in namespace {{ .Namespace | code "bash" }}</span></td></tr><tr><td id="file-render-issue-go-L21" data-line-number="21"></td><td id="file-render-issue-go-LC21"><span></span></td></tr><tr><td id="file-render-issue-go-L22" data-line-number="22"></td><td id="file-render-issue-go-LC22"><span><details></details></span></td></tr><tr><td id="file-render-issue-go-L23" data-line-number="23"></td><td id="file-render-issue-go-LC23"><span><summary><b>Logs</b></summary></span></td></tr><tr><td id="file-render-issue-go-L24" data-line-number="24"></td><td id="file-render-issue-go-LC24"><span>{{ .Logs | code "bash"}}</span></td></tr><tr><td id="file-render-issue-go-L25" data-line-number="25"></td><td id="file-render-issue-go-LC25"><span></span></td></tr><tr><td id="file-render-issue-go-L26" data-line-number="26"></td><td id="file-render-issue-go-LC26"><span></span></td></tr><tr><td id="file-render-issue-go-L27" data-line-number="27"></td><td id="file-render-issue-go-LC27"><span>### Cluster details</span></td></tr><tr><td id="file-render-issue-go-L28" data-line-number="28"></td><td id="file-render-issue-go-LC28"><span></span></td></tr><tr><td id="file-render-issue-go-L29" data-line-number="29"></td><td id="file-render-issue-go-LC29"><span>{{ .Version | code "yaml"}}</span></td></tr><tr><td id="file-render-issue-go-L30" data-line-number="30"></td><td id="file-render-issue-go-LC30"><span>`</span></td></tr><tr><td id="file-render-issue-go-L31" data-line-number="31"></td><td id="file-render-issue-go-LC31"></td></tr><tr><td id="file-render-issue-go-L32" data-line-number="32"></td><td id="file-render-issue-go-LC32"><span>func</span> <span>renderIssueBody</span>(<span>bodyTpl</span> <span>string</span>, <span>data</span> <span>IssueDetails</span>) (<span>string</span>, <span>error</span>) {</td></tr><tr><td id="file-render-issue-go-L33" data-line-number="33"></td><td id="file-render-issue-go-LC33"><span>if</span> <span>bodyTpl</span> <span>==</span> <span>""</span> {</td></tr><tr><td id="file-render-issue-go-L34" data-line-number="34"></td><td id="file-render-issue-go-LC34"><span>bodyTpl</span> <span>=</span> <span>defaultIssueBody</span></td></tr><tr><td id="file-render-issue-go-L35" data-line-number="35"></td><td id="file-render-issue-go-LC35">}</td></tr><tr><td id="file-render-issue-go-L36" data-line-number="36"></td><td id="file-render-issue-go-LC36"></td></tr><tr><td id="file-render-issue-go-L37" data-line-number="37"></td><td id="file-render-issue-go-LC37"><span>tmpl</span>, <span>_</span> <span>:=</span> <span>template</span>.<span>New</span>(<span>"issue-body"</span>).<span>Funcs</span>(template.<span>FuncMap</span>{</td></tr><tr><td id="file-render-issue-go-L38" data-line-number="38"></td><td id="file-render-issue-go-LC38"><span>"code"</span>: <span>func</span>(<span>syntax</span>, <span>in</span> <span>string</span>) <span>string</span> {</td></tr><tr><td id="file-render-issue-go-L39" data-line-number="39"></td><td id="file-render-issue-go-LC39"><span>return</span> <span>fmt</span>.<span>Sprintf</span>(<span>"<span>\n</span>```%s<span>\n</span>%s<span>\n</span>```<span>\n</span>"</span>, <span>syntax</span>, <span>in</span>)</td></tr><tr><td id="file-render-issue-go-L40" data-line-number="40"></td><td id="file-render-issue-go-LC40">},</td></tr><tr><td id="file-render-issue-go-L41" data-line-number="41"></td><td id="file-render-issue-go-LC41">}).<span>Parse</span>(<span>bodyTpl</span>)</td></tr><tr><td id="file-render-issue-go-L42" data-line-number="42"></td><td id="file-render-issue-go-LC42"></td></tr><tr><td id="file-render-issue-go-L43" data-line-number="43"></td><td id="file-render-issue-go-LC43"><span>var</span> <span>body</span> bytes.<span>Buffer</span></td></tr><tr><td id="file-render-issue-go-L44" data-line-number="44"></td><td id="file-render-issue-go-LC44"><span>tmpl</span>.<span>Execute</span>(<span>&amp;</span><span>body</span>, <span>data</span>)</td></tr><tr><td id="file-render-issue-go-L45" data-line-number="45"></td><td id="file-render-issue-go-LC45"></td></tr><tr><td id="file-render-issue-go-L46" data-line-number="46"></td><td id="file-render-issue-go-LC46"><span>return</span> <span>body</span>.<span>String</span>(), <span>nil</span></td></tr><tr><td id="file-render-issue-go-L47" data-line-number="47"></td><td id="file-render-issue-go-LC47">}</td></tr></tbody></table>

In this step, we use the issue template that the user specified in plugin configuration. I decided to use the [Go template](https://pkg.go.dev/text/template), as it fits perfectly into our template rendering flow.

‍

8.  Finally! Submitting an issue!

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="submit-issue.go"><tbody><tr><td id="file-submit-issue-go-L1" data-line-number="1"></td><td id="file-submit-issue-go-LC1"><span>package</span> main</td></tr><tr><td id="file-submit-issue-go-L2" data-line-number="2"></td><td id="file-submit-issue-go-LC2"></td></tr><tr><td id="file-submit-issue-go-L3" data-line-number="3"></td><td id="file-submit-issue-go-LC3"><span>import</span> (</td></tr><tr><td id="file-submit-issue-go-L4" data-line-number="4"></td><td id="file-submit-issue-go-LC4"><span>"context"</span></td></tr><tr><td id="file-submit-issue-go-L5" data-line-number="5"></td><td id="file-submit-issue-go-LC5"><span>"fmt"</span></td></tr><tr><td id="file-submit-issue-go-L6" data-line-number="6"></td><td id="file-submit-issue-go-LC6"></td></tr><tr><td id="file-submit-issue-go-L7" data-line-number="7"></td><td id="file-submit-issue-go-LC7"><span>"github.com/kubeshop/botkube/pkg/api/executor"</span></td></tr><tr><td id="file-submit-issue-go-L8" data-line-number="8"></td><td id="file-submit-issue-go-LC8"><span>"github.com/kubeshop/botkube/pkg/pluginx"</span></td></tr><tr><td id="file-submit-issue-go-L9" data-line-number="9"></td><td id="file-submit-issue-go-LC9">)</td></tr><tr><td id="file-submit-issue-go-L10" data-line-number="10"></td><td id="file-submit-issue-go-LC10"><span>// Execute returns a given command as a response.</span></td></tr><tr><td id="file-submit-issue-go-L11" data-line-number="11"></td><td id="file-submit-issue-go-LC11"><span>func</span> (<span>e</span> <span>*</span><span>GHExecutor</span>) <span>Execute</span>(<span>ctx</span> context.<span>Context</span>, <span>in</span> executor.<span>ExecuteInput</span>) (executor.<span>ExecuteOutput</span>, <span>error</span>) {</td></tr><tr><td id="file-submit-issue-go-L12" data-line-number="12"></td><td id="file-submit-issue-go-LC12"><span>// ...</span></td></tr><tr><td id="file-submit-issue-go-L13" data-line-number="13"></td><td id="file-submit-issue-go-LC13"><span>title</span> <span>:=</span> <span>fmt</span>.<span>Sprintf</span>(<span>"The `%s` malfunctions"</span>, <span>cmd</span>.<span>Create</span>.<span>Issue</span>.<span>Type</span>)</td></tr><tr><td id="file-submit-issue-go-L14" data-line-number="14"></td><td id="file-submit-issue-go-LC14"><span>issueURL</span>, <span>_</span> <span>:=</span> <span>createGitHubIssue</span>(<span>cfg</span>, <span>title</span>, <span>mdBody</span>)</td></tr><tr><td id="file-submit-issue-go-L15" data-line-number="15"></td><td id="file-submit-issue-go-LC15"></td></tr><tr><td id="file-submit-issue-go-L16" data-line-number="16"></td><td id="file-submit-issue-go-LC16"><span>return</span> executor.<span>ExecuteOutput</span>{</td></tr><tr><td id="file-submit-issue-go-L17" data-line-number="17"></td><td id="file-submit-issue-go-LC17"><span>Data</span>: <span>fmt</span>.<span>Sprintf</span>(<span>"New issue created successfully! 🎉<span>\n</span><span>\n</span>Issue URL: %s"</span>, <span>issueURL</span>),</td></tr><tr><td id="file-submit-issue-go-L18" data-line-number="18"></td><td id="file-submit-issue-go-LC18">}, <span>nil</span></td></tr><tr><td id="file-submit-issue-go-L19" data-line-number="19"></td><td id="file-submit-issue-go-LC19">}</td></tr><tr><td id="file-submit-issue-go-L20" data-line-number="20"></td><td id="file-submit-issue-go-LC20"></td></tr><tr><td id="file-submit-issue-go-L21" data-line-number="21"></td><td id="file-submit-issue-go-LC21"><span>func</span> <span>createGitHubIssue</span>(<span>cfg</span> <span>Config</span>, <span>title</span>, <span>mdBody</span> <span>string</span>) (<span>string</span>, <span>error</span>) {</td></tr><tr><td id="file-submit-issue-go-L22" data-line-number="22"></td><td id="file-submit-issue-go-LC22"><span>cmd</span> <span>:=</span> <span>fmt</span>.<span>Sprintf</span>(<span>"gh issue create --title %q --body '%s' --label bug -R %s"</span>, <span>title</span>, <span>mdBody</span>, <span>cfg</span>.<span>GitHub</span>.<span>Repository</span>)</td></tr><tr><td id="file-submit-issue-go-L23" data-line-number="23"></td><td id="file-submit-issue-go-LC23"></td></tr><tr><td id="file-submit-issue-go-L24" data-line-number="24"></td><td id="file-submit-issue-go-LC24"><span>envs</span> <span>:=</span> <span>map</span>[<span>string</span>]<span>string</span>{</td></tr><tr><td id="file-submit-issue-go-L25" data-line-number="25"></td><td id="file-submit-issue-go-LC25"><span>"GH_TOKEN"</span>: <span>cfg</span>.<span>GitHub</span>.<span>Token</span>,</td></tr><tr><td id="file-submit-issue-go-L26" data-line-number="26"></td><td id="file-submit-issue-go-LC26">}</td></tr><tr><td id="file-submit-issue-go-L27" data-line-number="27"></td><td id="file-submit-issue-go-LC27"></td></tr><tr><td id="file-submit-issue-go-L28" data-line-number="28"></td><td id="file-submit-issue-go-LC28"><span>return</span> <span>pluginx</span>.<span>ExecuteCommandWithEnvs</span>(<span>context</span>.<span>Background</span>(), <span>cmd</span>, <span>envs</span>)</td></tr><tr><td id="file-submit-issue-go-L29" data-line-number="29"></td><td id="file-submit-issue-go-LC29">}</td></tr></tbody></table>

GitHub provides a great [gh](https://cli.github.com/) CLI, which we use to submit our issue. To learn more about the CLI syntax, see their [manual](https://cli.github.com/manual/gh_issue_create).

The `gh` CLI doesn't accept fine-grained access tokens. As a workaround, you can use the [Go SDK](https://gist.github.com/mszostok/defa5a5390e87b4f011b986742f714d8).

‍

9.  The last part is to download our dependencies.

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="download-deps.go"><tbody><tr><td id="file-download-deps-go-L1" data-line-number="1"></td><td id="file-download-deps-go-LC1"><span>var</span> <span>depsDownloadLinks</span> <span>=</span> <span>map</span>[<span>string</span>]api.<span>Dependency</span>{</td></tr><tr><td id="file-download-deps-go-L2" data-line-number="2"></td><td id="file-download-deps-go-LC2"><span>"gh"</span>: {</td></tr><tr><td id="file-download-deps-go-L3" data-line-number="3"></td><td id="file-download-deps-go-LC3"><span>URLs</span>: <span>map</span>[<span>string</span>]<span>string</span>{</td></tr><tr><td id="file-download-deps-go-L4" data-line-number="4"></td><td id="file-download-deps-go-LC4"><span>"darwin/amd64"</span>: <span>"https://github.com/cli/cli/releases/download/v2.21.2/gh_2.21.2_macOS_amd64.tar.gz//gh_2.21.2_macOS_amd64/bin"</span>,</td></tr><tr><td id="file-download-deps-go-L5" data-line-number="5"></td><td id="file-download-deps-go-LC5"><span>"linux/amd64"</span>: <span>"https://github.com/cli/cli/releases/download/v2.21.2/gh_2.21.2_linux_amd64.tar.gz//gh_2.21.2_linux_amd64/bin"</span>,</td></tr><tr><td id="file-download-deps-go-L6" data-line-number="6"></td><td id="file-download-deps-go-LC6"><span>"linux/arm64"</span>: <span>"https://github.com/cli/cli/releases/download/v2.21.2/gh_2.21.2_linux_arm64.tar.gz//gh_2.21.2_linux_arm64/bin"</span>,</td></tr><tr><td id="file-download-deps-go-L7" data-line-number="7"></td><td id="file-download-deps-go-LC7"><span>"linux/386"</span>: <span>"https://github.com/cli/cli/releases/download/v2.21.2/gh_2.21.2_linux_386.tar.gz//gh_2.21.2_linux_386/bin"</span>,</td></tr><tr><td id="file-download-deps-go-L8" data-line-number="8"></td><td id="file-download-deps-go-LC8">},</td></tr><tr><td id="file-download-deps-go-L9" data-line-number="9"></td><td id="file-download-deps-go-LC9">},</td></tr><tr><td id="file-download-deps-go-L10" data-line-number="10"></td><td id="file-download-deps-go-LC10"><span>"kubectl"</span>: {</td></tr><tr><td id="file-download-deps-go-L11" data-line-number="11"></td><td id="file-download-deps-go-LC11"><span>URLs</span>: <span>map</span>[<span>string</span>]<span>string</span>{</td></tr><tr><td id="file-download-deps-go-L12" data-line-number="12"></td><td id="file-download-deps-go-LC12"><span>"darwin/amd64"</span>: <span>"https://dl.k8s.io/release/v1.26.0/bin/darwin/amd64/kubectl"</span>,</td></tr><tr><td id="file-download-deps-go-L13" data-line-number="13"></td><td id="file-download-deps-go-LC13"><span>"linux/amd64"</span>: <span>"https://dl.k8s.io/release/v1.26.0/bin/linux/amd64/kubectl"</span>,</td></tr><tr><td id="file-download-deps-go-L14" data-line-number="14"></td><td id="file-download-deps-go-LC14"><span>"linux/arm64"</span>: <span>"https://dl.k8s.io/release/v1.26.0/bin/linux/arm64/kubectl"</span>,</td></tr><tr><td id="file-download-deps-go-L15" data-line-number="15"></td><td id="file-download-deps-go-LC15"><span>"linux/386"</span>: <span>"https://dl.k8s.io/release/v1.26.0/bin/linux/386/kubectl"</span>,</td></tr><tr><td id="file-download-deps-go-L16" data-line-number="16"></td><td id="file-download-deps-go-LC16">},</td></tr><tr><td id="file-download-deps-go-L17" data-line-number="17"></td><td id="file-download-deps-go-LC17">},</td></tr><tr><td id="file-download-deps-go-L18" data-line-number="18"></td><td id="file-download-deps-go-LC18">}</td></tr><tr><td id="file-download-deps-go-L19" data-line-number="19"></td><td id="file-download-deps-go-LC19"></td></tr><tr><td id="file-download-deps-go-L20" data-line-number="20"></td><td id="file-download-deps-go-LC20"><span>func</span> <span>main</span>() {</td></tr><tr><td id="file-download-deps-go-L21" data-line-number="21"></td><td id="file-download-deps-go-LC21"><span>pluginx</span>.<span>DownloadDependencies</span>(<span>depsDownloadLinks</span>)</td></tr><tr><td id="file-download-deps-go-L22" data-line-number="22"></td><td id="file-download-deps-go-LC22"><span>//...</span></td></tr><tr><td id="file-download-deps-go-L23" data-line-number="23"></td><td id="file-download-deps-go-LC23">}</td></tr></tbody></table>

We already improved this step and in the 0.18 version Botkube will download defined [dependencies automatically](https://docs.botkube.io/next/plugin/dependencies). For now, you can use the `pluginx.DownloadDependencies` function to call our downloader explicitly. The syntax will remain the same.

‍

Congrats! The _gh_ plugin is finally implemented. Now, let's play a DevOps role! 😈 In the next section, I will show you how to build and release your brand-new executor plugin.

### Release the _gh_ executor

It's time to build your plugin. For that purpose, we will use [GoReleaser](https://goreleaser.com/). It simplifies building Go binaries for different architectures. The important thing is to produce the binaries for the architecture of the host platform where Botkube is running. Adjust the **goos**, **goarch**, and **goarm** properties based on this architecture.

‍ Add new build entry under _.goreleaser.yaml_: ‍

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="build.yml"><tbody><tr><td id="file-build-yml-L1" data-line-number="1"></td><td id="file-build-yml-LC1"><span>builds</span>:</td></tr><tr><td id="file-build-yml-L2" data-line-number="2"></td><td id="file-build-yml-LC2">- <span>id</span>: <span>gh</span></td></tr><tr><td id="file-build-yml-L3" data-line-number="3"></td><td id="file-build-yml-LC3"><span>main</span>: <span>cmd/gh/main.go</span></td></tr><tr><td id="file-build-yml-L4" data-line-number="4"></td><td id="file-build-yml-LC4"><span>binary</span>: <span>executor_gh_{{ .Os }}_{{ .Arch }}</span></td></tr><tr><td id="file-build-yml-L5" data-line-number="5"></td><td id="file-build-yml-LC5"></td></tr><tr><td id="file-build-yml-L6" data-line-number="6"></td><td id="file-build-yml-LC6"><span>no_unique_dist_dir</span>: <span>true</span></td></tr><tr><td id="file-build-yml-L7" data-line-number="7"></td><td id="file-build-yml-LC7"><span>env</span>:</td></tr><tr><td id="file-build-yml-L8" data-line-number="8"></td><td id="file-build-yml-LC8">- <span>CGO_ENABLED=0</span></td></tr><tr><td id="file-build-yml-L9" data-line-number="9"></td><td id="file-build-yml-LC9"><span>goos</span>:</td></tr><tr><td id="file-build-yml-L10" data-line-number="10"></td><td id="file-build-yml-LC10">- <span>linux</span></td></tr><tr><td id="file-build-yml-L11" data-line-number="11"></td><td id="file-build-yml-LC11">- <span>darwin</span></td></tr><tr><td id="file-build-yml-L12" data-line-number="12"></td><td id="file-build-yml-LC12"><span>goarch</span>:</td></tr><tr><td id="file-build-yml-L13" data-line-number="13"></td><td id="file-build-yml-LC13">- <span>amd64</span></td></tr><tr><td id="file-build-yml-L14" data-line-number="14"></td><td id="file-build-yml-LC14">- <span>arm64</span></td></tr><tr><td id="file-build-yml-L15" data-line-number="15"></td><td id="file-build-yml-LC15"><span>goarm</span>:</td></tr><tr><td id="file-build-yml-L16" data-line-number="16"></td><td id="file-build-yml-LC16">- <span>7</span></td></tr></tbody></table>

‍ Now, we need to distribute our plugins. As we mentioned earlier, a plugin repository can be any static file server. The [kubeshop/botkube-plugins-template](https://github.com/kubeshop/botkube-plugins-template) repository comes with two [GitHub Actions](https://github.com/features/actions):

1.  The [.github/workflows/release.yml](https://github.com/kubeshop/botkube-plugins-template/blob/main/.github/workflows/release.yml) action, which builds the plugin binaries and index file each time a new tag is pushed.
    
2.  The [.github/workflows/pages-release.yml](https://github.com/kubeshop/botkube-plugins-template/blob/main/.github/workflows/pages-release.yml) action, which updates GitHub Pages with plugin binaries and index file each time a new tag is pushed.
    

To cut a new release, you need to commit all your work and tag a new commit:

Next, let's push our changes and the new tag:

This triggers GitHub Action:

![Image 6](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63d817a66682780daf595ab4_release-job.png)

### What this automation does under the hood

This automation:

1.  Installs the latest GoReleaser tool.
    
2.  Builds all plugin binaries defined in the _.goreleaser.yaml_ file.
    
3.  Generates an index file using the [Botkube helper tool](https://docs.botkube.io/plugin/repo#generate-index-file).
    
4.  Generates a release description.
    
5.  Uses the [gh](https://cli.github.com/) CLI to create a new GitHub release.
    

### Use the _gh_ executor

In the description of a new GitHub release, you will see the repository URL that you can use within Botkube.

### Steps

1.  Follow one of our [installation guides](https://docs.botkube.io/installation). Once you reach the Botkube deployment step, add flags specified in the steps below to _helm install_.
    
2.  Export required environment variables:
    

Follow the official GitHub guide on how to create a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#creating-a-personal-access-token-classic). To be able to create GitHub issues, add the **repo** permission.

3.  Add the _gh_ executor related configuration:

4.  Depending on the selected platform, add the _plugin-based_ executor binding. For Slack, it looks like this:

If you follow all the steps above, you will have all the necessary flags allowing you to install Botkube with the **_gh_** executor!

Here's an example of a full command that you should have constructed for Slack installation:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Shell" data-tagsearch-path="run.sh"><tbody><tr><td id="file-run-sh-L1" data-line-number="1"></td><td id="file-run-sh-LC1"><span>export</span> SLACK_CHANNEL_NAME={channel_name}</td></tr><tr><td id="file-run-sh-L2" data-line-number="2"></td><td id="file-run-sh-LC2"><span>export</span> SLACK_APP_TOKEN={token}</td></tr><tr><td id="file-run-sh-L3" data-line-number="3"></td><td id="file-run-sh-LC3"><span>export</span> SLACK_BOT_TOKEN={token}</td></tr><tr><td id="file-run-sh-L4" data-line-number="4"></td><td id="file-run-sh-LC4"><span>export</span> PLUGINS_URL={plugin_index_url}</td></tr><tr><td id="file-run-sh-L5" data-line-number="5"></td><td id="file-run-sh-LC5"><span>export</span> REPOSITORY={repo} <span><span>#</span> format OWNER/REPO_NAME, e.g. kubeshop/botkube</span></td></tr><tr><td id="file-run-sh-L6" data-line-number="6"></td><td id="file-run-sh-LC6"><span>export</span> GITHUB_TOKEN={token}</td></tr><tr><td id="file-run-sh-L7" data-line-number="7"></td><td id="file-run-sh-LC7"></td></tr><tr><td id="file-run-sh-L8" data-line-number="8"></td><td id="file-run-sh-LC8">helm install --version v0.17.0 botkube --namespace botkube --create-namespace \</td></tr><tr><td id="file-run-sh-L9" data-line-number="9"></td><td id="file-run-sh-LC9">--set communications.default-group.socketSlack.enabled=true \</td></tr><tr><td id="file-run-sh-L10" data-line-number="10"></td><td id="file-run-sh-LC10">--set communications.default-group.socketSlack.channels.default.name=<span>${SLACK_CHANNEL_NAME}</span> \</td></tr><tr><td id="file-run-sh-L11" data-line-number="11"></td><td id="file-run-sh-LC11">--set communications.default-group.socketSlack.channels.default.bindings.executors={<span><span>"</span>plugin-based<span>"</span></span>} \</td></tr><tr><td id="file-run-sh-L12" data-line-number="12"></td><td id="file-run-sh-LC12">--set communications.default-group.socketSlack.appToken=<span>${SLACK_APP_TOKEN}</span> \</td></tr><tr><td id="file-run-sh-L13" data-line-number="13"></td><td id="file-run-sh-LC13">--set communications.default-group.socketSlack.botToken=<span>${SLACK_BOT_TOKEN}</span> \</td></tr><tr><td id="file-run-sh-L14" data-line-number="14"></td><td id="file-run-sh-LC14">--set <span><span>'</span>plugins.repositories.botkube-plugins.url<span>'</span></span>=<span>${PLUGINS_URL}</span> \</td></tr><tr><td id="file-run-sh-L15" data-line-number="15"></td><td id="file-run-sh-LC15">--set <span><span>'</span>executors.plugin-based.botkube-plugins/gh.config.github.repository<span>'</span></span>=<span>${REPOSITORY}</span> \</td></tr><tr><td id="file-run-sh-L16" data-line-number="16"></td><td id="file-run-sh-LC16">--set <span><span>'</span>executors.plugin-based.botkube-plugins/gh.config.github.token<span>'</span></span>=<span>${GITHUB_TOKEN}</span> \</td></tr><tr><td id="file-run-sh-L17" data-line-number="17"></td><td id="file-run-sh-LC17">botkube/botkube</td></tr></tbody></table>

### Testing

1.  Navigate to your communication channel.
    
2.  On a given channel, run: **@Botkube list executors**
    

![Image 7](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63d81ba395f87021e5d45685_list-exec.png)

It returns information about enabled _gh_ executor

3.  Create a failing Job:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="failing-job.yaml"><tbody><tr><td id="file-failing-job-yaml-L1" data-line-number="1"></td><td id="file-failing-job-yaml-LC1"><span>apiVersion</span>: <span>batch/v1</span></td></tr><tr><td id="file-failing-job-yaml-L2" data-line-number="2"></td><td id="file-failing-job-yaml-LC2"><span>kind</span>: <span>Job</span></td></tr><tr><td id="file-failing-job-yaml-L3" data-line-number="3"></td><td id="file-failing-job-yaml-LC3"><span>metadata</span>:</td></tr><tr><td id="file-failing-job-yaml-L4" data-line-number="4"></td><td id="file-failing-job-yaml-LC4"><span>name</span>: <span>oops</span></td></tr><tr><td id="file-failing-job-yaml-L5" data-line-number="5"></td><td id="file-failing-job-yaml-LC5"><span>spec</span>:</td></tr><tr><td id="file-failing-job-yaml-L6" data-line-number="6"></td><td id="file-failing-job-yaml-LC6"><span>backoffLimit</span>: <span>1</span></td></tr><tr><td id="file-failing-job-yaml-L7" data-line-number="7"></td><td id="file-failing-job-yaml-LC7"><span>template</span>:</td></tr><tr><td id="file-failing-job-yaml-L8" data-line-number="8"></td><td id="file-failing-job-yaml-LC8"><span>spec</span>:</td></tr><tr><td id="file-failing-job-yaml-L9" data-line-number="9"></td><td id="file-failing-job-yaml-LC9"><span>restartPolicy</span>: <span>Never</span></td></tr><tr><td id="file-failing-job-yaml-L10" data-line-number="10"></td><td id="file-failing-job-yaml-LC10"><span>containers</span>:</td></tr><tr><td id="file-failing-job-yaml-L11" data-line-number="11"></td><td id="file-failing-job-yaml-LC11">- <span>name</span>: <span>main</span></td></tr><tr><td id="file-failing-job-yaml-L12" data-line-number="12"></td><td id="file-failing-job-yaml-LC12"><span>image</span>: <span>docker.io/library/bash:5</span></td></tr><tr><td id="file-failing-job-yaml-L13" data-line-number="13"></td><td id="file-failing-job-yaml-LC13"><span>command</span>: <span>["bash"]</span></td></tr><tr><td id="file-failing-job-yaml-L14" data-line-number="14"></td><td id="file-failing-job-yaml-LC14"><span>args</span>:</td></tr><tr><td id="file-failing-job-yaml-L15" data-line-number="15"></td><td id="file-failing-job-yaml-LC15">- <span>-c</span></td></tr><tr><td id="file-failing-job-yaml-L16" data-line-number="16"></td><td id="file-failing-job-yaml-LC16">- <span>echo "Hello from failing job!" &amp;&amp; exit 42</span></td></tr></tbody></table>

‍

![Image 8](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/63d81bb595f870bd6ed45689_alert.png)

After a few seconds, you should see a new alert on your channel

4.  To create a GitHub issue for a given alert, run: **_@Botkube gh create issue job/oops_**‍

Summary
-------

Botkube executors are powerful because they can glue together three important parts: Kubernetes clusters, communicators, and tools of your choice. There would be nothing special about it if it wasn't, in fact, unburdening you of those implementation-specific details.

As you noticed, you can focus purely on your business logic. Without the need to use different chat libraries, know how to establish secure connection, or make your extension available only on specific channels. What's more, not only do you not have to learn it, but you don't have to support it either, as we do it for you.

Once Botkube is deployed, your extension will be available to you and your teammates in a given channel. There is no need to maintain your local setup. Thanks to that, you can also easily run executors on private clusters.

Botkube extensions can be used with other Botkube functionality too. It means that you can use them to create [automation](https://docs.botkube.io/configuration/action). We will shed more light on that in the next blog post. Stay tuned!

> Implement once, access everywhere (Slack, Discord, Mattermost, MS Teams).

How can I get involved?
-----------------------

The implemented plugin is as simple as possible. However, it is a great base for further extension based on your needs; for example: introduce your own Kubernetes annotation to route the notification to a specific repository, add a threshold to create issues only for constantly failing Pods, etc. The possibilities are endless, and we cannot wait to see what kind of great workflows you will create!

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap, get the features you’d like implemented.

There are plenty of options to contact us:

*   [GitHub issues](https://github.com/kubeshop/botkube/issues)
    
*   [Slack](https://join.botkube.io/)
    
*   or email our Product Leader at blair@kubeshop.io.
    

‍ Thank you for taking the time to learn about Botkube 🙌
