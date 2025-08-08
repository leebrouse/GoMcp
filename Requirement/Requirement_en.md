# 📄 Product Requirements Document (PRD)

**Project Name:** Intelligent LLM Tooling Service Platform (based on MCP Server architecture)

**Author:** Leebrouse

**Version:** 1.0

**Date:** 2025-08-05

---

## 1. Background

To improve unified management and invocation efficiency of AI tools, we will build a multi-agent system based on the MCP architecture. It supports registering and invoking multiple AI tools (e.g., LLM Q&A, document parsing, code review, vector search), integrates with vector databases, OpenTelemetry, and Kubernetes, and provides high scalability and observability.

---

## 2. Goals & Scope

### 🎯 Goals

- Provide a pluggable and configurable multi-agent platform that integrates various AI tools
- Support communication over the MCP protocol (e.g., interoperating with `mcphost` and `mcp-inspector`)
- Support auto-registration of tools with a factory pattern for modular tool management
- Enhance Q&A with RAG and vector databases

### ✅ Functional Scope

| Module | Description |
| ------ | ----------- |
| LLM Tooling Services | Multiple AI tool services such as `chatbox`, `codeReview`, and `summarize` |
| Tool Registration Factory | Use the factory pattern to uniformly register tools and inject handlers |
| Vector DB Integration | Use TiDB as backend; support document embedding and similarity search |
| RAG Flow Integration | Retrieval-Augmented Generation for document QA; support file upload, parsing, and knowledge injection |
| Configuration Management | Use Viper to manage configuration files; support env var overrides |
| Service Deployment | Use Makefile, Jenkins, ArgoCD, and K8s for automated build and deployment |
| Observability | Integrate OpenTelemetry for tracing and performance monitoring |

---

## 3. Core Architecture

```plaintext
┌────────────┐      ┌────────────┐      ┌─────────────┐
│  MCP Host  │─────▶│ MCP Server │─────▶│ Tool Handler│
└────────────┘      └────────────┘      └────┬────────┘
                                             │
       ┌────────────┐     ┌────────────┐     ▼
       │  TiDB │◀────│ Embedding  │   Tool Implementations (LLM, PDF, etc.)
       └────────────┘     └────────────┘
```

---

## 4. Feature Details

### 1) LLM Tools (chatbox / codeReview)

- Accept `prompt` / `path` parameters
- Call Gemini or OpenAI APIs
- Support switching models via configuration

### 2) Vector Database Service

- Support document embedding (`embed`) and vector search (`searchVector`)
- Integrate Qdrant (Docker container)
- Vectorize documents after adding them via `mcp resource`

### 3) Factory Registration Mechanism

- All tools are registered via a `ToolFactory` interface
- Tools are auto-discovered and loaded by the main service after initialization
- Facilitates plugin-style extension; adding new tools does not require changing core flow logic

### 4) Configuration Management (Viper)

- Load path: `./internal/common/config/global.yaml`
- Support environment variable overrides
- Example fields:

```yaml
mcpServer:
  llm:
    serverName: llm-server
    serverVersion: v1.0.0
```

---

## 5. Non-functional Requirements

| Item | Requirement |
| ---- | ----------- |
| Performance | Per-request response time ≤ 1s (excluding model latency) |
| Extensibility | Adding new Tools must not require changes to core service structure |
| Observability | Integrate OpenTelemetry for tracing and metrics |
| Security | Later integrate token verification and audit logging |
| Containerization | Build Docker images; support K8s / ArgoCD deployment |

---

## 6. Test Plan

| Test Type | Content |
| --------- | ------- |
| Unit Test | Use `go test` to cover all tool functions |
| Integration Test | Start `mcphost` + `mcp server` to test RPC |
| Tool Verification | Use `make test-mcp` to simulate tool invocation |
| RAG Test | Upload documents and test the `askDocs` tool |

---

## 7. Deployment Plan (CI/CD)

- Use Makefile to manage build and local test commands
- Jenkins + ArgoCD for automated build and release
- Kubernetes for orchestration

---

Below is a CI/CD architecture design integrating **Git + Jenkins + Docker + Kubernetes + Argo CD**. It fits the current `mcp server` project and supports future expansion into multi-service, cloud-native architecture.

---

### 📌 Diagram: CI/CD with Git + Jenkins + Docker + K8s + ArgoCD

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

### 🔧 Tech Stack Details

| Stage | Tool | Purpose |
| ----- | ---- | ------- |
| Development | GitHub / GitLab | Code hosting and change triggers |
| CI | Jenkins | Automated testing, build, image packaging, upload, and pushing deployment configs |
| Image | Docker | Container image packaging |
| Registry | DockerHub / ECR | Image storage |
| GitOps | Argo CD | Continuous delivery, monitors GitOps repo and syncs to K8s |
| Deployment | Kubernetes | Deployment, autoscaling, service discovery, logging & monitoring |

---

### 🚀 Suggested Directory Structure (GitOps Repo)

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

### ✅ Integration Recommendations

- **Jenkinsfile**: Add steps for build, test, image build, push, and updating `kustomization.yaml`.
- **Argo CD**: Connect to the GitOps repo and auto-deploy updates to K8s.
- **K8s**: Manage your MCP service with Helm or Kustomize.
- **Docker**: Keep images as small as possible (multi-stage builds), and push to a private or public registry.

---

If needed, I can generate:

1. `Jenkinsfile` example
2. `k8s` `deployment.yaml`
3. `ArgoCD Application` configuration
4. Best-practice `Dockerfile` template

Would you like me to proceed with these?

## 8. Milestones

| Date | Task |
| ---- | ---- |
| 08/05 ~ 08/08 | Tool module design and factory registration |
| 08/09 ~ 08/13 | Vector DB and RAG flow integration |
| 08/14 ~ 08/16 | Test cases and test pass |
| 08/17 | Deploy to K8s cluster, go-live |

---

If you want me to export this document to PDF, Markdown, or Word, let me know. Do you also want architecture diagrams, API definitions, or CI configs? I can further improve this document.
