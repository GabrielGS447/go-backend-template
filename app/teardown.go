package app

import (
	"context"
	"fmt"
	"time"

	"github.com/gabrielgs449/go-backend-template/database"
	"github.com/gofiber/fiber/v2"
)

func Teardown(server *fiber.App) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// stop server before database, otherwise ongoing requests will lose access to it.
	fmt.Println("Stopping server...")
	if err := server.Shutdown(); err != nil {
		panic(err)
	}

	fmt.Println("Closing database connection...")
	database.CloseMongoDB(ctx)
}
