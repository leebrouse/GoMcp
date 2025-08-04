package main
 
import (
    "context"
    "fmt"
 
    "github.com/mark3labs/mcp-go/mcp"
    "github.com/mark3labs/mcp-go/server"
)
 
func main() {
    // Create a new MCP server
    s := server.NewMCPServer(
        "Hello World Server",
        "1.0.0",
        server.WithToolCapabilities(true),
    )
 
    // Define a simple tool
    tool := mcp.NewTool("hello_world",
        mcp.WithDescription("Say hello to someone"),
        mcp.WithString("name",
            mcp.Required(),
            mcp.Description("Name of the person to greet"),
        ),
    )
 
    // Add tool handler
    s.AddTool(tool, helloHandler)
 
    // Start the stdio server
    if err := server.ServeStdio(s); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}
 
func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    arguments := request.GetArguments()
    name, ok := arguments["name"].(string)
    if !ok {
        return &mcp.CallToolResult{
            Content: []mcp.Content{
                mcp.TextContent{
                    Type: "text",
                    Text: "Error: name parameter is required and must be a string",
                },
            },
            IsError: true,
        }, nil
    }
 
    return &mcp.CallToolResult{
        Content: []mcp.Content{
            mcp.TextContent{
                Type: "text",
                Text: fmt.Sprintf("Hello, %s! 👋", name),
            },
        },
    }, nil
}