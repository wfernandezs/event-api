package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/models"
	"github.com/wfernandez/rest-api/utils"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}
	utils.JSONResponse(context, http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
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
	utils.JSONResponse(context, http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}
	utils.JSONResponse(context, http.StatusCreated, gin.H{"event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEvent(eventId)
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	if event.UserID != userId {
		utils.HandleError(context, http.StatusForbidden, err)
		return
	}


	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}
	utils.JSONResponse(context, http.StatusOK, gin.H{"event": updateEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.HandleError(context, http.StatusBadRequest, err)
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEvent(eventId)


	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}

	if event.UserID != userId {
		utils.HandleError(context, http.StatusForbidden, err)
		return
	}	

	var deleteEvent models.Event
	deleteEvent.ID = eventId
	err = deleteEvent.Delete()
	if err != nil {
		utils.HandleError(context, http.StatusInternalServerError, err)
		return
	}
	utils.JSONResponse(context, http.StatusOK, gin.H{"message": "Deleted"})
}