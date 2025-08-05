package main

import (
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/file/factory/tools"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/viper"
)

var (
	FileServerName    = viper.GetString("mcpServer.file.serverName")
	FileServerVersion = viper.GetString("mcpServer.file.serverVersion")
)

func main() {
	log.Println("Starting file server...")
	if err := StartFile(); err != nil {
		panic(err)
	}
}

func StartFile() error {

	log.Println("Starting MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		FileServerName,
		FileServerVersion,
		server.WithToolCapabilities(true),
	)

	// Create tool pool and auto register tools
	tools := tools.CreatToolPool()
	s.AddTools(tools...)

	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	return nil
}
