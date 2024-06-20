package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomaszub/go-todos-templ-htmx/service"
	"github.com/thomaszub/go-todos-templ-htmx/ui"
	"github.com/thomaszub/go-todos-templ-htmx/ui/templates"
)

func main() {
	port := 8080
	mux := http.NewServeMux()
	var h http.Handler = templ.NewCSSMiddleware(mux, templates.CheckboxStyle(), templates.DeleteBin())
	h = middleware.Compress(5)(h)
	h = middleware.Logger(h)

	s := service.NewToDos()
	c := ui.NewToDos(s)
	c.Register(mux)
	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), h))
}
