Title: Take Control of K8s Incident Management Alerts with Botkube and PagerDuty

URL Source: https://botkube.io/blog/take-control-of-k8s-incident-management-alerts-with-botkube-and-pagerduty

Published Time: Jun 07, 2024

Markdown Content:
![Image 1](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/636df3fb36b4e60920a3b1b2_hPLC9itV8zp-raGDFmvOZMfn2hV8RFcl237qzT8Wa1g.jpeg)

Kelly Revenaugh

Developer Relations Lead

Botkube

Customize Botkube alerts to automatically create PagerDuty tickets, streamline incident management, reduce manual work, and improve response times.

### Table of Contents

#### Start Using Botkube AI-Powered Assistant Today

#### Start Receiving Kubernetes Notifications Now

In the fast-paced world of DevOps and Site Reliability Engineering (SRE), staying on top of incidents and managing them efficiently is crucial. We are thrilled to announce the new integration between [Botkube and PagerDuty](https://botkube.io/integration/pagerduty-kubernetes-integration), designed to enhance your incident management workflow by automating the creation of tickets in PagerDuty based on customized monitoring alerts from Botkube.

Managing incidents quickly and effectively can be the difference between minor disruptions and major outages. Botkube is already a powerful tool for monitoring Kubernetes clusters and sending real-time alerts to your communication channels like [Slack](https://botkube.io/integration/slack) or [Microsoft Teams](https://botkube.io/integration/teams). PagerDuty, on the other hand, is a robust platform for managing and responding to incidents. By integrating these two tools, you can are providing a seamless way to ensure that critical alerts are not just seen but acted upon promptly.

![Image 2](https://cdn.prod.website-files.com/634fabb21508d6c9db9bc46f/6662fdc9be9496c201bd4997_Screenshot_Template_Botkube%20(2000%20x%201500%20px)%20(1).png)

‍

### Key Features of the Botkube and PagerDuty Integration

*   **Automated Ticket Creation**: With this integration, you can configure [Botkube to automatically create a ticket in PagerDuty](https://docs.botkube.io/installation/pagerduty/) whenever a specific alert or change is triggered. This ensures that your team is immediately aware of potential issues and can start addressing them without delay.
    
*   **Customizable Alerts**: You have [full control over the alerts](https://docs.botkube.io/installation/pagerduty/) that generate PagerDuty tickets. Customize the alerting filters, define which types of alerts should be escalated, and then configure the incident management protocols via PagerDuty. This means you only get notified about the most critical issues, reducing alert fatigue.
    
*   **Streamlined Workflow**: Once an alert is escalated to PagerDuty, your incident management process becomes more efficient. PagerDuty's powerful features, such as on-call schedules, escalation policies, and incident tracking, can be utilized to ensure a rapid and coordinated response. ‍
    

### Setting Up the Integration

[Setting up the Botkube and PagerDuty integration is straightforward.](https://docs.botkube.io/installation/pagerduty) Here’s a quick guide to get you started:

1.  **Configure Botkube & PagerDuty**:
    
    *   Ensure you have Botkube installed in your Kubernetes cluster.
        
    *   Update your Botkube configuration to include the PagerDuty platform.
        
    *   Create a new [PagerDuty Service](https://docs.botkube.io/) and share the Integration Key with Botkube.
        
2.  **Customize Alerts**:
    
    *   Define the alerts in Botkube that you want to escalate to PagerDuty. You can set filters based on namespaces, resources, and severity levels.
3.  **Set Up PagerDuty**:
    
    *   In PagerDuty, customize the incident rules and workflows to handle the tickets created by Botkube alerts. This includes setting up notification rules, escalation policies, and on-call schedules to ensure the right people are alerted at the right time.

### Get Started Today

The integration between Botkube and PagerDuty is a game-changer for teams looking to streamline their incident management processes quickly and easily

To get started with the Botkube and PagerDuty integration, check out our [detailed set up guide](https://docs.botkube.io/installation/pagerduty). If you’re new to Botkube, [sign up for a free account today](http://app.botkube.io/) to take advantage of the PagerDuty integration, our AI-powered Kubernetes Assistant, multi-cluster management and more.
