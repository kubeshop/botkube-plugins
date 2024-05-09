Title: Debugging | Botkube

URL Source: https://docs.botkube.io/plugin/debugging

Markdown Content:
Embedded logging allows you to access more information about the runtime operations of Botkube plugins.

To change the default log level, export a dedicated environment variable following this pattern `LOG_LEVEL_{pluginType}_{pluginRepo}_{pluginName}`, e.g., `LOG_LEVEL_EXECUTOR_BOTKUBE_KUBECTL`. The possible log level values are:

The plugin standard output is logged only if `debug` level is set.

info

The plugin name is normalized and all characters different from letters, digits, and the underscore (`_`) are replaced with underscore (`_`).

To change the log level for a given plugin directly in the Botkube deployment, specify `extraEnv` in the [values.yaml](https://github.com/kubeshop/botkube/blob/main/helm/botkube/values.yaml) file. For example:

    extraEnv:  - name: LOG_LEVEL_EXECUTOR_BOTKUBE_KUBECTL    value: "debug"
