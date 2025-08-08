package service

import "context"

type ResearchService interface {
	// using vector search to find the most similar one in the vector database (TiDB)
	SearchKnowledge(ctx context.Context, query string) (string, error)
}
