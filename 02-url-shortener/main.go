package _2_url_shortener

import (
	"fmt"
	"github.com/andreposman/gophercises/pkg/utils"
	"github.com/rs/zerolog/log"
	"net/http"
)

func URLShortener() {
	utils.PrintLineDivider(90)
	utils.PrintName("02", "URL Shortener")
	utils.PrintLineDivider(90)

	mux := createServer()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func createServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := fmt.Fprintln(w, "Hello world")

	log.Info().Msgf("Response: %s", string(resp))
}
