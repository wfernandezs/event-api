package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/db"
	"github.com/wfernandez/rest-api/routes"
	"github.com/wfernandez/rest-api/utils"
)

func main() {
	godotenv.Load() 
	utils.InitLogger()
	logger := utils.GetLogger()
	port := utils.GetEnv("PORT", "8080")

	logger.Info().Msg("Starting application")
	
	db.InitDB() // Initialize the database
	logger.Info().Msg("Database initialized")
	
	server := gin.Default()
	
	// Apply the error handler middleware globally
	server.Use(utils.ErrorHandlerMiddleware())
	
	routes.RegisterRoutes(server) // Register the routes
	logger.Info().Msg("Routes registered")
	
	logger.Info().Str("port", port).Msg("Server starting")
	server.Run(fmt.Sprintf(":%s", port))
}