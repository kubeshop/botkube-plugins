Title: Repository | Botkube

URL Source: https://docs.botkube.io/plugins/development/repo

Markdown Content:
A plugin repository is a place where you store your plugin binaries. This repository must be publicly available and supports downloading assets via HTTP(s). This document explains how to create Botkube plugin repositories by providing examples based on GitHub functionality. However, any static file server can be used, for instance: `s3`, `gcs`, etc.

This document describes how to set up such repository. If you use or plan to use GitHub you can adapt the [template repository](https://docs.botkube.io/plugins/development/quick-start) that has batteries included to start developing and hosting Botkube plugins right away.

Index file[​](https://docs.botkube.io/plugins/development/repo/#index-file "Direct link to Index file")
-------------------------------------------------------------------------------------------------------

Your plugin repository must contain at least one index file and one plugin binary. Depending on your needs and preferences, you can create one or more index files to categorize your plugins. You can host both the executor and source plugins in a single repository. You can also include them in the same index file.

In the index file, provide an entry for every plugin from your plugin repository. The index file must have the following syntax:

```
entries:  - name: { plugin_name }    type: { plugin_type } # executor or source    description: { plugin_description }    version: { plugin_version }    urls:      - url: { url_to_plugin_binary }        platform:          os: { plugin_operating_system } # darwin or linux          architecture: { plugin_architecture } # amd64 or arm64        dependencies: # optional dependencies          { dependency_name }:            url: { url_to_dependency_binary }
```

It is not required to host a plugin or dependency binary on the same server as the index file.

### Generate index file[​](https://docs.botkube.io/plugins/development/repo/#generate-index-file "Direct link to Generate index file")

You can create the index file by yourself our use our tool to generate it automatically based on the directory with plugins binaries. The binaries must be named according to the following pattern:

*   For executors, `executor_{plugin_name}_{os}_{arch}`; for example, `executor_kubectl_darwin_amd64`.
*   For sources, `source_{plugin_name}_{os}_{arch}`; for example, `source_kubernetes_darwin_amd64`.

**Steps**

1.  In your plugin repository, add `tools.go`:
    
    ```
    cat << EOF > tools.go//go:build toolspackage toolsimport (	 _ "github.com/kubeshop/botkube/hack")EOF
    ```
    
2.  Refresh dependencies:
    
3.  Build all your plugins. See [**Build plugin binaries**](https://docs.botkube.io/plugins/development/custom-executor).
    
4.  Generate an index file:
    
    ```
    go run github.com/kubeshop/botkube/hack -binaries-path "./dist" -url-base-path "https://example.com"
    ```
    
    info
    
    Replace the `-url-base-path` flag with the base path of your HTTP server. See [Hosting plugins](https://docs.botkube.io/plugins/development/repo/#host-plugins) for some examples.
    

Host plugins[​](https://docs.botkube.io/plugins/development/repo/#host-plugins "Direct link to Host plugins")
-------------------------------------------------------------------------------------------------------------

This section describes example ways for serving Botkube plugins.

### GitHub releases[​](https://docs.botkube.io/plugins/development/repo/#github-releases "Direct link to GitHub releases")

A GitHub release allows you to upload additional assets that are later accessible with a predictable URL. When you generate the index file, specify the `-url-base-path` flag as `https://github.com/{owner}/{repo}/releases/download/{release_tag}`, for example, `https://github.com/kubeshop/botkube/releases/download/v1.0.0`.

Once the plugin binaries are built and the index file is generated, you can create a GitHub release using [GitHub CLI](https://cli.github.com/). For example:

```
gh release create v1.0.0 \ ./dist/source_* \ ./dist/executor_* \ ./plugins-index.yaml
```

#### Automation[​](https://docs.botkube.io/plugins/development/repo/#automation "Direct link to Automation")

You can use [GitHub Actions](https://docs.github.com/en/actions) to publish Botkube plugins automatically each time a new tag is pushed. See the [`release` workflow](https://github.com/kubeshop/botkube-plugins-template/blob/main/.github/workflows/release.yml) on the `botkube-plugins-template` repository for the out-of-the-box solution, which you can use and modify if needed.

### GitHub pages[​](https://docs.botkube.io/plugins/development/repo/#github-pages "Direct link to GitHub pages")

GitHub allows you to serve static pages via GitHub Pages. When you generate the index file, specify the `-url-base-path` flag as `https://{user}.github.io/{repository}`, for example, `https://kubeshop.github.io/botkube-plugins`.

**Initial setup**

1.  Navigate to the Git repository with your plugins.
    
2.  Create the `gh-pages` branch:
    
    ```
    git switch --orphan gh-pagesgit commit --allow-empty -m "Initialization commit"git push -u origin gh-pages
    ```
    
3.  Follow [this](https://docs.github.com/en/pages/getting-started-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site#publishing-from-a-branch) guide to make sure your `gh-pages` branch is set as the source for GitHub Pages.
    

**Publishing steps**

1.  Clone `gh-pages` into `/tmp/botkube-plugins`:
    
    ```
    git clone -b gh-pages "https://github.com/{owner}/{repo}.git" /tmp/botkube-plugins
    ```
    
2.  Move built binaries and generated index file:
    
    ```
    mv dist/executor_* /tmp/botkube-plugins/mv dist/source_* /tmp/botkube-plugins/mv plugins-index.yaml /tmp/botkube-plugins
    ```
    
3.  Commit and push copied files:
    
    ```
    cd /tmp/botkube-pluginsgit add -Agit commit -m "Release Botkube plugins"git push
    ```
    
4.  Remove cloned `gh-pages`:
    
    ```
    cd -rm -rf /tmp/botkube-charts
    ```
    

In such setup, you can use your default branch to store your plugins code, and the `gh-pages` branch as a plugin repository.

#### Automation[​](https://docs.botkube.io/plugins/development/repo/#automation-1 "Direct link to Automation")

You can use [GitHub Actions](https://docs.github.com/en/actions) to publish Botkube plugins automatically each time a new tag is pushed. See the [`pages-release` workflow](https://github.com/kubeshop/botkube-plugins-template/blob/main/.github/workflows/pages-release.yml) on the `botkube-plugins-template` repository for the out-of-the-box solution, which you can use and modify if needed.

### Use hosted plugins[​](https://docs.botkube.io/plugins/development/repo/#use-hosted-plugins "Direct link to Use hosted plugins")

To use the plugins that you published, add your repository under `plugins` in the [values.yaml](https://github.com/kubeshop/botkube/blob/main/helm/botkube/values.yaml) file for a given Botkube deployment. For example:

```
plugins:  repositories:    repo-name:      url: https://example.com/plugins-index.yaml
```

Once the plugin repository is added, you can refer to it in the `executor` or `sources` section.

```
executors:  "plugins":    repo-name/executor-name@v1.0.0: # Plugin name syntax: {repo}/{plugin}[@{version}]. If version is not provided, the latest version from repository is used.      enabled: true      config: {} # Plugin's specific configuration.sources:  "plugins":    repo-name/source-name@v1.0.0: # Plugin name syntax: {repo}/{plugin}[@{version}]. If version is not provided, the latest version from repository is used.      enabled: true      config: {} # Plugin's specific configuration.
```
