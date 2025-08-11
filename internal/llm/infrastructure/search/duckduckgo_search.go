package search

import (
	"context"
	"fmt"

	"github.com/leebrouse/GoMcp/internal/llm/infrastructure"
	"github.com/robermar23/langchaingo/tools/duckduckgo"
)

//
type SearchEngin struct {
	maxResult int
	userAgent string
}

func NewDuckDuckGo(maxResult int, userAgent string) infrastructure.DuckDuckGo {
	return &SearchEngin{
		maxResult: maxResult,
		userAgent: userAgent,
	}
}

// call DuckDuckGoSearch to get reference from the public internat
func (se *SearchEngin) DuckDuckGoSearch(ctx context.Context, keyword string) (string, error) {
	// create DuckDuckGo search tool
	searchTool, err := duckduckgo.New(se.maxResult, se.userAgent)
	if err != nil {
		return fmt.Sprintf("Error init a new DuckDuckGo Search tool : %v", err), err
	}

	// execute searching
	result, err := searchTool.Call(ctx, keyword)
	if err != nil {
		return fmt.Sprintf("Fail to Search info by using DuckDuckGo: %v", err), err
	}

	// return result
	return result, nil
}
