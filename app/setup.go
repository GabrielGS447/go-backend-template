package app

import (
	"os"

	"github.com/gabrielgs449/go-backend-template/database"
	"github.com/gabrielgs449/go-backend-template/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Setup() (*echo.Echo, error) {
	if err := loadENV(); err != nil {
		return nil, err
	}

	if err := database.StartMongoDB(); err != nil {
		return nil, err
	}

	server := createServer()

	return server, nil
}

func loadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func createServer() *echo.Echo {
	// creates a new Echo instance
	server := echo.New()

	// attach middlewares
	server.Use(middleware.Recover())
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} - ${method} ${path} - ${latency_human}\n",
	}))

	router.AttachRoutes(server)

	// attach swagger
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	return server
}
