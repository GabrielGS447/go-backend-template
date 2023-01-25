package router

import (
	"github.com/gabrielgs449/go-backend-template/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func AttachRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)
	app.Get("/swagger/*", swagger.HandlerDefault)

	attachTodosRoutes(app)
}
