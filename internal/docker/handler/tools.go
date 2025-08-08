package handler

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/leebrouse/GoMcp/utils/custom"
	"github.com/mark3labs/mcp-go/mcp"
)

// docker handler
func DockerHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	image, ok := arguments["image"].(string)
	if !ok {
		return custom.NewTextResult("Error: image parameter is required and must be a string", true), nil
	}

	// run docker image 
	// tips：should write in the service layer
	err := exec.Command("docker", "run", image).Run()
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	return custom.NewTextResult(fmt.Sprintf("Docker image %s run successfully", image), false), nil
}
