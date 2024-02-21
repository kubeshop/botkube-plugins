package remote

import (
	"net/http"
	"time"

	"github.com/hasura/go-graphql-client"
)

const (
	defaultTimeout = 30 * time.Second
	//nolint:gosec // warns us about 'Potential hardcoded credentials' but there is no security issue here
	apiKeyHeaderName = "X-API-Key"
)

// Gql defines GraphQL client data structure.
type Gql struct {
	client   *graphql.Client
	deployID string
}

// NewDefaultGqlClient initializes GraphQL client with default options.
func NewDefaultGqlClient(remoteCfg Config) *Gql {
	httpCli := &http.Client{
		Transport: newAPIKeySecuredTransport(remoteCfg.APIKey),
		Timeout:   defaultTimeout,
	}

	return &Gql{
		client:   graphql.NewClient(remoteCfg.Endpoint, httpCli),
		deployID: remoteCfg.Identifier,
	}
}

// DeploymentID returns deployment ID.
func (g *Gql) DeploymentID() string {
	return g.deployID
}

// Client returns GraphQL client.
func (g *Gql) Client() *graphql.Client {
	return g.client
}

type apiKeySecuredTransport struct {
	apiKey    string
	transport *http.Transport
}

func newAPIKeySecuredTransport(apiKey string) *apiKeySecuredTransport {
	return &apiKeySecuredTransport{
		apiKey:    apiKey,
		transport: http.DefaultTransport.(*http.Transport).Clone(),
	}
}

// RoundTrip adds API key to request header and executes RoundTrip for the underlying transport.
func (t *apiKeySecuredTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.apiKey != "" {
		req.Header.Set(apiKeyHeaderName, t.apiKey)
	}
	return t.transport.RoundTrip(req)
}
