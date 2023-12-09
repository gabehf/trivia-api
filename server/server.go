package server

import (
	"os"

	"github.com/gabehf/trivia-api/trivia"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Q *trivia.Questions
}

func (s *Server) Init() {
	s.Q = new(trivia.Questions)
	s.Q.Init()
}

func Run() error {
	// init server struct
	s := new(Server)
	s.Init()

	// load trivia data
	file, err := os.Open("trivia.json")
	if err != nil {
		panic(err)
	}
	err = s.Q.Load(file)
	if err != nil {
		panic(err)
	}

	// create router and mount handlers
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/trivia", s.GetTrivia)
	e.GET("/guess", s.GetGuess)

	// start listening
	return e.Start(":3000")
}
