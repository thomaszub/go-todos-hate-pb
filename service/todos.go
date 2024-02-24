package service

import (
	"fmt"

	"github.com/thomaszub/go-todos-templ-htmx/model"
)

type ToDos struct {
	todos  []*model.ToDo
	nextId int
}

func NewToDos() *ToDos {
	return &ToDos{
		todos:  []*model.ToDo{},
		nextId: 1,
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

func (t *ToDos) Delete(id int) error {
	for i, todo := range t.todos {
		if id == todo.Id {
			t.todos = append(t.todos[:i], t.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("ToDo %d not found", id)
}

func (t *ToDos) Add(content string) model.ToDo {
	todo := model.ToDo{
		Id:      t.nextId,
		Content: content,
		Done:    false,
	}
	t.nextId += 1
	t.todos = append(t.todos, &todo)
	return todo
}
