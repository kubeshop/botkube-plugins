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

// Deployment returns deployment with Botkube configuration.
type Deployment struct {
	ID string
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

func (d *DeploymentClient) withRetries(fn func() error) error {
	return retry.Do(
		func() error {
			return fn()
		},
		retry.Attempts(maxRetries),
	)
}
