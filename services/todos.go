package services

import (
	"github.com/loganphillips792/golang-todo/model"
	"log/slog"
)

type Todos []*model.Todo

type TodosService struct {
	Log        *slog.Logger
}

func NewTodosService(log *slog.Logger) TodosService {
	return TodosService{
		Log: log,
	}
}

func (t TodosService) GetAllTodos() ([]*model.Todo, error)  {
	return []*model.Todo{
		{
			Id: 5,
			Text: "This is the first item on our list",
		},
	}, nil
}