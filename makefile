# start server
server:
	go run ./internal/example/tool.go

# client
client:
	go run ./internal/example/prompt.go

# start mcp server
start:
	mcp-inspector go run ./internal/cmd/main.go

# start mcphost
mcphost:
	mcphost --config ./internal/config/mcp.json -m google:gemini-2.5-flash