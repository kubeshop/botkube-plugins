package main

import (
	"context"
	"log"
	"os"
	"time"

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
	instructions := `You are an experienced DevOps engineer. You have deep
	understand how to operate a kubernetes cluster and troubleshoot running
	workloads in kubernetes. You have access to tools which can help you. Keep
	your answers short and on the subject. Do not get involved in unrelated
	topics.`

	return c.ModifyAssistant(ctx, id, openai.AssistantRequest{
		Model:        openai.GPT4TurboPreview,
		Instructions: &instructions,
		Tools:        openAITools,
	})
}
