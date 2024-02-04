package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo, handler *Handler) {
	e.GET("/", handler.GetAllTodos)
}