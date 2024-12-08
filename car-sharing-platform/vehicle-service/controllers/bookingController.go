package controllers

import (
	"car-sharing-platform/database"
	"car-sharing-platform/vehicle-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateBooking - POST /bookings (Using Raw SQL)
func CreateBooking(c *gin.Context) {
	var bookingData struct {
		UserID     int       `json:"user_id"`    // Accept userID as part of the input
		VehicleID  int       `json:"vehicle_id"` // Accept vehicleID as part of the input
		StartTime  time.Time `json:"start_time"` // Accept start_time from the input
		EndTime    time.Time `json:"end_time"`   // Accept end_time from the input
		TotalPrice float64   `json:"total_price"` // Accept total_price from the input
	}

	// Bind the incoming request body to bookingData
	if err := c.ShouldBindJSON(&bookingData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Raw SQL Query for Inserting Booking
	query := `
		INSERT INTO Bookings (UserID, VehicleID, StartTime, EndTime, TotalPrice)
		VALUES (?, ?, ?, ?, ?);
	`

	// Execute the Raw SQL Query
	result := database.DB.Exec(query, bookingData.UserID, bookingData.VehicleID, bookingData.StartTime, bookingData.EndTime, bookingData.TotalPrice)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking", "details": result.Error.Error()})
		return
	}

	// Respond with Success
	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully"})
}


// GetBookings - GET /bookings (Retrieve all bookings)
func GetBookings(c *gin.Context) {
	var bookings []models.Booking
	if err := database.DB.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// ModifyBooking - PUT /bookings/:id
func ModifyBooking(c *gin.Context) {
	// Get the booking ID from the URL parameter
	bookingID := c.Param("id")

	var bookingData struct {
		VehicleID  int       `json:"vehicle_id"`
		StartTime  time.Time `json:"start_time"`
	}

	// Bind the incoming request body to bookingData
	if err := c.ShouldBindJSON(&bookingData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Use raw SQL to update the booking
	query := `UPDATE Bookings SET VehicleID = ?, StartTime = ? WHERE ID = ?`
	if err := database.DB.Exec(query, bookingData.VehicleID, bookingData.StartTime, bookingID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to modify booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking modified successfully"})
}

// CancelBooking - DELETE /bookings/:id
func CancelBooking(c *gin.Context) {
	// Get the booking ID from the URL parameter
	bookingID := c.Param("id")

	// Fetch the current booking details to check the start time
	var booking models.Booking
	if err := database.DB.Where("ID = ?", bookingID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	// Use raw SQL to delete the booking
	query := `DELETE FROM Bookings WHERE ID = ?`
	if err := database.DB.Exec(query, bookingID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking cancelled successfully"})
}