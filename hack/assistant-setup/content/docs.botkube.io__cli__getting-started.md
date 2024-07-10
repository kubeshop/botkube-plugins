Title: Getting Started | Botkube

URL Source: https://docs.botkube.io/cli/getting-started

Markdown Content:
Getting Started | Botkube
===============
       

[Skip to main content](https://docs.botkube.io/cli/getting-started#__docusaurus_skipToContent_fallback)

**New to Botkube?** Get started fast (and free!) with the [Botkube Cloud Web App](https://app.botkube.io/).

[![Image 1: Botkube Logo](https://docs.botkube.io/images/botkube-black.svg)![Image 2: Botkube Logo](https://docs.botkube.io/images/botkube-white.svg) **Botkube**](https://docs.botkube.io/)[Documentation](https://docs.botkube.io/)[Community](https://docs.botkube.io/community/contribute/)

[1.12](https://docs.botkube.io/)

*   [Unreleased 🚧](https://docs.botkube.io/next/cli/getting-started)
*   [1.12](https://docs.botkube.io/cli/getting-started)
*   [1.11](https://docs.botkube.io/1.11/cli/getting-started)
*   [1.10](https://docs.botkube.io/1.10/cli/getting-started)
*   [1.9](https://docs.botkube.io/1.9/cli/getting-started)
*   [1.8](https://docs.botkube.io/1.8/cli/getting-started)
*   [1.7](https://docs.botkube.io/1.7/cli/getting-started)
*   * * *
    
*   [All versions](https://docs.botkube.io/versions)

[GitHub](https://github.com/kubeshop/botkube)[Slack](https://join.botkube.io/)

Search

*   [Installation](https://docs.botkube.io/)
    
*   [Tutorials and examples](https://docs.botkube.io/examples-and-tutorials/)
    
*   [Features](https://docs.botkube.io/features/event-notifications)
    
*   [Plugins](https://docs.botkube.io/plugins/)
    
*   [CLI](https://docs.botkube.io/cli/getting-started)
    
    *   [Getting Started](https://docs.botkube.io/cli/getting-started)
    *   [Migrating installation to Botkube Cloud](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud)
    *   [Commands](https://docs.botkube.io/cli/commands/botkube)
        
*   [Self-hosted Configuration](https://docs.botkube.io/self-hosted-configuration/)
    
*   [Troubleshooting](https://docs.botkube.io/troubleshooting/common-problems)
    
*   [Architecture](https://docs.botkube.io/architecture/)
    

*   [](https://docs.botkube.io/)
*   CLI
*   Getting Started

Version: 1.12

On this page

Getting Started
===============

Botkube includes a command-line interface (CLI) that you can use to interact with Botkube and Botkube Cloud from your terminal, or from a script.

Installation[​](https://docs.botkube.io/cli/getting-started#installation "Direct link to Installation")
-------------------------------------------------------------------------------------------------------

Select tab depending on the system you use:

*   Mac with Apple Silicon
*   Mac with Intel chip
*   Linux
*   Windows
*   Other

Use [Homebrew](https://brew.sh/) to install the latest Botkube CLI:

```
brew install kubeshop/botkube/botkube
```

* * *

Alternatively, download the Botkube CLI binary and move it to a directory under your `$PATH`:

```
curl -Lo botkube https://github.com/kubeshop/botkube/releases/download/v1.12.0/botkube-darwin-arm64chmod +x botkube && mv botkube /usr/local/bin/botkube
```

Use [Homebrew](https://brew.sh/) to install the latest Botkube CLI:

```
brew install kubeshop/botkube/botkube
```

* * *

Alternatively, download the Botkube CLI binary and move it to a directory under your `$PATH`:

```
curl -Lo botkube https://github.com/kubeshop/botkube/releases/download/v1.12.0/botkube-darwin-amd64chmod +x botkube && mv botkube /usr/local/bin/botkube
```

Download the Botkube CLI binary and move it to a directory under your `$PATH`:

```
curl -Lo botkube https://github.com/kubeshop/botkube/releases/download/v1.12.0/botkube-linux-amd64chmod +x botkube && mv botkube /usr/local/bin/botkube
```

note

You may need to use `sudo` to run the `mv` command as it moves the binary file to the `/usr/local/bin/` directory.

* * *

Alternatively, if you use [Homebrew](https://brew.sh/), you can use it to install the latest Botkube CLI:

```
brew install kubeshop/botkube/botkube
```

Use [curl](https://curl.se/windows/) to download the Botkube CLI binary:

```
curl -Lo botkube https://github.com/kubeshop/botkube/releases/download/v1.12.0/botkube-windows-amd64.exe
```

Move the binary to a directory under your `%PATH%`.

Use [curl](https://curl.se/) to download the Botkube CLI binary:

```
export OS="" # allowed values: darwin, linux, windowsexport ARCH="" # allowed values: amd64, arm64, armv7#export SUFFIX=".exe" # uncomment if OS is 'windows'curl -Lo botkube https://github.com/kubeshop/botkube/releases/download/v1.12.0/botkube-${OS}-${ARCH}${SUFFIX}
```

Move the binary to a directory under your `PATH`.

First use[​](https://docs.botkube.io/cli/getting-started#first-use "Direct link to First use")
----------------------------------------------------------------------------------------------

For the commands that are nested under `cloud` command you first need to authenticate with Botkube cloud by running:

```
botkube login
```

If credentials are valid, the output is:

```
Login Succeeded
```

All available commands you can simply discover by running `botkube --help` or `botkube <command> -h` to see the help output which corresponds to a given command.

Autocompletion[​](https://docs.botkube.io/cli/getting-started#autocompletion "Direct link to Autocompletion")
-------------------------------------------------------------------------------------------------------------

To learn how to enable autocompletion for your shell, run:

```
botkube completion --help
```

> **NOTE:** Be sure to **restart your shell** after installing autocompletion.

When you start typing a `botkube` command, press the `<tab>` character to show a list of available completions. Type `-<tab>` to show available flag completions.

[Edit this page](https://github.com/kubeshop/botkube-docs/edit/main/versioned_docs/version-1.12/cli/getting-started.mdx)

[Previous Debugging](https://docs.botkube.io/plugins/development/debugging)[Next Migrating installation to Botkube Cloud](https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud)

*   [Installation](https://docs.botkube.io/cli/getting-started#installation)
*   [First use](https://docs.botkube.io/cli/getting-started#first-use)
*   [Autocompletion](https://docs.botkube.io/cli/getting-started#autocompletion)

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
