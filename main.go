package main

import (
	_2_url_shortener "github.com/andreposman/gophercises/02-url-shortener"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	//_1_quizz.Quiz()
	_2_url_shortener.URLShortener()
}
