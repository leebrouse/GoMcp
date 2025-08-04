package main

import (
	"fmt"
	"log"

	mcpServer "github.com/leebrouse/GoMcp/internal/common/config/server"
	"github.com/leebrouse/GoMcp/internal/llm/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	log.Println("Starting llm server...")

	StartLLM()

}

func StartLLM() {

	log.Println("Starting MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		mcpServer.LLMServerName,
		mcpServer.LLMServerVersion,
		server.WithToolCapabilities(true),
	)

	// Define a simple tool
	log.Println("Adding chatbox tool...")
	chatbox := mcp.NewTool("chatbox",
		mcp.WithDescription("Send a prompt to the LLM"),
		mcp.WithString("prompt",
			mcp.Required(),
			mcp.Description("Prompt to send to the LLM"),
		),
	)


	codeReview := mcp.NewTool("codeReview",
		mcp.WithDescription("Review the code and provide a list of issues and suggestions for improvement"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the code to review"),
		),
	)


	// Add tool handler
	s.AddTool(chatbox, handler.ChatboxHandler)
	s.AddTool(codeReview, handler.CodeReviewHandler)


	// Start the stdio server
	// Start StreamableHTTP server
	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
