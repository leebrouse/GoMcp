package service

import "context"

type ResearchService interface {
	SearchKnowledge(ctx context.Context, query string) (string, error)
}
