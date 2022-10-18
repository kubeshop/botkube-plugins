# Map of enabled communication mediums. The `communications` property name is an alias for a given configuration.
#
# Format: communications.{alias}
communications:
  'default-group':
    # Settings for Slack
    slack:
      enabled: false
      channels:
        'alias':
          name: 'SLACK_CHANNEL'
          notification:
            # -- If true, the notifications are not sent to the channel. They can be enabled with `@Botkube` command anytime.
            disabled: false
          bindings:
            executors:
              - 'kubectl-read-only'
            sources:
              - 'k8s-events'
      token: "" # SLACK_API_TOKEN
      notification:
        type: short                             # Change notification type short/long you want to receive. Type is optional and default is short.

    # Settings for Slack with Socket Mode
    socketSlack:
      enabled: false
      channels:
        'alias':
          name: 'SLACK_CHANNEL'
          bindings:
            executors:
              - 'kubectl-read-only'
            sources:
              - 'k8s-events'
      botToken: "" # SLACK_BOT_TOKEN
      appToken: "" # SLACK_APP_TOKEN
      notification:
        type: short                             # Change notification type short/long you want to receive. Type is optional and default is short.

    # Settings for Mattermost
    mattermost:
      enabled: false
      url: 'MATTERMOST_SERVER_URL'              # URL where Mattermost is running. e.g https://example.com:9243
      token: 'MATTERMOST_TOKEN'                 # Personal Access token generated by Botkube user
      team: 'MATTERMOST_TEAM'                   # Mattermost Team to configure with Botkube
      botName: 'Botkube'                        # Bot name
      channels:
        'alias':
          name: 'MATTERMOST_CHANNEL'            # Mattermost Channel for receiving Botkube alerts:
          notification:
            # -- If true, the notifications are not sent to the channel. They can be enabled with `@Botkube` command anytime.
            disabled: false
          bindings:
            executors:
              - kubectl-read-only
            sources:
              - k8s-events
      notification:
        type: short                             # Change notification type short/long you want to receive. Type is optional and default is short.

    # Settings for MS Teams
    teams:
      enabled: false
      appID: 'APPLICATION_ID'
      appPassword: 'APPLICATION_PASSWORD'
      botName: 'Botkube'
      notification:
        type: short
      port: 3978

    # Settings for Discord
    discord:
      enabled: false
      token: 'DISCORD_TOKEN'                    # Botkube Bot Token
      botID: 'DISCORD_BOT_ID'                   # Botkube Application Client ID
      channels:
        'alias':
          id: 'DISCORD_CHANNEL_ID'            # Discord Channel id for receiving Botkube alerts:
          notification:
            # -- If true, the notifications are not sent to the channel. They can be enabled with `@Botkube` command anytime.
            disabled: false
          bindings:
            executors:
              - kubectl-read-only
            sources:
              - k8s-events
      notification:
        type: short                             # Change notification type short/long you want to receive. Type is optional and default is short.


    # Settings for ELS
    elasticsearch:
      enabled: false
      awsSigning:
        enabled: false                          # enable awsSigning using IAM for Elastisearch hosted on AWS, if true make sure AWS environment variables are set. Refer https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html
        awsRegion: 'us-east-1'                  # AWS region where Elasticsearch is deployed
        roleArn: ''                             # AWS IAM Role arn to assume for credentials, use this only if you dont want to use the EC2 instance role or not running on AWS instance
      server: 'ELASTICSEARCH_ADDRESS'           # e.g https://example.com:9243
      username: 'ELASTICSEARCH_USERNAME'        # Basic Auth
      password: 'ELASTICSEARCH_PASSWORD'
      skipTLSVerify: false                      # toggle verification of TLS certificate of the Elastic nodes. Verification is skipped when option is true. Enable to connect to clusters with self-signed certs
      # ELS index settings
      indices:
        'alias':
          name: botkube
          type: botkube-event
          shards: 1
          bindings:
            sources:
              - "k8s-events"
            # executors - not allowed in this case, ES is "sink" only.
    # Settings for Webhook
    webhook:
      enabled: false
      url: 'WEBHOOK_URL'                        # e.g https://example.com:80
      bindings:
        # -- Notification sources configuration for the webhook.
        sources:
          - k8s-events
