package audit

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/kubeshop/botkube/internal/graphql"
)

// AuditReporter defines interface for reporting audit events
type AuditReporter interface {
	ReportExecutorAuditEvent(ctx context.Context, e ExecutorAuditEvent) error
	ReportSourceAuditEvent(ctx context.Context, e SourceAuditEvent) error
}

// ExecutorAuditEvent contains audit event data
type ExecutorAuditEvent struct {
	CreatedAt    string
	PluginName   string
	PlatformUser string
	BotPlatform  BotPlatform
	Command      string
	Channel      string
}

// SourceAuditEvent contains audit event data
type SourceAuditEvent struct {
	CreatedAt  string
	PluginName string
	Event      string
	Bindings   []string
}

// NewAuditReporter creates new AuditReporter
func NewAuditReporter(logger logrus.FieldLogger, gql *graphql.Gql) AuditReporter {
	if _, provided := os.LookupEnv(graphql.GqlProviderIdentifierEnvKey); provided {
		return newGraphQLAuditReporter(logger.WithField("component", "GraphQLAuditReporter"), gql)
	}
	return newNoopAuditReporter(nil)
}
