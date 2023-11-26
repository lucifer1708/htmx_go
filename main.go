package main

import (
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	r := mux.NewRouter()
	r.Use(LoggingMiddleware)
	r.HandleFunc("/", h1)

	loggedRouter := LoggingMiddleware(r)

	// Start the server
	log.Info().Msg("Server starting on port 8000")
	http.ListenAndServe(":8000", loggedRouter)
}

type Dicts struct {
	Name  string
	Class string
}

func h1(w http.ResponseWriter, r *http.Request) {
	tmp1 := template.Must(template.ParseFiles("templates/index.html"))
	dicts := map[string][]Dicts{
		"dicts": {
			{Name: "sumitdhiman", Class: "12th"},
			{Name: "Hello", Class: "12th"},
			{Name: "Dinma", Class: "12th"},
			{Name: "cowstub", Class: "11th"},
		},
	}
	tmp1.Execute(w, dicts)
}
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.Path).
			Dur("duration", duration).
			Msg("HTTP Request")
	})
}
