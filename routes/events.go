package routes

import (
	"example/gingonic/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	email, exists := context.Get("email")
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User email not found"})
		return
	}

	userIdInterface, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found"})
		return
	}

	// Asserting the userId is of type float64
	userIdFloat, ok := userIdInterface.(float64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "User ID format is invalid"})
		return
	}

	// Convert float64 to int64 since JSON unmarshalling turns numbers into float64 by default
	userId := int64(userIdFloat)

	fmt.Println("User Email:", email, "UserID:", userId)

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse event data", "details": err.Error()})
		return
	}

	event.UserID = userId
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

	// Assuming GetEventById not only fetches the event but also returns it
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Event not found"})
		return
	}

	// Extract the logged-in user's ID from the context, set by your AuthMiddleware
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract user ID from context"})
		return
	}

	// Assuming the userId from the context is float64, convert it to int64 (as often needed when extracting from JSON)
	loggedInUserId := int64(userId.(float64))

	// Check if the logged-in user ID matches the event's user ID
	if event.UserID != loggedInUserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event"})
		return
	}

	var updatedEvent models.Event
	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(400, gin.H{"error": "Could not parse event data", "details": err.Error()})
		return
	}

	updatedEvent.ID = eventId

	if err := updatedEvent.Update(); err != nil {
		context.JSON(500, gin.H{"error": "Could not update event", "details": err.Error()})
		return
	}

	// Successfully updated the event
	context.JSON(200, gin.H{"message": "Updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message": "Event not found"})
		return
	}

	event.Delete()
	if err != nil {
		context.JSON(400, gin.H{"message": "Couldn't delete event"})
		return
	}

	context.JSON(200, gin.H{"message": "Deleted Successfully"})

}
