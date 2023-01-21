package router

import (
	"github.com/bmdavis419/go-backend-template/database"
	"github.com/bmdavis419/go-backend-template/handlers"
	"github.com/bmdavis419/go-backend-template/services"
	"github.com/gofiber/fiber/v2"
)

func AttachRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todosHandler := getTodosHandler()
	todos := app.Group("/todos")
	todos.Post("/", todosHandler.CreateTodo)
	todos.Get("/", todosHandler.GetAllTodos)
	todos.Get("/:id", todosHandler.GetTodoById)
	todos.Put("/:id", todosHandler.UpdateTodo)
	todos.Delete("/:id", todosHandler.DeleteTodo)
}

func getTodosHandler() handlers.TodosHandlerInterface {
	repository := database.NewTodosRepository()
	service := services.NewTodosService(repository)
	return handlers.NewTodoHandler(service)
}
