package main

import (
	"context"
	"fmt"
	"log"

	"github.com/robermar23/langchaingo/tools/duckduckgo"
)

func main() {
	ctx := context.Background()

	// 创建 DuckDuckGo 搜索工具（参数：最大结果数量，User-Agent）
	searchTool, _ := duckduckgo.New(5, "LangChainGo/1.0")

	// 执行搜索
	query := "Who won the Euro 2024?"
	result, err := searchTool.Call(ctx, query)
	if err != nil {
		log.Fatalf("搜索失败: %v", err)
	}

	// 打印搜索结果
	fmt.Println("搜索结果：")
	fmt.Println(result)
}
