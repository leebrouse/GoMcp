package server

import (
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/server/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	serverName    = "mcp-server"
	serverVersion = "1.0.0"
)

func Start() {

	log.Println("Starting MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		serverName,
		serverVersion,
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

	list := mcp.NewTool("list",
		mcp.WithDescription("展示当前目录下的文件"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the directory to list"),
		),
	)

	codeReview := mcp.NewTool("codeReview",
		mcp.WithDescription("Review the code and provide a list of issues and suggestions for improvement"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the code to review"),
		),
	)

	log.Println("Adding tool handlers...")
	// Add tool handler
	s.AddTool(chatbox, handler.ChatboxHandler)
	s.AddTool(list, handler.ListHandler)
	s.AddTool(codeReview, handler.CodeReviewHandler)

	// Start the stdio server
	// Start StreamableHTTP server
	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
