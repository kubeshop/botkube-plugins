package aibrain

import (
	"os"
	"path/filepath"
	"testing"

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
	out := msgAIAnswer("42.42", string(md))
	assertGolden(t, out.Sections[0].Base.Body.Plaintext, "slack.golden.md")

	// Teams
	out = msgAIAnswer("19:d25cbf7cbfa74d22b42a2918452e1153@thread.tacv2", string(md))
	assertGolden(t, out.BaseBody.Plaintext, "teams.golden.md")
}

func assertGolden(t *testing.T, actual, goldenPath string) {
	golden.Assert(t, actual, filepath.Join(t.Name(), goldenPath))
}
