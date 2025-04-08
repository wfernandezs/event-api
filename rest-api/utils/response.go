package utils

import (
	"github.com/gin-gonic/gin"
)


func HandleError(context *gin.Context, statusCode int, err error) {
    errorMessage := "An error occurred"
    
    if err != nil {
        errorMessage = err.Error()
    }
    
    context.JSON(statusCode, gin.H{"error": errorMessage})
}

func JSONResponse(context *gin.Context, statusCode int, payload gin.H) {
	context.JSON(statusCode, payload)
}
