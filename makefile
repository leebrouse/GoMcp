# start server
server:
	go run ./internal/example/tool.go

# client
client:
	go run ./internal/example/prompt.go

# start mcp server
start:
	mcp-inspector go run ./internal/common/config/mcp.json

# start mcphost
mcphost:
	mcphost --config ./internal/common/config/mcp.json -m google:gemini-2.5-flash

# set google api key
set-key:
	export GOOGLE_API_KEY='AIzaSyCKURVV8jEX3CsRu_4pysxmJm3IH4mr8VU'