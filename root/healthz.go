package root

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"log/slog"
)

var tpl = template.Must(template.ParseGlob("views/*.html"))

type Handler struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("healthz")
	fmt.Fprint(w, "ok")
}

func (h *Handler) Clicked(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("clicked")
	tpl.ExecuteTemplate(w, "clicked.html", map[string]any{
		"Msg": "wew",
	})
}

func (h *Handler) TriggerDelay(w http.ResponseWriter, r *http.Request) {
	// The data is passed in as q.
	q := r.URL.Query().Get("q")
	fmt.Println(q)
	fruits := []string{"apple", "banana", "cherry"}

	var result []string
	for _, f := range fruits {
		if strings.HasPrefix(f, q) {
			result = append(result, f)
		}
	}
	h.logger.Info("trigger_delay")

	tpl.ExecuteTemplate(w, "search.html", map[string]any{
		"Fruits": result,
		"Empty":  len(result) == 0,
		"Query":  q,
	})
}
