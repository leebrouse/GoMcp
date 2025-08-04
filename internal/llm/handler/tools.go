package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/leebrouse/GoMcp/internal/common/llm/gemini"
	"github.com/leebrouse/GoMcp/utils/custom"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/spf13/viper"
)

var (
	GeminiApiKey = viper.GetString("llm.gemini.apikey")
	GeminiModel  = viper.GetString("llm.gemini.model")
)

// chatbox handler
func ChatboxHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	prompt, ok := arguments["prompt"].(string)
	if !ok {
		return custom.NewTextResult("Error: prompt parameter is required and must be a string", true), nil
	}

	// TODO: Should read from the config file or the request
	llm := gemini.NewGeminiLLM(GeminiApiKey, GeminiModel)

	response, err := llm.GenerateText(ctx, prompt)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	return custom.NewTextResult(response, false), nil
}

// code review handler
func CodeReviewHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	path, ok := arguments["path"].(string)
	if !ok {
		return custom.NewTextResult("Error: path parameter is required and must be a string", true), nil
	}

	code, err := os.ReadFile(path)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	// init llm
	llm := gemini.NewGeminiLLM(GeminiApiKey, GeminiModel)

	prompt := fmt.Sprintf("Review the following code and provide a list of issues and suggestions for improvement. Return the results in a JSON object with the following fields: 'issues', 'suggestions', 'score'. The score should be a number between 0 and 100. The issues and suggestions should be an array of strings. The code is: %s", string(code))
	response, err := llm.GenerateText(ctx, prompt)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	return custom.NewTextResult(response, false), nil
}
