package service

import (
	"fmt"

	"github.com/thomaszub/go-todos-templ-htmx/model"
)

type ToDos struct {
	todos []*model.ToDo
}

func NewToDos() *ToDos {
	return &ToDos{
		todos: []*model.ToDo{
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
	todos := []model.ToDo{}
	for _, todo := range t.todos {
		todos = append(todos, *todo)
	}
	return todos
}

func (t *ToDos) SwapDone(id int) (model.ToDo, error) {
	for _, todo := range t.todos {
		if id == todo.Id {
			todo.Done = !todo.Done
			return *todo, nil
		}
	}
	return model.ToDo{}, fmt.Errorf("ToDo %d not found", id)
}
