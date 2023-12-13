package ui

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thomaszub/go-todos-templ-htmx/ui/templates"
)

//go:embed assets
var assets embed.FS

type TodosUI struct{}

func NewTodosUI() TodosUI {
	return TodosUI{}
}

func (t *TodosUI) Register(r chi.Router) {
	r.Get("/", t.Get)
	r.Route("/assets", func(r chi.Router) {
		r.Get("/*", http.FileServer(http.FS(assets)).ServeHTTP)
	})
}

func (t *TodosUI) Get(w http.ResponseWriter, r *http.Request) {
	templates.Todos().Render(r.Context(), w)
}
