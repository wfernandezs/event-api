package main

import (
	// Adding missing import
	// Adding missing import

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/db"
	"github.com/wfernandez/rest-api/routes"
)

func main() {
		db.InitDB() // Initialize the database
    server := gin.Default()
		routes.RegisterRoutes(server) // Register the routes
    server.Run(":8080")
}