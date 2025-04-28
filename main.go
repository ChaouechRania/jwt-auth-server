package main

import (
	"jwt-auth-server/config"
	"jwt-auth-server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load environment variables
	config.LoadConfig()
	// Connect to the database
	config.ConnectDatabase()

	// Set up the router
	r := gin.Default()
	routes.SetupRouter(r)

	r.Run(":8080")

}
