package routes

import (
	middleware "example/gingonic/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterdRoutes(server *gin.Engine) {
	protected := server.Group("/").Use(middleware.AuthMiddleware())
	server.POST("/signup", signup)
	server.POST("/login", loginHandler)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	protected.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

}
