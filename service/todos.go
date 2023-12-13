package service

import "github.com/thomaszub/go-todos-templ-htmx/model"

type ToDos struct {
	todos []model.ToDo
}

func NewToDos() *ToDos {
	return &ToDos{
		todos: []model.ToDo{
			{
				Id:      1,
				Content: "Milk",
				Done:    false,
			},
			{
				Id:      2,
				Content: "Bread",
				Done:    true,
			},
		},
	}
}

func (t *ToDos) GetAll() []model.ToDo {
	return t.todos
}
