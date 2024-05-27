Title: Chat Operations (ChatOps) Explained by Botkube

URL Source: https://botkube.io/learn/chatops

Markdown Content:
ChatOps is a term that describes the practice of using chat tools and bots to automate tasks and workflows in software development and operations. ChatOps enables teams to collaborate more effectively, share knowledge, and increase transparency and accountability. It can also help teams adopt a DevOps culture of continuous improvement and feedback.

One example of a chat tool that supports ChatOps is Botkube, which integrates with Kubernetes and allows users to monitor, debug, and manage their clusters from chat platforms like Slack or Mattermost. Botkube's Product Lead, Blair Rampling, wrote an [article](https://thenewstack.io/chatops-where-automation-collaboration-and-devops-culture-meet/) on TheNewStack that explains more about ChatOps and how Botkube can help teams leverage it.

What Activities in K8s can be Automated with ChatOps?
-----------------------------------------------------

One domain that has a lot of potential for ChatOps automation is Kubernetes. Kubernetes is an open-source system for automating deployment, scaling, and management of containerized applications. Kubernetes allows users to run applications across multiple servers without worrying about the underlying infrastructure. However, Kubernetes also involves a lot of complex tasks that require a lot of commands and configurations.

Some examples of activities related to Kubernetes that can be automated are:

### Cluster Creation

Instead of manually setting up a cluster of servers to run Kubernetes applications, automation tools can create clusters automatically using cloud services or local machines. This can save time and resources and ensure consistency and security of the cluster environment.

### Application Deployment

Instead of manually deploying applications to a Kubernetes cluster using command-line interface (CLI) commands or configuration files (YAML), automation tools can deploy applications automatically using graphical user interface (GUI) tools or declarative methods. This can simplify the deployment process and reduce human errors and failures.

### Application Scaling

Instead of manually scaling applications up or down based on demand or performance metrics using CLI commands or configuration files (YAML), automation tools can scale applications automatically using horizontal pod autoscaler (HPA) or vertical pod autoscaler (VPA). This can optimize the resource utilization and availability of the applications.

### Appplication Monitoring

Instead of manually monitoring applications for health status, performance metrics, logs, events, etc., using CLI commands or configuration files (YAML), automation tools can monitor applications automatically using dashboard tools or alerting systems. This can improve the visibility and troubleshooting of the applications.

One way to make Kubernetes activities easier to automate is to use tools that allow users to create aliases or short phrases to reference long CLI commands. Maria put together a [tutorial](https://botkube.io/blog/command-line-magic-simplify-your-life-with-custom-kubernetes-kubectrl-aliases-on-botkube) covering how set up aliases using a ChatOps automation tool like Botkube.

Rising Trend in combining ChatOps with AIOps
--------------------------------------------

After moving Kubernetes operation alerts into a chat channel, it is natural for the team to want to be able to take action on these said events. We talked earlier about automating k8s related tasks, but being able to preform true operations for Kubernetes requires the use of [kubectl commands](https://botkube.io/learn/kubectl-cheat-sheet). Botkube has always had a Kubectl plugin to allow ChatOps users to preform, but what if you a developer alone in the chat channel when the DevOps team is OOO? This developer does not know much about operating the Kubernetes cluster. This is where [AIOps](https://botkube.io/learn/use-cases-for-aiops-platform) comes into help with operations and almost become the next iteration of ChatOps.

In the mentioned example above, a developer with little know how of Kubernetes, because the monitoring was set up to alert the channel in Slack so the developer knows he needs to do something. Now he could go back and look at the chat history if the cluster management channel to see how previous SREs or DevOps engineers solved the issue, but now with AIOps, he can simply ask for a solution to the error. Botkube's AI assistant acts as a perfect example of this with allowing a one button click on the k8s error notification that came in to ask AI what the cause was. If it is something Botkube's kubectl permissions are able to solve, such as changing namespaces, then the assistant will offer to run the appropriate command to solve the issue.

Bonus Video of Our Team Explaining what ChatOps for Kubernetes Looks Like[‍](https://www.youtube.com/@KunalKushwaha)
--------------------------------------------------------------------------------------------------------------------

Kunal Kushwaha hosted us on his show for aspiring software & DevOps engineers. Maria, Botkube's Developer Advocate, speaks with him about how to set up ChatOps for monitoring your Kuberentes Cluster. She also mentions how many users have already set this up and using Botkube for team collaboration around operations management.
