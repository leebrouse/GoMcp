# 📄 需求文档（PRD）

**项目名称：** 智能 LLM 工具服务平台（基于 MCP Server 架构）
**作者：** \[你的名字]
**版本：** 1.0
**日期：** 2025-08-05

---

## 一、项目背景

为了提升人工智能工具的统一管理和调用效率，构建一个基于 MCP 架构的多智能体系统，支持注册、调用多个 AI 工具（如 LLM 问答、文档解析、代码审查、向量搜索等），可与向量数据库、OpenTelemetry、Kubernetes 等集成，具备高度可扩展性和可观测性。

---

## 二、目标与范围

### 🎯 目标

* 提供一个可插件化、可配置的多智能体平台，支持多种 AI 工具集成
* 支持通过 MCP 协议通信（例如：`mcphost` 与 `mcp-inspector` 对接）
* 支持自动注册工具，采用工厂模式进行工具模块化管理
* 支持通过 RAG 和向量数据库增强问答能力

### ✅ 功能范围

| 模块       | 描述                                               |
| -------- | ------------------------------------------------ |
| LLM 工具服务 | `chatbox`, `codeReview`, `summarize` 等多种 AI 工具服务 |
| 工具注册工厂   | 使用工厂模式统一注册工具并注入处理函数                              |
| 向量数据库接入  | 使用 Qdrant 作为后端，支持文档嵌入、相似搜索等功能                    |
| RAG 流程集成 | 支持基于文档问答的检索增强生成（RAG），支持文件上传解析和知识注入               |
| 配置管理     | 使用 Viper 管理配置文件，支持环境变量覆盖                         |
| 服务部署     | 使用 Makefile、Jenkins、ArgoCD、K8s 实现自动构建与部署         |
| 可观测性接入   | 接入 OpenTelemetry，进行链路追踪与性能监控                     |

---

## 三、核心技术架构

```plaintext
┌────────────┐      ┌────────────┐      ┌─────────────┐
│  MCP Host  │─────▶│ MCP Server │─────▶│ Tool Handler│
└────────────┘      └────────────┘      └────┬────────┘
                                             │
       ┌────────────┐     ┌────────────┐     ▼
       │  Qdrant DB │◀────│ Embedding  │   Tool 实现（LLM、PDF 等）
       └────────────┘     └────────────┘
```

---

## 四、功能详细说明

### 1. LLM 工具（chatbox / codeReview）

* 接收 `prompt` / `path` 参数
* 调用 Gemini 或 OpenAI API
* 支持多种模型切换（通过配置）

### 2. 向量数据库服务

* 支持文档嵌入（`embed`）和向量搜索（`searchVector`）
* 接入 Qdrant（Docker 容器）
* 文档资源通过 `mcp resource` 添加后向量化

### 3. 工厂注册机制

* 所有工具通过 `ToolFactory` 接口注册
* 工具初始化后可被主服务自动识别与加载
* 便于插件化扩展，新增工具无需修改主流程逻辑

### 4. 配置管理（Viper）

* 加载路径：`./internal/common/config/global.yaml`
* 支持环境变量覆盖
* 关键字段示例：

  ```yaml
  mcpServer:
    llm:
      serverName: llm-server
      serverVersion: v1.0.0
  ```

---

## 五、非功能需求

| 项目    | 要求                                |
| ----- | --------------------------------- |
| 性能    | 每个请求响应时间 ≤ 1s（模型能力除外）             |
| 扩展性   | 新增 Tool 无需改动主服务代码结构               |
| 可观测性  | 接入 OpenTelemetry 支持链路追踪、指标上报      |
| 安全性   | 后期集成 Token 校验、日志审计                |
| 容器化部署 | 使用 Docker 构建镜像，支持 K8s / ArgoCD 部署 |

---

## 六、测试方案

| 测试类型   | 内容                                 |
| ------ | ---------------------------------- |
| 单元测试   | 使用 `go test` 覆盖所有工具函数              |
| 集成测试   | 启动 `mcphost` + `mcp server` 测试 RPC |
| 工具验证   | `make test-mcp` 模拟工具调用             |
| RAG 测试 | 上传文档后测试 `askDocs` 工具               |

---

## 七、部署方案（CI/CD）

