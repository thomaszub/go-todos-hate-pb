package controller

import (
	"net/http"

	"github.com/thomaszub/go-todos-templ-htmx/templates"
)

type TodosController struct{}

func NewTodos() TodosController {
	return TodosController{}
}

func (t *TodosController) Get(w http.ResponseWriter, r *http.Request) {
	templates.Todos().Render(r.Context(), w)
}
