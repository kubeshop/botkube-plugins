package auth

import (
	"context"

	"github.com/kubeshop/botkube-cloud-plugins/internal/remote"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/executor"
)

// ProtectedExecutor protects Executor usage without active Cloud connection.
type ProtectedExecutor struct {
	underlying             executor.Executor
	openSourceBlockage     bool
	initConnectionBlockage bool
	authChecker            *remote.ConnChecker
}

// NewProtectedExecutor returns wrapped Executor instance with Cloud connection checker.
func NewProtectedExecutor(exec executor.Executor) executor.Executor {
	cfg, ok := remote.GetConfig()
	if !ok {
		return &ProtectedExecutor{
			underlying:         exec,
			openSourceBlockage: true,
		}
	}

	authChecker := remote.NewConnChecker(cfg)
	err := authChecker.IsConnectedWithCould()
	if err != nil {
		return &ProtectedExecutor{
			underlying:             exec,
			initConnectionBlockage: true,
		}
	}
	noop := func() {}
	authChecker.AsyncSupervisor(noop)

	return &ProtectedExecutor{
		underlying:  exec,
		authChecker: authChecker,
	}
}

// Execute provides executor functionality only with active Cloud connection.
func (e *ProtectedExecutor) Execute(ctx context.Context, input executor.ExecuteInput) (executor.ExecuteOutput, error) {
	if e.isBlocked() {
		return executor.ExecuteOutput{
			Message: unauthorizedMessage(e.openSourceBlockage),
		}, nil
	}

	return e.underlying.Execute(ctx, input)
}

// Metadata returns plugin metadata even without active Cloud connection.
func (e *ProtectedExecutor) Metadata(ctx context.Context) (api.MetadataOutput, error) {
	return e.underlying.Metadata(ctx)
}

// Help provides help functionality only with active Cloud connection.
func (e *ProtectedExecutor) Help(ctx context.Context) (api.Message, error) {
	if e.isBlocked() {
		return unauthorizedMessage(e.openSourceBlockage), nil
	}
	return e.underlying.Help(ctx)
}

func (e *ProtectedExecutor) isBlocked() bool {
	return e.openSourceBlockage || e.initConnectionBlockage || e.authChecker.ReachedPermanentBlockage()
}
