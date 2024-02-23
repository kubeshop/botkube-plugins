package main

import (
	_ "embed"

	"github.com/kubeshop/botkube-cloud-plugins/internal/auth"

	"github.com/hashicorp/go-plugin"

	aibrain "github.com/kubeshop/botkube-cloud-plugins/internal/source/ai-brain"
	"github.com/kubeshop/botkube/pkg/api/source"
)

// version is set via ldflags by GoReleaser.
var version = "dev"

func main() {
	source.Serve(map[string]plugin.Plugin{
		aibrain.PluginName: &source.Plugin{
			Source: auth.NewProtectedSource(aibrain.NewSource(version)),
		},
	})
}
