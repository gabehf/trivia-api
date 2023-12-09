package trivia_test

import (
	"bytes"
	"reflect"
	"slices"
	"testing"

	"github.com/gabehf/trivia-api/trivia"
)

func TestGetRandomQuestion(t *testing.T) {
	// on OK path, GetTrivia must return the question in our test data
	tq, _ := Q.GetRandomQuestion("world history")
	if tq == nil {
		t.Fatal("trivia question must not be nil")
	}
	if !reflect.DeepEqual(tq, expect) {
		t.Errorf("returned question does not match expectation, got %v", tq)
	}

	// with no category specified, GetTrivia must pick a random category and fetch a question
	// with only one question in our test data, it is the same question from before
	tq, _ = Q.GetRandomQuestion("")
	if tq == nil {
		t.Fatal("trivia question must not be nil")
	}
	if !reflect.DeepEqual(tq, expect) {
		t.Errorf("returned question does not match expectation, got %v", tq)
	}

	// on FAIL path, GetTrivia must return nil to indicate no questions are found
	tq, _ = Q.GetRandomQuestion("Geography")
	if tq != nil {
		t.Errorf("expected nil, got %v", tq)
	}
}

func TestGetQuestionById(t *testing.T) {
	// on OK path, GetTrivia must return the question in our test data
	tq := Q.GetQuestionById("world History|0")
	if tq == nil {
		t.Fatal("trivia question must not be nil")
	}
	if !reflect.DeepEqual(tq, expect) {
		t.Errorf("returned question does not match expectation, got %v", tq)
	}

	// FAIL path: malformed id
	tq = Q.GetQuestionById("hey")
	if tq != nil {
		t.Errorf("expected nil, got %v", tq)
	}
	// FAIL path: invalid category
	tq = Q.GetQuestionById("hey|0")
	if tq != nil {
		t.Errorf("expected nil, got %v", tq)
	}
	// FAIL path: invalid index
	tq = Q.GetQuestionById("world history|9")
	if tq != nil {
		t.Errorf("expected nil, got %v", tq)
	}
}

func TestLoad(t *testing.T) {
	json := []byte(`
		{
			"world history": [
				{
					"category": "World History",
					"question": "How many years did the 100 years war last?",
					"answer": "116",
					"format": "MultipleChoice",
					"choices": [
						"116",
						"87",
						"12",
						"205"
					]
				},
				{
					"category": "World History",
					"question": "True or False: John Wilkes Booth assassinated Abraham Lincoln.",
					"answer": "True",
					"format": "TrueFalse"
				}
			],
			"geography": [
				{
					"category": "Geography",
					"question": "What is the capital city of Japan?",
					"answer": "Tokyo",
					"format": "MultipleChoice",
					"choices": [
						"Beijing",
						"Seoul",
						"Bangkok",
						"Tokyo"
					]
				},
				{
					"category": "Geography",
					"question": "True or False: The Amazon Rainforest is located in Africa.",
					"answer": "False",
					"format": "TrueFalse"
				}
			]
		}
	`)
	expectCategories := []string{"world history", "geography"}
	expectQuestions := map[string][]trivia.Question{
		"world history": {
			{
				Question: "How many years did the 100 years war last?",
				Format:   "MultipleChoice",
				Answer:   "116",
				Category: "World History",
				Choices: []string{
					"116",
					"87",
					"12",
					"205",
				},
			},
			{
				Question: "True or False: John Wilkes Booth assassinated Abraham Lincoln.",
				Format:   "TrueFalse",
				Answer:   "True",
				Category: "World History",
			},
		},
		"geography": {
			{
				Question: "What is the capital city of Japan?",
				Format:   "MultipleChoice",
				Answer:   "Tokyo",
				Category: "Geography",
				Choices:  []string{"Beijing", "Seoul", "Bangkok", "Tokyo"},
			},
			{
				Question: "True or False: The Amazon Rainforest is located in Africa.",
				Format:   "TrueFalse",
				Answer:   "False",
				Category: "Geography",
			},
		},
	}
	qq := new(trivia.Questions)
	qq.Init()
	err := qq.Load(bytes.NewReader(json))
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}
	for _, cat := range expectCategories {
		if !slices.Contains(qq.Categories, cat) {
			t.Errorf("expected category %s not present", cat)
		}
	}
	if !reflect.DeepEqual(qq.M, expectQuestions) {
		t.Errorf("unexpected question map, got %v", qq.M)
	}
}
