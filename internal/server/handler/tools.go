package handler

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/leebrouse/GoMcp/internal/llm/gemini"
	"github.com/mark3labs/mcp-go/mcp"
)

// TODO: Should read from the config file or the request
const (
	geminiApiKey = "AIzaSyCKURVV8jEX3CsRu_4pysxmJm3IH4mr8VU"
	geminiModel  = "gemini-2.5-flash"
)

// chatbox handler
func ChatboxHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	prompt, ok := arguments["prompt"].(string)
	if !ok {
		return newTextResult("Error: prompt parameter is required and must be a string", true), nil
	}

	// TODO: Should read from the config file or the request
	llm := gemini.NewGeminiLLM(geminiApiKey, geminiModel)

	response, err := llm.GenerateText(ctx, prompt)
	if err != nil {
		return newTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	return newTextResult(response, false), nil
}

// list handler
func ListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	path, ok := arguments["path"].(string)
	if !ok {
		return newTextResult("Error: path parameter is required and must be a string", true), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return newTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return newTextResult("Files:\n"+strings.Join(fileNames, "\n"), false), nil
}

// Helper function to create a new text result
func newTextResult(text string, isError bool) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: text,
			},
		},
		IsError: isError,
	}
}
