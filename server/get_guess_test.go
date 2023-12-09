package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabehf/trivia-api/server"
	"github.com/labstack/echo/v4"
)

func TestGuessHandler(t *testing.T) {
	jsonBody := []byte("{\"question_id\":\"World History|0\",\"guess\":\"seven\"}")

	// OK path: json body
	e := echo.New()
	req := httptest.NewRequest("GET", "/guess", bytes.NewReader(jsonBody))
	req.Header["Content-Type"] = []string{"application/json"}
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err := S.GetGuess(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.Code)
	}

	result := new(server.GetGuessResponse)
	err = json.Unmarshal(res.Body.Bytes(), result)
	if err != nil {
		t.Error("malformed json response")
	}
	if result.QuestionId != "World History|0" {
		t.Errorf("expected question_id 'World History|0', got '%s'", result.QuestionId)
	}
	if result.Correct != true {
		t.Errorf("expected correct to be true, got false")
	}

	// OK path: urlencoded body
	e = echo.New()
	req = httptest.NewRequest("GET", "/guess?question_id=World+History%7C0&guess=Seven", nil)
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	err = S.GetGuess(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", res.Code)
	}

	result = new(server.GetGuessResponse)
	err = json.Unmarshal(res.Body.Bytes(), result)
	if err != nil {
		t.Error("malformed json response")
	}
	if result.QuestionId != "World History|0" {
		t.Errorf("expected question_id 'World History|0', got '%s'", result.QuestionId)
	}
	if result.Correct != true {
		t.Errorf("expected correct to be true, got false")
	}

	// FAIL path: invalid question id
	e = echo.New()
	req = httptest.NewRequest("GET", "/guess?question_id=hey&guess=Seven", nil)
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	err = S.GetGuess(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusNotFound {
		t.Errorf("expected status 400 Bad Request, got %d", res.Code)
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
	if errResult.Data["question_id"] == "" {
		t.Errorf("expected error information in data[question_id], got \"\"")
	}

	// FAIL path: missing params
	e = echo.New()
	req = httptest.NewRequest("GET", "/guess", nil)
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	res = httptest.NewRecorder()
	c = e.NewContext(req, res)
	err = S.GetGuess(c)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if res.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 Bad Request, got %d", res.Code)
	}
	errResult = struct {
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
	if errResult.Data["question_id"] == "" {
		t.Errorf("expected error information in data[question_id], got \"\"")
	}
	if errResult.Data["guess"] == "" {
		t.Errorf("expected error information in data[guess], got \"\"")
	}
}
