package controllers

import (
	"car-sharing-platform/database"
	"car-sharing-platform/vehicle-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAvailableVehicles - GET /vehicles
func GetAvailableVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	if err := database.DB.Where("status = ?", "Available").Find(&vehicles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vehicles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"vehicles": vehicles})
}
