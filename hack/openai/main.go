package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/MakeNowJust/heredoc"
	openai "github.com/sashabaranov/go-openai"
)

const (
	prodAssistantID = "asst_eMM9QaWLi6cajHE4PdG1yU53"
	devAssistantID  = "asst_ejVrAgjhhvCw6jGFYq5JyBqj"
)

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	env := flag.String("env", "dev", "Environment to update. Allow values: 'dev' or 'prod'.")
	flag.Parse()

	var assistantID string
	switch *env {
	case "dev":
		assistantID = devAssistantID
	case "prod":
		assistantID = prodAssistantID
	default:
		log.Fatalf("Invalid environment: %s", *env)
	}

	// Update assistant with latest tools.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("Updating %s assistant with ID %s", *env, assistantID)
	_, err := updateAssistant(ctx, client, assistantID)
	if err != nil {
		log.Fatal(err)
	}
}

func updateAssistant(ctx context.Context, c *openai.Client, id string) (openai.Assistant, error) {
	instructions := heredoc.Doc(`
Your role involves deeply understanding how to operate and troubleshoot Kubernetes clusters and their workloads. 
You possess extensive expertise in Kubernetes and cloud-native networking. 
Take kubernetes cluster and namespace configuration, such as security policies, events, during troubleshooting.
Employ a Chain of Thought (CoT) process for diagnosing and resolving cluster issues. 
Ensure your explanations are simplified for clarity to non-technical users. 
Utilize available tools for diagnosing the specific cluster in question. Focus your responses to be concise and relevant.
Keep responses concise, within 2000 characters. Provide extended answers only upon request.

Make sure you fetch Botkube Agent configuration to answer questions about Botkube or channel configuration.
`)

	return c.ModifyAssistant(ctx, id, openai.AssistantRequest{
		Model:        openai.GPT4Turbo,
		Instructions: &instructions,
		Tools:        openAITools,
	})
}
