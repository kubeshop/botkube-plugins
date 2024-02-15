package remote

import (
	"context"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/hasura/go-graphql-client"
)

// GraphQLClient defines GraphQL client.
type GraphQLClient interface {
	Client() *graphql.Client
	DeploymentID() string
}

// DeploymentClient defines GraphQL client for Deployment.
type DeploymentClient struct {
	client GraphQLClient
}

// NewDeploymentClient initializes GraphQL client.
func NewDeploymentClient(client GraphQLClient) *DeploymentClient {
	return &DeploymentClient{client: client}
}

// Deployment returns deployment with Botkube configuration.
type Deployment struct {
	ID string
}

// IsConnectedWithCould returns whether connected to Botkube Cloud
func (d *DeploymentClient) IsConnectedWithCould() bool {
	var query struct {
		Deployment struct {
			ID string
		} `graphql:"deployment(id: $id)"`
	}
	deployID := d.client.DeploymentID()
	variables := map[string]interface{}{
		"id": graphql.ID(deployID),
	}
	err := d.withRetries(context.Background(), 5, func() error {
		return d.client.Client().Query(context.Background(), &query, variables)
	})
	if err != nil {
		return false
	}

	return query.Deployment.ID == d.client.DeploymentID()
}

func (d *DeploymentClient) withRetries(ctx context.Context, maxRetries int, fn func() error) error {
	return retry.Do(
		func() error {
			return fn()
		},
		retry.DelayType(func(uint, error, *retry.Config) time.Duration {
			return 200 * time.Microsecond
		}),
		retry.Attempts(uint(maxRetries)), // infinite, we cancel that by our own
		retry.Context(ctx),
	)
}
