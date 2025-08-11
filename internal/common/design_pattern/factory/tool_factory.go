package factory

import (
	"fmt"
	"sync"
)

// ToolType 工具类型枚举
type ToolType string

const (
	ToolTypeLLM    ToolType = "llm"
	ToolTypeDocker ToolType = "docker"
	ToolTypeFile   ToolType = "file"
	ToolTypeSearch ToolType = "search"
)

// ToolConfig 工具配置接口
type ToolConfig interface {
	GetType() ToolType
	Validate() error
}

// Tool 工具接口
type Tool interface {
	Execute(input interface{}) (interface{}, error)
	GetName() string
	GetDescription() string
}

// ToolFactory 工具工厂
type ToolFactory struct {
	creators map[ToolType]ToolCreator
	mu       sync.RWMutex
}

// ToolCreator 工具创建器接口
type ToolCreator func(config ToolConfig) (Tool, error)

// NewToolFactory 创建工具工厂
func NewToolFactory() *ToolFactory {
	return &ToolFactory{
		creators: make(map[ToolType]ToolCreator),
	}
}

// RegisterCreator 注册工具创建器
func (f *ToolFactory) RegisterCreator(toolType ToolType, creator ToolCreator) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.creators[toolType] = creator
}

// CreateTool 创建工具
func (f *ToolFactory) CreateTool(config ToolConfig) (Tool, error) {
	f.mu.RLock()
	creator, exists := f.creators[config.GetType()]
	f.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("unsupported tool type: %s", config.GetType())
	}

	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return creator(config)
}

// LLMConfig LLM工具配置
type LLMConfig struct {
	ModelName string `json:"model_name"`
	APIKey    string `json:"api_key"`
	MaxTokens int    `json:"max_tokens"`
}

func (c *LLMConfig) GetType() ToolType {
	return ToolTypeLLM
}

func (c *LLMConfig) Validate() error {
	if c.ModelName == "" {
		return fmt.Errorf("model name is required")
	}
	if c.APIKey == "" {
		return fmt.Errorf("API key is required")
	}
	return nil
}

// DockerConfig Docker工具配置
type DockerConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UseTLS   bool   `json:"use_tls"`
	CertPath string `json:"cert_path"`
}

func (c *DockerConfig) GetType() ToolType {
	return ToolTypeDocker
}

func (c *DockerConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port <= 0 {
		return fmt.Errorf("port must be positive")
	}
	return nil
}
