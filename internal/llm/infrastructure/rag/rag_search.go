// using TiDB as vector database
package rag

import (
	"context"

	"github.com/leebrouse/GoMcp/internal/llm/infrastructure"
)

type RAG struct{}

func NewRAG() infrastructure.RAGRetriever {
	return &RAG{}
}

// Vector search
func (r *RAG) VectorSearch(ctx context.Context, query string) ([]infrastructure.SearchResult, error) {
	//Todo: VectorSearch
	return nil, nil
}

// Full-text search
func (r *RAG) FullTextSearch(ctx context.Context, query string) ([]infrastructure.SearchResult, error) {
	//Todo: FullTextSearch
	return nil, nil
}

// Mix-text search
func (r *RAG) MixSearch(ctx context.Context, query string) ([]infrastructure.SearchResult, error) {
	//Todo: MixTextSearch
	return nil, nil
}
