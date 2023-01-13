package main

import (
	"log"
	"os"

	"github.com/bmdavis419/go-backend-template/app"
	_ "github.com/bmdavis419/go-backend-template/docs"
)

// @title The Better Backend Template
// @version 0.1
// @description An example template of a Golang backend API using Fiber and MongoDB
// @contact.name Ben Davis
// @license.name MIT
// @host localhost:8080
// @BasePath /
func main() {
	// defer app teardown
	defer app.Teardown()

	// setup app
	server, err := app.Setup()
	if err != nil {
		panic(err)
	}

	// get the port and start server
	port := os.Getenv("PORT")
	log.Fatal(server.Listen(":" + port))
}
