package server

import (
	"fmt"
	"log"

	"github.com/leebrouse/GoMcp/internal/server/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	FileServerName    = "file-server"
	FileServerVersion = "1.0.0"
)

func StartFile() {

	log.Println("Starting MCP server...")
	// Create a new MCP server
	s := server.NewMCPServer(
		FileServerName,
		FileServerVersion,
		server.WithToolCapabilities(true),
	)


	list := mcp.NewTool("list",
		mcp.WithDescription("展示当前目录下的文件"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the directory to list"),
		),
	)

	codeReview := mcp.NewTool("codeReview",
		mcp.WithDescription("Review the code and provide a list of issues and suggestions for improvement"),
		mcp.WithString("path",
			mcp.Required(),
			mcp.Description("Path to the code to review"),
		),
	)

	// fileUnderstanding := mcp.NewTool("fileUnderstanding",
	// 	mcp.WithDescription("Understand the file and provide a list of issues and suggestions for improvement"),
	// 	mcp.WithString("path",
	// 		mcp.Required(),
	// 		mcp.Description("Path to the file to understand"),
	// 	),
	// )

	log.Println("Adding tool handlers...")
	// Add tool handler
	s.AddTool(list, handler.ListHandler)
	s.AddTool(codeReview, handler.CodeReviewHandler)
	// s.AddTool(fileUnderstanding, handler.FileUnderstandingHandler)
	// Start the stdio server
	// Start StreamableHTTP server
	log.Println("Starting MCP server...")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
