package main

import (
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default() for engine
	server := gin.Default()

	// setting handler for GET request
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// Run() to run server to listening to every request
	server.Run(":8080") // localhost:8080
}

// context comes from handler function / controller
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	// gin.H is MAP
	// context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// shouldBindJSON = fmt.Scan
	// store body to event variable = request body must require same with Event model Sructure
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
