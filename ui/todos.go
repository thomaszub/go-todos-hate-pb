package ui

import (
	"embed"
	"fmt"
	"net/http"
	"strconv"

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
	r.Patch("/{id}/done", t.SwapDone)
	r.Route("/assets", func(r chi.Router) {
		r.Get("/*", http.FileServer(http.FS(assets)).ServeHTTP)
	})
}

func (t *ToDos) Get(w http.ResponseWriter, r *http.Request) {
	templates.Todos(t.service.GetAll()).Render(r.Context(), w)
}

func (t *ToDos) SwapDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	todo, err := t.service.SwapDone(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("No To-Do found for id %d", id)))
	}
	templates.Checkbox(todo)
}
