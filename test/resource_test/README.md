<!-- mcp-inspector go run main.go -->
# MCP 资源服务器示例

这是一个基于 [MCP (Model Context Protocol)](https://modelcontextprotocol.io/) 的资源服务器示例，演示如何创建一个简单的文件资源服务。

## 功能特性

- **静态文件资源**: 提供 README.md 文件作为可访问的资源
- **MIME 类型支持**: 自动识别并设置正确的 MIME 类型 (text/markdown)
- **标准 MCP 协议**: 完全兼容 MCP 协议规范
- **简单易用**: 通过标准输入/输出进行通信

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 运行服务器

```bash
go run resource.go
```

服务器将通过标准输入/输出启动，等待 MCP 客户端的连接。

### 3. 使用客户端连接

```bash
# 使用 MCP 客户端连接到此服务器
mcp-client connect --server "go run resource.go"
```

## 代码结构

### 主要组件

- **MCPServer**: 创建 MCP 服务器实例
- **Resource**: 定义可访问的资源 (README.md)
- **Resource Handler**: 处理资源读取请求

### 关键代码

```go
// 创建 MCP 服务器
s := server.NewMCPServer("File Server", "1.0.0",
    server.WithResourceCapabilities(true, true),
)

// 添加静态文件资源
s.AddResource(
    mcp.NewResource(
        "file://README.md",
        "Project README",
        mcp.WithResourceDescription("Main project documentation"),
        mcp.WithMIMEType("text/markdown"),
    ),
    handleReadmeFile,
)
```

## 资源定义

| 属性 | 值 | 描述 |
|------|-----|------|
| URI | `file://README.md` | 资源的唯一标识符 |
| 名称 | `Project README` | 资源的显示名称 |
| 描述 | `Main project documentation` | 资源的详细描述 |
| MIME 类型 | `text/markdown` | 内容的媒体类型 |

## 扩展指南

### 添加更多资源

```go
// 添加新的资源
s.AddResource(
    mcp.NewResource(
        "file://config.yaml",
        "Configuration File",
        mcp.WithResourceDescription("Application configuration"),
        mcp.WithMIMEType("application/x-yaml"),
    ),
    handleConfigFile,
)
```

### 自定义资源处理器

```go
func handleConfigFile(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
    content, err := os.ReadFile("config.yaml")
    if err != nil {
        return nil, fmt.Errorf("failed to read config: %w", err)
    }

    return []mcp.ResourceContents{
        mcp.TextResourceContents{
            URI:      request.Params.URI,
            MIMEType: "application/x-yaml",
            Text:     string(content),
        },
    }, nil
}
```

## 故障排除

### 常见问题

1. **依赖问题**: 确保已安装 `github.com/mark3labs/mcp-go`
2. **文件不存在**: 确保 README.md 文件存在于项目根目录
3. **权限问题**: 确保有读取文件的权限

### 调试模式

```bash
# 启用详细日志
DEBUG=1 go run resource.go
```

## 相关链接

- [MCP 协议规范](https://modelcontextprotocol.io/)
- [MCP Go SDK](https://github.com/mark3labs/mcp-go)
- [项目主页](https://github.com/your-org/oasisdb)

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](../LICENSE) 文件。
