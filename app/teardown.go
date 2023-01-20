package app

import (
	"fmt"

	"github.com/bmdavis419/go-backend-template/database"
	"github.com/gofiber/fiber/v2"
)

func Teardown(server *fiber.App) {
	fmt.Println("Stopping server...")
	server.Shutdown() // stop server before database!!!
	fmt.Println("Closing database connection...")
	database.CloseMongoDB()
}
