package authorized

import (
	"context"

	"github.com/kubeshop/botkube-cloud-plugins/internal/remote"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
)

type unauthorizedSource struct {
	metadata api.MetadataOutput
}

func NewSource(source source.Source) source.Source {
	deployClient := remote.NewDeploymentClient(remote.NewDefaultGqlClient())
	ok := deployClient.IsConnectedWithCould()
	if !ok {
		metadata, _ := source.Metadata(context.Background())
		return &unauthorizedSource{
			metadata: metadata,
		}
	}

	return source
}

func (s unauthorizedSource) Stream(context.Context, source.StreamInput) (source.StreamOutput, error) {
	panic("couldn't run the plugin: invalid license")
}

func (s unauthorizedSource) Metadata(context.Context) (api.MetadataOutput, error) {
	return s.metadata, nil
}

func (s unauthorizedSource) HandleExternalRequest(context.Context, source.ExternalRequestInput) (source.ExternalRequestOutput, error) {
	panic("couldn't run the plugin: invalid license")
}
