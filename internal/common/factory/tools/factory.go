package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

type ToolFactory func() mcp.Tool

var registry = map[string]ToolFactory{}

func RegisterTool(name string, factory ToolFactory) {
	registry[name] = factory
}

func GetTool(name string) mcp.Tool {
	if factory, exists := registry[name]; exists {
		return factory()
	}
	panic("tool not registered: " + name)
}

func GetAllTools() []mcp.Tool {
	var tools []mcp.Tool
	for _, factory := range registry {
		tools = append(tools, factory())
	}
	return tools
}
