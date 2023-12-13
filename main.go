package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomaszub/go-todos-templ-htmx/ui"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	c := ui.NewTodosUI()
	c.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
