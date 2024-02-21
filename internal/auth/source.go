package auth

import (
	"context"
	"log"

	"github.com/kubeshop/botkube-cloud-plugins/internal/remote"

	"github.com/kubeshop/botkube/pkg/api"
	"github.com/kubeshop/botkube/pkg/api/source"
)

// ProtectedSource protects source usage without active cloud connection.
type ProtectedSource struct {
	underlying             source.Source
	openSourceBlockage     bool
	initConnectionBlockage bool
}

// NewProtectedSource returns wrapped Source instance with Cloud conn checker.
func NewProtectedSource(source source.Source) source.Source {
	cfg, ok := remote.GetConfig()
	if !ok {
		return &ProtectedSource{
			underlying:         source,
			openSourceBlockage: true,
		}
	}
	authChecker := remote.NewConnChecker(cfg)
	err := authChecker.IsConnectedWithCould()
	if err != nil {
		return &ProtectedSource{
			underlying:             source,
			initConnectionBlockage: true,
		}
	}

	onPermanentBlockage := func() {
		// we may already start multiple streams (`Stream()`),
		// so we want to kill the plugin process to stop all of them.
		// Once restarted, and we will be still blocked, a proper message will be sent to all configured channels.
		log.Fatal("Failed to connect with Cloud")
	}
	authChecker.AsyncSupervisor(onPermanentBlockage)
	return &ProtectedSource{
		underlying: source,
	}
}

// Stream provides stream functionality only with active cloud connection.
func (s *ProtectedSource) Stream(ctx context.Context, in source.StreamInput) (source.StreamOutput, error) {
	if s.isBlocked() {
		out := source.StreamOutput{
			Event: make(chan source.Event, 1),
		}
		out.Event <- source.Event{
			Message: unauthorizedMessage(s.openSourceBlockage),
		}

		return out, nil
	}

	return s.underlying.Stream(ctx, in)
}

// Metadata returns plugin metadata even without active cloud connection.
func (s *ProtectedSource) Metadata(ctx context.Context) (api.MetadataOutput, error) {
	return s.underlying.Metadata(ctx)
}

// HandleExternalRequest provides external request functionality only with active cloud connection.
func (s *ProtectedSource) HandleExternalRequest(ctx context.Context, in source.ExternalRequestInput) (source.ExternalRequestOutput, error) {
	if s.isBlocked() {
		out := source.ExternalRequestOutput{
			Event: source.Event{
				Message: unauthorizedMessage(s.openSourceBlockage),
			},
		}
		return out, nil
	}

	return s.underlying.HandleExternalRequest(ctx, in)
}

func (s *ProtectedSource) isBlocked() bool {
	return s.openSourceBlockage || s.initConnectionBlockage
}
