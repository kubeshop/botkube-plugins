Title: RBAC | Botkube

URL Source: https://docs.botkube.io/configuration/rbac

Markdown Content:
Botkube allows plugins to access Kubernetes API by defining [RBAC](https://en.wikipedia.org/wiki/Role-based_access_control) rules as part of plugin `context.rbac` configuration. Kubeconfig generation needs to be requested by defining `context.rbac` property.

Based on this configuration Botkube generates a temporary [kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/) with user and/or group impersonation. This kubeconfig is available to plugins in the `Execute` and `Stream` contexts.

Architecture[​](#architecture"DirectlinktoArchitecture")
------------------------------------------------------------

Botkube uses its own cluster credentials to generate a temporary kubeconfig, and the kubeconfig uses [user/group impersonation](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#user-impersonation).

For source plugins, the kubeconfig is generated once - during plugin startup. For executor plugins, the kubeconfig is generated every time a command is sent to the plugin, which allows greater flexibility, such as including the name of the channel the command was sent from.

![Image 1: diagram](https://docs.botkube.io/assets/images/botkube-read-only-717ed01cf9fa5e6621f2a09c7b29a32d.svg)

Configuration[​](#configuration"DirectlinktoConfiguration")
---------------------------------------------------------------

For each executor and source plugin, you can define a `context.rbac` configuration. This config is used to generate a dedicated kubeconfig.

executors:  "kubectl":    botkube/kubectl@v1:      enabled: true      # ...      context:        rbac:          user:            type: Static # Static or ChannelName            static: # applicable only for "Static" user mapping type              value: botkube-internal-static-user            prefix: "" # optional prefix for user name; useful especially for channel name mapping          group:            type: Static # Static or ChannelName            static: # applicable only for "Static" group mapping type              values:                - "my-group1"                - "my-group2"            prefix: "" # optional prefix for all group names; useful especially for channel name mapping

### Mapping types[​](#mapping-types"DirectlinktoMappingtypes")

For both user and group, the following mapping types are supported:

*   `Static`

For user, it uses a single static value. For group, it uses a list of static values. The value is prepended with an optional prefix.

*   `ChannelName`

Channel name is used as subject for user or group impersonation. The channel name is prepended with an optional prefix. This mapping is only available for executor plugins.


### Default configuration[​](#default-configuration"DirectlinktoDefaultconfiguration")

When a given plugin have `context.rbac` property undefined, Botkube doesn't generate a kubeconfig for this plugin. To request kubeconfig generation, define `context.rbac` property with empty value:

executors:  "kubectl":    botkube/kubectl@v1:      enabled: true      # ...      context:        rbac: {} # enable kubeconfig generation

However, such configuration will generate a kubeconfig with empty impersonation config, which effectively means an anonymous access to the Kubernetes API.

During Botkube installation, Botkube generates Kubernetes ClusterRole and ClusterRoleBinding resources with read-only access for the default group `botkube-plugins-default`. This group is used by default across the `values.yaml` for all default plugins.

rbac:  # ...  groups:    "botkube-plugins-default":      create: true      rules:        - apiGroups: ["*"]          resources: ["*"]          verbs: ["get", "watch", "list"]

See the [`values.yaml`](https://github.com/kubeshop/botkube/blob/v1.10.0/helm/botkube/values.yaml) for more details.

#### Defaults for user mapping when group mapping is used[​](#defaults-for-user-mapping-when-group-mapping-is-used"DirectlinktoDefaultsforusermappingwhengroupmappingisused")

Kubernetes requires user for group impersonation. That's why when a group mapping is user without `context.rbac.user` mapping defined, Botkube uses `botkube-internal-static-user` user name for impersonation. For example, when the following configuration is used:

executors:  "kubectl":    botkube/kubectl@v1:      enabled: true      # ...      context:        rbac:          # no user mapping defined          group:            type: Static            static:              value: botkube-plugins-default

It is equivalent to:

executors:  "kubectl":    botkube/kubectl@v1:      enabled: true      # ...      context:        rbac:          user:            type: Static            static:              value: botkube-internal-static-user          group:            type: Static            static:              value: botkube-plugins-default

#### Defaults for Botkube Cloud[​](#defaults-for-botkube-cloud"DirectlinktoDefaultsforBotkubeCloud")

When configuring plugin on Botkube Cloud, the "Default" permissions mean that the `botkube-plugins-default` group will be used, which have read-only access to Kubernetes API and is configured during Botkube installation. See the [Default configuration](#default-configuration) section.

![Image 2: Cloud RBAC defaults](https://docs.botkube.io/assets/images/cloud-rbac-default-9ffb455fa5b9f7e43b81da8a4b91f868.png)

Examples[​](#examples"DirectlinktoExamples")
------------------------------------------------

This paragraph contains examples of RBAC configuration for different use cases.

tip

You can use `rbac.groups` or `extraObjects` overrides during Botkube installation to create custom RBAC resources. See the [`values.yaml`](https://github.com/kubeshop/botkube/blob/v1.10.0/helm/botkube/values.yaml) for more details.

### Kubectl executor with read-only Pod access based on static group mapping[​](#kubectl-executor-with-read-only-pod-access-based-on-static-group-mapping"DirectlinktoKubectlexecutorwithread-onlyPodaccessbasedonstaticgroupmapping")

In this example an executor plugin is defined with static RBAC that maps to group `read-pods`.

1.  Consider the following Botkube config:

# ...executors:  "kubectl-read-only":    botkube/kubectl@v1:      enabled: true      # ...      context:        rbac:          group:            type: Static            static:              values: [read-pods]


Let's assume this plugin is bound to at least one channel.

1.  Consider the following Kubernetes RBAC configuration:

apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRolemetadata:  name: kubectl-read-podsrules:  - apiGroups: [""]    resources: ["pods"]    verbs: ["get", "watch", "list"]---apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRoleBindingmetadata:  name: kubectl-read-podsroleRef:  apiGroup: rbac.authorization.k8s.io  kind: ClusterRole  name: kubectl-read-podssubjects:  - kind: Group    name: read-pods # <-- this is the group name used in Botkube config    apiGroup: rbac.authorization.k8s.io


In a result, when this executor plugin is invoked, Botkube generates a kubeconfig impersonating group `read-pods` and passes it to the plugin. The plugin then can authenticate with the API server with identity of group `read-pods`. In that way, the plugin can use read-only operations on Pods.

### Kubernetes source plugin with read-only access based on static user mapping[​](#kubernetes-source-plugin-with-read-only-access-based-on-static-user-mapping"DirectlinktoKubernetessourcepluginwithread-onlyaccessbasedonstaticusermapping")

In this example a single source plugin is defined with static RBAC that maps to user `kubernetes-read-only`.

1.  Consider the following Botkube config:

sources:  "kubernetes":    botkube/kubernetes@v1:      enabled: true      # ...      context:        rbac:          user:            type: Static            static:              value: kubernetes-read-only

2.  Consider the following Kubernetes RBAC configuration:

apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRolemetadata:  name: readerrules:  - apiGroups: ["*"]    resources: ["*"]    verbs: ["get", "watch", "list"]---apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRoleBindingmetadata:  name: readerroleRef:  apiGroup: rbac.authorization.k8s.io  kind: ClusterRole  name: readersubjects:  - kind: User    name: kubernetes-read-only # <-- this is the username used in Botkube config    apiGroup: rbac.authorization.k8s.io


In a result, the source plugin can access all Kubernetes resources with read-only permissions.

### Kubectl executor plugin with different permissions based on channel name mapping[​](#kubectl-executor-plugin-with-different-permissions-based-on-channel-name-mapping"DirectlinktoKubectlexecutorpluginwithdifferentpermissionsbasedonchannelnamemapping")

In this example **kubectl** executor plugin is configured with channel name mapping and bound to two channels, `ch-1` and `ch-2`. In Kubernetes RBAC resources, group `ch-1` is given write access, while group `ch-2` is given only read access.

1.  Consider the following Botkube config:

executors:  "kubectl":    botkube/kubectl@v1:      # ...      enabled: true      context:        rbac:          group:            type: ChannelNamecommunications:  "default-group":    socketSlack:      enabled: true      # ...      channels:        "ch-1":          name: ch-1          bindings:            executors:              - kubectl        "ch-2":          name: ch-2          bindings:            executors:              - kubectl# ...

2.  Consider the following Kubernetes RBAC configuration:

apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRolemetadata:  name: editorrules:  - apiGroups: ["*"]    resources: ["*"]    verbs: ["get", "watch", "list", "update", "create", "delete"]---apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRoleBindingmetadata:  name: editorroleRef:  apiGroup: rbac.authorization.k8s.io  kind: ClusterRole  name: editorsubjects:  - kind: Group    name: ch-1 # <-- channel name used in Botkube config    apiGroup: rbac.authorization.k8s.io---apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRolemetadata:  name: read-onlyrules:  - apiGroups: ["*"]    resources: ["*"]    verbs: ["get", "watch", "list"]---apiVersion: rbac.authorization.k8s.io/v1kind: ClusterRoleBindingmetadata:  name: read-onlyroleRef:  apiGroup: rbac.authorization.k8s.io  kind: ClusterRole  name: read-onlysubjects:  - kind: Group    name: ch-2 # <-- channel name used in Botkube config    apiGroup: rbac.authorization.k8s.io


In a result, users in channel `ch-1` can execute all kubectl commands, while users in channel `ch-2` can only execute read-only commands.

Limitations[​](#limitations"DirectlinktoLimitations")
---------------------------------------------------------

This paragraph contains limitations of the current implementation.

### Shared file system[​](#shared-file-system"DirectlinktoSharedfilesystem")

Botkube runs plugin processes in the same container within the same Pod. Therefore, all plugins share the same file system.

If you're a plugin developer and decide to write kubeconfig to the file system, be aware that it can be accessible by all plugins in the container.

### RBAC configuration merging[​](#rbac-configuration-merging"DirectlinktoRBACconfigurationmerging")

The same executor plugins with different RBAC configuration cannot be bound to the same channel. This is validated during Botkube startup and will result in an error.

For example, the following configuration is invalid:

communications:  "default-group":    socketSlack:      enabled: true      # ...      channels:        "ch-1":          name: ch-1          bindings:            executors:              - kubectl              - kubectl-read-onlyexecutors:  "kubectl":    botkube/kubectl@v1:      enabled: true      # ...      context:        rbac: # Different RBAC configuration          group:            type: ChannelName  "kubectl-read-only":    botkube/kubectl@v1:      enabled: true      # ...      context:        rbac: # Different RBAC configuration          user:            type: Static            static:              value: kubectl-read-only

### Supported RBAC mappings[​](#supported-rbac-mappings"DirectlinktoSupportedRBACmappings")

While Executor plugins support multiple mapping types, there are the following limitations:

*   Source plugins support only the `Static` mapping.
*   Automated [actions](https://docs.botkube.io/usage/automated-actions) support only the `Static` mapping.

Troubleshooting[​](#troubleshooting"DirectlinktoTroubleshooting")
---------------------------------------------------------------------

In most cases troubleshooting Botkube RBAC issues means [troubleshooting Kubernetes RBAC](https://kubernetes.io/docs/reference/access-authn-authz/authorization/#checking-api-access), where `kubectl auth` command can help.

If you see the following error:

Error: create: failed to create: secrets is forbidden: User "botkube-internal-static-user" cannot create resource "secrets" in API group "" in the namespace "default"

that means the RBAC rules configured for a given plugin are insufficient in a given context.

Firstly, ensure what user/group is used for impersonation. To do that, check your configuration against the mapping description from the [Configuration](#configuration) section.

### Checking available actions for a given user/group[​](#checking-available-actions-for-a-given-usergroup"DirectlinktoCheckingavailableactionsforagivenuser/group")

After obtaining proper user and group, use the following command to list all available actions for a given user and/or group:

kubectl auth can-i --as {user} --as-group {group} --list

For example, to list all available actions for user `botkube-internal-static-user` and group `private-channel` use:

kubectl auth can-i --as botkube-internal-static-user --as-group private-channel --list

### Checking if a given user/group can perform a given action[​](#checking-if-a-given-usergroup-can-perform-a-given-action"DirectlinktoCheckingifagivenuser/groupcanperformagivenaction")

To verify if a given user and/or group can perform a given action, use:

kubectl auth can-i get pod -n botkube --as {user} --as-group {group}

For example, to verify if user `botkube-internal-static-user` and group `private-channel` can get Secret in namespace `botkube` use:

kubectl auth can-i get secret -n botkube --as botkube-internal-static-user --as-group private-channel

Plugin development[​](#plugin-development"DirectlinktoPlugindevelopment")
------------------------------------------------------------------------------

If you are a plugin developer and want to learn how to use generated kubeconfig in the plugin codebase, refer to [Using kubeconfig](https://docs.botkube.io/plugin/using-kubeconfig) document.
