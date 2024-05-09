Title: MS Teams Bot development | Botkube

URL Source: https://docs.botkube.io/community/contribute/ms-teams-develop

Markdown Content:
*   [](https://docs.botkube.io/)
*   [Contribute](https://docs.botkube.io/community/contribute/)
*   MS Teams Bot development

MS Teams Bot development
------------------------

Microsoft Teams is an entirely cloud-based product. Because of this, Botkube must be publicly accessible via an HTTPS endpoint.

For a local Kubernetes cluster, you can use the tunneling software, for example [`ngrok`](https://ngrok.com/). It creates an externally addressable URL for a port you open locally on your machine.

Steps[​](#steps "Direct link to Steps")
---------------------------------------

1.  Install [`ngrok`](https://ngrok.com/download).
2.  Start `ngrok`:
3.  Install Botkube according to the [MS Teams tutorial](https://docs.botkube.io/installation/teams/), but apply the following changes:
    1.  Ignore the [Prerequisites](https://docs.botkube.io/installation/teams/#prerequisites) section and skip the Ingress controller installation.
    2.  In the [B. Deploy Botkube controller](https://docs.botkube.io/installation/teams/#b-deploy-botkube-controller) section, skip all parts related to the certificate and Ingress setup.
    3.  In the [Add a Bot to the App](https://docs.botkube.io/installation/teams/#add-the-bot-feature-to-the-app) section, specify the base URL from step 2 for the **Messaging Endpoint** property. Use the following URL format: `{ngrok_url}/bots/teams/v1/messages`, for example, `https://177b-37-30-104-55.eu.ngrok.io/bots/teams/v1/messages`.

If you stop and restart `ngrok`, the URL changes. In such a case, you need to update the **Messaging Endpoint** property and reinstall the bot into your Team.

4.  After Botkube installation, forward the local port to the Kubernetes Botkube Service:
    
        kubectl port-forward -n botkube svc/botkube 3978
    

[](https://docs.botkube.io/community/contribute/release)[Elasticsearch notifier development](https://docs.botkube.io/community/contribute/elasticsearch-develop)

*   [Steps](#steps)
