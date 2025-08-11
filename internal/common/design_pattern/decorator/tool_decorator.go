package decorator

import (
	"context"
	"fmt"
	"time"
)

// Tool 基础工具接口
type Tool interface {
	Execute(ctx context.Context, input interface{}) (interface{}, error)
	GetName() string
}

// BaseTool 基础工具实现
type BaseTool struct {
	name string
}

func NewBaseTool(name string) *BaseTool {
	return &BaseTool{name: name}
}

func (t *BaseTool) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	return fmt.Sprintf("Base tool %s executed with input: %v", t.name, input), nil
}

func (t *BaseTool) GetName() string {
	return t.name
}

// ToolDecorator 工具装饰器接口
type ToolDecorator interface {
	Tool
	SetTool(tool Tool)
}

// LoggingDecorator 日志装饰器
type LoggingDecorator struct {
	tool Tool
}

func NewLoggingDecorator(tool Tool) *LoggingDecorator {
	return &LoggingDecorator{tool: tool}
}

func (d *LoggingDecorator) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	start := time.Now()
	fmt.Printf("[LOG] Starting execution of tool: %s\n", d.tool.GetName())

	result, err := d.tool.Execute(ctx, input)

	duration := time.Since(start)
	if err != nil {
		fmt.Printf("[LOG] Tool %s failed after %v: %v\n", d.tool.GetName(), duration, err)
	} else {
		fmt.Printf("[LOG] Tool %s completed successfully in %v\n", d.tool.GetName(), duration)
	}

	return result, err
}

func (d *LoggingDecorator) GetName() string {
	return d.tool.GetName()
}

func (d *LoggingDecorator) SetTool(tool Tool) {
	d.tool = tool
}

// RetryDecorator 重试装饰器
type RetryDecorator struct {
	tool       Tool
	maxRetries int
	backoff    time.Duration
}

func NewRetryDecorator(tool Tool, maxRetries int, backoff time.Duration) *RetryDecorator {
	return &RetryDecorator{
		tool:       tool,
		maxRetries: maxRetries,
		backoff:    backoff,
	}
}

func (d *RetryDecorator) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	var lastErr error

	for attempt := 0; attempt <= d.maxRetries; attempt++ {
		if attempt > 0 {
			fmt.Printf("[RETRY] Attempt %d for tool %s\n", attempt, d.tool.GetName())
			time.Sleep(d.backoff * time.Duration(attempt))
		}

		result, err := d.tool.Execute(ctx, input)
		if err == nil {
			return result, nil
		}

		lastErr = err
		fmt.Printf("[RETRY] Tool %s failed on attempt %d: %v\n", d.tool.GetName(), attempt+1, err)
	}

	return nil, fmt.Errorf("tool %s failed after %d attempts: %w", d.tool.GetName(), d.maxRetries+1, lastErr)
}

func (d *RetryDecorator) GetName() string {
	return d.tool.GetName()
}

func (d *RetryDecorator) SetTool(tool Tool) {
	d.tool = tool
}

// CacheDecorator 缓存装饰器
type CacheDecorator struct {
	tool  Tool
	cache map[string]interface{}
}

func NewCacheDecorator(tool Tool) *CacheDecorator {
	return &CacheDecorator{
		tool:  tool,
		cache: make(map[string]interface{}),
	}
}

func (d *CacheDecorator) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	// 简单的缓存键生成（实际应用中应该使用更复杂的哈希算法）
	cacheKey := fmt.Sprintf("%s_%v", d.tool.GetName(), input)

	// 检查缓存
	if cached, exists := d.cache[cacheKey]; exists {
		fmt.Printf("[CACHE] Cache hit for tool %s\n", d.tool.GetName())
		return cached, nil
	}

	// 执行工具
	result, err := d.tool.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	// 缓存结果
	d.cache[cacheKey] = result
	fmt.Printf("[CACHE] Cached result for tool %s\n", d.tool.GetName())

	return result, nil
}

func (d *CacheDecorator) GetName() string {
	return d.tool.GetName()
}

func (d *CacheDecorator) SetTool(tool Tool) {
	d.tool = tool
}

// MetricsDecorator 指标装饰器
type MetricsDecorator struct {
	tool    Tool
	metrics map[string]int64
}

func NewMetricsDecorator(tool Tool) *MetricsDecorator {
	return &MetricsDecorator{
		tool:    tool,
		metrics: make(map[string]int64),
	}
}

func (d *MetricsDecorator) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	start := time.Now()

	result, err := d.tool.Execute(ctx, input)

	duration := time.Since(start)
	d.metrics["execution_count"]++
	d.metrics["total_duration_ms"] += duration.Milliseconds()

	if err != nil {
		d.metrics["error_count"]++
	} else {
		d.metrics["success_count"]++
	}

	fmt.Printf("[METRICS] Tool %s - Executions: %d, Success: %d, Errors: %d, Avg Duration: %dms\n",
		d.tool.GetName(),
		d.metrics["execution_count"],
		d.metrics["success_count"],
		d.metrics["error_count"],
		d.metrics["total_duration_ms"]/d.metrics["execution_count"])

	return result, err
}

func (d *MetricsDecorator) GetName() string {
	return d.tool.GetName()
}

func (d *MetricsDecorator) SetTool(tool Tool) {
	d.tool = tool
}

func (d *MetricsDecorator) GetMetrics() map[string]int64 {
	return d.metrics
}
