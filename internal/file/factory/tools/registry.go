package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool("file", newFileTool)
}

/*file mcp server */
// file tool
func newFileTool() mcp.Tool {
	return mcp.NewTool("file",
		mcp.WithDescription("File tool"),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to the file")),
	)
}

// To do more tools .......
