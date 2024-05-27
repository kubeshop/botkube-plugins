Title: Unable to Get Local Issuer Certificate Solved!

URL Source: https://botkube.io/learn/fix-the-unable-to-get-local-issuer-certificate-error-in-kubernetes

Markdown Content:
What is the "ssl certificate problem unable to get local issuer certificate" error?
-----------------------------------------------------------------------------------

The <code>ssl certificate problem unable to get local issuer certificate</code> error is a security error that occurs when a Kubernetes cluster is configured to use a self-signed certificate. A self-signed certificate is a certificate that is not signed by a trusted certificate authority. This means that the certificate cannot be verified by the client, which prevents the client from establishing a secure connection to the server.

This error sometimes can be shortened to "ssl git error". It is the Git Error that plagues local clusters on setup. It is hard to get self service security certificates perfect, but hopefully this page can be a good starting point. While security certificates are not unique to K8s, it is a common error that DevOps engineers face when deploying Kubernetes.

What causes the "ssl certificate problem unable to get local issuer certificate" error?
---------------------------------------------------------------------------------------

The <code>ssl certificate problem unable to get local issuer certificate</code> error is caused by the misconfiguration of the SSL certificate on the Kubernetes cluster. When a client attempts to connect to the cluster, the client will not be able to verify the certificate because it is not signed by a trusted certificate authority. This will result in the error message <code>ssl certificate problem unable to get local issuer certificate</code>.

**\*Quick Tip:** Sometimes detecting the error message is the hardest part, most of the time requiring sifting through cluster logs using the command line interface. We created Botkube to assist with this labor intensive process. Having Botkube in a cluster will give developers two advantages to troubleshooting this error:

1.  Slack or Teams alert to a shared channel when this error occurs to allow for immediate action.
2.  Ability to automate the log pulling and filtering all directly from the shared group channel!

See what else [Botkube Cloud can do to help errors and alerts.](https://botkube.io/features)

How can you fix the "ssl certificate problem unable to get local issuer certificate" errors?
--------------------------------------------------------------------------------------------

There are two ways to fix the <code>ssl certificate problem unable to get local issuer certificate</code> errors:

1.  You can add the self-signed certificate to the trusted certificate store on the client. This will allow the client to verify the certificate and establish a secure connection to the cluster.
2.  You can use a certificate signed by a trusted certificate authority. This will ensure that the certificate can be verified by the client and that the connection to the cluster is secure.

How to prevent "ssl certificate problem unable to get local issuer certificate" errors?
---------------------------------------------------------------------------------------

To prevent <code>ssl certificate problem unable to get local issuer certificate</code> errors, you should use a certificate signed by a trusted certificate authority. You can also add the self-signed certificate to the trusted certificate store on the client.

Here are the steps on how to add a self-signed certificate to the trusted certificate store on a Linux machine:

1.  Open the file <code>/etc/ssl/certs/ca-certificates.crt</code>.
2.  Add the self-signed certificate to the file.
3.  Save the file.
4.  Restart the web browser.

Here are the steps on how to install a certificate signed by a trusted certificate authority:

1.  Obtain the certificate from the certificate authority.
2.  Import the certificate into the trusted certificate store on the client.
3.  Restart the web browser.

I hope this article helps you fix the <code>ssl certificate problem unable to get local issuer certificate</code> error in Kubernetes. Be sure to check out the other [K8s error articles](https://botkube.io/learn-main-topic/errors) that try to cover other common errors that developers run into while orchestrating Kubernetes.

One final tip, **do not be afraid to search for tooling that helps with troubleshooting of common errors**. Botkube's AI assistant is a great example of a tool that helps with K8s specific troubleshooting tasks. [Try out Botkube for free](https://app.botkube.io/) to get started with collaborative troubleshooting directly in your communications platform.

Solving the Error using Botkube's AI Troubleshooting Assistant
--------------------------------------------------------------

If you do end up using Botkube's Kubernetes [AI troubleshooting assistant](https://botkube.io/blog/explore-the-new-era-of-aiops-with-botkubes-ai-assistant), follow the below steps to help troubleshoot the SSL Certificate error:

### Step 1 - Start Prompt

Type in **@botkube ai** into the chat platform to start communicating with our helpful AI assistant

![Image 1: Prompting AI to check Kubernetes related SSL issue](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/663e739b18498714e3ffd411_Botkube%20AI%20prompt.png)

### Step 2 - Find where the Issue Occurs within the Cluster  

Ask the AI which namespace, pod, or service is causing the 'SSL Certificate Problem: Unable to Get Local Issuer Certificate Error' in my cluster to find where the issue is coming from. With Botube being aware of the k8s cluster it is installed in, it allows our AI Assistant to scan all the resources being used in the cluster to find the source.

![Image 2: AI finding the cause of Kubernetes error](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/663e7600646f66c7c3f9ff46_SSL%20AI%20question%201.png)

### Step 3 - Ask How to Fix the Error in that Location

After the AI response tells you where this error is occurring, it is time to ask it how to fix the unable to get certificate error.

![Image 3: Asking how to fix an SSL located in a Kubernetes Pod](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/663e7a1d79a3f8aa5b90904c_AI%20prompt%202.png)

### Step 4 - Solve the Issue in DNS

Now unlike some other issues, such as the [OOMKILLED error](https://botkube.io/learn/what-is-oomkilled), Kubectl commands cannot directly fix this issue. So Botkube's assitant cannot take this one fully to a completed issue solved, but by now you show know exactly which domain is causing the issue and in which part of the cluster. Now you will have to go to your DNS provider and renew, reapply, or correct the SSL certificate tied to the domain of the service. The most common domain name service providers are: GoDaddy, Cloudflare, Namecheap, or Google Domains. You will need admin access to the DNS and simply reverify the certificate of the service causing the issue
