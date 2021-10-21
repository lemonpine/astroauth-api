package routers

import (
	"astroauth-api/database"
	"astroauth-api/middleware"
	"astroauth-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AppUserRouter(router *gin.Engine) {
	appuser := router.Group("/app")

	appuser.POST("/register", middleware.CheckApp(), AppRegister)

	appuser.Use(middleware.CheckApp(), middleware.BasicAuth())
	{
		appuser.POST("/login", AppLogin)
	}
}

func AppRegister(c *gin.Context) {
	var rUser models.AppUser
	c.ShouldBindJSON(&rUser)

	//Validate user input

	if rUser.License == "" || rUser.License == " " {
		c.JSON(200, gin.H{
			"message": "license missing",
		})
		return
	}

	//Check if license exists
	var license models.License
	license.AppID = rUser.AppID
	license.License = rUser.License
	if err := database.DB.Where("license=? AND app_id=?", license.License, license.AppID).First(&license).Error; err != nil {
		c.JSON(200, models.Error{Message: "License invalid"})
		return
	}

	//Check if email is available
	if err := database.DB.Where("email=? AND app_id=?", rUser.Email, rUser.AppID).First(&rUser).Error; err == nil {
		c.JSON(200, models.Error{Message: "Email not available"})
		return
	}

	//Check if username is availabe
	if err := database.DB.Where("username=? AND app_id=?", rUser.Username, rUser.AppID).First(&rUser).Error; err == nil {
		c.JSON(200, models.Error{Message: "Username not available"})
		return
	}
	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rUser.Password), 8)
	if err != nil {
		c.JSON(500, models.Error{Message: "Internal server error"})
		return
	}
	rUser.Password = string(hashedPassword)

	//Add user to DB
	database.DB.Create(&rUser)
}

func AppLogin(c *gin.Context) {
	var rUser models.AppUser
	c.ShouldBindJSON(&rUser)

}
