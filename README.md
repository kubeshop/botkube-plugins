# Botkube Cloud Plugins

This repository shows Botkube Cloud plugins.

## Requirements

- [Go](https://golang.org/doc/install) >= 1.21
- [GoReleaser](https://goreleaser.com/) >= 1.21
- [`golangci-lint`](https://golangci-lint.run/) >= 1.55

## Development

**Prerequisite**

- You are able to start the [Botkube Agent](https://github.com/kubeshop/botkube/blob/main/CONTRIBUTING.md#build-and-run-locally).

**Steps**

1. Start the local plugins server to serve binaries from [`dist`](dist) folder:

   ```bash
   make serve-local-plugins
   ```

   > **Tip**
   > If Botkube runs inside the k3d cluster, export the `PLUGIN_SERVER_HOST=http://host.k3d.internal` environment variable.

2. Export the Botkube plugins cache directory:

   ```bash
   export BOTKUBE_PLUGINS_CACHE__DIR="/tmp/plugins"
   ```

3. Add a `cloud-plugins` entry for your Agent plugins repository:

   ```yaml
   plugins:
     repositories:
       cloud-plugins:
         url: http://localhost:3010/botkube.yaml
   ```

4. In another terminal window, run:

   ```bash
   # rebuild plugins only for the current GOOS and GOARCH
   make build-plugins-single &&
   # remove cached plugins
   rm -rf $BOTKUBE_PLUGINS_CACHE__DIR &&
   # start Botkube to download fresh plugins
   ./botkube-agent
   ```

   Each time you make a change to the [source](cmd/source) or [executors](cmd/executor) plugins, rerun the above command.

   > **Tip**
   > To build specific plugin binaries, use `PLUGIN_TARGETS`. For example, `PLUGIN_TARGETS="helm,argocd" make build-plugins-single`.

## Release

### Fetch content for OpenAI assistant

The AI plugin uses Botkube content (website, docs, blog posts, etc.). To refresh it, follow the steps:

1. Navigate to the [` Fetch content for OpenAI assistant` GitHub Actions workflow](https://github.com/kubeshop/botkube-cloud-plugins/actions/workflows/ai-assistant-fetch-content.yml).
1. Optionally check the "Purge all content" checkbox.
1. Trigger the pipeline.


The content is shared between dev and prod AI plugin, and is synchronized during the [plugin relese](#release-plugins-for-botkube-agent).

### Release plugins for Botkube Agent

The latest plugins from `main` are released automatically to the bucket `botkube-cloud-plugins-latest`.
To release a new production version of the plugins:

1. Navigate to the [`Trigger release` GitHub Actions workflow](https://github.com/kubeshop/botkube-cloud-plugins/actions/workflows/release.yml).
1. Provide the target version in the `version` input. Do not forget about the `v` prefix.

   For example, if you want to release version `1.2.3`, you should provide `version: v1.2.3`.

1. Trigger the pipeline.

For both latest and production plugins, the OpenAI assistant for AI plugin is reconfigured automatically concurrently to the plugins build.
