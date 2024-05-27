Title: Contribute | Botkube

URL Source: https://docs.botkube.io/community/contribute/

Markdown Content:
*   [](https://docs.botkube.io/)
*   Contribute

How to Contribute to Botkube
----------------------------

We'd love your help!

Botkube is [MIT Licensed](https://github.com/kubeshop/botkube/blob/main/LICENSE) and accepts contributions via GitHub pull requests. This document outlines conventions on development workflow, commit message formatting, contact points and other resources to make it easier to get your contributions accepted.

We gratefully welcome improvements to [documentation](https://docs.botkube.io/ "Go to documentation site") as well as to code.

Contributing to documentation[​](#contributing-to-documentation "Direct link to Contributing to documentation")
---------------------------------------------------------------------------------------------------------------

Follow the [botkube-docs/CONTRIBUTING.md](https://github.com/kubeshop/botkube-docs/blob/main/CONTRIBUTING.md) file to learn how to contribute to documentation.

Build and run Botkube from source code[​](#build-and-run-botkube-from-source-code "Direct link to Build and run Botkube from source code")
------------------------------------------------------------------------------------------------------------------------------------------

This section describes how to build and run Botkube from source code.

### Prerequisite[​](#prerequisite "Direct link to Prerequisite")

*   [Go](https://go.dev/), at least 1.18
    
*   `make`
    
*   [Docker](https://docs.docker.com/install/)
    
*   Kubernetes cluster, at least 1.21
    
*   Cloned Botkube repository
    
    Use the following command to clone it:
    
        git clone https://github.com/kubeshop/botkube.git
    

### Build and install on Kubernetes[​](#build-and-install-on-kubernetes "Direct link to Build and install on Kubernetes")

1.  Build Botkube and create a new container image tagged as `ghcr.io/kubeshop/botkube:v9.99.9-dev`. Choose one option:
    
    *   **Single target build for your local K8s cluster**
        
        This is ideal for running Botkube on a local cluster, e.g. using [kind](https://kind.sigs.k8s.io/) or [`minikube`](https://minikube.sigs.k8s.io/docs/).
        
        Remember to set the `IMAGE_PLATFORM` env var to your target architecture. By default, the build targets `linux/amd64`.
        
        For example, the command below builds the `linux/arm64` target:
        
            IMAGE_PLATFORM=linux/arm64 make container-image-singledocker tag ghcr.io/kubeshop/botkube:v9.99.9-dev {your_account}/botkube:v9.99.9-devdocker push {your_account}/botkube:v9.99.9-dev
        
        Where `{your_account}` is Docker hub or any other registry provider account to which you can push the image.
        
    *   **Multi-arch target builds for any K8s cluster**
        
        This is ideal for running Botkube on remote clusters.
        
        When tagging your dev image take care to add your target image architecture as a suffix. For example, in the command below we added `-amd64` as our target architecture.
        
        This ensures the image will run correctly on the target K8s cluster.
        
        > **Note** This command takes some time to run as it builds the images for multiple architectures.
        
            make container-imagedocker tag ghcr.io/kubeshop/botkube:v9.99.9-dev-amd64 {your_account}/botkube:v9.99.9-devdocker push {your_account}/botkube:v9.99.9-dev
        
        Where `{your_account}` is Docker hub or any other registry provider account to which you can push the image.
        
2.  Install Botkube with any of communication platform configured, according to [the installation instructions](https://docs.botkube.io/installation/). During the Helm chart installation step, set the following flags:
    
        export IMAGE_REGISTRY="{imageRegistry}" # e.g. docker.ioexport IMAGE_PULL_POLICY="{pullPolicy}" # e.g. Always or IfNotPresent--set image.registry=${IMAGE_REGISTRY} \--set image.repository={your_account}/botkube \--set image.tag=v9.99.9-dev \--set image.pullPolicy=${IMAGE_PULL_POLICY}
    
    Check [values.yaml](https://github.com/kubeshop/botkube/blob/main/helm/botkube/values.yaml) for default options.
    

### Build and run locally[​](#build-and-run-locally "Direct link to Build and run locally")

For faster development, you can also build and run Botkube outside K8s cluster.

1.  Build Botkube local binary:
    
        # Fetch the dependenciesgo mod download# Build the binarygo build -o botkube-agent ./cmd/botkube-agent/
    
2.  Create a local configuration file to override default values. For example, set communication credentials, specify cluster name, and disable analytics:
    
        cat <<EOF > local_config.yamlcommunications:  default-group:    socketSlack:      enabled: true      channels:        default:          name: random      appToken: "xapp-xxxx"      botToken: "xoxb-xxxx"configWatcher:   enabled: falsesettings:  clusterName: "labs"analytics:  # -- If true, sending anonymous analytics is disabled. To learn what date we collect,  # see [Privacy Policy](https://botkube.io/privacy#privacy-policy).  disable: trueEOF
    
    To learn more about configuration, visit [https://docs.botkube.io/configuration/](https://docs.botkube.io/configuration/).
    
3.  Export paths to configuration files. The priority will be given to the last (right-most) file specified.
    
        export BOTKUBE_CONFIG_PATHS="$(pwd)/helm/botkube/values.yaml,$(pwd)/local_config.yaml"
    
4.  Export the path to Kubeconfig:
    
        export BOTKUBE_SETTINGS_KUBECONFIG=/Users/$USER/.kube/config # set custom path if necessary
    
5.  Make sure you are able to access your Kubernetes cluster:
    
        Kubernetes master is running at https://192.168.39.233:8443CoreDNS is running at https://192.168.39.233:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy...
    
6.  Run Botkube agent binary:
    

#### Develop Botkube plugins[​](#develop-botkube-plugins "Direct link to Develop Botkube plugins")

**Prerequisite**

*   Being able to start the Botkube binary locally.
*   [GoReleaser](https://goreleaser.com/install/)

**Steps**

1.  Start fake plugins server to serve binaries from `dist` folder:
    
        go run test/helpers/plugin_server.go
    
    > **Note** If Botkube runs inside the k3d cluster, export the `PLUGIN_SERVER_HOST=http://host.k3d.internal` environment variable.
    
2.  Export Botkube plugins cache directory:
    
        export BOTKUBE_PLUGINS_CACHE__DIR="/tmp/plugins"
    
3.  In other terminal window, run:
    
        # rebuild plugins only for current GOOS and GOARCHmake build-plugins-single &&# remove cached pluginsrm -rf $BOTKUBE_PLUGINS_CACHE__DIR &&# start botkube to download fresh plugins./botkube-agent
    
    > **Note** Each time you make a change to the [source](https://github.com/kubeshop/botkube/blob/main/cmd/source) or [executors](https://github.com/kubeshop/botkube/blob/main/cmd/executor) plugins re-run the above command.
    
    > **Note** To build specific plugin binaries, use `PLUGIN_TARGETS`. For example `PLUGIN_TARGETS="x, kubectl" make build-plugins-single`.
    

Making A Change[​](#making-a-change "Direct link to Making A Change")
---------------------------------------------------------------------

*   Before making any significant changes, please [open an issue](https://github.com/kubeshop/botkube/issues). Discussing your proposed changes ahead of time will make the contribution process smooth for everyone.
    
*   Once we've discussed your changes, and you've got your code ready, make sure that the build steps mentioned above pass. Open your pull request against the [`main`](https://github.com/kubeshop/botkube/tree/main) branch.
    
    To learn how to do it, follow the **Contribute** section in the [Git workflow guide](https://github.com/kubeshop/botkube/blob/main/git-workflow.md).
    
*   To avoid build failures in CI, install [`golangci-lint`](https://golangci-lint.run/welcome/install/) and run:
    
        # From project root directorymake lint-fix
    
    This will run the `golangci-lint` tool to lint the Go code.
    

### Run the e2e tests[​](#run-the-e2e-tests "Direct link to Run the e2e tests")

Here [are the details you need](https://github.com/kubeshop/botkube/blob/main/test/README.md) to set up and run the e2e tests.

### Create a Pull Request[​](#create-a-pull-request "Direct link to Create a Pull Request")

*   Make sure your pull request has [good commit messages](https://chris.beams.io/posts/git-commit/):
    
    *   Separate subject from body with a blank line
    *   Limit the subject line to 50 characters
    *   Capitalize the subject line
    *   Do not end the subject line with a period
    *   Use the imperative mood in the subject line
    *   Wrap the body at 72 characters
    *   Use the body to explain _what_ and _why_ instead of _how_
*   Try to squash unimportant commits and rebase your changes on to the `main` branch, this will make sure we have clean log of changes.
    

Support Channels[​](#support-channels "Direct link to Support Channels")
------------------------------------------------------------------------

Join the Botkube-related discussion on Slack!

Create your Slack account on [Botkube](https://join.botkube.io/) workspace.

To report bug or feature, use [GitHub issues](https://github.com/kubeshop/botkube/issues/new/choose).

[Elasticsearch notifier development](https://docs.botkube.io/community/contribute/elasticsearch-develop)

*   [Contributing to documentation](#contributing-to-documentation)
*   [Build and run Botkube from source code](#build-and-run-botkube-from-source-code)
    *   [Prerequisite](#prerequisite)
    *   [Build and install on Kubernetes](#build-and-install-on-kubernetes)
    *   [Build and run locally](#build-and-run-locally)
*   [Making A Change](#making-a-change)
    *   [Run the e2e tests](#run-the-e2e-tests)
    *   [Create a Pull Request](#create-a-pull-request)
*   [Support Channels](#support-channels)
