# start server
server:
	go run ./internal/example/tool.go

# client
client:
	go run ./internal/example/prompt.go

# start mcp server
start:
	mcp-inspector go run ./internal/cmd/main.go