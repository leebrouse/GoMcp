package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/common/llm"
	"github.com/leebrouse/GoMcp/internal/common/model"
	"google.golang.org/genai"
)

type GeminiLLM struct {
	model.LlmModel
}

// NewGeminiLLM creates a new Gemini LLM
func NewGeminiLLM(apiKey string, modelName string, embedder string) llm.LLM {
	return &GeminiLLM{
		LlmModel: model.LlmModel{
			ApiKey:   apiKey,
			Model:    modelName,
			Embedder: embedder,
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

// Embeding prompt form string to byte for the RAG system
// Tips: Role means——
func (llm *GeminiLLM) Embeding(ctx context.Context, prompt string, role genai.Role) (string, error) {
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Tips (EN): Role denotes the speaker/identity of this content, e.g., RoleUser (user), RoleModel (model).
	// In generation tasks, it affects instruction priority and safety policies; in embedding tasks, it's mostly metadata and typically does not affect vector results.
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, role),
	}

	result, err := client.Models.EmbedContent(ctx,
		llm.Embedder,
		contents,
		nil,
	)
	if err != nil {
		return "", err
	}

	embeddings, err := json.MarshalIndent(result.Embeddings, "", "  ")
	if err != nil {
		return "", err
	}

	return string(embeddings), nil
}
