Title: Quick start | Botkube

URL Source: https://docs.botkube.io/plugins/development/quick-start

Markdown Content:
*   [](https://docs.botkube.io/)
*   [Plugins](https://docs.botkube.io/plugins/)
*   [Plugin development](https://docs.botkube.io/plugins/development/)
*   Quick start

Version: 1.12

Botkube provides a quick start repository to start developing Botkube [source](https://docs.botkube.io/architecture/#source) and [executor](https://docs.botkube.io/architecture/#executor) plugins in Go. It has all batteries included; example plugins:

*   The [`echo`](https://github.com/kubeshop/botkube-plugins-template/blob/main/cmd/echo/main.go) executor that sends back the command that was specified,
*   The [`ticker`](https://github.com/kubeshop/botkube-plugins-template/blob/main/cmd/ticker/main.go) source that emits an event each time the configured time duration elapses,
*   The [`forwarder`](https://github.com/kubeshop/botkube-plugins-template/blob/main/cmd/forwarder/main.go) source that echos the message sent as an incoming webhook request.

and two example release jobs:

*   [GitHub releases](https://github.com/kubeshop/botkube-plugins-template/blob/main/.github/workflows/release.yml)
*   [GitHub Pages](https://github.com/kubeshop/botkube-plugins-template/blob/main/.github/workflows/pages-release.yml)

Use template[​](https://docs.botkube.io/plugins/development/quick-start/#use-template "Direct link to Use template")
--------------------------------------------------------------------------------------------------------------------

1.  Navigate to [`botkube-plugins-template`](https://github.com/kubeshop/botkube-plugins-template).
    
2.  Click **"Use this template"**, next **"Create a new repository"**
    
    ![Image 1: Create Repo](https://docs.botkube.io/assets/images/use-tpl-abd758819e831ddf629b4e4f42e9a452.png)
    
    This creates your own plugin repository with a single commit.
    
3.  After a few seconds, the `.github/workflows/setup.yml` job will create a new commit. This job removes Kubeshop related files, such as:
    
    *   `CONTRIBUTING.md`
    *   `CODE_OF_CONDUCT.md`
    *   `LICENSE`
    *   `SECURITY.md`
    *   `.github/CODEOWNERS`
    
    Additionally, it updates links in README.md to point to your own repository instead of Kubeshop's one. In case of job failure, you need to make and commit those changes manually.
    
    This job runs only once, later you can remove it or disable it.
    
4.  Decide which release job you prefer, the GitHub releases, or GitHub pages. Once decided, remove one the above workflow in your new repository:
    
    *   GitHub releases, defined at `.github/workflows/release.yml` in your GitHub repository.
    *   GitHub Pages, defined at `.github/workflows/pages-release.yml` in your GitHub repository.
    
    If you pick the GitHub Pages, a further configuration is needed:
    
    1.  In project settings, enable [publishing GitHub Pages via GitHub Action](https://docs.github.com/en/pages/getting-started-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site#publishing-with-a-custom-github-actions-workflow).
    2.  Adjust the `github-pages` deployment environment. You can either:
        *   remove the deployment environment if you don't need any protection rules
        *   add an environment protection rule so that you can deploy pages on each new tag. If you use tagging like `vX.Y.Z`, you can add the `v*` rule. As an alternative, can select **All branches** from the dropdown.
5.  Add LICENSE file and update the README.md file to describe the plugins that you created.
    

[Previous Plugin development](https://docs.botkube.io/plugins/development/)[Next Custom executor](https://docs.botkube.io/plugins/development/custom-executor)
