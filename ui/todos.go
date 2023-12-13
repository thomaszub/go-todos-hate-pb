package ui

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thomaszub/go-todos-templ-htmx/service"
	"github.com/thomaszub/go-todos-templ-htmx/ui/templates"
)

//go:embed assets
var assets embed.FS

type ToDos struct {
	service *service.ToDos
}

func NewToDos(service *service.ToDos) ToDos {
	return ToDos{service}
}

func (t *ToDos) Register(r chi.Router) {
	r.Get("/", t.Get)
	r.Route("/assets", func(r chi.Router) {
		r.Get("/*", http.FileServer(http.FS(assets)).ServeHTTP)
	})
}

func (t *ToDos) Get(w http.ResponseWriter, r *http.Request) {
	templates.Todos(t.service.GetAll()).Render(r.Context(), w)
}
