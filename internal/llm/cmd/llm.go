package main

import (
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/llm/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/viper"
)

var (
	LLMServerName    = viper.GetString("mcpServer.llm.serverName")
	LLMServerVersion = viper.GetString("mcpServer.llm.serverVersion")
)

func main() {
	log.Println("Starting llm server...")

	StartLLM()

}

// StartLLM starts the LLM MCP server
func StartLLM() {

	log.Println("Starting MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		LLMServerName,
		LLMServerVersion,
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
