package router

import (
	"github.com/gabrielgs449/go-backend-template/database"
	"github.com/gabrielgs449/go-backend-template/handlers"
	"github.com/gabrielgs449/go-backend-template/services"
	"github.com/gofiber/fiber/v2"
)

func attachTodosRoutes(app *fiber.App) {
	todosHandler := makeTodosHandler()
	todos := app.Group("/todos")
	todos.Post("/", todosHandler.CreateTodo)
	todos.Get("/", todosHandler.GetAllTodos)
	todos.Get("/:id", todosHandler.GetTodoById)
	todos.Patch("/:id", todosHandler.UpdateTodo)
	todos.Delete("/:id", todosHandler.DeleteTodo)
}

func makeTodosHandler() handlers.TodosHandlerInterface {
	repository := database.NewTodosRepository()
	service := services.NewTodosService(repository)
	return handlers.NewTodoHandler(service)
}
