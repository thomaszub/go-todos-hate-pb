package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/thomaszub/go-todos-hate-pb/service"
	"github.com/thomaszub/go-todos-hate-pb/ui"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		s := service.NewToDos(e.App.Dao(), e.App.Logger())
		c := ui.NewToDos(s)
		c.Register(e.Router)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
