package trivia_test

import (
	"testing"

	"github.com/gabehf/trivia-api/trivia"
)

var Q trivia.Questions
var expect *trivia.Question

func TestMain(m *testing.M) {
	Q.Init()
	Q.Categories = []string{"world history"}
	Q.M = map[string][]trivia.Question{
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
	expect = &trivia.Question{
		Question: "The ancient city of Rome was built on how many hills?",
		Category: "World History",
		Format:   "MultipleChoice",
		Choices: []string{
			"Eight",
			"Four",
			"Nine",
			"Seven",
		},
		Answer: "Seven",
	}
	m.Run()
}
