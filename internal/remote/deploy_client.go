package remote

import (
	"context"
	"fmt"

	"github.com/avast/retry-go/v4"
	"github.com/hasura/go-graphql-client"
)

const maxRetries = 5

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
func NewDeploymentClient(cfg Config) *DeploymentClient {
	return &DeploymentClient{client: NewDefaultGqlClient(cfg)}
}

// IsConnectedWithCould returns whether connected to Botkube Cloud
func (d *DeploymentClient) IsConnectedWithCould() error {
	var query struct {
		Deployment struct {
			ID string
		} `graphql:"deployment(id: $id)"`
	}
	deployID := d.client.DeploymentID()
	variables := map[string]interface{}{
		"id": graphql.ID(deployID),
	}
	err := d.withRetries(func() error {
		return d.client.Client().Query(context.Background(), &query, variables)
	})

	if err != nil {
		return err
	}

	if query.Deployment.ID != deployID {
		return fmt.Errorf("instance with id %q is not recognized by Cloud", deployID)
	}

	return nil
}

// GetConfigOutput returns deployment with Botkube configuration.
type GetConfigOutput struct {
	ResourceVersion int
	YAMLConfig      string
}

// GetConfig retrieves deployment configuration.
func (d *DeploymentClient) GetConfig(ctx context.Context) (GetConfigOutput, error) {
	var query struct {
		Deployment GetConfigOutput `graphql:"deployment(id: $id)"`
	}
	deployID := d.client.DeploymentID()
	variables := map[string]interface{}{
		"id": graphql.ID(deployID),
	}
	err := d.client.Client().Query(ctx, &query, variables)
	if err != nil {
		return GetConfigOutput{}, fmt.Errorf("while getting config with resource version for %q: %w", deployID, err)
	}
	return query.Deployment, nil
}

func (d *DeploymentClient) withRetries(fn func() error) error {
	return retry.Do(
		func() error {
			return fn()
		},
		retry.Attempts(maxRetries),
	)
}
