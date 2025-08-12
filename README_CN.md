# 🚀 GoMcp - 智能 LLM 工具服务平台

[![Go Version](https://img.shields.io/badge/Go-1.24.4+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](docker-compose.yaml)

> 基于 MCP (Model Context Protocol) 架构的多智能体系统，支持注册、调用多个 AI 工具，具备高度可扩展性和可观测性。

![基础设施架构图](images/infra.png)
## 🎯 项目简介

GoMcp 是一个基于 MCP 协议构建的智能 LLM 工具服务平台，旨在提供统一的多智能体工具管理和调用服务。项目采用 Go 语言开发，支持多种 AI 工具集成，包括 LLM 问答、文档解析、代码审查、向量搜索等功能。

### 主要优势

- 🔌 **插件化架构**：支持工具的动态注册和扩展
- 🚀 **高性能**：基于 Go 语言，响应时间 ≤ 1s
- 🔍 **可观测性**：集成 OpenTelemetry 支持链路追踪
- 🐳 **容器化部署**：支持 Docker 和 Kubernetes 部署
- 📊 **向量数据库**：集成 TiDB 支持文档嵌入和相似搜索

## ✨ 核心特性

### 🤖 AI 工具服务
- **LLM 问答**：支持 Gemini、OpenAI 等多种模型
- **代码审查**：自动代码质量分析和建议
- **文档摘要**：智能文档内容总结
- **RAG 增强**：基于文档的检索增强生成

### 🏭 工具管理
- **工厂模式**：统一工具注册和管理机制
- **自动发现**：工具自动注册和加载
- **配置管理**：基于 Viper 的灵活配置系统

### 🗄️ 数据存储
- **向量数据库**：TiDB 集成，支持文档嵌入
- **相似搜索**：高效的向量相似度计算
- **资源管理**：支持多种文档格式解析

### 🔧 开发运维
- **链路追踪**：OpenTelemetry 集成
- **容器化**：Docker 支持
- **CI/CD**：Jenkins + ArgoCD + Kubernetes
- **监控告警**：完整的可观测性方案

## 🏗️ 技术架构
### 核心组件

- **MCP Server**：协议通信层，处理客户端请求
- **Tool Factory**：工具注册工厂，管理所有可用工具
- **LLM Service**：大语言模型服务，支持多种 AI 模型
- **File Service**：文件处理服务，支持文档解析
- **Vector Service**：向量数据库服务，支持相似搜索
- **Config Service**：配置管理服务，基于 Viper

## 🚀 快速开始

### 环境要求

- Go 1.24.4+
- Docker & Docker Compose
- MySQL/TiDB (可选，用于向量数据库)

### 安装 Go 环境

#### 方法一：官方安装包（推荐）

1. **下载 Go 1.24.4+**
   - 访问 [Go 官方下载页面](https://golang.org/dl/)
   - 选择适合您系统的版本：
     - **Linux**: `go1.24.4.linux-amd64.tar.gz`
     - **macOS**: `go1.24.4.darwin-amd64.tar.gz`
     - **Windows**: `go1.24.4.windows-amd64.msi`

2. **Linux/macOS 安装**
```bash
# 下载并解压
wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz

# 配置环境变量
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# 验证安装
go version
```

3. **Windows 安装**
   - 下载 `.msi` 安装包
   - 双击运行安装程序
   - 按照向导完成安装
   - 打开命令提示符验证：`go version`

#### 方法二：包管理器安装

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install golang-go
```

**CentOS/RHEL:**
```bash
sudo yum install golang
# 或使用 dnf (新版本)
sudo dnf install golang
```

**macOS (使用 Homebrew):**
```bash
brew install go
```

#### 方法三：使用 g (Go 版本管理器)

```bash
# 安装 g
curl -sSL https://git.io/g-install | sh -s

# 安装指定版本的 Go
g install 1.24.4

# 使用指定版本
g use 1.24.4
```

#### 验证安装

```bash
# 检查 Go 版本
go version

# 检查 Go 环境
go env

# 设置 GOPROXY（国内用户推荐）
go env -w GOPROXY=https://goproxy.cn,direct
```

### 安装步骤

1. **克隆项目**
```bash
git clone https://github.com/leebrouse/GoMcp.git
cd GoMcp
```

2. **安装依赖**
```bash
# install dependence
go mod download

# install mcphost
go install github.com/mark3labs/mcphost@latest

# add your Gemini api key
export GOOGLE_API_KEY='your-api-key'

```

3. **配置环境**
```bash
cp internal/common/config/mcp.json.example internal/common/config/mcp.json
# 编辑配置文件，设置 API 密钥等
```

4. **启动服务**
```bash
# 使用 Makefile
make mcphost

# 或直接运行
mcphost --config ./internal/common/config/mcp.json -m google:gemini-2.5-flash
```

![Run_Example](images/example.png)
### Docker 部署 ( Not complete yet )

```bash
# 构建镜像
make docker-build

# 启动服务
docker-compose up -d
```

## 📦 功能模块

### 1. LLM 工具服务 (`internal/llm/`)

提供多种 AI 工具服务：

- **ChatBox**：通用对话服务
- **CodeReview**：代码审查服务
- **Summarize**：文档摘要服务

### 2. 文件处理服务 (`internal/file/`)

支持多种文件格式处理：

- **PDF 解析**：文档内容提取
- **文本处理**：文本清洗和预处理
- **格式转换**：支持多种格式转换

### 3. 通用组件 (`internal/common/`)

- **配置管理**：基于 Viper 的配置系统
- **数据模型**：统一的数据结构定义
- **设计模式**：常用设计模式实现
- **OpenTelemetry**：可观测性集成

### 4. 工具工厂 (`internal/llm/factory/`)

实现工具的动态注册和管理：

```go
// 工具注册示例
func init() {
    factory.RegisterTool("chatbox", &ChatBoxTool{})
    factory.RegisterTool("codeReview", &CodeReviewTool{})
}
```

## 📚 API 文档

### MCP 协议接口

项目基于 MCP 协议提供以下接口：

#### 1. 工具列表 (`listTools`)

获取所有可用工具列表：

```json
{
  "tools": [
    {
      "name": "chatbox",
      "description": "通用对话服务",
      "inputSchema": {
        "type": "object",
        "properties": {
          "prompt": {
            "type": "string",
            "description": "用户输入的问题"
          }
        }
      }
    }
  ]
}
```

#### 2. 工具调用 (`callTool`)

调用指定工具：

```json
{
  "name": "chatbox",
  "arguments": {
    "prompt": "你好，请介绍一下 GoMcp 项目"
  }
}
```

#### 3. 资源管理 (`listResources`)

管理文档资源：

```json
{
  "resources": [
    {
      "uri": "file:///path/to/document.pdf",
      "name": "项目文档",
      "description": "GoMcp 项目说明文档"
    }
  ]
}
```

### 配置示例

```yaml
mcpServer:
  llm:
    serverName: "llm-server"
    serverVersion: "v1.0.0"
    models:
      - name: "gemini-2.5-flash"
        provider: "google"
        apiKey: "${GEMINI_API_KEY}"
      - name: "gpt-4"
        provider: "openai"
        apiKey: "${OPENAI_API_KEY}"
  
  vector:
    database:
      type: "tidb"
      host: "localhost"
      port: 4000
      database: "vector_db"
```

## 🐳 部署指南 (Not complete yet)

### Kubernetes 部署

1. **创建命名空间**
```bash
kubectl create namespace mcp-system
```

2. **部署应用**
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

3. **配置 ConfigMap**
```bash
kubectl apply -f k8s/configmap.yaml
```

### CI/CD 流程

项目支持完整的 CI/CD 流程：

1. **代码提交** → GitHub/GitLab
2. **自动测试** → Jenkins
3. **构建镜像** → Docker
4. **推送镜像** → Docker Registry
5. **自动部署** → ArgoCD + Kubernetes

## 🛠️ 开发指南

### 项目结构

```
GoMcp/
├── cmd/                    # 主程序入口
├── internal/              # 内部包
│   ├── common/           # 通用组件
│   │   ├── config/       # 配置管理
│   │   ├── model/        # 数据模型
│   │   └── opentelemetry/ # 可观测性
│   ├── llm/              # LLM 服务
│   │   ├── domain/       # 领域模型
│   │   ├── service/      # 业务服务
│   │   ├── handler/      # 请求处理
│   │   └── factory/      # 工具工厂
│   └── file/             # 文件服务
├── example/              # 示例代码
├── test/                 # 测试用例
├── Requirement/          # 需求文档
└── images/              # 文档图片
```

### 添加新工具

1. **实现工具接口**
```go
type Tool interface {
    Name() string
    Description() string
    Execute(ctx context.Context, args map[string]interface{}) (interface{}, error)
}
```

2. **注册工具**
```go
func init() {
    factory.RegisterTool("myTool", &MyTool{})
}
```

3. **编写测试**
```go
func TestMyTool(t *testing.T) {
    // 测试代码
}
```

### 运行测试

```bash
# 运行所有测试
make test

# 运行 MCP 集成测试
make test-mcp

# 运行特定测试
go test -v ./test/llm_function_test/
```

## 🤝 贡献指南

我们欢迎所有形式的贡献！

### 贡献流程

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 开发规范

- 遵循 Go 语言编码规范
- 添加必要的测试用例
- 更新相关文档
- 确保 CI/CD 流程通过

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [MCP Protocol](https://modelcontextprotocol.io/) - 模型上下文协议
- [Mark3Labs](https://github.com/mark3labs/mcp-go) - Go MCP 实现
- [LangChain Go](https://github.com/robermar23/langchaingo) - Go 语言 AI 框架
- [Viper](https://github.com/spf13/viper) - 配置管理库
- [GORM](https://gorm.io/) - Go ORM 库

## 📞 联系我们

- **作者**: Leebrouse
- **项目地址**: https://github.com/leebrouse/GoMcp
- **问题反馈**: [GitHub Issues](https://github.com/leebrouse/GoMcp/issues)

---

⭐ 如果这个项目对您有帮助，请给我们一个星标！
