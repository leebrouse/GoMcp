package main

import (
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/file/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/viper"
)

var (
	FileServerName    = viper.GetString("mcpServer.file.serverName")
	FileServerVersion = viper.GetString("mcpServer.file.serverVersion")
)

func main() {
	log.Println("Starting file server...")
	StartFile()

}

func StartFile() {

	log.Println("Starting MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		FileServerName,
		FileServerVersion,
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

	// Add tool handler
	// Tip should use tools
	s.AddTool(list, handler.ListHandler)

	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
