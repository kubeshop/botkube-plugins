Title: OOMKilled Explained: Unraveling Kubernetes Memory Limit Errors

URL Source: https://botkube.io/learn/what-is-oomkilled

Markdown Content:
OOMKilled is an error that occurs when a container or pod in Kubernetes uses more memory than it is allowed. This can happen for a variety of reasons, such as a memory leak, a bug in the application, or a spike in traffic. When a container or pod is terminated due to OOMKilled, it will not be able to recover and will need to be restarted.

**How to View the Error Message**
---------------------------------

If you see an OOMKilled error, you can view the error message by running the following command:

`kubectl logs <pod_name>`

The error message will show you the name of the container that was terminated and the reason for the termination.

### Help with pulling Logs

Botkube can help pull logs directly from Slack or Microsoft Teams to help speed up the troubleshooting process. No need to remember the [kubectl command](https://botkube.io/learn/kubectl-cheat-sheet) to pull logs, simply select logs from the drop down that appears on the error message received from the cluster. Botkube will also allow you to filter through the logs to quickly find error messages.

![Image 1: Pulling k8s logs from Slack](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/664789d8541102b509a99a81_Pulling%20Logs%20from%20Slack.png)

**What is Error Code 137?**
---------------------------

Error Code 137 is the exit code that is returned when a container or pod is terminated due to OOMKilled. You can see the error code by running the following command:

`kubectl get pod <pod_name> -o yaml`

The error code will be listed in the "Status" section of the pod's YAML definition.

**How to Prevent OOMKilled**
----------------------------

There are a few things that can be done to prevent OOMKilled in Kubernetes. One is to set memory limits on containers and pods. This will prevent containers from using more memory than they are allowed. Another is to use a memory profiler to identify and fix memory leaks. Finally, it is important to monitor your Kubernetes cluster for signs of OOMKilled errors. If you see an OOMKilled error, you can take steps to troubleshoot the issue and prevent it from happening again.

### Need to Check Pod Memory?

Botkube's new AI powered assistant allows you to quickly check the memory size of the pods connected to your cluster. You could do it the manual way of running kubectl commands or searching through the yaml file to find limits, but the assistant is much faster. Simply ask the Chat Bot: '**@Botkube** ai what are the memory limits of my pods?'

![Image 2: Asking AI for help with memory limits](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/667dc01d73830c573baaa1cb_asking%20ai%20for%20memory%20limit2.png)

**How to Troubleshoot OOMKilled**
---------------------------------

If you see an OOMKilled error, you can take the following steps to troubleshoot the issue:

1.  Identify the container or pod that was terminated.
2.  Check the memory usage of the container or pod.
3.  Look for any errors in the container or pod logs.
4.  Update the container or pod image.
5.  Increase the memory limit for the container or pod.

**Conclusion**
--------------

OOMKilled is an error that can cause significant disruptions to your Kubernetes applications. By following the tips in this article, you can help to prevent OOMKilled and keep your applications running smoothly.

### **Additional Tips for Troubleshooting OOMKilled**

Here are some additional tips for troubleshooting OOMKilled errors:

*   Use a tool like Prometheus to collect metrics about your Kubernetes cluster.
*   Use a tool like Botkube to see the error messages directly in Slack or Teams and use AI to suggest a troubleshooting action directly from the Slack channel.

### Solving OOMKilled with AI for Kubernetes Pods

Botkube's new AI assistant revolutionizes how you handle OOMKilled errors in your Kubernetes pods. This intelligent feature proactively scans your cluster to pinpoint the exact pod responsible for the error. It then offers actionable solutions: either tailored guidance on optimizing your workload to reduce memory consumption or the ability to use Botkube's integrated Kubectl commands to increase memory allocation directly. With this AI-powered tool, you can swiftly resolve OOMKilled errors, saving valuable time and ensuring your Kubernetes applications run smoothly.
