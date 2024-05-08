package aibrain

import (
	"testing"

	"github.com/kubeshop/botkube/pkg/loggerx"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
)

func TestTrimCitations(t *testing.T) {
	// given
	in := &openai.MessageText{
		Value: `Botkube supports a variety of source plugins, which are essentially event sources that can be integrated into Botkube. These plugins allow Botkube to receive events or notifications from different tools and services, which can then be sent to communication channels based on the configured source bindings. Here are some of the source plugins available for Botkube:

1. **Kubernetes native events**: This has been a part of Botkube since its inception, allowing it to listen for and react to events occurring within a Kubernetes cluster.
2. **Helm**: As an executor plugin, Helm can also be configured to send notifications about Helm chart deployments.
3. **Prometheus**: Planned as a new source plugin, allowing Botkube to receive alerts from Prometheus.
4. **Argo CD**: This plugin enables Botkube to interact with Argo CD, providing notifications and allowing commands to be executed within Botkube.
5. **GitHub Events**: This plugin streams events from selected GitHub repositories to your communication platform via Botkube.
6. **Keptn**: Integrates with Keptn to orchestrate load tests and validate results, sending notifications to communication platforms.

These plugins enhance Botkube's functionality, making it a versatile tool for monitoring and interacting with various systems directly from your communication platform【4:0†source】.`,

		Annotations: []interface{}{
			map[string]interface{}{
				"end_index": 1345.0,
				"file_citation": map[string]interface{}{
					"file_id": "file-Gfa1ahzAvp7UB1hToG3Udzj0",
				},
				"start_index": 1333.0,
				"text":        "【4:0†source】",
				"type":        "file_citation"},
		},
	}
	expected := `Botkube supports a variety of source plugins, which are essentially event sources that can be integrated into Botkube. These plugins allow Botkube to receive events or notifications from different tools and services, which can then be sent to communication channels based on the configured source bindings. Here are some of the source plugins available for Botkube:

1. **Kubernetes native events**: This has been a part of Botkube since its inception, allowing it to listen for and react to events occurring within a Kubernetes cluster.
2. **Helm**: As an executor plugin, Helm can also be configured to send notifications about Helm chart deployments.
3. **Prometheus**: Planned as a new source plugin, allowing Botkube to receive alerts from Prometheus.
4. **Argo CD**: This plugin enables Botkube to interact with Argo CD, providing notifications and allowing commands to be executed within Botkube.
5. **GitHub Events**: This plugin streams events from selected GitHub repositories to your communication platform via Botkube.
6. **Keptn**: Integrates with Keptn to orchestrate load tests and validate results, sending notifications to communication platforms.

These plugins enhance Botkube's functionality, making it a versatile tool for monitoring and interacting with various systems directly from your communication platform.`

	a := &assistant{}

	// when
	out := a.trimCitationsIfPresent(loggerx.NewNoop(), in)

	// then

	assert.Equal(t, expected, out)
}
