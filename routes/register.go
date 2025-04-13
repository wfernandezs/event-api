package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/models"
	"github.com/wfernandez/rest-api/utils"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}
	event, err := models.GetEvent(eventId)

	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	err = event.Register(userId)
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}
	utils.JSONResponse(context, http.StatusOK, gin.H{"message": "Registration successful", "event": event})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}
	event, err := models.GetEvent(eventId)

	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}
	utils.JSONResponse(context, http.StatusOK, gin.H{"message": "Registration cancelled", "event": event})
}
