package controllers

import (
	"car-sharing-platform/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateInvoice - POST /billing/invoice
func CreateInvoice(c *gin.Context) {
	var invoiceData struct {
		UserID     int       `json:"user_id"`
		VehicleID  int       `json:"vehicle_id"`
		StartTime  time.Time `json:"start_time"`
		EndTime    time.Time `json:"end_time"`
		TotalPrice float64   `json:"total_price"`
	}

	// Bind the incoming request body to invoiceData
	if err := c.ShouldBindJSON(&invoiceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create a new invoice entry
	query := `INSERT INTO Invoices (UserID, VehicleID, StartTime, EndTime, TotalPrice)
	          VALUES (?, ?, ?, ?, ?)`

	// Execute the SQL query to insert the invoice data
	if err := database.DB.Exec(query, invoiceData.UserID, invoiceData.VehicleID, invoiceData.StartTime, invoiceData.EndTime, invoiceData.TotalPrice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice created successfully"})
}

