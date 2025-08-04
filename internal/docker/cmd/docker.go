package main

import (
	"fmt"
	"log"

	mcpServer "github.com/leebrouse/GoMcp/internal/common/config/server"
	"github.com/leebrouse/GoMcp/internal/docker/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {

	log.Println("Starting docker server...")
	StartDocker()

}

func StartDocker() {

	log.Println("Starting Docker MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		mcpServer.DockerServerName,
		mcpServer.DockerServerVersion,
		server.WithToolCapabilities(true),
	)

	// Define a simple tool
	log.Println("Adding docker tool...")
	docker := mcp.NewTool("docker",
		mcp.WithDescription("Run a docker image"),
		mcp.WithString("image",
			mcp.Required(),
			mcp.Description("Image to run"),
		),
	)

	log.Println("Adding tool handlers...")
	s.AddTool(docker, handler.DockerHandler)

	// Start the stdio server
	// Start StreamableHTTP server
	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
