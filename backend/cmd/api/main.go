package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/iskandersierra/go-todo/backend/pkg/todos"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	app.GET("/todo", todos.HandleTodoList)
	app.GET("/todo/:id", todos.HandleTodoDetails)
    app.POST("/todo", todos.HandleCreateTodo)
    app.PUT("/todo/:id", todos.HandleUpdateTodo)
    app.PUT("/todo/:id/done", todos.HandleDoneTodo)
    app.PUT("/todo/:id/undone", todos.HandleUndoneTodo)
    app.DELETE("/todo/:id", todos.HandleDeleteTodo)

	err := app.Start(":8080")
	app.Logger.Fatal(err)
}
