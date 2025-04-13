package utils

import (
	"github.com/gin-gonic/gin"
)

// HandleError creates an AppError and processes it
func HandleError(context *gin.Context, statusCode int, err error) {
	appErr := AppError{
		StatusCode: statusCode,
		Message:    err.Error(),
		Error:      err,
	}
	HandleAppError(context, appErr)
}

// JSONResponse sends a standard JSON response with the given status code and payload
func JSONResponse(context *gin.Context, statusCode int, payload gin.H) {
	context.JSON(statusCode, payload)
}

// NewAppError creates a new application error with the provided details
func NewAppError(statusCode int, message string, err error) AppError {
	return AppError{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}
}
