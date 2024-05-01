package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/alextanhongpin/go-alpine-htmx/root"
)

var (
	// Apply the base template, then the content.
	home    = template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
	profile = template.Must(template.ParseFiles("templates/base.html", "templates/profile.html"))
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mux := http.NewServeMux()

	r := root.New(logger)
	mux.HandleFunc("GET /healthz", r.Healthz)
	mux.HandleFunc("POST /clicked", r.Clicked)
	mux.HandleFunc("GET /trigger_delay", r.TriggerDelay)
	mux.HandleFunc("POST /response", r.Response)
	mux.HandleFunc("GET /users/me", r.Me)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		home.Execute(w, nil)
	})

	mux.HandleFunc("GET /profile", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}
		profile.Execute(w, map[string]any{
			"Name": name,
		})
	})

	logger.Info("listening to port *:8080. press ctrl+c to cancel.")
	http.ListenAndServe(":8080", mux)
}
