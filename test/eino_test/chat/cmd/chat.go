package main

import (
	"context"
	"log"

	"eino/chat"
)

func main() {
	ctx := context.Background()

	// 使用模版创建messages
	log.Printf("===create messages===\n")
	messages := chat.CreateMessagesFromTemplate()
	log.Printf("messages: %+v\n\n", messages)

	// 创建llm
	log.Printf("===create llm===\n")
	cm := chat.CreateGeminiChatModel(ctx)
	// cm := createOllamaChatModel(ctx)
	log.Printf("create llm success\n\n")

	log.Printf("===llm generate===\n")
	result := chat.Generate(ctx, cm, messages)
	log.Printf("result: %+v\n\n", result)

	log.Printf("===llm stream generate===\n")
	streamResult := chat.Stream(ctx, cm, messages)
	chat.ReportStream(streamResult)
}
