package main

import (
	"log"

	"github.com/leebrouse/GoMcp/internal/llm/handler"
	"github.com/leebrouse/GoMcp/internal/common/factory/tools" // 注册所有工具
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

	s := server.NewMCPServer(
		LLMServerName,
		LLMServerVersion,
		server.WithToolCapabilities(true),
	)

	// 自动注册所有工具并添加对应处理器
	for _, tool := range tools.GetAllTools() {
		var handlerFunc server.ToolHandlerFunc

		switch tool.GetName() {
		case "chatbox":
			handlerFunc = handler.ChatboxHandler
		case "codeReview":
			handlerFunc = handler.CodeReviewHandler
		default:
			log.Printf("No handler found for tool: %s", tool.GetName())
			continue
		}

		s.AddTool(tool, handlerFunc)
		log.Printf("Tool registered: %s", tool.GetName())
	}

	// 启动服务
	if err := server.ServeStdio(s); err != nil {
		return err
	}

	return nil
}
