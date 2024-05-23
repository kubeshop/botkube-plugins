package aibrain

import (
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
	"golang.org/x/exp/slices"
)

// https://platform.openai.com/docs/api-reference/runs/object#runs/object-tools
var userFriendlyToolNames = map[string]string{
	"file_search": "Botkube content search", // for now it is only that, later when we will have the "Bring your own docs" functionality, we should rename it

	"function/botkubeGetStartupAgentConfiguration": "Botkube agent configuration",
	"function/botkubeGetAgentStatus":               "Botkube agent status",
	"function/kubectlGetResource":                  "kubectl get",
	"function/kubectlGetEvents":                    "kubectl get events",
	"function/kubectlDescribeResource":             "kubectl describe",
	"function/kubectlTopPods":                      "kubectl top pods",
	"function/kubectlTopNodes":                     "kubectl top nodes",
	"function/kubectlLogs":                         "kubectl logs",

	"code_interpreter": "Code Interpreter", // we don't use this one, at least yet
}

func getFriendlyToolCallsFromRunSteps(runSteps []openai.RunStep) map[string]struct{} {
	toolCalls := make(map[string]struct{})
	for _, step := range runSteps {
		if step.StepDetails.Type != openai.RunStepTypeToolCalls {
			continue
		}

		for _, t := range step.StepDetails.ToolCalls {
			key := getUserFriendlyToolNameKey(t)
			outName, ok := userFriendlyToolNames[key]
			if !ok {
				continue
			}

			toolCalls[outName] = struct{}{}
		}
	}

	return toolCalls
}

func getUserFriendlyToolNameKey(toolCall openai.ToolCall) string {
	var keySuffix string
	if toolCall.Function.Name != "" {
		keySuffix = fmt.Sprintf("/%s", toolCall.Function.Name)
	}

	key := fmt.Sprintf("%s%s", toolCall.Type, keySuffix)
	return key
}

func printUsedTools(usedTools map[string]struct{}) string {
	if len(usedTools) == 0 {
		return ""
	}

	var outputToolNames []string
	for name := range usedTools {
		outputToolNames = append(outputToolNames, name)
	}
	slices.Sort(outputToolNames)

	return fmt.Sprintf("ℹ️ **Sources used:** %s", strings.Join(outputToolNames, " | "))
}
