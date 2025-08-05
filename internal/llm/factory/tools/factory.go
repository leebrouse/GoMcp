package tools

import (
	"log"

	"github.com/leebrouse/GoMcp/internal/llm/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// create Tool pool
func CreatToolPool() []server.ServerTool {

	var tools []server.ServerTool

	// auto register all tools and add corresponding handler
	for _, tool := range GetAllTools() {
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

		serverTool := server.ServerTool{
			Tool:    tool,
			Handler: handlerFunc,
		}
		tools = append(tools, serverTool)

		log.Printf("Tool registered: %s", tool.GetName())
	}

	return tools
}

// Tool factory
type ToolFactory func() mcp.Tool

var registry = map[string]ToolFactory{}

func RegisterTool(name string, factory ToolFactory) {
	registry[name] = factory
}

// get tool by name
func GetTool(name string) mcp.Tool {
	if factory, exists := registry[name]; exists {
		return factory()
	}
	panic("tool not registered: " + name)
}

// get all tools
func GetAllTools() []mcp.Tool {
	var tools []mcp.Tool
	for _, factory := range registry {
		tools = append(tools, factory())
	}
	return tools
}
