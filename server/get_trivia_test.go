package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gabehf/trivia-api/server"
	"github.com/labstack/echo/v4"
)

func TestTriviaHandler(t *testing.T) {
	jsonBody := []byte("{\"category\":\"World History\"}")

	expect := server.GetTriviaResponse{
		Question: "The ancient city of Rome was built on how many hills?",
		Format:   "MultipleChoice",
		Category: "World History",
	}

	// OK path: json body
	e := echo.New()
	req := httptest.NewRequest("GET", "/trivia", bytes.NewReader(jsonBody))
	req.Header["Content-Type"] = []string{"application/json"}
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err := S.GetTrivia(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.Code)
	}

	result := new(server.GetTriviaResponse)
	err = json.Unmarshal(res.Body.Bytes(), result)
	if err != nil {
		t.Error("malformed json response")
	}
	if result.Question != expect.Question {
		t.Errorf("expected question '%s', got '%s'", expect.Question, result.Question)
	}
	if result.Format != expect.Format {
		t.Errorf("expected format %s, got %s", expect.Format, result.Format)
	}
	if !strings.EqualFold(expect.Category, result.Category) {
		t.Errorf("expected category %s, got %s", expect.Category, result.Category)
	}

	// OK path: urlencoded body
	e = echo.New()
	req = httptest.NewRequest("GET", "/trivia?category=World+History", nil)
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	err = S.GetTrivia(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.Code)
	}
	expect = server.GetTriviaResponse{
		Question: "The ancient city of Rome was built on how many hills?",
		Format:   "MultipleChoice",
		Category: "World History",
	}
	result = new(server.GetTriviaResponse)
	err = json.Unmarshal(res.Body.Bytes(), result)
	if err != nil {
		t.Error("malformed json response")
	}
	if result.Question != expect.Question {
		t.Errorf("expected question '%s', got '%s'", expect.Question, result.Question)
	}
	if result.Format != expect.Format {
		t.Errorf("expected format %s, got %s", expect.Format, result.Format)
	}
	if !strings.EqualFold(expect.Category, result.Category) {
		t.Errorf("expected category %s, got %s", expect.Category, result.Category)
	}

	// OK path: no body (random category)
	e = echo.New()
	req = httptest.NewRequest("GET", "/trivia", nil)
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	err = S.GetTrivia(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.Code)
	}
	expect = server.GetTriviaResponse{
		Question: "The ancient city of Rome was built on how many hills?",
		Format:   "MultipleChoice",
		Category: "World History",
	}
	result = new(server.GetTriviaResponse)
	err = json.Unmarshal(res.Body.Bytes(), result)
	if err != nil {
		t.Error("malformed json response")
	}
	if result.Question != expect.Question {
		t.Errorf("expected question '%s', got '%s'", expect.Question, result.Question)
	}
	if result.Format != expect.Format {
		t.Errorf("expected format %s, got %s", expect.Format, result.Format)
	}
	if !strings.EqualFold(expect.Category, result.Category) {
		t.Errorf("expected category %s, got %s", expect.Category, result.Category)
	}

	// FAIL path: invalid category
	e = echo.New()
	req = httptest.NewRequest("GET", "/trivia?category=70s+Music", nil)
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	err = S.GetTrivia(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusNotFound {
		t.Errorf("expected status 404 Not Found, got %d", res.Code)
	}
	errResult := struct {
		Error bool
		Data  map[string]string
	}{}
	err = json.Unmarshal(res.Body.Bytes(), &errResult)
	if err != nil {
		t.Error("malformed json response")
	}
	if !errResult.Error {
		t.Error("expected error to be true, got false")
	}
	if errResult.Data["category"] == "" {
		t.Errorf("expected error information in data[category], got \"\"")
	}
}
