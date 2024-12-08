package main

import (
	"car-sharing-platform/database"
	"car-sharing-platform/user-service/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Start initializes and runs the user service
func Start() {
	// Initialize the database connection
	database.ConnectDatabase()

	// Initialize Gin router
	router := gin.Default()

	// Initialize routes
	routes.InitializeRoutes(router)

	// Start the server on port 8081
	log.Println("User-service running on http://localhost:8081")
	router.Run(":8081")
}

// main is used to run the service standalone
func main() {
	// Check if the service is running as standalone
	if len(os.Args) == 1 {
		log.Println("Running user-service in standalone mode.")
		Start()
	}
}
