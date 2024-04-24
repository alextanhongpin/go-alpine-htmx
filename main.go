package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/alextanhongpin/go-alpine-htmx/root"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mux := http.NewServeMux()

	r := root.New(logger)
	mux.HandleFunc("GET /healthz", r.Healthz)
	mux.HandleFunc("POST /clicked", r.Clicked)
	mux.HandleFunc("GET /trigger_delay", r.TriggerDelay)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	logger.Info("listening to port *:8080. press ctrl+c to cancel.")
	http.ListenAndServe(":8080", mux)
}
