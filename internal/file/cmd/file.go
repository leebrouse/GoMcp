package main

import (
	"fmt"
	"log"

	mcpServer "github.com/leebrouse/GoMcp/internal/common/config/server"
	"github.com/leebrouse/GoMcp/internal/file/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	log.Println("Starting file server...")
	StartFile()

}

func StartFile() {

	log.Println("Starting MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		mcpServer.FileServerName,
		mcpServer.FileServerVersion,
		server.WithToolCapabilities(true),
	)

	list := mcp.NewTool("list",
		mcp.WithDescription("展示当前目录下的文件"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the directory to list"),
		),
	)

	log.Println("Adding tool handlers...")
	s.AddTool(list, handler.ListHandler)

	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
