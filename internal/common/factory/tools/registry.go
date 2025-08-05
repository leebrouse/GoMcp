package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool("chatbox", newChatboxTool)
	RegisterTool("codeReview", newCodeReviewTool)
}

// chatbox tool
func newChatboxTool() mcp.Tool {
	return mcp.NewTool("chatbox",
		mcp.WithDescription("Send a prompt to the LLM"),
		mcp.WithString("prompt", mcp.Required(), mcp.Description("Prompt to send to the LLM")),
	)
}

// code review tool
func newCodeReviewTool() mcp.Tool {
	return mcp.NewTool("codeReview",
		mcp.WithDescription("Review the code and provide suggestions"),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to the code")),
	)
}
