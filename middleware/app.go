package middleware

import (
	"astroauth-api/database"
	"astroauth-api/models"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CheckApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rUser models.AppUser
		c.ShouldBindJSON(&rUser)

		var app models.App

		if err := database.DB.Where("app_id=?", rUser.AppID).First(&app).Error; err != nil {
			c.JSON(200, models.Error{Message: "Application not found"})
			c.Abort()
		}

	}
}
