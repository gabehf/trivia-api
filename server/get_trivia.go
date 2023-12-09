package server

import (
	"math/rand"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GetTriviaRequest struct {
	Category string `json:"category" query:"category"`
}
type GetTriviaResponse struct {
	QuestionId string            `json:"question_id"`
	Question   string            `json:"question"`
	Category   string            `json:"category"`
	Format     string            `json:"format"`
	Choices    map[string]string `json:"choices,omitempty"`
}

type ErrorResponse struct {
	Error   bool              `json:"error"`
	Data    map[string]string `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
}

func (s *Server) GetTrivia(e echo.Context) error {
	req := new(GetTriviaRequest)
	e.Bind(req)

	question, qIndex := s.Q.GetRandomQuestion(req.Category)
	if question == nil {
		return e.JSON(404, &ErrorResponse{
			Error: true,
			Data: map[string]string{
				"category": "category is invalid",
			},
		})
	}
	// randomly order answer choices if the format is multiple choice
	if question.Format == "MultipleChoice" && question.Choices != nil {
		rand.Shuffle(len(question.Choices), func(i, j int) {
			question.Choices[i], question.Choices[j] = question.Choices[j], question.Choices[i]
		})
		// enforce that multiple choice questions must have four choices
		// if not, there must be an error in our data somewhere that we need
		// to fix
		if len(question.Choices) != 4 {
			return e.JSON(500, &ErrorResponse{
				Error:   true,
				Message: "internal server error",
			})
		}
	}

	// build and return response
	tq := new(GetTriviaResponse)
	tq.QuestionId = question.Category + "|" + strconv.Itoa(qIndex)
	tq.Category = question.Category
	tq.Format = question.Format
	tq.Question = question.Question
	if tq.Format == "MultipleChoice" {
		tq.Choices = map[string]string{
			"A": question.Choices[0],
			"B": question.Choices[1],
			"C": question.Choices[2],
			"D": question.Choices[3],
		}
	}
	return e.JSONPretty(200, tq, "  ")
}
