package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool("chatbox", newChatboxTool)
	RegisterTool("codeReview", newCodeReviewTool)
	RegisterTool("readDocument", newReadDocument)
}

/*llm mcp server */
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

// Document reader tool
func newReadDocument() mcp.Tool {
	return mcp.NewTool("readDocument",
		mcp.WithDescription("read Document and do operation by the prompt"),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to the document or PDF")),
		mcp.WithString("prompt", mcp.Required(), mcp.Description("what do user want to get or do from the document")),
	)
}

//ToDo: embedder tool
