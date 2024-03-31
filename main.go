package main

import (
	"github.com/gin-gonic/gin"
	"example/gingonic/models"
	"example/gingonic/db"
)

func main() {
	db.InitDB()
	r := gin.Default()
	r.GET("/events", getEvents)
	r.POST("/events", createEvent)
	r.Run()
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch events. try again later."})
		return
	}
	context.JSON(200, events)
	
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
        context.JSON(400, gin.H{"error": "Could not parse data", "details": err.Error()})
        return
    }

	event.UserID = 1
	if err := event.Save(); err != nil {
		context.JSON(500, gin.H{"message": "Could not create event. Try again later.", "details": err.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "Event Created", "event": event})
}
