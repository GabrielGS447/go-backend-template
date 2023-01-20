package app

import (
	"fmt"

	"github.com/bmdavis419/go-backend-template/database"
	"github.com/gofiber/fiber/v2"
)

func Teardown(server *fiber.App) {
	// stop server before database, otherwise ongoing requests will lose access to it.
	fmt.Println("Stopping server...")
	if err := server.Shutdown(); err != nil {
		panic(err)
	}

	fmt.Println("Closing database connection...")
	database.CloseMongoDB()
}
