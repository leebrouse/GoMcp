package main

import (
	"log"

	"github.com/leebrouse/GoMcp/internal/server"
)

func main() {
	log.Println("Starting server...")

	server.StartLLM()

}
