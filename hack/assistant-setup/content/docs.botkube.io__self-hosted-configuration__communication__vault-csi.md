Title: Tokens from Vault via CSI driver | Botkube

URL Source: https://docs.botkube.io/self-hosted-configuration/communication/vault-csi

Markdown Content:
info

This document is applicable only for self-hosted installations. For Botkube Cloud installations, the Botkube Agent configuration, along with all secrets, are managed via the [Botkube Cloud dashboard](https://app.botkube.io/).

This tutorial shows how to install Botkube which takes the configuration from Vault instance.

Prerequisites[​](#prerequisites "Direct link to Prerequisites")
---------------------------------------------------------------

*   Kubernetes Clusters that supports CSI.
    
    For example, to run [K3s](https://k3s.io/) using [Lima](https://github.com/lima-vm/lima), run:
    
        limactl start template://k3s
    
*   [Socket Slack integration installed](https://docs.botkube.io/installation/slack/socket-slack).
    
*   [`helm`](https://helm.sh/docs/intro/install/) v3 installed.
    
*   [`kubectl`](https://kubernetes.io/docs/tasks/tools/#kubectl) installed.
    

Steps[​](#steps "Direct link to Steps")
---------------------------------------

This instruction guides you through the installation of Botkube and Vault on a Kubernetes cluster and configuring them together.

1.  Install Vault with CSI enabled:
    
        helm repo add hashicorp https://helm.releases.hashicorp.comhelm repo updatehelm install vault hashicorp/vault --namespace default \    --set "server.dev.enabled=true" \    --set "injector.enabled=false" \    --set "csi.enabled=true"
    
2.  Add Slack token in Vault:
    
        # Exec to podkubectl exec -n default -it vault-0 -- /bin/sh
    
        # Write the token to Vaultvault kv put -mount=secret slack-app-token token=xapp-...vault kv put -mount=secret slack-bot-token token=xoxb-...
    
3.  Enable Vault's Kubernetes authentication:
    
        vault auth enable kubernetesvault write auth/kubernetes/config \    kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443"
    
        vault policy write internal-app - <<EOFpath "secret/data/slack-app-token" {  capabilities = ["read"]}path "secret/data/slack-bot-token" {  capabilities = ["read"]}EOF
    
        vault write auth/kubernetes/role/database \    bound_service_account_names=botkube-sa \    bound_service_account_namespaces=default \    policies=internal-app \    ttl=20m
    
        # Exit from the Vault Podexit
    
4.  Install the Secrets Store CSI driver:
    
        helm repo add secrets-store-csi-driver https://kubernetes-sigs.github.io/secrets-store-csi-driver/chartshelm install csi secrets-store-csi-driver/secrets-store-csi-driver \--namespace default \--set syncSecret.enabled=true
    
5.  Create BotKue installation parameters:
    
        cat > /tmp/values.yaml << ENDOFFILEextraObjects:  - apiVersion: secrets-store.csi.x-k8s.io/v1    kind: SecretProviderClass    metadata:      name: vault-database    spec:      provider: vault      parameters:        roleName: "database"        vaultAddress: "http://vault.default:8200"        objects: |          - objectName: "slack-app-token"            secretPath: "secret/data/slack-app-token"            secretKey: "token"          - objectName: "slack-bot-token"            secretPath: "secret/data/slack-bot-token"            secretKey: "token"      secretObjects:      - secretName: communication-slack        type: Opaque        data:        - objectName: "slack-app-token"          key: "slack-app-token"        - objectName: "slack-bot-token"          key: "slack-bot-token"communications:  'default-group':  # Settings for SocketSlack    socketSlack:      enabled: true      channels: {} # configure your channels      # botToken - specified via env variable      # appToken - specified via env variableextraEnv:  - name: BOTKUBE_COMMUNICATIONS_DEFAULT-GROUP_SOCKET__SLACK_APP__TOKEN    valueFrom:      secretKeyRef:        name: communication-slack        key: slack-app-token  - name: BOTKUBE_COMMUNICATIONS_DEFAULT-GROUP_SOCKET__SLACK_BOT__TOKEN    valueFrom:      secretKeyRef:        name: communication-slack        key: slack-bot-tokenextraVolumeMounts:  - name: secrets-store-inline    mountPath: "/mnt/secrets-store"    readOnly: trueextraVolumes:  - name: secrets-store-inline    csi:      driver: secrets-store.csi.k8s.io      readOnly: true      volumeAttributes:        secretProviderClass: "vault-database"ENDOFFILE
    
6.  Install Botkube:
    
        helm install botkube --namespace default \-f /tmp/values.yaml \./helm/botkube
