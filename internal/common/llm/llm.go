package llm

import (
	"context"

	"google.golang.org/genai"
)

// llm interface
type LLM interface {
	// GenerateText generates text using the LLM
	GenerateText(ctx context.Context, prompt string) (string, error)

	// Read local PDF and do operation
	ReadDocument(ctx context.Context, path string, prompt string) (string, error)

	// Embeding text to vector val
	Embeding(ctx context.Context, prompt string, role genai.Role) ([]float32, error)

	// To do more functions like image generation, etc....
}
