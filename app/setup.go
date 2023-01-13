package app

import (
	"os"

	"github.com/bmdavis419/go-backend-template/database"
	"github.com/bmdavis419/go-backend-template/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

func Setup() (*fiber.App, error) {
	// load env
	err := loadENV()
	if err != nil {
		return nil, err
	}

	// start database
	err = database.StartMongoDB()
	if err != nil {
		return nil, err
	}

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	addSwaggerRoute(app)

	return app, nil
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

func addSwaggerRoute(app *fiber.App) {
	// setup swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
}
