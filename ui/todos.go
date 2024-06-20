package ui

import (
	"embed"
	"fmt"
	"net/http"
	"strconv"

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

func (t *ToDos) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /", t.Get)
	mux.HandleFunc("PATCH /{id}/done", t.SwapDone)
	mux.HandleFunc("DELETE /{id}", t.Delete)
	mux.Handle("GET /assets/{rest...}", http.FileServer(http.FS(assets)))
	mux.HandleFunc("POST /", t.Add)
}

func (t *ToDos) Get(w http.ResponseWriter, r *http.Request) {
	templates.Todos(t.service.GetAll()).Render(r.Context(), w)
}

func (t *ToDos) SwapDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	todo, err := t.service.SwapDone(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("No To-Do found for id %d", id)))
		return
	}
	templates.Todo(todo).Render(r.Context(), w)
}

func (t *ToDos) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = t.service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("No To-Do found for id %d", id)))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *ToDos) Add(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("newtodo")
	if content == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No To-Do content is set"))
		return
	}
	todo := t.service.Add(content)
	templates.Todo(todo).Render(r.Context(), w)
}
