Title: Creating the Botkube Flux plugin for Day 2 operations

URL Source: https://botkube.io/blog/creating-the-botkube-flux-plugin-for-day-2-operations

Published Time: Sep 07, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/636df3ec19d99029e7cfb5bd_n-vUM0Lz3yXA1WLyfxEMUxXQpUDEiKv9EmEe4NnCbCI.png)

Mateusz Szostok

Software Engineer

Botkube

From this deep dive you'll know the Zapier-like aspect of Botkube that connects K8s clusters, GitHub, and the Flux CLI—all of this to make you fall in love with your Day 2 operations.

### Table of Contents

*   [Introduction](#introduction)
*   [The Evolution of Flux Executor Plugin](#the-evolution-of-flux-executor-plugin)
*   [Interactivity and Decision-Making](#interactivity-and-decision-making)
*   [Manual Approach vs. Botkube Flux Plugin](#manual-approach-vs-botkube-flux-plugin)
*   [Behind the Scenes: Developing the Botkube Flux Plugin](#behind-the-scenes-developing-the-botkube-flux-plugin)
*   [Conclusion](#conclusion)

#### Manage your Kubernetes Clusters Directly in Slack and Microsoft Teams!

#### Get started with Botkube Cloud

Hey there, tech enthusiasts! Today we're diving into the world of GitOps with the brand-new Botkube Flux plugin.

Introduction
------------

In Botkube, we move towards GitOps by developing extensions like the new [Flux plug-in](https://botkube.io/integration/botkube-flux-kubernetes-integration). We all know the importance of GitOps workflow automation—it's the backbone of modern Kubernetes management. But sometimes, we encounter situations where automation is needed but hasn't been implemented yet. That's where Botkube steps in.

The Botkube Flux plugin is designed to bridge the gap between your Kubernetes clusters and GitHub repositories. It enables you to execute commands directly from Slack using interactive forms. Plus, it keeps you in the loop by notifying you about GitHub pull requests that are altering your kustomization files. It provides a dedicated button to show a diff report between the pull request and your current cluster state.

In this blog post, we will reveal the cards and dive deep into the thought process and implementation details of the Flux plugin. You'll get to know the Zapier-like aspect of Botkube that connects Kubernetes clusters, GitHub, and the Flux CLI – all of this to make you fall in love with your Day 2 operations.

The Evolution of Flux Executor Plugin
-------------------------------------

Now, let's get into the nitty-gritty of this plugin's journey. Picture yourself as part of a team managing Kubernetes applications with multiple pull requests. Our goal is to integrate with Flux CD to simplify Day 2 operations for Flux users. Let's take a closer look at the user story:

1.  Be able to run Flux CLI commands from any communication platform, just like you do in your terminal:

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f9944f5ec10d12756e4d2b_flux-get-source-git.png)

‍

2.  Next, let's make it mobile friendly. The secret ingredient is interactivity like buttons and select menus. ‍

![Image 3](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f9949ad7623bfae8357d71_flux-get-source-git-btns.png)

Typing or copy-pasting a long names doesn't work well. Now, you have a handy Flux client right in your pocket, ready with just a few clicks. And we are just half-way there 😈‍

‍

3.  Here comes the last, but unique, part that makes the difference: **support for Day 2 operations**.

In our case, we stitched together three important parts: Kubernetes cluster, GitHub platform, and Flux CLI. As a result, we've provided a streamlined experience for generating a diff report in the context of GitHub pull requests and the current cluster state.

![Image 4](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f994c33a6f72e35362f50c_flux-diff.png)

🎁 As you may notice, the diff report notification includes some useful actions out-of-the-box:

*   Posting the diff report as a GitHub comment.
*   Approving the pull request.
*   Viewing the pull request.

4\. Now, when we're happy about the result, we are still missing one more part to **automate our day 2 operation.**

‍

Even though the diffing flow integrates with GitHub, it still requires two manual steps:

*   discovering that a new pull request was created
*   constructing a Flux related command

That's how the GitHub Events source was born. Now we can set up a complete workflow to:

1.  Watch for GitHub pull requests that changes files in `kustomize` directory. Alternatively, we can use label selectors.
2.  Get notification on Slack about new pull request.
3.  Render and embed event-aware button to run a diff report.

![Image 5](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f994dd9581964e18ef36bd_pr-event.png)

‍

Now, you may think that what we achieve in those 4 steps it's great but will be hard to configure. Is it? We hope that included YAML sample proves that it is not:

*   Flux Executor, with optional GitHub token:

*   GitHub Events workflow:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="gh-events-cfg.yaml"><tbody><tr><td id="file-gh-events-cfg-yaml-L1" data-line-number="1"></td><td id="file-gh-events-cfg-yaml-LC1"><span>botkube/github-events</span>: <span><span>#</span> GitHub Events</span></td></tr><tr><td id="file-gh-events-cfg-yaml-L2" data-line-number="2"></td><td id="file-gh-events-cfg-yaml-LC2"><span>github</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L3" data-line-number="3"></td><td id="file-gh-events-cfg-yaml-LC3"><span>auth</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L4" data-line-number="4"></td><td id="file-gh-events-cfg-yaml-LC4"><span>accessToken</span>: <span><span>"</span>ghp_<span>"</span></span></td></tr><tr><td id="file-gh-events-cfg-yaml-L5" data-line-number="5"></td><td id="file-gh-events-cfg-yaml-LC5"><span>repositories</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L6" data-line-number="6"></td><td id="file-gh-events-cfg-yaml-LC6">- <span>name</span>: <span>mszostok/podinfo</span></td></tr><tr><td id="file-gh-events-cfg-yaml-L7" data-line-number="7"></td><td id="file-gh-events-cfg-yaml-LC7"><span>on</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L8" data-line-number="8"></td><td id="file-gh-events-cfg-yaml-LC8"><span>pullRequests</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L9" data-line-number="9"></td><td id="file-gh-events-cfg-yaml-LC9">- <span>types</span>: <span>["open"]</span></td></tr><tr><td id="file-gh-events-cfg-yaml-L10" data-line-number="10"></td><td id="file-gh-events-cfg-yaml-LC10"><span>paths</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L11" data-line-number="11"></td><td id="file-gh-events-cfg-yaml-LC11"><span>include</span>: <span>["kustomize/.*"]</span></td></tr><tr><td id="file-gh-events-cfg-yaml-L12" data-line-number="12"></td><td id="file-gh-events-cfg-yaml-LC12"><span>notificationTemplate</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L13" data-line-number="13"></td><td id="file-gh-events-cfg-yaml-LC13"><span>extraButtons</span>:</td></tr><tr><td id="file-gh-events-cfg-yaml-L14" data-line-number="14"></td><td id="file-gh-events-cfg-yaml-LC14">- <span>displayName</span>: <span><span>"</span>Flux Diff<span>"</span></span></td></tr><tr><td id="file-gh-events-cfg-yaml-L15" data-line-number="15"></td><td id="file-gh-events-cfg-yaml-LC15"><span>commandTpl</span>: <span><span>"</span>flux diff ks podinfo --path ./kustomize --github-ref {{ .HTMLURL }} <span>"</span></span></td></tr></tbody></table>

We've kept the syntax generic, allowing you to configure different command buttons under specific types of events.

Interactivity and Decision-Making
---------------------------------

Thanks to automated notifications with event-aware buttons, you can easily generate reports and share them with your team or contributors.

While posting diff reports can be fully automated, you might want to do it intentionally by clicking a button. Why? Because the report may contain sensitive information that you don't want to fully disclose to external contributors.

Using, for example, GitHUb Action, unfortunately, doesn't give us this freedom. We tried to automate the whole workflow while still keeping you as the decision-maker when it comes to sharing the report or directly approving the PR.

With that being said, nothing blocks us from adding in the future support for AI assistance. Imagine an AI that reviews the diff report and decides whether to proceed with automated approval. Are you ready for AIOps? Exciting times ahead!

Manual Approach vs. Botkube Flux Plugin
---------------------------------------

While you were reading the first part of the Flux plugin evolution, did you consider what kind of manual steps would be required without the plugin? Let's break it down:

1.  Checking GitHub repository for a new pull requests.
2.  **(One-time operation)** Downloading and installing Flux CLI on you localhost.
3.  Manually connecting to the related Kubernetes cluster.
4.  **(One-time operation)** Cloning the repository.
5.  Checking out the pull request.
6.  Constructing a Flux command.
7.  Sharing the diff report on Slack/GitHub.

Even if we exclude the one-time operations, we're left with 5 steps for each new pull request. Lots of manual steps mean lots of room for human errors. Plus, all that jumping between different sites and context switching can impact your productivity. It's much better to focus on the main aspect, which is the review, and let automation handle the rest.

Behind the Scenes: Developing the Botkube Flux Plugin
-----------------------------------------------------

The development of the Botkube Flux Executor plugin involved several key aspects:

1.  🕹️ **Interactivity**: We leveraged the [`exec`](https://docs.botkube.io/usage/executor/exec#table-parser) plugin developed in previous release, making adding interactivity almost out-of-the-box. The `exec` plugin allows you to port any CLI into the communication platform window.
    
    In this case, we reused it as Go SDK. Here is the blueprint that describes translation of CLI table output into an interactive message: ‍
    

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="blueprint.yaml"><tbody><tr><td id="file-blueprint-yaml-L1" data-line-number="1"></td><td id="file-blueprint-yaml-LC1">- <span>trigger</span>:</td></tr><tr><td id="file-blueprint-yaml-L2" data-line-number="2"></td><td id="file-blueprint-yaml-LC2"><span>command</span>:</td></tr><tr><td id="file-blueprint-yaml-L3" data-line-number="3"></td><td id="file-blueprint-yaml-LC3"><span>regex</span>: <span><span>"</span>flux get sources (buc</span></td></tr><tr><td id="file-blueprint-yaml-L4" data-line-number="4"></td><td id="file-blueprint-yaml-LC4"><span>type: <span>"</span></span><span>parser:table:space"</span></td></tr><tr><td id="file-blueprint-yaml-L5" data-line-number="5"></td><td id="file-blueprint-yaml-LC5"><span>message</span>:</td></tr><tr><td id="file-blueprint-yaml-L6" data-line-number="6"></td><td id="file-blueprint-yaml-LC6"><span>selects</span>:</td></tr><tr><td id="file-blueprint-yaml-L7" data-line-number="7"></td><td id="file-blueprint-yaml-LC7">- <span>name</span>: <span><span>"</span>Item<span>"</span></span></td></tr><tr><td id="file-blueprint-yaml-L8" data-line-number="8"></td><td id="file-blueprint-yaml-LC8"><span>keyTpl</span>: <span><span>"</span>{{ .Name }}<span>"</span></span></td></tr></tbody></table>

This makes the implementation much simpler, so we could focus more on native support for the diffing flow.

‍

2.  🗃️ **Storing blueprints**: The blueprints used for converting CLI output into interactive messages are not stored directly in the Go code. Instead, we store them as YAML files in a dedicated folder so that IDEs, the GitHub platform, and others can provide us with valid syntax highlighting and formatting. Next, we use Go embed functionality, which is really handy here. With just three lines of code, Go embeds the entire directory with our manifests: ‍

Thanks to embedding it, we can distribute it as a single plugin binary, and we don't need to make any external download calls on startup.

‍

3.  🔍 & 🔐 **Auto-discovering GitHub repos**: In order to discover related GitHub repository, we need to get Flux custom resource. We used the [`controller-runtime`](https://github.com/kubernetes-sigs/controller-runtime/blob/main/pkg/client/client.go) client, which supports Go types natively. This eliminated the need to work with the unstructured type, making things smoother and less error-prone. This is backed by dedicated plugin **RBAC** impersonation that we introduced a couple releases ago.
    
4.  🔄 **Cloning and checking out PR**: Checking out a pull request can be tricky, especially when dealing with external contributors and their forks. Instead of reinventing the wheel, we integrated the widely-known `gh` CLI. It was easy to add an external dependency just by defining: ‍
    

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Go" data-tagsearch-path="deps.go"><tbody><tr><td id="file-deps-go-L1" data-line-number="1"></td><td id="file-deps-go-LC1"><span>"gh"</span>: {</td></tr><tr><td id="file-deps-go-L2" data-line-number="2"></td><td id="file-deps-go-LC2"><span>URLs</span>: <span>map</span>[<span>string</span>]<span>string</span>{</td></tr><tr><td id="file-deps-go-L3" data-line-number="3"></td><td id="file-deps-go-LC3"><span>"darwin/amd64"</span>: <span>"https://github.com/cli/cli/releases/download/v2.32.1/gh_2.32.1_macOS_amd64.zip//gh_2.32.1_macOS_amd64/bin"</span>,</td></tr><tr><td id="file-deps-go-L4" data-line-number="4"></td><td id="file-deps-go-LC4"><span>"linux/amd64"</span>: <span>"https://github.com/cli/cli/releases/download/v2.32.1/gh_2.32.1_linux_amd64.tar.gz//gh_2.32.1_linux_amd64/bin"</span>,</td></tr><tr><td id="file-deps-go-L5" data-line-number="5"></td><td id="file-deps-go-LC5"><span>// etc.</span></td></tr><tr><td id="file-deps-go-L6" data-line-number="6"></td><td id="file-deps-go-LC6">},</td></tr><tr><td id="file-deps-go-L7" data-line-number="7"></td><td id="file-deps-go-LC7">},</td></tr></tbody></table>

Under the hood we use \`go-getter\` library which has a lot of great features. If you need to download assets from different sources, we recommend that library your projects as well.

‍ ‍

The trickiest part was to develop GitHub Events source. The best way is to use GitHub App with the webhook approach. However, we didn't want to require you to have an external endpoint exposed from your cluster.

We started with [GitHub repository events endpoint](https://docs.github.com/en/rest/activity/events?apiVersion=2022-11-28#list-repository-events). But it turned out that even though it serves events that we are interested in, it was not meant to be used for the real-time use-cases. We still integrate with the `events` API, but it's recommended for event subscription where time is not that important. For example, getting notification about new stars on your GitHub repositories: ‍

![Image 6](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64f994ee3a6f72e3536300f4_stars.png)

‍ To achieve our e2e goal, we decided to develop a custom polling mechanism that uses [pull request endpoint](https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-pull-requests). The polling mechanism forces us to be more rational about the number of calls to fit into defined rate limits. We decided on two things:

1.  [Conditional requests](https://docs.github.com/en/rest/overview/resources-in-the-rest-api?apiVersion=2022-11-28#conditional-requests) because receiving a 304 response doesn't count against token rate limit.
2.  Adding support for GitHub App tokens. By using GitHub Apps, you can increase your maximum rate limits because multiple GitHub Apps are independent and do not share the rate limits. Where, using multiple Personal Access Tokens (PATs) for the same account will result in sharing the same rate limit.

In the future, we can consider adding a token rotator that automatically switches tokens before hitting the rate limit.

For the [Botkube web app](https://app.botkube.io/) we will consider native integration using GitHub App, to reduce friction with the initial setup for Flux and GitHub Events plugins.

Conclusion
----------

The [Botkube Flux plug-in](https://botkube.io/blog/introducing-botkubes-integration-with-flux) offers valuable solutions for streamlining GitOps workflows. Its capabilities, including mobile-friendly interactivity and automated Day 2 operations support, can significantly enhance your Kubernetes management.

We encourage you to explore the Botkube Flux plugin and consider integrating it into your workflows. Don't hesitate to share your feedback and ideas about Botkube. Feel free to reach out to us on [Slack](https://join.botkube.io/) or [Twitter](http://twitter.com/botkube_io).

Thank you for taking the time to learn about Botkube 🙌
