package ui

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/thomaszub/go-todos-hate-pb/service"
	"github.com/thomaszub/go-todos-hate-pb/ui/templates"
)

//go:embed assets
var assets embed.FS

type ToDos struct {
	service *service.ToDos
}

func NewToDos(service *service.ToDos) ToDos {
	return ToDos{service}
}

func (t *ToDos) Register(router *echo.Echo) {
	router.Use(echo.WrapMiddleware(func(h http.Handler) http.Handler {
		return templ.NewCSSMiddleware(h, templates.CheckboxStyle(), templates.DeleteBin())
	}))
	router.GET("/", t.Get)
	router.PATCH("/:id/done", t.SwapDone)
	router.DELETE("/:id", t.Delete)
	router.StaticFS("/assets", echo.MustSubFS(assets, "assets"))
	router.POST("/", t.Add)
}

func (t *ToDos) Get(c echo.Context) error {
	todos, err := t.service.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return renderTemplate(c, http.StatusOK, templates.Todos(todos))
}

func (t *ToDos) SwapDone(c echo.Context) error {
	id := c.PathParam("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "no To-Do id is set")
	}
	todo, err := t.service.SwapDone(id)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("no To-Do found for id %s", id))
	}
	return renderTemplate(c, http.StatusOK, templates.Todo(todo))
}

func (t *ToDos) Delete(c echo.Context) error {
	id := c.PathParam("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "no To-Do id is set")
	}
	err := t.service.Delete(id)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("no To-Do found for id %s", id))
	}
	return c.String(http.StatusOK, "")
}

func (t *ToDos) Add(c echo.Context) error {
	content := c.FormValue("newtodo")
	if content == "" {
		return c.String(http.StatusBadRequest, "no To-Do content is set")
	}
	todo, err := t.service.Add(content)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return renderTemplate(c, http.StatusOK, templates.Todo(todo))
}

func renderTemplate(ctx echo.Context, status int, temp templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	temp.Render(ctx.Request().Context(), buf)
	return ctx.HTML(status, buf.String())
}
