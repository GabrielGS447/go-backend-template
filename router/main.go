package router

import (
	"github.com/gabrielgs449/go-backend-template/handlers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func AttachRoutes(server *echo.Echo) {
	server.GET("/health", handlers.HandleHealthCheck)
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	attachTodosRoutes(server)
}
