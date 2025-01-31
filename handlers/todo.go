package handlers

import (
	"RestToDo/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var todos []models.Todo
var lastID int

func GetTodos(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusOK,
		todos,
	)
}

func GetTodo(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for _, todo := range todos {
		if todo.ID == id {
			return ctx.JSON(http.StatusOK, todo)
		}
	}
	return ctx.JSON(http.StatusNotFound, "Todo not found")
}

func CreateTodo(ctx echo.Context) error {
	todo := new(models.Todo)
	if err := ctx.Bind(todo); err != nil {
		return err
	}
	lastID++
	todo.ID = lastID
	todos = append(todos, *todo)
	return ctx.JSON(http.StatusCreated, todo)
}

func UpdateTodo(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	updatedTodo := new(models.Todo)
	if err := ctx.Bind(updatedTodo); err != nil {
		return err
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Completed = updatedTodo.Completed
			return ctx.JSON(http.StatusOK, todo)
		}
	}
	return ctx.JSON(http.StatusNotFound, "Todo not found")
}

func DeleteTodo(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return ctx.NoContent(http.StatusNoContent)
		}
	}
	return ctx.JSON(http.StatusNotFound, "Todo not found")
}
