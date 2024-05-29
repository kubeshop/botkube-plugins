Title: Command Line Magic: Simplify Your Life with Custom Aliases

URL Source: https://botkube.io/blog/command-line-magic-simplify-your-life-with-custom-kubernetes-kubectrl-aliases-on-botkube

Published Time: Mar 09, 2023

Markdown Content:
![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ed63e5b48fed17e54625_SE6Pjp9PW9TaOwePHJXRaxaLQgYdT2HX_5PYASmvIx8.jpeg)

Maria Ashby

Developer Advocate

Botkube

With custom aliases, you can create shortcuts for those long and confusing Kubernetes commands, making it easy to manage your clusters and focus on what really matters - your code.

### Table of Contents

*   [So how do you do it?](#so-how-do-you-do-it-)
*   [Syntax](#syntax)

#### Start Using Botkube AI Assistant Today

Are you tired of typing out long and complex Kubernetes commands, only to realize you've made a typo and have to start over? Do you wish there was a way to make managing your Kubernetes clusters easier and more efficient? Fortunately, [custom kubectl aliases](https://docs.botkube.io/usage/executor/#aliases) on Botkube are here to save the day. With custom aliases, you can create shortcuts for those long and confusing Kubernetes commands, making it easy to manage your clusters and focus on what really matters - **your code**.

Whether you're a seasoned Kubernetes pro or just starting out, custom aliases on Botkube can help you work more efficiently and productively. In this blog, we'll show you why custom aliases with Botkube should be in every developer's toolkit.

### Prerequisites :

*   Access to a Kubernetes cluster
    
*   Botkube [installed](https://docs.botkube.io/) and configured
    

So how do you do it?
--------------------

Alias are a shortcut for a longer command or just a part of it. It can be defined for all commands, including [executor plugins](https://docs.botkube.io/usage/executor/) and built-in Botkube commands. This powerful tool lets you create shortcuts for longer, more complex commands, making them quicker and easier to execute. And the best part? You can create aliases for all commands, including built-in Botkube commands and executor plugins. When you use an alias, it replaces the original command with the underlying command before executing it, saving you precious seconds and reducing the risk of errors.

To save on time, Botkube had implemented the `k` and `kc` aliases for the kubectl command, much like DevOps admins like to do on the command line. But we took it a step further: **now you can configure your own custom aliases for any command in Botkube**. Not only can you alias a single command like kubectl, you can create an alias for the full command, including options and flags.

In the example shown below, kgp is an alias for the full kubectl get pods command.

![Image 2](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6408ec7df6773520fc340602_sLERrE-WA2Iv0EldUqYb-eeTU_dmKcSc9eK3k6ryJguJX8MuZyReFo14bHFCWumC846c87NEyVpyjLs3bJBImbK_aF_0iH4k6JCgoQHl0hrLRWzBdnZ5Y8Hg8AMICY4tRhyP06K9v539W8TuW9mTvvY.png)

Once you have configured aliases, you can use them interchangeably with a full command. For example:

**k** as kubectl,

**kgp** as `kubectl get pods`,

**kgpa** as `kubectl get pods -A`,

**hh** as `helm history`,

**a** as list actions, the built-in Botkube command,

and so on.

Aliases are defined globally for the whole Botkube installation. To see which aliases are available for current conversation, run `@Botkube list aliases`.

Syntax
------

![Image 3](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/642d8fd3d44f31b2f4b28059_Screenshot%202023-04-05%20at%208.11.31%20AM.png)

So there you have it- **the power of custom aliases on Botkube**. By creating personalized shortcuts for your most-used commands, you'll be able to work smarter, not harder, and get things done more efficiently than ever before. Plus, with the ability to create aliases for all commands, including built-in Botkube commands and executor plugins, the possibilities are truly endless. So, if you want to streamline your workflow and become a more efficient developer, custom alias feature on Botkube is definitely worth exploring

As always, we want to hear your feedback and ideas about Botkube. Help us plan the Botkube roadmap and get the features you’d like implemented.

There are plenty of options to contact us:

*   [GitHub issues](https://github.com/kubeshop/botkube/issues)
    
*   [Slack](https://join.botkube.io/)
    
*   or email our Product Leader at blair@kubeshop.io.
