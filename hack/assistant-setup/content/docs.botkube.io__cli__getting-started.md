Title: Getting Started | Botkube

URL Source: https://docs.botkube.io/cli/getting-started

Markdown Content:
Version: 1.10

Botkube includes a command-line interface (CLI) that you can use to interact with Botkube and Botkube Cloud from your terminal, or from a script.

Installation[​](#installation "Direct link to Installation")
------------------------------------------------------------

Select tab depending on the system you use:

*   Mac with Apple Silicon
*   Mac with Intel chip
*   Linux
*   Windows
*   Other

Use [Homebrew](https://brew.sh/) to install the latest Botkube CLI:

    brew install kubeshop/botkube/botkube

* * *

Alternatively, download the Botkube CLI binary and move it to a directory under your `$PATH`:

    curl -Lo botkube https://github.com/kubeshop/botkube/releases/download/v1.10.0/botkube-darwin-arm64chmod +x botkube && mv botkube /usr/local/bin/botkube

First use[​](#first-use "Direct link to First use")
---------------------------------------------------

For the commands that are nested under `cloud` command you first need to authenticate with Botkube cloud by running:

If credentials are valid, the output is:

All available commands you can simply discover by running `botkube --help` or `botkube <command> -h` to see the help output which corresponds to a given command.

Autocompletion[​](#autocompletion "Direct link to Autocompletion")
------------------------------------------------------------------

To learn how to enable autocompletion for your shell, run:

    botkube completion --help

> **NOTE:** Be sure to **restart your shell** after installing autocompletion.

When you start typing a `botkube` command, press the `<tab>` character to show a list of available completions. Type `-<tab>` to show available flag completions.
