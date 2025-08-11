package infrastructure

import "context"

// 搜索结果统一结构
type SearchResult struct {
	ID      string  // 文档或片段 ID
	Content string  // 原始文本
	Score   float64 // 相似度或 BM25 分值
}

type RAGRetriever interface {
	// Vector Search
	VectorSearch(ctx context.Context, query string) ([]SearchResult, error)

	// FullText Search
	FullTextSearch(ctx context.Context, query string) ([]SearchResult, error)

	// Mix Search
	MixSearch(ctx context.Context, query string) ([]SearchResult, error) // 拼写已修正
}
