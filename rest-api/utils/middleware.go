package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppError represents a structured error that can be returned by handlers
type AppError struct {
	StatusCode int
	Message    string
	Error      error
}

// ErrorHandlerMiddleware catches panics and provides a consistent error response
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Execute subsequent middleware/handler
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				log.Printf("Panic occurred: %v", err)

				// Return a 500 internal server error
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "An unexpected error occurred",
				})

				// Abort the request
				c.Abort()
			}
		}()

		c.Next()
	}
}

// HandleAppError processes structured application errors
func HandleAppError(c *gin.Context, appErr AppError) {
	// Log the error if available
	if appErr.Error != nil {
		log.Printf("Error: %v", appErr.Error)
	}

	// Return JSON response with appropriate status code
	c.JSON(appErr.StatusCode, gin.H{
		"error": appErr.Message,
	})
}