package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/leebrouse/GoMcp/internal/common/llm/gemini"
	"github.com/leebrouse/GoMcp/internal/llm/infrastructure/search"
	"github.com/leebrouse/GoMcp/utils/custom"
	"github.com/leebrouse/GoMcp/utils/helper"
	"github.com/mark3labs/mcp-go/mcp"
)

// private config
// Tips: The config should be set in the global.yaml not in the code, TEST to resolve it yet
const (
	GeminiApiKey        = "AIzaSyCKURVV8jEX3CsRu_4pysxmJm3IH4mr8VU"
	GeminiModel         = "gemini-2.0-flash"
	GeminiEmbedder      = "gemini-embedding-001"
	DuckDuckGoUserAgent = "LangChainGo/1.0"
)

// chatbox handler
func ChatboxHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	prompt, ok := arguments["prompt"].(string)
	if !ok {
		return custom.NewTextResult("Error: prompt parameter is required and must be a string", true), nil
	}

	// TODO: Should read from the config file or the request
	llm := gemini.NewGeminiLLM(GeminiApiKey, GeminiModel, GeminiEmbedder)

	response, err := llm.GenerateText(ctx, prompt)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	return custom.NewTextResult(response, false), nil
}

// Enhancing "chatbox handler" because it can geting search information from the DuckDuckGo
func DuckDuckGoChatboxHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	keyword, ok := arguments["keyword"].(string)
	if !ok {
		return custom.NewTextResult("Error: keyword parameter is required and must be a string", true), nil
	}

	// Service: To do research process

	// create DuckDuckGo client
	duckduckgo := search.NewDuckDuckGo(5, "LangChainGo/1.0")

	// search information from the DuckDuckGo and get reference
	reference, err := duckduckgo.DuckDuckGoSearch(ctx, keyword)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	// create gemini client
	llm := gemini.NewGeminiLLM(GeminiApiKey, GeminiModel, GeminiEmbedder)

	// creata a new prompt attached to the DuckDuckGo reference
	prompt, err := helper.ParseMarkDown("internal/llm/factory/prompt/duckduckgo_en.md", reference)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error to get new prompt: %v", err), true), nil
	}

	// prompt := fmt.Sprintf(
	// 	"你是一名专业的研究助理。\n\n"+
	// 		"下面这段文字是 DuckDuckGo 返回的原始参考内容，请仅基于它进行回答，不要添加任何超出原文的信息。\n"+
	// 		"如果原文无法回答用户问题，请直接说明“资料不足”。\n\n"+
	// 		"参考内容：\n%s\n\n"+
	// 		"请用简洁、连贯的中文总结上述内容，并指出其中最重要的 3 个要点。",
	// 	reference,
	// )
	// generate a text by using the gemini api

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
	llm := gemini.NewGeminiLLM(GeminiApiKey, GeminiModel, GeminiEmbedder)

	prompt := fmt.Sprintf("Review the following code and provide a list of issues and suggestions for improvement. Return the results in a JSON object with the following fields: 'issues', 'suggestions', 'score'. The score should be a number between 0 and 100. The issues and suggestions should be an array of strings. The code is: %s", string(code))
	response, err := llm.GenerateText(ctx, prompt)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	return custom.NewTextResult(response, false), nil
}

// Read document handler
func ReadDocumentHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	//get path
	path, ok := arguments["path"].(string)
	if !ok {
		return custom.NewTextResult("Error: path parameter is required and must be a string", true), nil
	}

	//get prompt
	prompt, ok := arguments["prompt"].(string)
	if !ok {
		return custom.NewTextResult("Error: prompt parameter is required and must be a string", true), nil
	}

	// init llm
	llm := gemini.NewGeminiLLM(GeminiApiKey, GeminiModel, GeminiEmbedder)

	// call ReadDocument function by interface
	response, err := llm.ReadDocument(ctx, path, prompt)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	return custom.NewTextResult(response, false), nil
}

//Todo: embad prompt for to vector search from the TiDB
