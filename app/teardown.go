package app

import (
	"context"
	"fmt"
	"time"

	"github.com/gabrielgs449/go-backend-template/database"
	"github.com/labstack/echo/v4"
)

func Teardown(server *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// stop server before database, otherwise ongoing requests will lose access to it.
	fmt.Println("Stopping server...")
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Closing database connection...")
	database.CloseMongoDB(ctx)
}
