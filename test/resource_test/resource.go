package main

// mcp-inspector go run resource.go
import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("File Server", "1.0.0",
		server.WithResourceCapabilities(true, true),
	)

	// Add a static file resource
	s.AddResource(
		mcp.NewResource(
			"./README.md",
			"Project README",
			mcp.WithResourceDescription("Main project documentation"),
			mcp.WithMIMEType("text/markdown"),
		),
		handleReadmeFile,
	)

	server.ServeStdio(s)
}

// resource handler
func handleReadmeFile(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	content, err := os.ReadFile("README.md")
	if err != nil {
		return nil, fmt.Errorf("failed to read README: %w", err)
	}

	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      request.Params.URI,
			MIMEType: "text/markdown",
			Text:     string(content),
		},
	}, nil
}
