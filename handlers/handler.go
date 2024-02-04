package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/loganphillips792/golang-todo/model"
	"github.com/loganphillips792/golang-todo/services"
	"github.com/loganphillips792/golang-todo/templates"
	"log/slog"

)

type Handler struct { 
	Log        *slog.Logger
	TodosService *services.TodosService
}

// type TodoService interface {
// 	GetAllTodos() ([]*model.Todo, error)
// }

func NewHandler(log *slog.Logger, t services.TodosService) *Handler {
	return &Handler{
		Log: log,
		TodosService: &t,
	}
}

func (h *Handler) GetAllTodos(c echo.Context) error {

	// get todos
	//todos, _ := h.TodosService.GetAllTodos()

	return nil
}


// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func IndexHandler(c echo.Context) error {
	return Render(c, http.StatusOK, templates.Index([]*model.Todo{}))
}