package main

import (
	"net/http" // Adding missing import

	"github.com/gin-gonic/gin"
	"github.com/wfernandez/rest-api/db"
	"github.com/wfernandez/rest-api/models"
)

func main() {
		db.InitDB() // Initialize the database
    server := gin.Default()
    server.GET("/events", getEvents)
    
		server.POST("/events", createEvent)
    server.Run(":8080")
}

func getEvents(context *gin.Context) {
	  events := models.GetAllEvents()
		context.JSON(http.StatusOK, gin.H{
				"events": events,
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
		event.UserID = 1
		event.ID = len(models.GetAllEvents()) + 1
		event.Save()
		context.JSON(http.StatusCreated, gin.H{"event": event})
}