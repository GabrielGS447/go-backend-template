package router

import (
	"github.com/bmdavis419/go-backend-template/handlers"
	"github.com/bmdavis419/go-backend-template/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todos := app.Group("/todos")
	todos.Post("/", handlers.CreateTodo)
	todos.Get("/", handlers.GetAllTodos)
	todos.Get("/:id", middlewares.ParseObjectId, handlers.GetTodoById)
	todos.Put("/:id", middlewares.ParseObjectId, handlers.UpdateTodo)
	todos.Delete("/:id", middlewares.ParseObjectId, handlers.DeleteTodo)
}
