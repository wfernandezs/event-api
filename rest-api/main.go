package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/db"
	"github.com/wfernandez/rest-api/models"
	"github.com/wfernandez/rest-api/routes"
	"github.com/wfernandez/rest-api/utils"
)

func main() {
	utils.LoadEnv()
	utils.InitLogger()
	logger := utils.GetLogger()
	port := utils.GetEnv("PORT", "8080")

	logger.Info().Msg("Starting application")
	
	// Initialize database
	gormDB := db.GetInstance().DB
	
	// List all models being migrated
	modelNames := models.ListRegisteredModels()
	logger.Info().Msgf("Migrating models: %s", strings.Join(modelNames, ", "))
	
	// Auto-migrate all registered models
	if err := gormDB.AutoMigrate(models.GetRegisteredModels()...); err != nil {
		logger.Fatal().Err(err).Msg("Failed to run migrations")
	}
	
	logger.Info().Msg("Database migrations completed successfully")
	
	server := gin.Default()
	
	// Apply the error handler middleware globally
	server.Use(utils.ErrorHandlerMiddleware())
	
	routes.RegisterRoutes(server) // Register the routes
	logger.Info().Msg("Routes registered")
	
	logger.Info().Str("port", port).Msg("Server starting")
	server.Run(fmt.Sprintf(":%s", port))
}