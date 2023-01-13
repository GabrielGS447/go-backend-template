package app

import (
	"github.com/bmdavis419/go-backend-template/config"
	"github.com/bmdavis419/go-backend-template/database"
	"github.com/bmdavis419/go-backend-template/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Setup() (*fiber.App, error) {
	// load env
	err := config.LoadENV()
	if err != nil {
		return nil, err
	}

	// start database
	err = database.StartMongoDB()
	if err != nil {
		return nil, err
	}

	// defer closing database
	defer database.CloseMongoDB()

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
	config.AddSwaggerRoutes(app)

	return app, nil
}
