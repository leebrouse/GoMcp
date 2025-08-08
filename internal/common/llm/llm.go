package llm

import (
	"context"

	"google.golang.org/genai"
)

// llm interface
type LLM interface {
	// GenerateText generates text using the LLM
	GenerateText(ctx context.Context, prompt string) (string, error)

	Embeding(ctx context.Context, prompt string, role genai.Role) (string, error)
	// To do more functions like image generation, etc....
}
