package controllers

import (
	"car-sharing-platform/database"
	"car-sharing-platform/user-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccount - Fetch account details
func GetAccount(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":           user.Name,
		"Email":          user.Email,
		"Phone":          user.Phone,
		"MembershipTier": user.MembershipTier,
		"TotalSpending":  user.TotalSpending,
	})
}

// UpdateProfile - Update user profile
func UpdateProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

    var input struct {
        Name  string `form:"name" binding:"required"`
        Phone string `form:"phone" binding:"required"`
    }

    if err := c.ShouldBind(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func GetRentalHistory(c *gin.Context) {
    session, err := store.Get(c.Request, "user-session")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
        return
    }

    userID := session.Values["userID"]
    if userID == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var rentalHistory []models.RentalHistory
    err = database.DB.Where("UserID = ?", userID).Find(&rentalHistory).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve rental history"})
        return
    }

    c.JSON(http.StatusOK, rentalHistory)
}
