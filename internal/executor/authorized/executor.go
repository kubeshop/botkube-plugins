package authorized

import (
	"context"

	"github.com/kubeshop/botkube-cloud-plugins/internal/remote"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/executor"
)

type unauthorizedExecutor struct {
	metadata api.MetadataOutput
}

func NewExecutor(exec executor.Executor) executor.Executor {
	deployClient := remote.NewDeploymentClient(remote.NewDefaultGqlClient())
	ok := deployClient.IsConnectedWithCould()
	if !ok {
		metadata, _ := exec.Metadata(context.Background())
		return &unauthorizedExecutor{
			metadata: metadata,
		}
	}

	return exec
}

func (e unauthorizedExecutor) Execute(context.Context, executor.ExecuteInput) (executor.ExecuteOutput, error) {
	panic("couldn't run the plugin: invalid license")
}

func (e unauthorizedExecutor) Metadata(context.Context) (api.MetadataOutput, error) {
	return e.metadata, nil
}

func (e unauthorizedExecutor) Help(context.Context) (api.Message, error) {
	panic("couldn't run the plugin: invalid license")
}
