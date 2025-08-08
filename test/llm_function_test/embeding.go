package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	contents := []*genai.Content{
		// Tips: Role 表示这段内容的“发言者/身份”，如 RoleUser（用户）、RoleModel（模型）。
		// 在生成任务中会影响指令优先级与安全策略；在嵌入(Embedding)任务中多为元信息，通常不影响向量结果。
		genai.NewContentFromText("What is the meaning of life?", genai.RoleUser),
	}
	result, err := client.Models.EmbedContent(ctx,
		"gemini-embedding-001",
		contents,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	embeddings, err := json.MarshalIndent(result.Embeddings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(embeddings))

	if len(result.Embeddings) > 0 {
		fmt.Println("dim:", len(result.Embeddings[0].Values))
	}
}
