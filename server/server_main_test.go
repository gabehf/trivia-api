package server_test

import (
	"testing"

	"github.com/gabehf/trivia-api/server"
	"github.com/gabehf/trivia-api/trivia"
)

var S *server.Server

func TestMain(m *testing.M) {
	S = new(server.Server)
	S.Init()
	S.Q.Init()
	S.Q.Categories = []string{"world history"}
	S.Q.M = map[string][]trivia.Question{
		"world history": {
			{
				Question: "The ancient city of Rome was built on how many hills?",
				Format:   "MultipleChoice",
				Category: "World History",
				Choices: []string{
					"Eight",
					"Four",
					"Nine",
					"Seven",
				},
				Answer: "Seven",
			},
		},
	}
	m.Run()
}
