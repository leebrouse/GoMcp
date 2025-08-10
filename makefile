# ========================
# 📦 Project Configuration
# ========================

APP_NAME       := mcphost
CONFIG_PATH    := ./internal/common/config/mcp.json
MODEL          := google:gemini-2.5-flash
GO_FILES       := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Default target
.DEFAULT_GOAL := help

# ========================
# 📋 Help Menu
# ========================

.PHONY: help
help:
	@echo ""
	@echo "🛠️  Available commands:"
	@echo "  make build          Build the Go project"
	@echo "  make run            Start mcphost (use default model)"
	@echo "  make server         Start example tool server"
	@echo "  make client         Start example client"
	@echo "  make test           Run all Go tests"
	@echo "  make docker-build   Build Docker image"
	@echo "  make test-mcp       Run integration tests"
	@echo ""


.PHONY: mcphost
mcphost:
	@echo "🚀 Starting $(APP_NAME)..."
	$(APP_NAME) --config $(CONFIG_PATH) -m $(MODEL)

# ========================
# 🎯 Examples
# ========================

.PHONY: server
server:
	@echo "🟢 Starting tool example server..."
	go run ./internal/example/tool.go

.PHONY: client
client:
	@echo "🟢 Starting client example..."
	go run ./internal/example/prompt.go

# ========================
# 🧪 Tests
# ========================

.PHONY: test
test:
	@echo "🧪 Running all Go unit tests..."
	go test -v ./...

.PHONY: test-mcp
test-mcp:
	@echo "🧪 Testing MCP tool calls..."
	go test -v ./test/...

# ========================
# 🐳 Docker support
# ========================

.PHONY: docker-build
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t $(APP_NAME):latest .

