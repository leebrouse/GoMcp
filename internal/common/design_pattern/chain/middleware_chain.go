package chain

import (
	"context"
	"fmt"
)

// Request 请求结构
type Request struct {
	Data    interface{}            `json:"data"`
	Headers map[string]string      `json:"headers"`
	Params  map[string]interface{} `json:"params"`
}

// Response 响应结构
type Response struct {
	Data    interface{}       `json:"data"`
	Headers map[string]string `json:"headers"`
	Status  int               `json:"status"`
	Error   error             `json:"error,omitempty"`
}

// Handler 处理器接口
type Handler interface {
	Handle(ctx context.Context, req *Request) (*Response, error)
	SetNext(handler Handler)
}

// BaseHandler 基础处理器
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// AuthenticationHandler 认证处理器
type AuthenticationHandler struct {
	BaseHandler
	apiKey string
}

func NewAuthenticationHandler(apiKey string) *AuthenticationHandler {
	return &AuthenticationHandler{apiKey: apiKey}
}

func (h *AuthenticationHandler) Handle(ctx context.Context, req *Request) (*Response, error) {
	// 检查API密钥
	if req.Headers["X-API-Key"] != h.apiKey {
		return &Response{
			Status: 401,
			Error:  fmt.Errorf("unauthorized: invalid API key"),
		}, nil
	}

	fmt.Println("[AUTH] Authentication passed")

	// 调用下一个处理器
	if h.next != nil {
		return h.next.Handle(ctx, req)
	}

	return &Response{Status: 200}, nil
}

// RateLimitHandler 限流处理器
type RateLimitHandler struct {
	BaseHandler
	limit    int
	requests map[string]int
}

func NewRateLimitHandler(limit int) *RateLimitHandler {
	return &RateLimitHandler{
		limit:    limit,
		requests: make(map[string]int),
	}
}

func (h *RateLimitHandler) Handle(ctx context.Context, req *Request) (*Response, error) {
	clientID := req.Headers["X-Client-ID"]
	if clientID == "" {
		clientID = "default"
	}

	// 简单的内存限流（实际应用中应该使用Redis等）
	if h.requests[clientID] >= h.limit {
		return &Response{
			Status: 429,
			Error:  fmt.Errorf("rate limit exceeded"),
		}, nil
	}

	h.requests[clientID]++
	fmt.Printf("[RATE_LIMIT] Client %s: %d/%d requests\n", clientID, h.requests[clientID], h.limit)

	// 调用下一个处理器
	if h.next != nil {
		return h.next.Handle(ctx, req)
	}

	return &Response{Status: 200}, nil
}

// ValidationHandler 验证处理器
type ValidationHandler struct {
	BaseHandler
}

func NewValidationHandler() *ValidationHandler {
	return &ValidationHandler{}
}

func (h *ValidationHandler) Handle(ctx context.Context, req *Request) (*Response, error) {
	// 验证请求数据
	if req.Data == nil {
		return &Response{
			Status: 400,
			Error:  fmt.Errorf("request data is required"),
		}, nil
	}

	fmt.Println("[VALIDATION] Request validation passed")

	// 调用下一个处理器
	if h.next != nil {
		return h.next.Handle(ctx, req)
	}

	return &Response{Status: 200}, nil
}

// LoggingHandler 日志处理器
type LoggingHandler struct {
	BaseHandler
}

func NewLoggingHandler() *LoggingHandler {
	return &LoggingHandler{}
}

func (h *LoggingHandler) Handle(ctx context.Context, req *Request) (*Response, error) {
	fmt.Printf("[LOGGING] Processing request: %+v\n", req.Data)

	// 调用下一个处理器
	if h.next != nil {
		resp, err := h.next.Handle(ctx, req)
		if err != nil {
			fmt.Printf("[LOGGING] Request failed: %v\n", err)
		} else {
			fmt.Printf("[LOGGING] Request completed with status: %d\n", resp.Status)
		}
		return resp, err
	}

	return &Response{Status: 200}, nil
}

// BusinessLogicHandler 业务逻辑处理器
type BusinessLogicHandler struct {
	BaseHandler
}

func NewBusinessLogicHandler() *BusinessLogicHandler {
	return &BusinessLogicHandler{}
}

func (h *BusinessLogicHandler) Handle(ctx context.Context, req *Request) (*Response, error) {
	// 执行业务逻辑
	fmt.Println("[BUSINESS] Executing business logic")

	// 模拟业务处理
	result := fmt.Sprintf("Processed: %v", req.Data)

	return &Response{
		Status: 200,
		Data:   result,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

// ChainBuilder 链构建器
type ChainBuilder struct {
	handlers []Handler
}

func NewChainBuilder() *ChainBuilder {
	return &ChainBuilder{
		handlers: make([]Handler, 0),
	}
}

func (cb *ChainBuilder) AddHandler(handler Handler) *ChainBuilder {
	cb.handlers = append(cb.handlers, handler)
	return cb
}

func (cb *ChainBuilder) Build() Handler {
	if len(cb.handlers) == 0 {
		return nil
	}

	// 构建责任链
	for i := 0; i < len(cb.handlers)-1; i++ {
		cb.handlers[i].SetNext(cb.handlers[i+1])
	}

	return cb.handlers[0]
}

// 使用示例
func ExampleUsage() {
	// 构建处理链
	chain := NewChainBuilder().
		AddHandler(NewLoggingHandler()).
		AddHandler(NewAuthenticationHandler("secret-key")).
		AddHandler(NewRateLimitHandler(10)).
		AddHandler(NewValidationHandler()).
		AddHandler(NewBusinessLogicHandler()).
		Build()

	// 创建请求
	req := &Request{
		Data: "test data",
		Headers: map[string]string{
			"X-API-Key":   "secret-key",
			"X-Client-ID": "client1",
		},
		Params: map[string]interface{}{
			"param1": "value1",
		},
	}

	// 处理请求
	resp, err := chain.Handle(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", resp)
}
