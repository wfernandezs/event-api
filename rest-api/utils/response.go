package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleError(context *gin.Context, statusCode int, err error) {
	context.JSON(statusCode, gin.H{"error": err.Error()})
}

func JSONResponse(context *gin.Context, statusCode int, payload gin.H) {
	context.JSON(statusCode, payload)
}