* 使用 Makefile 管理构建和本地测试命令
* Jenkins + ArgoCD 自动构建与发布
* 使用 Kubernetes 进行容器编排

---

以下是一个将 **Git + Jenkins + Docker + Kubernetes + Argo CD** 整合起来的 **CI/CD 架构图设计说明**。它适用于你当前的 `mcp server` 项目，并支持后续扩展为多服务、Cloud Native 架构。

---

### 📌 架构图说明：CI/CD with Git + Jenkins + Docker + K8s + ArgoCD

```
                     +----------------------+
                     |      Developer       |
                     |   Push Code to Git   |
                     +----------+-----------+
                                |
                                ▼
                       +------------------+
                       |      GitHub      |
                       |(Git Repository)  |
                       +--------+---------+
                                |
                Webhook Triggers Jenkins Build
                                |
                                ▼
                      +---------------------+
                      |      Jenkins        |
                      |   (CI Pipeline)     |
                      +---------------------+
                      | 1. Checkout Code    |
                      | 2. Run Tests        |
                      | 3. Build Binary     |
                      | 4. Build Docker Img |
                      | 5. Push to Registry |
                      | 6. Update Manifests |
                      +---------+-----------+
                                |
                                ▼
                   +---------------------------+
                   |     Docker Registry       |
                   |   (e.g. DockerHub/ECR)    |
                   +---------------------------+
                                |
                                ▼
                     +---------------------+
                     |     GitOps Repo     |
                     |  (Helm/Manifests)   |
                     +---------------------+
                                |
                        Auto-Sync Triggered
                                |
                                ▼
                     +----------------------+
                     |      Argo CD         |
                     |   (CD Controller)    |
                     +----------+-----------+
                                |
                                ▼
                      +------------------+
                      |    Kubernetes    |
                      |    (Target Env)  |
                      +------------------+
                      | 1. Pull Image     |
                      | 2. Deploy MCP     |
                      | 3. Auto Scale     |
                      +------------------+
```

---

### 🔧 技术栈明细

| 阶段     | 工具              | 作用                        |
| ------ | --------------- | ------------------------- |
| 开发     | GitHub / GitLab | 代码托管与变更触发                 |
| CI     | Jenkins         | 自动测试、构建、镜像打包、上传、推送部署配置等   |
| 镜像     | Docker          | 容器镜像打包                    |
| 镜像仓库   | DockerHub / ECR | 镜像存储                      |
| GitOps | Argo CD         | 持续部署，监控 GitOps 仓库并同步至 K8s |
| 部署     | Kubernetes      | 部署、扩缩容、服务发现、日志监控等         |

---

### 🚀 示例目录结构建议（GitOps Repo）

```bash
gitops/
├── apps/
│   └── mcp-server/
│       ├── kustomization.yaml
│       └── deployment.yaml
│       └── service.yaml
└── base/
    └── mcp-server/
        ├── deployment.yaml
        └── configmap.yaml
        └── kustomization.yaml
```

---

### ✅ 接入建议

* **Jenkinsfile**：
  添加构建、测试、镜像构建、推送、更新 `kustomization.yaml` 的步骤。
* **Argo CD**：
  连接 GitOps Repo，自动监听更新并部署到 K8s。
* **K8s**：
  使用 Helm 或 Kustomize 管理你的 mcp 服务。
* **Docker**：
  镜像尽可能小（使用多阶段构建），并推送到私有或公有仓库。

---

如果你需要，我可以为你生成：

1. `Jenkinsfile` 示例
2. `k8s` 的 `deployment.yaml`
3. `ArgoCD Application` 配置
4. `Dockerfile` 最佳实践模板

是否需要我继续生成这些内容？


## 八、里程碑

| 日期             | 任务              |
| -------------- | --------------- |
| 08/05 \~ 08/08 | 工具模块设计与注册工厂实现   |
| 08/09 \~ 08/13 | 向量数据库与 RAG 流程集成 |
| 08/14 \~ 08/16 | 测试用例编写与测试通过     |
| 08/17          | 部署至 K8s 集群，正式上线 |

---

如果你希望我把这份文档导出为 PDF、Markdown 或 Word 文档，也可以告诉我。你还想加上功能图表、接口定义或 CI 配置吗？我可以进一步完善。
