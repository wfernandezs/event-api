package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/utils"
)

func Authenticate(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}
	token := authHeader[7:]
	userId, errToken := utils.VerifyToken(token)

	if errToken != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	context.Set("userId", userId)

	context.Next()

}