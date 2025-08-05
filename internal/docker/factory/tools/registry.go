package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

func init() {
	RegisterTool("docker", newDockerTool)
}

/*docker mcp server */
// docker tool
func newDockerTool() mcp.Tool {
	return mcp.NewTool("docker",
		mcp.WithDescription("Docker tool"),
		mcp.WithString("prompt", mcp.Required(), mcp.Description("Prompt to send to the LLM")),
	)
}

// To do more tools .......