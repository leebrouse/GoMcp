package strategy

import (
	"context"
)

// LLMStrategy 定义LLM策略接口
type LLMStrategy interface {
	Generate(ctx context.Context, prompt string) (string, error)
	GetModelName() string
}

// GeminiStrategy Gemini模型策略
type GeminiStrategy struct {
	apiKey string
}

func NewGeminiStrategy(apiKey string) *GeminiStrategy {
	return &GeminiStrategy{apiKey: apiKey}
}

func (g *GeminiStrategy) Generate(ctx context.Context, prompt string) (string, error) {
	// 实现Gemini API调用
	return "Gemini response: " + prompt, nil
}

func (g *GeminiStrategy) GetModelName() string {
	return "gemini-2.5-flash"
}

// OpenAIStrategy OpenAI模型策略
type OpenAIStrategy struct {
	apiKey string
}

func NewOpenAIStrategy(apiKey string) *OpenAIStrategy {
	return &OpenAIStrategy{apiKey: apiKey}
}

func (o *OpenAIStrategy) Generate(ctx context.Context, prompt string) (string, error) {
	// 实现OpenAI API调用
	return "OpenAI response: " + prompt, nil
}

func (o *OpenAIStrategy) GetModelName() string {
	return "gpt-4"
}

// LLMContext LLM上下文，管理策略切换
type LLMContext struct {
	strategy LLMStrategy
}

func NewLLMContext(strategy LLMStrategy) *LLMContext {
	return &LLMContext{strategy: strategy}
}

func (c *LLMContext) SetStrategy(strategy LLMStrategy) {
	c.strategy = strategy
}

func (c *LLMContext) Execute(ctx context.Context, prompt string) (string, error) {
	return c.strategy.Generate(ctx, prompt)
}
