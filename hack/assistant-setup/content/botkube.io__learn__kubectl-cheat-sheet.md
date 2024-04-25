Title: Kubectl Cheat Sheet: Essential Commands at Your Fingertips

URL Source: https://botkube.io/learn/kubectl-cheat-sheet

Markdown Content:
**What is Kubectl?**
--------------------

`Kubectl` is a command-line tool used to deploy and manage applications on a Kubernetes cluster. It allows you to interact with the Kubernetes API to create, update, and delete resources such as pods, services, and deployments. kubectl can be used to manage both local and remote Kubernetes clusters and is an essential tool for anyone working with Kubernetes.

**How to Install Kubectl into Kubernetes**
------------------------------------------

### **Manual Installation Method**

If you want to use kubectl, the command-line tool for Kubernetes, you need to install it on your machine. Depending on your operating system, there are different ways to do that. Some operating systems, such as Windows, Linux, and Mac, have local specific options that let you download and install kubectl directly. Other solutions, such as Minikube, allow you to install the tool set into your Kubernetes cluster without needing a local installation.

Either way, installation is quick and easy, and you can find detailed steps for each option on the \[official Kubernetes documentation\]([https://kubernetes.io/docs/tasks/tools/install-kubectl/](https://kubernetes.io/docs/tasks/tools/install-kubectl/)).

### **Botkube’s Setup Wizard for Kubectl Easy Installation**

This Kubectl cheatsheet would not be complete without showing the new quickest way to install Kubectl commands on your cluster. With Botkube’s setup wizard, simply select that you want to add Kubectl into your cluster on the Step 3 plugin page. Feel free to add any other helpful K8s plugins like Helm. Once the setup wizard is complete, users can access Kubectl commands, such as kubectl restart pod, from their K8s cluster.

![Image 1: Easily install kubectl into kubernetes clusters with one click on the Botkube Cluster setup wizard](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64d154ccef984e82336dae1a_UDXteleWUlLJ1h495wr-eU7OqNyx3C-_aON-kSFgRVCK_35_iIzuouiTIHDYyo8ERPM0wCxseEROlkkkkVZVDJNmSJnm1JhA53HDTMGkkUGeDLEl5jKVpVaNciIhllLYqpsYfuza79QwLhH0cp1UE0Q.png)

**Tip #1 - How to check kubectl version?**
------------------------------------------

To check the version of `kubectl`, run the following command:

`kubectl version`

This will display the version information for both the client and server components of `kubectl`.

**Tip #2 - What are kubectl logs?**
-----------------------------------

`kubectl logs` is a command used to print the logs of a container in a pod. This is useful for debugging purposes and for getting information about what a container is doing. By default, it will print the logs from the first container in the specified pod. You can also specify a container name to print logs from a specific container within the pod.

Logs is one area of the Kubernetes DevOps stack to get correct and make easily available to everyone on the team. Botkube has learned from our own [internal need for automed Kubernetes log pulling](https://botkube.io/blog/optimizing-collaboration-and-notifications-with-the-botkube-argocd-plugin#how-botkube-uses-argocd-for-gitops-management) how important it is that everyone has access and knows how to pull the logs they need. Botkube cloud enables this developer self service platform directly in a group chat channel.

Botkube helps Automate K8s log management in three ways:

1.  **Alerts** \- Notification that a pod is having an error directly to preferred chat platform; Slack, Discord, Teams, or Mattermost.
2.  **One-Click Log Management** \- Drop down of Kubectl log commands to run without switching back to the terminal.
3.  **Multipurpose Log Commands** - Receive log trail in the shared channel for everyone on the team to see and learn from the troubleshooting process.

![Image 2: Running Kubectl Logs command from Slack to describe a pod error](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/6540047570b8a1cc0cd1d65c_SpSgE1gwyiRtteaKC-Spqsx8es_eBln7hmaQbAGhoyAG3WHx0lgEDOHghYjdqgXLo7X78d_is5SFVUUwpx4Tg6qFsQd3DP6XUB9Nqcobt-EaYs9TVblIS92-Jg9ZOUKcZaLAlt7PSUJYgZGaP3vyKRM.png)

If this automated log pulling process is something that your K8s development team may need, [give it a try](https://app.botkube.io/) free of charge today!

**Tip # 3 - Deleting with kubectl delete pod**
----------------------------------------------

`kubectl delete pod` is a command used to delete a pod. When you delete a pod, Kubernetes will create a new one to replace it. This is useful for updating pods with new configurations or images. If you want to delete the pod and prevent Kubernetes from creating a new one, you can use the --force flag.

**Tip #4 - Running kubectl -k to Deploy Kustomization File**
------------------------------------------------------------

`kubectl -k` is a command used to deploy Kubernetes resources using a kustomization file. A kustomization file is a YAML file that defines a set of Kubernetes resources and how they should be deployed. It allows you to customize your deployments without modifying the original YAML files.

**Tip #5 - Copying Files with kubectl cp**
------------------------------------------

`kubectl cp` is a command used to copy files and directories between a container and the user's local file system. This is useful for transferring files to and from a container for debugging purposes or to extract data from a container.

To copy a file from a container to the local file system, use the following command:

`kubectl cp <namespace>/<pod>:<container-path> <local-path>`

To copy a file from the local file system to a container, use the following command:

`kubectl cp <local-path> <namespace>/<pod>:<container-path>`

**Tip #6 - Run Container Commands with kubectl exec**
-----------------------------------------------------

`kubectl exec` is a command used to run a command inside a container in a pod. This is useful for debugging purposes or for running commands that are not included in the container image. By default, it will run the command in the first container of the specified pod. You can also specify a container name to run the command in a specific container within the pod.

**Tip #7 - Port Forwarding with kubectl Port Forward**
------------------------------------------------------

`kubectl port` forward is a command used to forward a local port to a port on a pod. This is useful for accessing services running inside a pod directly from your local machine. For example, if a pod is running a web server on `port 80`, you can use `kubectl port forward` to forward `port 80` on the pod to `port 8080` on your local machine. Then, you can access the web server by navigating to `http://localhost:8080` in your web browser.

To forward a port, use the following command:

`kubectl port-forward <pod-name> <local-port>:<pod-port>`

`<pod-name>` is the name of the pod you want to forward the port from. `<local-port>` is the port on your machine that you want to forward to. `<pod-port>` is the port on the pod that you want to forward from.

**Tip #8 - Submitting a File using the kubectl Apply Command**
--------------------------------------------------------------

`kubectl apply` is a command used to apply a configuration file to a Kubernetes cluster. It is the recommended way to manage Kubernetes resources and can be used to create, update, or delete resources. When you apply a configuration file, Kubernetes will compare the desired state of the resources in the file to the current state of the resources in the cluster and make any necessary changes to bring the resources to the desired state.

\*`kubectl apply -f` typically requires a YAML or JSON filename afterward to instruct what resources need to be built.

When you run `kubectl apply -f <filename>`, Kubernetes will read the specified configuration file and perform the following actions:

1.  If the resource(s) defined in the file do not already exist in the cluster, Kubernetes will create them.
2.  If the resource(s) already exist in the cluster, Kubernetes will update them to match the desired state specified in the file. This is where the "apply" command differs from "kubectl create" – it can both create and update resources.
3.  If there are resource definitions in the file that are no longer needed, Kubernetes will delete those resources from the cluster.

**Tip # 9 - What does _kubectl create Namespace_ do?**
------------------------------------------------------

`kubectl create namespace` is a command used to create a new namespace in a Kubernetes cluster. Namespaces are a way to divide a Kubernetes cluster into multiple virtual clusters. They provide a scope for names, so resources with the same name can coexist in different namespaces. Namespaces can be used to separate different environments, such as production and development, or different teams working on the same cluster. To create a new namespace, use the following command:

`kubectl create namespace <namespace-name>`

Replace `<namespace-name>` with the desired name of the new namespace.

**Tip #10 - Troubleshooting with kubectl Debug Command**
--------------------------------------------------------

`kubectl debug` is a command used to enter a debug container that is co-located with a target container in the same pod. This is useful for debugging purposes, as it allows you to examine the environment of a container and debug issues that may be occurring. When you enter the debug container, you will have access to the same file system and network namespace as the target container, allowing you to run diagnostic commands and examine logs.

To enter the debug container, use the following command:

`kubectl debug <pod-name> -c <container-name>`

`<pod-name>` is the name of the pod that contains the target container. `<container-name>` is the name of the target container that you want to debug.

**How Botkube makes Kubectl Easier**
------------------------------------

Botkube simplifies the use of `kubectl` in many areas, but three stand out as particularly beneficial. Let us break down how Botkube helps make using kubectl commands easy.

1.  **Run Kubectl Commands from Slack**

Firstly, Botkube enables users to run `kubectl` commands directly from Slack or Teams, eliminating the need to open a separate command line interface. This feature enables developers, site reliability managers, and other employees in the group chat to run necessary `kubectl` commands with ease.

![Image 3: Ran Kubectl get pods command from slack](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64d154cb1fc2678bcaa56f98_fbbDuo-oKj7sih4vmPrvRSy0yvKRbUtmy9fXEqKJpPOGxymF-frlxtR4oa90ANbmcy6g33uaxLtBYUfk9GGlazqWAZbQXPlo8eAOTWhhTAtBNxYWb51mRwQZRKtzQ1SBorlKAlSbpoOU_W9_EeCSRlc.png)

2.  **Get Kubectl help (Even Chatgpt Kubectl Help)**

As seen in the image below, Botkube even gives an option to get kubectl help directly from Slack. Our AI assistant will bring up helpful Kubectl commands and describe a little what they do. This ability combined with the new [Kubernetes Doctor Plugin](https://botkube.io/blog/use-chatgpt-to-troubleshoot-kubernetes-errors-with-botkubes-doctor), that allows users to further query ChatGPT about the correct Kubectl command to run, should make anyone a Kubectl expert quick!

![Image 4: Kubectl commands running in a slack channel](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64d154cb0b2e6541c7bf6681_tJUUtb_nGsLwC-FoHmfUBedMJRKPEdswkO5vIxsKpimjVmU-FGIIoL3vVw29Qf2WXtC6524Js6Id8RWfMkg_rkdUcmCjbQnLgfoZIlJcOroO24iy6sCWlCyYKl3kVhcDP9LvDK6Pm_ZasxgaTU2K5K0.png)

3.  **Set Kubectl Command Aliases**

Finally, Botkube introduces the ability to create command aliases, allowing users to assign quick phrases or letters to commonly used kubectl commands. This feature eliminates the need to remember long and complex four-word commands, streamlining the Kubectl process for all users. These features make Botkube an essential tool for any team looking to simplify their Kubectl workflow and increase productivity. Read more about [Kubectl aliases on our blog](https://botkube.io/blog/command-line-magic-simplify-your-life-with-custom-kubernetes-kubectrl-aliases-on-botkube).

![Image 5: creating a kubectl alias with Botkube's web interface](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/64d154cb0dca04c16dae44db_fjOpVIQetAJ-b8hhrWV8fqy3H63TJPAW4zdIkRjc5uh0mlK5hvgU_YUAGOq7OJQXjtxzQTDURFH5tx9-JLL2NFvKNMrEQwbOH2oeHjZsrzaPvzD1iY5cJe8L4dF0tIqnpmdW86WyS2KNhrpzV-ouuJw.png)

‍
