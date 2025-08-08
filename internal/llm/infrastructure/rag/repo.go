// using TiDB as vector database 
package rag

type RAGRepository struct{}

func (r *RAGRepository) Search(query string) string {
	// 向量搜索逻辑
	return "search result"
}
