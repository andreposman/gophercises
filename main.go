package main

import (
	_1_quizz "github.com/andreposman/gophercises/01-quiz"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	_1_quizz.Quiz()
}
