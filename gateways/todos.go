package gateways

import (
	"github.com/loganphillips792/golang-todo/model"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

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

func(gateway TodosGateway) GetAllTodos() []*model.Todo {
	gateway.Logger.Info("GetAllTodos")

	query := "SELECT * FROM todos"

	gateway.Logger.Info(
		"Running sql statement",
		"SQL", query,
	)

	rows, err := t.DbConn.Query(query)

	if err != nil {
		gateway.Logger.Info(err)
		gateway.Logger.Error(err)
	}

	defer rows.Close()

	var todos []*model.Todo

	for rows.Next() {
		var todo *model.Todo
		err := rows.Scan(&todo.Id, &todo.Text, &todo.Checked)

		if err != nil {
			gateway.Logger.Error(err)
		}
		todos = append(todos, todo)
	}
	return todos
}