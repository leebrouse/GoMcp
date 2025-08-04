package handler

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/leebrouse/GoMcp/utils/custom"
	"github.com/mark3labs/mcp-go/mcp"
)

// list handler
func ListHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	arguments := request.GetArguments()
	path, ok := arguments["path"].(string)
	if !ok {
		return custom.NewTextResult("Error: path parameter is required and must be a string", true), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return custom.NewTextResult(fmt.Sprintf("Error: %v", err), true), nil
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return custom.NewTextResult("Files:\n"+strings.Join(fileNames, "\n"), false), nil
}
