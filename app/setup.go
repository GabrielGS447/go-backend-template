package app

import (
	"os"

	"github.com/gabrielgs449/go-backend-template/database"
	"github.com/gabrielgs449/go-backend-template/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

func Setup() (*fiber.App, error) {
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

func createServer() *fiber.App {
	// creates a new Fiber instance
	server := fiber.New()

	// attach middlewares
	server.Use(recover.New())
	server.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path} - ${latency}\n",
	}))

	router.AttachRoutes(server)

	// attach swagger
	server.Get("/swagger/*", swagger.HandlerDefault)

	return server
}
