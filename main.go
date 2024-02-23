package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomaszub/go-todos-templ-htmx/service"
	"github.com/thomaszub/go-todos-templ-htmx/ui"
	"github.com/thomaszub/go-todos-templ-htmx/ui/templates"
)

func main() {
	port := 8080
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))
	h := templ.NewCSSMiddleware(r, templates.CheckboxStyle(), templates.DeleteBin())

	s := service.NewToDos()
	c := ui.NewToDos(s)
	c.Register(r)
	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), h))
}
