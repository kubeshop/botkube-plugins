package auth

import "github.com/kubeshop/botkube/pkg/api"

func unauthorizedMessage(isOpenSourceInstance bool) api.Message {
	btnBuilder := api.NewMessageButtonBuilder()

	if isOpenSourceInstance {
		return api.Message{
			Sections: []api.Section{
				{
					Base: api.Base{
						Header:      "Cloud-Only Plugin Enabled",
						Description: "This plugin is available only on Botkube Cloud. To use it, migrate your Botkube Agent instance to Botkube Cloud.",
					},
					Buttons: []api.Button{
						btnBuilder.ForURL("Migrate Installation", "https://docs.botkube.io/cli/migrating-installation-to-botkube-cloud", api.ButtonStylePrimary),
						btnBuilder.ForURL("Open Botkube Cloud", "https://app.botkube.io"),
					},
				},
			},
		}
	}

	return api.Message{
		Sections: []api.Section{
			{
				Base: api.Base{
					Header:      "Lost Connection with Botkube Cloud",
					Description: "This plugin requires an active connection to Botkube Cloud. Ensure that your Agent has access to Botkube Cloud with valid credentials. Check Agent logs for more information.",
				},
				Buttons: []api.Button{
					btnBuilder.ForURL("See Troubleshooting Guide", "https://docs.botkube.io/operation/common-problems"),
					btnBuilder.ForURL("How to Get Agent Logs", "https://docs.botkube.io/operation/diagnostics#agent-logs"),
				},
			},
		},
	}
}
