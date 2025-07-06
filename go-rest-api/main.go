package main

import (
	"go-rest-api/db"
	"go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// Default() for engine
	server := gin.Default()

	routes.RegisterRoutes(server)

	// Run() to run server to listening to every request
	server.Run(":8080") // localhost:8080
}
