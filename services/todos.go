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

func NewTodosService(log *slog.Logger, gateway gateways.SqlLiteGateway) TodosService {
	return TodosService{
		Log: log,
		DbGateway: gateway,
	}
}

func (t TodosService) GetAllTodos() ([]*model.Todo, error)  {

	todos, _ := t.DbGateway.GetAllTodos()

	return todos, nil
}