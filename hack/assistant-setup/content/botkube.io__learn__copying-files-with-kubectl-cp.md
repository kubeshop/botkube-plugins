Title: Master File Transfers in Kubernetes with kubectl cp

URL Source: https://botkube.io/learn/copying-files-with-kubectl-cp

Markdown Content:
Kubernetes is a powerful tool for managing and deploying containerized applications. One of the many useful features of Kubernetes is the ability to copy files to and from pods using the kubectl cp command. In this article, we will explore how to use this command to copy files from a pod to your local machine and vice versa.

**The kubectl cp command**
--------------------------

The kubectl cp command allows you to copy files and directories to and from containers in a Kubernetes pod. It works similarly to the cp command in Linux, but with the added functionality of being able to copy files to and from remote containers.

**Copying files from a pod to your local machine**
--------------------------------------------------

To copy a file from a pod to your local machine, you will need to know the name of the pod and the path to the file you want to copy. You can use the kubectl get pods command to list all the pods in your cluster and find the name of the pod you want to copy from.

Once you have the pod name, you can use the following command to copy the file to your local machine:

kubectl cp :

For example, if you want to copy a file named app.log from a pod named my-pod to your current working directory, you would use the following command:

kubectl cp my-pod:/var/log/app.log .

This will copy the file to your current working directory.

**Copying files from your local machine to a pod**
--------------------------------------------------

To copy a file from your local machine to a pod, you will need to know the name of the pod and the path to the directory where you want to copy the file. You can use the kubectl get pods command to list all the pods in your cluster and find the name of the pod you want to copy to.

Once you have the pod name, you can use the following command to copy the file to the pod:

kubectl cp :

For example, if you want to copy a file named config.yaml from your current working directory to a pod named my-pod in the /etc/config directory, you would use the following command:

kubectl cp config.yaml my-pod:/etc/config

This will copy the file to the specified directory in the pod.

**Copying directories**
-----------------------

The kubectl cp command also allows you to copy entire directories to and from pods. To copy a directory, simply add the -r flag to the command. For example:

kubectl cp -r my-pod:/var/log .

This will copy the entire /var/log directory from the pod to your current working directory.

**Conclusion**
--------------

The `kubectl cp` command is a valuable asset in your Kubernetes toolkit. It streamlines file management within your pods, making it easy to transfer logs, configuration files, and other data.

To further enhance your Kubernetes experience, consider integrating a tool like Botkube. Not only does Botkube provide other valuable features, but it also allows you to execute `kubectl` commands (including `kubectl cp`) directly from your Slack or Microsoft Teams workspace. Learn more about setting up the `kubectl` executor in Botkube's documentation: [https://docs.botkube.io/configuration/executor/kubectl](https://docs.botkube.io/configuration/executor/kubectl)

Additionally, don't forget to take advantage of Botkube's other helpful resources, such as the [kubectl cheat sheet](https://botkube.io/learn/kubectl-cheat-sheet), to optimize your Kubernetes workflow.

‍
