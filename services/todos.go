package services

import (
	"github.com/loganphillips792/golang-todo/model"
	"github.com/loganphillips792/golang-todo/gateways"
	"log/slog"
)

type Todos []*model.Todo

type TodosService struct {
	Log        *slog.Logger
	DbGateway gateways.SqlLiteGateway
}

func NewTodosService(log *slog.Logger) TodosService {
	return TodosService{
		Log: log,
	}
}

func (t TodosService) GetAllTodos() ([]*model.Todo, error)  {

	todos, _ := t.DbGateway.GetAllTodos()

	return todos, nil
}