package routes

import (
	"github.com/gin-gonic/gin"
	"example/gingonic/models"
	"strconv"
	"net/http"
	"fmt"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. try again later."}) // 500
		return
	}
	context.JSON(200, events)
	
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could not fetch event. try again later."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Could not fetch event. try again later."})
		return
	}
	context.JSON(200, event)
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

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Invalid event ID"})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Event not found"})
		return
	}

	var updatedEvent models.Event
	if err = context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event data", "details": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	fmt.Println(updatedEvent)
	if err = updatedEvent.Update(); err != nil {
		context.JSON(500, gin.H{"error": "Could not update event", "details": err.Error()})
		return
	}

	// Successfully updated the event
	context.JSON(200, gin.H{"message": "Updated successfully"})
}