package chat

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/gemini"
	"github.com/cloudwego/eino/components/model"
	"google.golang.org/genai"
)

// CreateGeminiChatModel creates a Gemini chat model
func CreateGeminiChatModel(ctx context.Context) model.ToolCallingChatModel {
	//1. init Gemini client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GOOGLE_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("create gemini client failed, err=%v", err)
	}

	//2. create chat model
	chatModel, err := gemini.NewChatModel(ctx, &gemini.Config{
		Client: client,
		Model:  "gemini-2.0-flash",
	})
	if err != nil {
		log.Fatalf("create gemini chat model failed, err=%v", err)
	}

	//3. return chat model
	return chatModel
}
