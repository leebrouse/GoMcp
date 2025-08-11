package infrastructure

import "context"

type DuckDuckGo interface {
	// search info from the duckduckgo
	DuckDuckGoSearch(ctx context.Context, prompt string) (string, error)

	//Todo: Extend more functions .........
}
