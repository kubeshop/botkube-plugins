package aibrain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/golden"
)

// TestConvertProperlyAIAnswer tests that we can convert proper AI answer to Slack or Teams message format.
//
// To update golden files run:
//
//	go  test ./internal/source/ai-brain/... -run=TestConvertProperlyAIAnswer -update
func TestConvertProperlyAIAnswer(t *testing.T) {
	// given
	md, err := os.ReadFile(filepath.Join("testdata", t.Name(), "aiAnswer.md"))
	require.NoError(t, err)

	// Slack
	out := msgAIAnswer("42.42", string(md), "")
	assertGolden(t, out[0].Sections[0].Base.Body.Plaintext, "slack.golden.md")

	// Teams
	out = msgAIAnswer("19:d25cbf7cbfa74d22b42a2918452e1153@thread.tacv2", string(md), "")
	assertGolden(t, out[0].BaseBody.Plaintext, "teams.golden.md")
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
			expected: "This is a longer sentence that needs to be trunca...",
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
			result := ellipticalTruncate(tc.input, maxPromptLen)

			// then
			assert.Equal(t, tc.expected, result)
		})
	}
}
