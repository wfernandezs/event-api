package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type AppError struct {
	StatusCode int
	Message    string
	Error      error
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger := GetLogger()
				
				var stackTracer stackTracer
				var errObj error
				
				switch e := err.(type) {
				case error:
					errObj = e
					logger.Error().Err(e).Msg("Panic occurred")
				default:
					errObj = fmt.Errorf("%v", err)
					logger.Error().Interface("panic", err).Msg("Panic occurred")
				}
				
				if errors.As(errObj, &stackTracer) {
					logger.Error().Msg(fmt.Sprintf("Stack trace:\n%+v", stackTracer))
				}

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "An unexpected error occurred",
				})

				c.Abort()
			}
		}()

		c.Next()
	}
}
type stackTracer interface {
	StackTrace() errors.StackTrace
}

func HandleAppError(c *gin.Context, appErr AppError) {
	logger := GetLogger()
	
	if appErr.Error != nil {
		var stackTracer stackTracer
		if errors.As(appErr.Error, &stackTracer) {
			logger.Error().Err(appErr.Error).Msg(appErr.Message)
			logger.Error().Msg(fmt.Sprintf("Stack trace:\n%+v", stackTracer))
		} else {
			logger.Error().Err(appErr.Error).Msg(appErr.Message)
		}
	}

	c.JSON(appErr.StatusCode, gin.H{
		"error": appErr.Message,
	})
}