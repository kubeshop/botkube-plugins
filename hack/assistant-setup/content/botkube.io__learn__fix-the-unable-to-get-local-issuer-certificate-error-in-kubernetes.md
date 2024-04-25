Title: Unable to Get Local Issuer Certificate Solved!

URL Source: https://botkube.io/learn/fix-the-unable-to-get-local-issuer-certificate-error-in-kubernetes

Markdown Content:
What is the "ssl certificate problem unable to get local issuer certificate" error?
-----------------------------------------------------------------------------------

The `ssl certificate problem unable to get local issuer certificate` error is a security error that occurs when a Kubernetes cluster is configured to use a self-signed certificate. A self-signed certificate is a certificate that is not signed by a trusted certificate authority. This means that the certificate cannot be verified by the client, which prevents the client from establishing a secure connection to the server.

This error sometimes can be shortened to "ssl git error". It is the Git Error that plagues local clusters on setup. It is hard to get self service security certificates perfect, but hopefully this page can be a good starting point. While security certificates are not unique to K8s, it is a common error that DevOps engineers face when deploying Kubernetes.

What causes the "ssl certificate problem unable to get local issuer certificate" error?
---------------------------------------------------------------------------------------

The `ssl certificate problem unable to get local issuer certificate` error is caused by the misconfiguration of the SSL certificate on the Kubernetes cluster. When a client attempts to connect to the cluster, the client will not be able to verify the certificate because it is not signed by a trusted certificate authority. This will result in the error message `ssl certificate problem unable to get local issuer certificate`.

**\*Quick Tip:** Sometimes detecting the error message is the hardest part, most of the time requiring sifting through cluster logs using the command line interface. We created Botkube to assist with this labor intensive process. Having Botkube in a cluster will give developers two advantages to troubleshooting this error:

1.  Slack or Teams alert to a shared channel when this error occurs to allow for immediate action.
2.  Ability to automate the log pulling and filtering all directly from the shared group channel!

See what else [Botkube Cloud can do to help errors and alerts.](https://botkube.io/features)

How can you fix the "ssl certificate problem unable to get local issuer certificate" errors?
--------------------------------------------------------------------------------------------

There are two ways to fix the `ssl certificate problem unable to get local issuer certificate` errors:

1.  You can add the self-signed certificate to the trusted certificate store on the client. This will allow the client to verify the certificate and establish a secure connection to the cluster.
2.  You can use a certificate signed by a trusted certificate authority. This will ensure that the certificate can be verified by the client and that the connection to the cluster is secure.

How to prevent "ssl certificate problem unable to get local issuer certificate" errors?
---------------------------------------------------------------------------------------

To prevent `ssl certificate problem unable to get local issuer certificate` errors, you should use a certificate signed by a trusted certificate authority. You can also add the self-signed certificate to the trusted certificate store on the client.

Here are the steps on how to add a self-signed certificate to the trusted certificate store on a Linux machine:

1.  Open the file `/etc/ssl/certs/ca-certificates.crt`.
2.  Add the self-signed certificate to the file.
3.  Save the file.
4.  Restart the web browser.

Here are the steps on how to install a certificate signed by a trusted certificate authority:

1.  Obtain the certificate from the certificate authority.
2.  Import the certificate into the trusted certificate store on the client.
3.  Restart the web browser.

I hope this article helps you fix the `ssl certificate problem unable to get local issuer certificate` error in Kubernetes. Be sure to check out the other [K8s error articles](https://botkube.io/learn-main-topic/errors) that try to cover other common errors that developers run into while orchestrating Kubernetes.

One final tip, **do not be afraid to search for tooling that helps with troubleshooting of common errors**. Botkube's AI assistant is a great example of a tool that helps with K8s specific troubleshooting tasks. [Try out Botkube for free](http://app.botkube.io/) to get started with collaborative troubleshooting directly in your communications platform.
