package server

import (
	"fmt"

	"github.com/leebrouse/GoMcp/internal/server/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Start() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Hello World Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Define a simple tool
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	// Add tool handler
	s.AddTool(tool, handler.HelloHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
