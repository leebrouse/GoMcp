package llm

import "context"

// llm interface
type LLM interface {
	// GenerateText generates text using the LLM
	GenerateText(ctx context.Context, prompt string) (string, error)

	// To do more functions like image generation, etc....
}
