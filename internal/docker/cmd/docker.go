package main

import (
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/docker/factory/tools"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/viper"
)

var (
	DockerServerName    = viper.GetString("mcpServer.docker.serverName")
	DockerServerVersion = viper.GetString("mcpServer.docker.serverVersion")
)

func main() {

	log.Println("Starting docker server...")
	if err := StartDocker(); err != nil {
		panic(err)
	}

}

// StartDocker starts the Docker MCP server
func StartDocker() error {

	log.Println("Starting Docker MCP server...")

	// Init docker MCP server
	s := server.NewMCPServer(
		DockerServerName,
		DockerServerVersion,
		server.WithToolCapabilities(true),
	)

	// Create tool pool and auto register tools
	tools := tools.CreatToolPool()
	s.AddTools(tools...)

	// Start the stdio server
	// Start StreamableHTTP server
	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	return nil
}
