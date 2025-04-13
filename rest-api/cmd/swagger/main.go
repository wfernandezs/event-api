package main

import (
	"log"
	"os"

	"github.com/wfernandez/rest-api/swagger"
)

// A standalone program to run the Swagger documentation server
func main() {
	swaggerPort := os.Getenv("SWAGGER_PORT")
	if swaggerPort == "" {
		swaggerPort = "8081"
	}

	// Start the swagger server
	log.Fatal(swagger.StartServer(swaggerPort))
}