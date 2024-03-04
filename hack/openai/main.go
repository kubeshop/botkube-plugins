package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MakeNowJust/heredoc"
	openai "github.com/sashabaranov/go-openai"
)

const (
	assistantID = "asst_eMM9QaWLi6cajHE4PdG1yU53" // Botkube
)

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// Update assistant with latest tools.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := updateAssistant(ctx, client, assistantID)
	if err != nil {
		log.Fatal(err)
	}
}

func updateAssistant(ctx context.Context, c *openai.Client, id string) (openai.Assistant, error) {
	instructions := heredoc.Doc(`
			Your role involves deeply understanding how to operate and troubleshoot Kubernetes clusters and their workloads. 
			You possess extensive expertise in Kubernetes and cloud-native networking. 
			Employ a Chain of Thought (CoT) process for diagnosing and resolving cluster issues. 
			Ensure your explanations are simplified for clarity to non-technical users. 
			Utilize available tools for diagnosing the specific cluster in question. Focus your responses to be concise and relevant.`)

	return c.ModifyAssistant(ctx, id, openai.AssistantRequest{
		Model:        openai.GPT4TurboPreview,
		Instructions: &instructions,
		Tools:        openAITools,
	})
}