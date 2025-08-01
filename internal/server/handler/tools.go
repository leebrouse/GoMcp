package handler

import (
	"context"
	"fmt"

	"github.com/leebrouse/GoMcp/internal/llm/gemini"
	"github.com/mark3labs/mcp-go/mcp"
)

// TODO: Should read from the config file or the request
const (
	geminiApiKey = "AIzaSyCKURVV8jEX3CsRu_4pysxmJm3IH4mr8VU"
	geminiModel  = "gemini-2.5-flash"
)

func HelloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	name, ok := arguments["name"].(string)
	if !ok {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Error: name parameter is required and must be a string",
				},
			},
			IsError: true,
		}, nil
	}

	llm := gemini.NewGeminiLLM(geminiApiKey, geminiModel)

	response, err := llm.GenerateText(ctx, name)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error: %v", err),
				},
			},
			IsError: true,	
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: response,
			},
		},
	}, nil
}
