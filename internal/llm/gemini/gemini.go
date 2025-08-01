package gemini

import (
	"context"
	"fmt"

	"github.com/leebrouse/GoMcp/internal/common/model"
	"google.golang.org/genai"
)

type GeminiLLM struct {
	model.LlmModel
}

// NewGeminiLLM creates a new Gemini LLM
func NewGeminiLLM(apiKey string, modelName string) *GeminiLLM {
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
