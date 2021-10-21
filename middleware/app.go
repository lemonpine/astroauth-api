package middleware

import (
	"astroauth-api/database"
	"astroauth-api/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rUser models.AppUser
		c.ShouldBindJSON(&rUser)

		//Check if basic auth authorization header is present
		username, password, err := c.Request.BasicAuth()
		if !err {
			c.JSON(401, models.Error{Message: "Unauthorized"})
			c.Abort()
			return
		}

		if username == "" || password == "" {
			c.JSON(200, models.Error{Message: "Username or password incorrect"})
			c.Abort()
			return
		}

		//Check if user exists
		var DBUser models.AppUser
		if err := database.DB.Where("username=? AND app_id=?", username, rUser.AppID).First(&DBUser).Error; err != nil {
			c.JSON(200, models.Error{Message: "Username or password incorrect"})
			c.Abort()
			return
		}

		//Check password
		err2 := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(password))
		if err2 != nil {
			c.JSON(200, models.Error{Message: "Username or password incorrect"})
			c.Abort()
			return
		}
	}
}

func CheckApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var rUser models.AppUser
		c.ShouldBindJSON(&rUser)

		var app models.App

		if err := database.DB.Where("app_id=?", rUser.AppID).First(&app).Error; err != nil {
			c.JSON(404, models.Error{Message: "Application not found"})
			c.Abort()
			return
		}

	}
}
