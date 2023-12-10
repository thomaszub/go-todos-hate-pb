package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomaszub/go-todos-templ-htmx/controller"
)

func main() {
	c := controller.NewTodos()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", c.Get)

	log.Fatal(http.ListenAndServe(":8080", r))
}
