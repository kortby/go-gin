package main

import (
	"github.com/gin-gonic/gin"
	"example/gingonic/db"
	"example/gingonic/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterdRoutes(server)
	server.Run(":8080")
}
