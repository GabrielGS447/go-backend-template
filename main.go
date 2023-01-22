package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabrielgs449/go-backend-template/app"
	_ "github.com/gabrielgs449/go-backend-template/docs"
	"github.com/gofiber/fiber/v2"
)

// @title The Better Backend Template
// @version 0.1
// @description An example template of a Golang backend API using Fiber and MongoDB
// @contact.name Ben Davis
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	// setup app
	server, err := app.Setup()
	if err != nil {
		panic(err)
	}

	// start server in a goroutine so it doesn't block
	go start(server)

	// wait for a signal to gracefully shutdown
	waitForShutdownSignal()

	// teardown app
	app.Teardown(server)

	fmt.Println("Application gracefully stopped.")

	os.Exit(0)
}

func start(server *fiber.App) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "60384"
	}

	if err := server.Listen(":" + port); err != nil {
		panic(err)
	}
}

func waitForShutdownSignal() {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
