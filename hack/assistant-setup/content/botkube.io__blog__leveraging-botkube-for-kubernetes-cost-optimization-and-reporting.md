Title: Leveraging Botkube for Kubernetes Cost Optimization and Reporting

URL Source: https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting

Published Time: Jan 29, 2024

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6669d3e5f7e06a9b425573d6_rohit-ghumare-photo.jpg)

Rohit Ghumare

Google Developer Expert – Google Cloud, CNCF Ambassador

Botkube

This tutorial will guide you through using Botkube to monitor resource usage and optimize costs in Kubernetes. Read on to learn more.

### Table of Contents

*   [Introduction](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#introduction)
*   [What is Botkube?](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#what-is-botkube-)
*   [Prerequisites](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#prerequisites)
*   [Setting Up Resource Monitoring in Botkube](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#setting-up-resource-monitoring-in-botkube)
*   [Creating Alerts for Resource Usage Anomalies](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#creating-alerts-for-resource-usage-anomalies)
*   [Generating Cost Reports](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#generating-cost-reports)
*   [Analyzing and Acting on the Data](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#analyzing-and-acting-on-the-data)
*   [Conclusion](https://botkube.io/blog/leveraging-botkube-for-kubernetes-cost-optimization-and-reporting#conclusion)

#### Start Using Botkube AI-Powered Assistant Today

**Introduction**
----------------

As Kubernetes environments grow in complexity and scale, managing costs becomes increasingly crucial. This tutorial will guide you through using Botkube to monitor resource usage and optimize costs in Kubernetes. By the end of this tutorial, you'll be able to set up alerts and reports to track and manage your Kubernetes resource utilization effectively.

What is Botkube?
----------------

Botkube stands out as a vital messaging solution for keeping a close eye on and troubleshooting Kubernetes clusters. Seamlessly integrating with popular platforms such as [Slack](https://botkube.io/integration/slack), [Microsoft Teams](https://botkube.io/integration/teams), and Mattermost, it keeps you in the loop with instant alerts and updates on the status of Kubernetes resources. Tailor Botkube will keep tabs on particular resources and utilize its handy feature to run kubectl operations with ease, making cluster management a breeze. It's a game-changer for DevOps teams, offering them the agility to tackle issues swiftly, keep the health of clusters in check, and smooth out the overall process of operating Kubernetes.

‍

![Image 2](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/65b72fc399ffea33dc180476_botklube_diagram.png)

botkube architecture diagram

.

The above architecture represents what Botkube is capable of and how it integrates with plugins and communications channels. We will focus on the monitoring side of this architecture more for this particular article; I’m talking about `executor`s and `events`.

Botkube documentation includes a [diagnostics page](https://docs.botkube.io/operation/diagnostics) for the debugging issues in Helm whenever you’re stuck, so bookmark it for now.

**Prerequisites**
-----------------

*   A Kubernetes cluster.
*   Botkube is installed and configured in your cluster.
*   Basic familiarity with Kubernetes resource management.
*   We recommend checking the [previous blog](https://botkube.io/blog/streamlining-kubernetes-management-a-closer-look-at-botkube-chatops) for a detailed guide on Botkube Installation on the Kubernetes cluster for better understanding and learning.

**Setting Up Resource Monitoring in Botkube**
---------------------------------------------

### A. **Configuring Botkube to Monitor Specific Resources**

Configure BotKube to monitor resources that significantly impact costs, such as Pods, Deployments, and StatefulSets. Modify the BotKube configuration to include these resources and set the desired level of alerting.

> **`If you’ve installed Botkube using the helm as given in the blog listed above, you must find the values.yaml file that helps you to configure the changes into Botkube as per your requirement.`**

‍

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="values.yaml"><tbody><tr><td id="file-values-yaml-L1" data-line-number="1"></td><td id="file-values-yaml-LC1"><span>resources</span>:</td></tr><tr><td id="file-values-yaml-L2" data-line-number="2"></td><td id="file-values-yaml-LC2">- <span>name</span>: <span>pod</span></td></tr><tr><td id="file-values-yaml-L3" data-line-number="3"></td><td id="file-values-yaml-LC3"><span>namespaces</span>:</td></tr><tr><td id="file-values-yaml-L4" data-line-number="4"></td><td id="file-values-yaml-LC4"><span>include</span>:</td></tr><tr><td id="file-values-yaml-L5" data-line-number="5"></td><td id="file-values-yaml-LC5">- <span>all</span></td></tr><tr><td id="file-values-yaml-L6" data-line-number="6"></td><td id="file-values-yaml-LC6"><span>events</span>:</td></tr><tr><td id="file-values-yaml-L7" data-line-number="7"></td><td id="file-values-yaml-LC7">- <span>create</span></td></tr><tr><td id="file-values-yaml-L8" data-line-number="8"></td><td id="file-values-yaml-LC8">- <span>delete</span></td></tr><tr><td id="file-values-yaml-L9" data-line-number="9"></td><td id="file-values-yaml-LC9">- <span>error</span></td></tr><tr><td id="file-values-yaml-L10" data-line-number="10"></td><td id="file-values-yaml-LC10">- <span>name</span>: <span>deployment</span></td></tr><tr><td id="file-values-yaml-L11" data-line-number="11"></td><td id="file-values-yaml-LC11"><span>namespaces</span>:</td></tr><tr><td id="file-values-yaml-L12" data-line-number="12"></td><td id="file-values-yaml-LC12"><span>include</span>:</td></tr><tr><td id="file-values-yaml-L13" data-line-number="13"></td><td id="file-values-yaml-LC13">- <span>all</span></td></tr><tr><td id="file-values-yaml-L14" data-line-number="14"></td><td id="file-values-yaml-LC14"><span>events</span>:</td></tr><tr><td id="file-values-yaml-L15" data-line-number="15"></td><td id="file-values-yaml-LC15">- <span>create</span></td></tr><tr><td id="file-values-yaml-L16" data-line-number="16"></td><td id="file-values-yaml-LC16">- <span>delete</span></td></tr><tr><td id="file-values-yaml-L17" data-line-number="17"></td><td id="file-values-yaml-LC17">- <span>error</span></td></tr></tbody></table>

### B. **Enabling Resource Specification Checks**

Enable checks for resource specifications like requests and limits to ensure optimal allocation of CPU and memory resources. This helps in preventing over-provisioning and under-provisioning of resources.

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="resource-specifications-botkube.yaml"><tbody><tr><td id="file-resource-specifications-botkube-yaml-L1" data-line-number="1"></td><td id="file-resource-specifications-botkube-yaml-LC1"><span>settings</span>:</td></tr><tr><td id="file-resource-specifications-botkube-yaml-L2" data-line-number="2"></td><td id="file-resource-specifications-botkube-yaml-LC2"><span>resources</span>:</td></tr><tr><td id="file-resource-specifications-botkube-yaml-L3" data-line-number="3"></td><td id="file-resource-specifications-botkube-yaml-LC3"><span>requests</span>:</td></tr><tr><td id="file-resource-specifications-botkube-yaml-L4" data-line-number="4"></td><td id="file-resource-specifications-botkube-yaml-LC4"><span>cpu</span>: <span><span>"</span>100m<span>"</span></span></td></tr><tr><td id="file-resource-specifications-botkube-yaml-L5" data-line-number="5"></td><td id="file-resource-specifications-botkube-yaml-LC5"><span>memory</span>: <span><span>"</span>100Mi<span>"</span></span></td></tr><tr><td id="file-resource-specifications-botkube-yaml-L6" data-line-number="6"></td><td id="file-resource-specifications-botkube-yaml-LC6"><span>limits</span>:</td></tr><tr><td id="file-resource-specifications-botkube-yaml-L7" data-line-number="7"></td><td id="file-resource-specifications-botkube-yaml-LC7"><span>cpu</span>: <span><span>"</span>200m<span>"</span></span></td></tr><tr><td id="file-resource-specifications-botkube-yaml-L8" data-line-number="8"></td><td id="file-resource-specifications-botkube-yaml-LC8"><span>memory</span>: <span><span>"</span>200Mi</span></td></tr></tbody></table>

**Creating Alerts for Resource Usage Anomalies**
------------------------------------------------

### **1\. Configuring Alerts for Over-Utilization**

Set up alerts in Botkube for scenarios where resource utilization exceeds a specified threshold. These alerts can notify you when resources are over-utilized, indicating potential cost inefficiencies.

‍

### **2\. Alerts for Under-Utilization**

Similarly, configure alerts for under-utilized resources, which can help identify opportunities for scaling down and cost savings.

**Generating Cost Reports**
---------------------------

### **1\. Setting Up Periodic Reporting**

To set up periodic reporting in BotKube, you must configure a custom script or plugin that periodically gathers resource usage data and sends a report. Let's assume we have a script named `generate_cost_report.sh` that compiles the report. We will schedule this script to run at regular intervals using a Kubernetes CronJob.

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="YAML" data-tagsearch-path="periodic-reporting-botkube.yaml"><tbody><tr><td id="file-periodic-reporting-botkube-yaml-L1" data-line-number="1"></td><td id="file-periodic-reporting-botkube-yaml-LC1"><span>apiVersion</span>: <span>batch/v1</span></td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L2" data-line-number="2"></td><td id="file-periodic-reporting-botkube-yaml-LC2"><span>kind</span>: <span>CronJob</span></td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L3" data-line-number="3"></td><td id="file-periodic-reporting-botkube-yaml-LC3"><span>metadata</span>:</td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L4" data-line-number="4"></td><td id="file-periodic-reporting-botkube-yaml-LC4"><span>name</span>: <span>cost-report-cronjob</span></td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L5" data-line-number="5"></td><td id="file-periodic-reporting-botkube-yaml-LC5"><span>spec</span>:</td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L6" data-line-number="6"></td><td id="file-periodic-reporting-botkube-yaml-LC6"><span>schedule</span>: <span><span>"</span>0 0 * * *<span>"</span></span> <span><span>#</span> This schedule is for daily execution at midnight</span></td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L7" data-line-number="7"></td><td id="file-periodic-reporting-botkube-yaml-LC7"><span>jobTemplate</span>:</td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L8" data-line-number="8"></td><td id="file-periodic-reporting-botkube-yaml-LC8"><span>spec</span>:</td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L9" data-line-number="9"></td><td id="file-periodic-reporting-botkube-yaml-LC9"><span>template</span>:</td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L10" data-line-number="10"></td><td id="file-periodic-reporting-botkube-yaml-LC10"><span>spec</span>:</td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L11" data-line-number="11"></td><td id="file-periodic-reporting-botkube-yaml-LC11"><span>containers</span>:</td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L12" data-line-number="12"></td><td id="file-periodic-reporting-botkube-yaml-LC12">- <span>name</span>: <span>cost-report</span></td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L13" data-line-number="13"></td><td id="file-periodic-reporting-botkube-yaml-LC13"><span>image</span>: <span>your-report-generator-image</span></td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L14" data-line-number="14"></td><td id="file-periodic-reporting-botkube-yaml-LC14"><span>command</span>: <span>["/bin/bash", "-c", "./generate_cost_report.sh"]</span></td></tr><tr><td id="file-periodic-reporting-botkube-yaml-L15" data-line-number="15"></td><td id="file-periodic-reporting-botkube-yaml-LC15"><span>restartPolicy</span>: <span>OnFailure</span></td></tr></tbody></table>

In this configuration, replace `your-report-generator-image` with the Docker image that contains your script. The `generate_cost_report.sh` should output the report in a format Botkube can send to your communication platform.

### **2\. Custom Reporting Scripts**

Here's a simple Bash script example to generate a cost report. This script could be expanded to pull more detailed data as per your requirements.

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Shell" data-tagsearch-path="reporting-script.shell"><tbody><tr><td id="file-reporting-script-shell-L1" data-line-number="1"></td><td id="file-reporting-script-shell-LC1"><span><span>#!</span>/bin/bash</span></td></tr><tr><td id="file-reporting-script-shell-L2" data-line-number="2"></td><td id="file-reporting-script-shell-LC2"></td></tr><tr><td id="file-reporting-script-shell-L3" data-line-number="3"></td><td id="file-reporting-script-shell-LC3"><span><span>#</span> generate_cost_report.sh</span></td></tr><tr><td id="file-reporting-script-shell-L4" data-line-number="4"></td><td id="file-reporting-script-shell-LC4"></td></tr><tr><td id="file-reporting-script-shell-L5" data-line-number="5"></td><td id="file-reporting-script-shell-LC5"><span><span>#</span> Fetch resource usage details</span></td></tr><tr><td id="file-reporting-script-shell-L6" data-line-number="6"></td><td id="file-reporting-script-shell-LC6">pod_usage=<span><span>$(</span>kubectl top pod --all-namespaces --sort-by=<span><span>'</span>cpu<span>'</span></span> <span>|</span> head -5<span>)</span></span></td></tr><tr><td id="file-reporting-script-shell-L7" data-line-number="7"></td><td id="file-reporting-script-shell-LC7">node_usage=<span><span>$(</span>kubectl top node <span>|</span> head -5<span>)</span></span></td></tr><tr><td id="file-reporting-script-shell-L8" data-line-number="8"></td><td id="file-reporting-script-shell-LC8"></td></tr><tr><td id="file-reporting-script-shell-L9" data-line-number="9"></td><td id="file-reporting-script-shell-LC9"><span><span>#</span> Format the report</span></td></tr><tr><td id="file-reporting-script-shell-L10" data-line-number="10"></td><td id="file-reporting-script-shell-LC10">report=<span><span>"</span>Kubernetes Cost Report:\n\n<span>"</span></span></td></tr><tr><td id="file-reporting-script-shell-L11" data-line-number="11"></td><td id="file-reporting-script-shell-LC11">report+=<span><span>"</span>Top 5 CPU Consuming Pods:\n<span>${pod_usage}</span>\n\n<span>"</span></span></td></tr><tr><td id="file-reporting-script-shell-L12" data-line-number="12"></td><td id="file-reporting-script-shell-LC12">report+=<span><span>"</span>Top 5 CPU Consuming Nodes:\n<span>${node_usage}</span><span>"</span></span></td></tr><tr><td id="file-reporting-script-shell-L13" data-line-number="13"></td><td id="file-reporting-script-shell-LC13"></td></tr><tr><td id="file-reporting-script-shell-L14" data-line-number="14"></td><td id="file-reporting-script-shell-LC14"><span><span>#</span> Send the report via BotKube</span></td></tr><tr><td id="file-reporting-script-shell-L15" data-line-number="15"></td><td id="file-reporting-script-shell-LC15"><span>echo</span> <span><span>"</span><span>$report</span><span>"</span></span> <span>|</span> botkube send -n botkube -c botkube-communication-channel</td></tr><tr><td id="file-reporting-script-shell-L16" data-line-number="16"></td><td id="file-reporting-script-shell-LC16"></td></tr></tbody></table>

‍```
‍
```

**Analyzing and Acting on the Data**
------------------------------------

### **1\. Analyzing Reports**

While the script handles the generation and sending of the reports, analysis is more of a manual process. You would typically read through the reports to identify trends and anomalies in resource usage.

### **2\. Implementing Cost-Saving Measures**

Based on the analysis, you might write scripts to automate scaling actions. Here's a simple script to scale down a deployment if it's underutilized:

<table data-hpc="" data-tab-size="8" data-paste-markdown-skip="" data-tagsearch-lang="Shell" data-tagsearch-path="check-scale-down.shell"><tbody><tr><td id="file-check-scale-down-shell-L1" data-line-number="1"></td><td id="file-check-scale-down-shell-LC1"><span><span>#!</span>/bin/bash</span></td></tr><tr><td id="file-check-scale-down-shell-L2" data-line-number="2"></td><td id="file-check-scale-down-shell-LC2"></td></tr><tr><td id="file-check-scale-down-shell-L3" data-line-number="3"></td><td id="file-check-scale-down-shell-LC3"><span><span>#</span> check_and_scale_down.sh</span></td></tr><tr><td id="file-check-scale-down-shell-L4" data-line-number="4"></td><td id="file-check-scale-down-shell-LC4"></td></tr><tr><td id="file-check-scale-down-shell-L5" data-line-number="5"></td><td id="file-check-scale-down-shell-LC5"><span><span>#</span> Define threshold and deployment</span></td></tr><tr><td id="file-check-scale-down-shell-L6" data-line-number="6"></td><td id="file-check-scale-down-shell-LC6">cpu_threshold=30 <span><span>#</span> 30%</span></td></tr><tr><td id="file-check-scale-down-shell-L7" data-line-number="7"></td><td id="file-check-scale-down-shell-LC7">deployment_name=<span><span>"</span>your-deployment<span>"</span></span></td></tr><tr><td id="file-check-scale-down-shell-L8" data-line-number="8"></td><td id="file-check-scale-down-shell-LC8">namespace=<span><span>"</span>your-namespace<span>"</span></span></td></tr><tr><td id="file-check-scale-down-shell-L9" data-line-number="9"></td><td id="file-check-scale-down-shell-LC9"></td></tr><tr><td id="file-check-scale-down-shell-L10" data-line-number="10"></td><td id="file-check-scale-down-shell-LC10"><span><span>#</span> Fetch current CPU usage (in percentage)</span></td></tr><tr><td id="file-check-scale-down-shell-L11" data-line-number="11"></td><td id="file-check-scale-down-shell-LC11">current_cpu_usage=<span><span>$(</span>kubectl top pod -n <span>$namespace</span> <span>|</span> grep <span>$deployment_name</span> <span>|</span> awk <span><span>'</span>{print $3}<span>'</span></span> <span>|</span> sed <span><span>'</span>s/%//<span>'</span></span><span>)</span></span></td></tr><tr><td id="file-check-scale-down-shell-L12" data-line-number="12"></td><td id="file-check-scale-down-shell-LC12"></td></tr><tr><td id="file-check-scale-down-shell-L13" data-line-number="13"></td><td id="file-check-scale-down-shell-LC13"><span><span>#</span> Check if current usage is below threshold and scale down if it is</span></td></tr><tr><td id="file-check-scale-down-shell-L14" data-line-number="14"></td><td id="file-check-scale-down-shell-LC14"><span>if</span> [ <span><span>"</span><span>$current_cpu_usage</span><span>"</span></span> <span>-lt</span> <span><span>"</span><span>$cpu_threshold</span><span>"</span></span> ]<span>;</span> <span>then</span></td></tr><tr><td id="file-check-scale-down-shell-L15" data-line-number="15"></td><td id="file-check-scale-down-shell-LC15"><span>echo</span> <span><span>"</span>Current CPU usage (<span>$current_cpu_usage</span>%) is below threshold (<span>$cpu_threshold</span>%). Scaling down...<span>"</span></span></td></tr><tr><td id="file-check-scale-down-shell-L16" data-line-number="16"></td><td id="file-check-scale-down-shell-LC16">kubectl scale deployment <span>$deployment_name</span> --replicas=1 -n <span>$namespace</span></td></tr><tr><td id="file-check-scale-down-shell-L17" data-line-number="17"></td><td id="file-check-scale-down-shell-LC17"><span>else</span></td></tr><tr><td id="file-check-scale-down-shell-L18" data-line-number="18"></td><td id="file-check-scale-down-shell-LC18"><span>echo</span> <span><span>"</span>Current CPU usage is above threshold. No scaling performed.<span>"</span></span></td></tr><tr><td id="file-check-scale-down-shell-L19" data-line-number="19"></td><td id="file-check-scale-down-shell-LC19"><span>fi</span></td></tr></tbody></table>

`‍`The script `check_and_scale_down.sh` has been created. This script checks the CPU usage of a specific deployment and scales it down if the usage is below the defined threshold. You can customize the `cpu_threshold<strong>`, `deployment_name`, and `namespace` variables as per your requirements.

> 💡 This script is a basic example and serves as a starting point. Depending on your cluster's complexity and specific needs, you may need to expand or modify it. Remember to test any automation scripts in a controlled environment before deploying them in production.

**Conclusion**
--------------

Following these steps, you can effectively use Botkube to monitor, report, and optimize your Kubernetes resource usage for better cost management. Regular monitoring and proactive management are key to maintaining an efficient and cost-effective Kubernetes environment.

Remember, the configurations and scripts can be further customized to suit your specific needs and the complexity of your Kubernetes environment.

For more details, Join the Botkube-related discussion on Slack! Create your Slack account on [Botkube](https://join.botkube.io/) workspace today. To report bugs or features, use [GitHub issues](https://github.com/kubeshop/botkube/issues/new/choose).
