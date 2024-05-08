Title: Unable to Get Local Issuer Certificate Solved!

URL Source: https://botkube.io/learn/fix-the-unable-to-get-local-issuer-certificate-error-in-kubernetes

Markdown Content:
Fix "SSL Certificate Problem: Unable to Get Local Issuer Certificate" Error
---------------------------------------------------------------------------

![Image 1](https://assets-global.website-files.com/634fabb21508d6c9db9bc46f/65fa18185ef2ff8b0ce59fce_LEARN_TN_Errors%20(4).png)

![Image 2](https://assets-global.website-files.com/plugins/Basic/assets/placeholder.60f9b1840c.svg)

Table of Contents
-----------------

*   [What is the "ssl certificate problem unable to get local issuer certificate" error?](#what-is-the-ssl-certificate-problem-unable-to-get-local-issuer-certificate-error--2)
*   [What causes the "ssl certificate problem unable to get local issuer certificate" error?](#what-causes-the-ssl-certificate-problem-unable-to-get-local-issuer-certificate-error--2)
*   [How can you fix the "ssl certificate problem unable to get local issuer certificate" errors?](#how-can-you-fix-the-ssl-certificate-problem-unable-to-get-local-issuer-certificate-errors--2)
*   [How to prevent "ssl certificate problem unable to get local issuer certificate" errors?](#how-to-prevent-ssl-certificate-problem-unable-to-get-local-issuer-certificate-errors--2)

#### Get started with Botkube Cloud

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

### About Botkube

Botkube is a collaborative troubleshooting tool designed specifically for Kubernetes users. With Botkube, you can seamlessly receive and act on alerts directly within your preferred messaging and collaboration platforms like Slack, Microsoft Teams, Discord, and Mattermost. In addition, Botkube enables you to automate actions based on events, run kubectl and Helm commands, receive recommendations for best practices and much more. [Get started with Botkube for free.](https://app.botkube.io/)
