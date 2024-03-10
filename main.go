package main

import (
	"context"
	"fmt"

	"database/sql"
	"errors"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/loganphillips792/golang-todo/config"
	"github.com/loganphillips792/golang-todo/gateways"
	"github.com/loganphillips792/golang-todo/handlers"
	"github.com/loganphillips792/golang-todo/services"
	"github.com/loganphillips792/golang-todo/templates"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg, configError := config.Init()

	if configError != nil {
		slog.Error("error when reading config file")
	}

	var logger *slog.Logger

	if cfg.AppEnvironment == "development" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	db := initializeDatabase()
	defer db.Close()


	e := echo.New()


	// initialize gateway to pass to service
	gateway := gateways.NewTodosDatabaseGateway(logger, db)

	// initialize service to pass to handlerx
	service := services.NewTodosService(logger, gateway)
	
	handler := handlers.NewHandler(logger, service)

	// handlers.SetupRoutes(e, handler)

	e.GET("/", func(c echo.Context) error {
		todos, _ := handler.TodosService.GetAllTodos()
		component := templates.Index(todos)
		return component.Render(context.Background(), c.Response().Writer)
	})


	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":3000"))
}

// // This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
// func Render(ctx echo.Context, statusCode int, t templ.Component) error {
// 	ctx.Response().Writer.WriteHeader(statusCode)
// 	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
// 	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
// }

// func IndexHandler(c echo.Context) error {
// 	return Render(c, http.StatusOK, templates.Index([]*model.Todo{}))
// }

func initializeDatabase() *sqlx.DB {
	slog.Info("Initializing SQL Lite database...")

	file, openFileErr := os.Open("data.db")

	if openFileErr != nil {
		slog.Info(openFileErr.Error())
	}

	if errors.Is(openFileErr, os.ErrNotExist) {
		file, _ = os.Create("data.db")
	}

	file.Close()

	db, err := sql.Open("sqlite3", "data.db")

	if err != nil {
		slog.Error(err.Error())
	}

	sqlxDb := sqlx.NewDb(db, "sqlite3")

	// create tables and seed data
	if errors.Is(openFileErr, os.ErrNotExist) {
		c, err := os.ReadFile("seed.sql")

		if err != nil {
			slog.Error(err.Error())
		}

		sql := string(c)

		slog.Info(sql)
		
		_, err = db.Exec(sql)

		if err != nil {
			slog.Info(err.Error())
			os.Exit(1)
		}

	}

	return sqlxDb
}