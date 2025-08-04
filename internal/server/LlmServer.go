package server

import (
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/server/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	LLMServerName    = "llm-server"
	LLMServerVersion = "1.0.0"
)

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

	// docUnderstanding := mcp.NewTool("docUnderstanding",
	// 	mcp.WithDescription("Send a prompt to the LLM"),
	// 	mcp.WithString("prompt",
	// 		mcp.Required(),
	// 		mcp.Description("Prompt to send to the LLM"),
	// 	),
	// 	mcp.WithString("fileUrl",
	// 		mcp.Required(),
	// 		mcp.Description("File URL to send to the LLM"),
	// 	),
	// )

	// Add tool handler
	s.AddTool(chatbox, handler.ChatboxHandler)
	// s.AddTool(docUnderstanding, handler.DocumentUnderstandingHandler)

	// Start the stdio server
	// Start StreamableHTTP server
	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
