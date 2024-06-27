package service

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
	"github.com/thomaszub/go-todos-hate-pb/model"
)

type ToDos struct {
	dao    *daos.Dao
	logger *slog.Logger
}

func NewToDos(dao *daos.Dao, logger *slog.Logger) *ToDos {
	return &ToDos{
		dao:    dao,
		logger: logger,
	}
}

func (t *ToDos) GetAll() ([]model.ToDo, error) {
	todos := []model.ToDo{}
	records, err := t.dao.FindRecordsByExpr("todos")
	fmt.Printf("%v", records)
	if err != nil {
		t.logger.Error("could not fetch from todos collection", "error", err.Error())
		return todos, errors.New("failed to fetch ToDos")
	}

	for _, record := range records {
		todos = append(todos, modelFromRecord(record))
	}
	return todos, nil
}

func (t *ToDos) SwapDone(id string) (model.ToDo, error) {
	record, err := t.dao.FindRecordById("todos", id)
	if err != nil {
		t.logger.Error("could not fetch ToDo", "id", id, "error", err.Error())
		return model.ToDo{}, fmt.Errorf("ToDo %s not found", id)
	}

	record.Set("done", !record.GetBool("done"))

	if err := t.dao.SaveRecord(record); err != nil {
		t.logger.Error("could not swap done for ToDo", "id", id, "error", err.Error())
		return model.ToDo{}, errors.New("failed to swap ToDo")
	}
	return modelFromRecord(record), nil
}

func (t *ToDos) Delete(id string) error {
	record, err := t.dao.FindRecordById("todos", id)
	if err != nil {
		t.logger.Error("could not fetch ToDo", "id", id, "error", err.Error())
		return fmt.Errorf("ToDo %s not found", id)
	}

	t.dao.DeleteRecord(record)
	return nil
}

func (t *ToDos) Add(content string) (model.ToDo, error) {
	coll, err := t.dao.FindCollectionByNameOrId("todos")
	if err != nil {
		t.logger.Error("could not fetch todos collection", "error", err.Error())
		return model.ToDo{}, errors.New("failed to add ToDo")
	}
	record := models.NewRecord(coll)
	record.Load(map[string]interface{}{
		"content": content,
		"done":    false,
	})
	if err = t.dao.SaveRecord(record); err != nil {
		t.logger.Error("Could not save ToDo", "error", err.Error())
		return model.ToDo{}, errors.New("failed to add ToDo")
	}
	return modelFromRecord(record), nil
}

func modelFromRecord(record *models.Record) model.ToDo {
	return model.ToDo{
		Id:      record.GetString("id"),
		Content: record.GetString("content"),
		Done:    record.GetBool("done"),
	}
}
