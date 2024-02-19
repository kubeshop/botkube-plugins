package flux

import (
	"github.com/kubeshop/botkube/pkg/config"
	pluginx "github.com/kubeshop/botkube/pkg/plugin"
)

// Config holds Flux executor configuration.
type Config struct {
	Logger config.Logger `yaml:"log"`
	GitHub struct {
		Auth struct {
			// The GitHub access token.
			// Instruction for creating a token can be found here: https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/#creating-a-token.
			// When not provided some functionality may not work. For example, adding a comment under pull request.
			AccessToken string `yaml:"accessToken"`
		} `yaml:"auth"`
	} `yaml:"github"`

	// Fields not exposed to the user in the JSON schema
	TmpDir pluginx.TmpDir `yaml:"tmpDir"`
}
