package routes

import (
	"car-sharing-platform/user-service/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Public Routes
	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)
	router.GET("/auth/logout", controllers.Logout)

	// Protected Routes
	protected := router.Group("/")
	protected.Use(controllers.AuthMiddleware())
	{
		protected.GET("/account", controllers.GetAccount)
		protected.PUT("/account", controllers.UpdateProfile)
		protected.GET("/rental-history", controllers.GetRentalHistory)

	}
}
