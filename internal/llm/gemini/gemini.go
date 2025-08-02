package gemini

import (
	"context"
	"fmt"

	"github.com/leebrouse/GoMcp/internal/common/model"
	"github.com/leebrouse/GoMcp/internal/llm"
	"github.com/mark3labs/mcp-go/mcp"
	"google.golang.org/genai"
)

type GeminiLLM struct {
	model.LlmModel
}

// NewGeminiLLM creates a new Gemini LLM
func NewGeminiLLM(apiKey string, modelName string) llm.LLM {
	return &GeminiLLM{
		LlmModel: model.LlmModel{
			ApiKey: apiKey,
			Model:  modelName,
		},
	}
}

// GenerateText generates text using the Gemini model
func (llm *GeminiLLM) GenerateText(ctx context.Context, prompt string) (string, error) {

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  llm.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return fmt.Sprintf("Error initializing client: %v", err), err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		llm.Model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return fmt.Sprintf("Error generating content: %v", err), err
	}

	return result.Text(), nil

}

// GenerateWithTools generates result using the Gemini model with tools by the rest api
func (llm *GeminiLLM) GenerateWithTools(ctx context.Context, prompt string, tools []mcp.Tool) (string, error) {
	// client, err := genai.NewClient(ctx, &genai.ClientConfig{
	// 	APIKey:  llm.ApiKey,
	// 	Backend: genai.BackendGeminiAPI,
	// })
	// if err != nil {
	// 	return fmt.Sprintf("Error initializing client: %v", err), err
	// }

	// result, err := client.Models.(
	// 	ctx,
	// 	llm.Model,
	// 	genai.Text(prompt),
	// 	nil,
	// )
	// if err != nil {
	// 	return fmt.Sprintf("Error generating content: %v", err), err
	// }

	// return result.Text(), nil
	return "", nil
}
