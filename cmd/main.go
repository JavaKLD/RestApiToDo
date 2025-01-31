package main

import (
	"RestToDo/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todos", handlers.GetTodos)
	e.GET("/todos/:id", handlers.GetTodo)
	e.POST("/todos", handlers.CreateTodo)
	e.PUT("/todos/:id", handlers.UpdateTodo)
	e.DELETE("/todos/:id", handlers.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
