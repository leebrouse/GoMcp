# ========================
# 📦 项目基本配置
# ========================

APP_NAME       := mcphost
CONFIG_PATH    := ./internal/common/config/mcp.json
MODEL          := google:gemini-2.5-flash
GO_FILES       := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# 默认目标
.DEFAULT_GOAL := help

# ========================
# 📋 帮助菜单
# ========================

.PHONY: help
help:
	@echo ""
	@echo "🛠️  可用命令："
	@echo "  make build          编译 Go 项目"
	@echo "  make run            启动 mcphost（使用默认模型）"
	@echo "  make server         启动示例工具服务"
	@echo "  make client         启动示例客户端"
	@echo "  make test           执行所有 Go 测试"
	@echo "  make docker-build   构建 Docker 镜像"
	@echo "  make test-mcp       执行集成测试"
	@echo ""


.PHONY: mcphost
mcphost:
	@echo "🚀 启动 $(APP_NAME)..."
	$(APP_NAME) --config $(CONFIG_PATH) -m $(MODEL)

# ========================
# 🎯 示例程序
# ========================

.PHONY: server
server:
	@echo "🟢 启动 tool 示例 server..."
	go run ./internal/example/tool.go

.PHONY: client
client:
	@echo "🟢 启动 client 示例..."
	go run ./internal/example/prompt.go

# ========================
# 🧪 测试
# ========================

.PHONY: test
test:
	@echo "🧪 运行所有 Go 单元测试..."
	go test -v ./...

.PHONY: test-mcp
test-mcp:
	@echo "🧪 测试 MCP 工具调用..."
	go test -v ./test/...

# ========================
# 🐳 Docker 支持
# ========================

.PHONY: docker-build
docker-build:
	@echo "🐳 构建 Docker 镜像..."
	docker build -t $(APP_NAME):latest .
