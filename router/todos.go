package router

import (
	"github.com/gabrielgs449/go-backend-template/database"
	"github.com/gabrielgs449/go-backend-template/handlers"
	"github.com/gabrielgs449/go-backend-template/services"
	"github.com/labstack/echo/v4"
)

func attachTodosRoutes(server *echo.Echo) {
	todosHandler := makeTodosHandler()
	todos := server.Group("/todos")
	todos.POST("", todosHandler.CreateTodo)
	todos.GET("", todosHandler.GetAllTodos)
	todos.GET("/:id", todosHandler.GetTodoById)
	todos.PATCH("/:id", todosHandler.UpdateTodo)
	todos.DELETE("/:id", todosHandler.DeleteTodo)
}

func makeTodosHandler() handlers.TodosHandlerInterface {
	repository := database.NewTodosRepository()
	service := services.NewTodosService(repository)
	return handlers.NewTodoHandler(service)
}
