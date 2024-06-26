package gateways

import (
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/loganphillips792/golang-todo/model"
)

type SqlLiteGateway interface {
	GetAllTodos() ([]*model.Todo, error)
}

type TodosGateway struct {
	Logger        *slog.Logger
	DbConn *sqlx.DB
}

func NewTodosDatabaseGateway(log *slog.Logger, db *sqlx.DB) TodosGateway {
	return TodosGateway {
		Logger: log,
		DbConn: db,
	}
}

func(gateway TodosGateway) GetAllTodos() ([]*model.Todo, error) {
	gateway.Logger.Info("GetAllTodos")

	query := "SELECT * FROM todos"

	gateway.Logger.Info(
		"Running sql statement",
		"SQL", query,
	)

	rows, err := gateway.DbConn.Query(query)

	if err != nil {
		gateway.Logger.Info(err.Error())
		gateway.Logger.Error(err.Error())
	}

	defer rows.Close()

	var todos []*model.Todo

	for rows.Next() {
		var todo model.Todo
		fmt.Printf("%v\n", todo)
		err := rows.Scan(&todo.Id, &todo.Text, &todo.Checked)

		if err != nil {
			gateway.Logger.Error(err.Error())
		}
		todos = append(todos, todo)
	}
	return todos, nil
}