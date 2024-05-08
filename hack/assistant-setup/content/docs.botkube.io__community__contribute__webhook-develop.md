Title: Outgoing Webhook development | Botkube

URL Source: https://docs.botkube.io/community/contribute/webhook-develop

Markdown Content:
*   [](https://docs.botkube.io/)
*   [Contribute](https://docs.botkube.io/community/contribute/)
*   Outgoing Webhook development

Outgoing Webhook development
----------------------------

To develop the outgoing webhook integration, you can use [Echo-Server](https://github.com/Ealenn/Echo-Server), which logs all incoming requests.

Steps[​](#steps"DirectlinktoSteps")
---------------------------------------

1.  Install Echo server with Helm:

helm repo add ealenn https://ealenn.github.io/chartshelm repo update ealennhelm install echo-server ealenn/echo-server --set application.logs.ignore.ping=true --set application.enable.environment=false --wait

2.  Go through the [Outgoing Webhook Botkube installation](https://docs.botkube.io/installation/webhook/) instruction. Provide `http://echo-server.default` as `WEBHOOK_URL`.

3.  Watch the logs:

kubectl logs -l app.kubernetes.io/name=echo-server -f

Every incoming request will be logged there.


[Elasticsearch notifier development](https://docs.botkube.io/community/contribute/elasticsearch-develop)[](https://docs.botkube.io/community/credits)

*   [Steps](#steps)
