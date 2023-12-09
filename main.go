package main

import (
	"log"

	"github.com/gabehf/trivia-api/server"
)

func main() {
	log.Println("Trivia API listening on http://127.0.0.1:3000")
	log.Fatal(server.Run())
}
