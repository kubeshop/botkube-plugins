Title: Basic diagnostics | Botkube

URL Source: https://docs.botkube.io/operation/diagnostics

Markdown Content:
Here you can find the list of basic diagnostic actions that may help you look for bug causes.

In order to run the suggested commands, make sure that you have installed:

*   [`helm`](https://helm.sh/docs/intro/install/)
*   [`kubectl`](https://kubernetes.io/docs/tasks/tools/)

Agent[​](#agent"DirectlinktoAgent")
---------------------------------------

This section describes [Botkube agent](https://docs.botkube.io/architecture/) related diagnostic.

### Agent version[​](#agent-version"DirectlinktoAgentversion")

The easiest way to check the Agent version is to get the Docker image:

kubectl get deploy botkube -n botkube -o=jsonpath="{'Used images\n'}{range .spec.template.spec.containers[*]}{.name}{':\t'}{.image}{'\n'}{end}" You should get an output similar to this: Used imagesbotkube: ghcr.io/kubeshop/botkube:v1.5.0cfg-watcher: ghcr.io/kubeshop/k8s-sidecar:in-cluster-config The `botkube` is the agent image. The container image tag (`v1.5.0`) is the version in which it was deployed on the cluster. ### Agent health[​](#agent-health"DirectlinktoAgenthealth")

To check if the Agent Pods are in the `Running` state, run:

kubectl get pod -n botkube -l app=botkube

All the containers from Pods should be in the `Running` status. Restarts' number higher than one may also indicate problems, e.g. not enough resource, lack of permissions, network timeouts, etc.

### Agent logs[​](#agent-logs"DirectlinktoAgentlogs")

If the Botkube Agent is [healthy](#agent-health), you should be able to track any bug by checking the logs. To check the logs, run:

kubectl logs -n botkube -l app=botkube -c botkube

To get all logs specify `--tail=-1`, otherwise only 10 last lines are displayed.

To check the logs since a given time, use the `--since-time` or `--since` flag, for example:

--since-time=2020-03-30T10:02:08Z

### Agent configuration[​](#agent-configuration"DirectlinktoAgentconfiguration")

Select a tab to use a tool of your choice for getting Botkube configuration:

*   Botkube CLI
*   kubectl

note

The `botkube config get` command is available from the `v1.4.0` version.

Install [Botkube CLI](https://docs.botkube.io/cli/getting-started#installation) and run:

botkube config get > /tmp/bk-config.yaml

### Agent restart[​](#agent-restart"DirectlinktoAgentrestart")

When Pods are unhealthy, or if the operation processing is stuck, you can restart the Pod using this command:

kubectl delete po -n botkube -l app=botkube

### Agent debug logging[​](#agent-debug-logging"DirectlinktoAgentdebuglogging")

In order to change the logging level to `debug`, run:

helm upgrade botkube botkube/botkube -n botkube --set settings.log.level="debug" --reuse-values

If the Botkube agent Pod isn't restarted automatically, [restart it manually](#agent-restart).

### Check configured plugin repositories[​](#check-configured-plugin-repositories"DirectlinktoCheckconfiguredpluginrepositories")

Select a tab to use a tool of your choice for checking plugin repository configuration:

*   yq
*   jq
*   grep

Install [`yq`](https://github.com/mikefarah/yq) and run:

helm get values botkube --all -oyaml | yq '.plugins'