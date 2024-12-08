package routes

import (
	"car-sharing-platform/vehicle-service/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Vehicle Routes
	router.GET("/vehicles", controllers.GetAvailableVehicles) // Get available vehicles

	// Booking Routes
	router.POST("/bookings", controllers.CreateBooking)    // Create booking
	router.GET("/bookings", controllers.GetBookings)        // Get all bookings
	router.PUT("/bookings/:id", controllers.ModifyBooking)  // Modify booking (update vehicle and start time)
	router.DELETE("/bookings/:id", controllers.CancelBooking) // Cancel booking
	
	// Billing Routes
	router.POST("/billing/invoice", controllers.CreateInvoice) // Create invoice
}
