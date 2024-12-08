package main

import (
	"car-sharing-platform/database"
	"car-sharing-platform/vehicle-service/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	// Initialize the database connection
	database.ConnectDatabase()

	// Initialize Gin router
	router := gin.Default()

	// Add routes
	routes.InitializeRoutes(router)

	// Start the server on port 8082
	log.Println("Vehicle-service running on http://localhost:8082")
	router.Run(":8082")
}

func main() {
	// Run the vehicle service
	Start()
}
