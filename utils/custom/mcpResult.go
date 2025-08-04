package custom

import "github.com/mark3labs/mcp-go/mcp"

// Helper function to create a new text result
func NewTextResult(text string, isError bool) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: text,
			},
		},
		IsError: isError,
	}
}
