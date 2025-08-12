package service

import "context"

// llm mcp server interface 
type LLMService interface {







	// using vector search to find the most similar one in the vector database (TiDB)
	SearchKnowledge(ctx context.Context, query string) (string, error)
}

// file mcp server interface 

// docker mcp server interface 

//Todo： extend more mcp server interface 