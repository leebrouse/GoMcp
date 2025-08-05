package main

import (
	"log"

	"github.com/leebrouse/GoMcp/internal/llm/factory/tools"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/viper"
)

var (
	LLMServerName    = viper.GetString("mcpServer.llm.serverName")
	LLMServerVersion = viper.GetString("mcpServer.llm.serverVersion")
)

func main() {
	log.Println("Starting LLM server...")
	if err := StartLLM(); err != nil {
		panic(err)
	}
}

func StartLLM() error {
	log.Println("Creating MCP server...")

	// Init llm MCP server
	s := server.NewMCPServer(
		LLMServerName,
		LLMServerVersion,
		server.WithToolCapabilities(true),
	)

	// Create tool pool and auto register tools
	tools := tools.CreatToolPool()
	s.AddTools(tools...)

	// Start llm MCP server
	if err := server.ServeStdio(s); err != nil {
		return err
	}

	return nil
}
