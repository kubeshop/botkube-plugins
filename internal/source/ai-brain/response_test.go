package aibrain

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"
)

// TestConvertProperlyAIAnswer tests that we can convert proper AI answer to Slack or Teams message format.
//
// To update golden files run:
//
//	go test ./internal/source/ai-brain/... -run=TestConvertProperlyAIAnswer -update
func TestConvertProperlyAIAnswer(t *testing.T) {
	// given
	md, err := os.ReadFile(filepath.Join("testdata", t.Name(), "aiAnswer.md"))
	require.NoError(t, err)

	// Slack
	out := msgAIAnswer(openai.Run{}, &Payload{
		MessageID: "42.42",
		Prompt:    "This is a test",
	}, string(md), nil)
	assertGolden(t, out.Sections[0].Base.Body.Plaintext, "slack.golden.md")

	// Teams
	out = msgAIAnswer(openai.Run{}, &Payload{
		MessageID: "19:d25cbf7cbfa74d22b42a2918452e1153@thread.tacv2",
		Prompt:    "This is a test",
	}, string(md), nil)
	assertGolden(t, out.BaseBody.Plaintext, "teams.golden.md")
}

// TestConvertProperlyAIAnswerWithTools tests that we can convert proper AI answer to Slack or Teams message format.
//
// To update golden files run:
//
//	go test ./internal/source/ai-brain/... -run=TestConvertProperlyAIAnswerWithTools -update
func TestConvertProperlyAIAnswerWithTools(t *testing.T) {
	// given
	md, err := os.ReadFile(filepath.Join("testdata", t.Name(), "aiAnswer.md"))
	require.NoError(t, err)

	toolCalls := []openai.RunStep{
		{
			StepDetails: openai.StepDetails{
				Type: openai.RunStepTypeToolCalls,
				ToolCalls: []openai.ToolCall{
					{
						Type: openai.ToolType("file_search"),
					},
					{
						Type: openai.ToolTypeFunction,
						Function: openai.FunctionCall{
							Name:      "botkubeGetStartupAgentConfiguration",
							Arguments: "foo",
						},
					},
				},
			},
		},
		{
			StepDetails: openai.StepDetails{
				Type: openai.RunStepTypeToolCalls,
				ToolCalls: []openai.ToolCall{
					{
						Type: openai.ToolTypeFunction,
						Function: openai.FunctionCall{
							Name:      "botkubeGetAgentStatus",
							Arguments: "foo",
						},
					},
					{
						Type: openai.ToolTypeFunction,
						Function: openai.FunctionCall{
							Name:      "kubectlGetResource",
							Arguments: "foo",
						},
					},
					{
						Type: openai.ToolTypeFunction,
						Function: openai.FunctionCall{
							Name:      "kubectlLogs",
							Arguments: "foo",
						},
					},
				},
			},
		},
	}
	convertedToolCalls := getFriendlyToolCallsFromRunSteps(toolCalls)

	// Slack
	out := msgAIAnswer(openai.Run{}, &Payload{
		MessageID: "42.42",
		Prompt:    "This is a test",
	}, string(md), convertedToolCalls)

	outBytes, err := json.MarshalIndent(out, "", "  ")
	require.NoError(t, err)
	assertGolden(t, string(outBytes), "slack-tools.golden.json")

	// Teams
	out = msgAIAnswer(openai.Run{}, &Payload{
		MessageID: "19:d25cbf7cbfa74d22b42a2918452e1153@thread.tacv2",
		Prompt:    "This is a test",
	}, string(md), convertedToolCalls)
	outBytes, err = json.MarshalIndent(out, "", "  ")
	require.NoError(t, err)
	assertGolden(t, string(outBytes), "teams-tools.golden.json")

	// Discord and others
	out = msgAIAnswer(openai.Run{}, &Payload{
		MessageID: "",
		Prompt:    "This is a test",
	}, string(md), convertedToolCalls)
	outBytes, err = json.MarshalIndent(out, "", "  ")
	require.NoError(t, err)
	assertGolden(t, string(outBytes), "discord-tools.golden.json")
}

func assertGolden(t *testing.T, actual, goldenPath string) {
	golden.Assert(t, actual, filepath.Join(t.Name(), goldenPath))
}

func TestEllipticalTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Short sentence",
			input:    "This is a short sentence.",
			expected: "This is a short sentence.",
		},
		{
			name:     "Longer sentence",
			input:    "This is a longer sentence that needs to be truncated.",
			expected: "This is a longer sentence that needs to be truncat...",
		},
		{
			name:     "Exact 50 characters sentence",
			input:    "This sentence has exactly 50 characters used in it",
			expected: "This sentence has exactly 50 characters used in it",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// when
			result := ellipticalTruncate(tc.input, 50)

			// then
			assert.Equal(t, tc.expected, result)
		})
	}
}
