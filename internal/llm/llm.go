package llm

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// llm interface
type LLM interface {
	// GenerateText generates text using the LLM
	GenerateText(ctx context.Context, prompt string) (string, error)

	// GenerateWithTools generates text using the LLM with tools
	GenerateWithTools(ctx context.Context, prompt string, tools []mcp.Tool) (string, error)
	// To do more functions like image generation, etc....
}
