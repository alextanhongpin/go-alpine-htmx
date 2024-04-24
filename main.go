package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("healthz")
		fmt.Fprint(w, "ok")
	})
	mux.HandleFunc("POST /clicked", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("clicked")
		template.Must(template.New("clicked").Parse(`<h1>{{.Msg}}</h1>`)).Execute(w, map[string]any{
			"Msg": "clicked",
		})
	})

	mux.HandleFunc("GET /trigger_delay", func(w http.ResponseWriter, r *http.Request) {
		// The data is passed in as q.
		q := r.URL.Query().Get("q")
		fmt.Println(q)
		fruits := []string{"apple", "banana", "cherry"}

		var sb strings.Builder
		for _, f := range fruits {
			if strings.HasPrefix(f, q) {
				sb.WriteString(fmt.Sprintf(`<li>%s</li>`, f))
			}
		}
		logger.Info("trigger_delay")

		fmt.Fprint(w, sb.String())
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	logger.Info("listening to port *:8080. press ctrl+c to cancel.")
	http.ListenAndServe(":8080", mux)
}
