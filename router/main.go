package router

import (
	"github.com/bmdavis419/go-backend-template/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todos := app.Group("/todos")
	todos.Post("/", handlers.CreateTodo)
	todos.Get("/", handlers.GetAllTodos)
	todos.Get("/:id", handlers.GetTodoById)
	todos.Put("/:id", handlers.UpdateTodo)
	todos.Delete("/:id", handlers.DeleteTodo)
}
