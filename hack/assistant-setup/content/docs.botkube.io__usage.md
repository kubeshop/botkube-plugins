Title: Usage | Botkube

URL Source: https://docs.botkube.io/usage/

Markdown Content:
Botkube allows you to monitor your Kubernetes cluster, as well as debug and troubleshoot your workloads.

See all available commands[​](#see-all-available-commands "Direct link to See all available commands")
------------------------------------------------------------------------------------------------------

Run `@Botkube help` to find information about the supported commands.

Check Botkube status[​](#check-botkube-status "Direct link to Check Botkube status")
------------------------------------------------------------------------------------

Run `@Botkube ping` to the channel where Botkube is added. The Botkube will respond you with the **pong** message from all the configured clusters.

For [multi-cluster configuration](https://docs.botkube.io/usage/executor/#specify-cluster-name), use the `--cluster-name` flag to get response from the cluster mentioned in the flag. Else check the deployment in Kubernetes cluster in the **botkube** namespace.
