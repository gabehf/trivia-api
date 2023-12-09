package server

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type GetGuessRequest struct {
	QuestionId string `json:"question_id" query:"question_id"`
	Guess      string `json:"guess" query:"guess"`
}
type GetGuessResponse struct {
	QuestionId string `json:"question_id"`
	Correct    bool   `json:"correct"`
}

func (s *Server) GetGuess(e echo.Context) error {
	req := new(GetGuessRequest)
	e.Bind(req)

	// ensure required parameters exist
	errs := make(map[string]string, 0)
	if req.Guess == "" {
		errs["guess"] = "required parameter missing"
	}
	if req.QuestionId == "" {
		errs["question_id"] = "required parameter missing"
	}
	if len(errs) > 0 {
		return e.JSON(400, &ErrorResponse{
			Error: true,
			Data:  errs,
		})
	}

	question := s.Q.GetQuestionById(req.QuestionId)
	if question == nil {
		errs["question_id"] = "invalid or malformed"
		return e.JSON(404, &ErrorResponse{
			Error: true,
			Data:  errs,
		})
	}

	// validate answer with case insensitive string compare
	correct := strings.EqualFold(question.Answer, req.Guess)

	return e.JSONPretty(200, &GetGuessResponse{req.QuestionId, correct}, "  ")
}
