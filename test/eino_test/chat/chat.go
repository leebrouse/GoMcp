package chat

import (
	"context"
	"log"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

// Generate generates a message from a chat model
func Generate(ctx context.Context, llm model.ToolCallingChatModel, in []*schema.Message) *schema.Message {
	result, err := llm.Generate(ctx, in)
	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}

// Stream streams a message from a chat model
func Stream(ctx context.Context, llm model.ToolCallingChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message] {
	result, err := llm.Stream(ctx, in)
	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}
