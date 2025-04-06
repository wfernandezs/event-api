package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/models"
)

func getEvents(context *gin.Context) {
	  events, err := models.GetAllEvents()
		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}
		context.JSON(http.StatusOK, gin.H{
				"events": events,
		})
}

func getEvent(context *gin.Context) {
		eventId, err := strconv.ParseInt((context.Param("id")), 10, 64)
		if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
						"error": "Invalid event ID",
				})
				return
		}
	  event, err := models.GetEvent(eventId)
		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}
		context.JSON(http.StatusOK, gin.H{
				"event": event,
		})
}

func createEvent(context *gin.Context) {
		var event models.Event
		err := context.ShouldBindJSON(&event)
		if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
				})
				return
		}

		err = event.Save()
		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}
		context.JSON(http.StatusCreated, gin.H{"event": event})
}

func updateEvent(context *gin.Context) {
		eventId, err := strconv.ParseInt((context.Param("id")), 10, 64)
		if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
						"error": "Invalid event ID",
				})
				return
		}

		_ , err = models.GetEvent(eventId)
		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}

		var updateEvent models.Event
		err = context.ShouldBindJSON(&updateEvent)

		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}

		updateEvent.ID = eventId
		err = updateEvent.Update()

		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}		
		context.JSON(http.StatusOK, gin.H{
				"event": updateEvent,
		})
}

func deleteEvent(context *gin.Context) {
		eventId, err := strconv.ParseInt((context.Param("id")), 10, 64)
		if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
						"error": "Invalid event ID",
				})
				return
		}

		_ , err = models.GetEvent(eventId)

		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}

		var deleteEvent models.Event
		deleteEvent.ID = eventId
		err = deleteEvent.Delete()

		if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
				})
				return
		}		
		context.JSON(http.StatusOK, gin.H{
				"message": "Deleted",
		})
}