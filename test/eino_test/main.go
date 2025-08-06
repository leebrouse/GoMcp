package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	// 使用模版创建messages
	log.Printf("===create messages===\n")
	messages := CreateMessagesFromTemplate()
	log.Printf("messages: %+v\n\n", messages)

	// 创建llm
	log.Printf("===create llm===\n")
	cm := CreateGeminiChatModel(ctx)
	// cm := createOllamaChatModel(ctx)
	log.Printf("create llm success\n\n")

	log.Printf("===llm generate===\n")
	result := Generate(ctx, cm, messages)
	log.Printf("result: %+v\n\n", result)

	log.Printf("===llm stream generate===\n")
	streamResult := Stream(ctx, cm, messages)
	ReportStream(streamResult)
}
