package routers

import (
	"astroauth-api/database"
	"astroauth-api/middleware"
	"astroauth-api/models"
	"time"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func AppUserRouter(router *gin.Engine) {
	appuser := router.Group("/app")

	appuser.POST("/register", middleware.CheckApp(), AppRegister)

	appuser.Use(middleware.CheckApp(), middleware.AppBasicAuth())
	{
		appuser.POST("/login", AppLogin)
	}
}

func AppRegister(c *gin.Context) {
	var rUser models.AppUser
	c.ShouldBindBodyWith(&rUser, binding.JSON)

	//Validate user input
	err := rUser.Validate()
	if err != nil {
		c.JSON(200, gin.H{"error": err})
		return
	}

	if rUser.License == "" || rUser.License == " " {
		c.JSON(200, gin.H{
			"message": "license missing",
		})
		return
	}

	//get length to set expiry of user, checks if license exists, checks if its used
	var LicenseLength uint
	var LicenseLevel uint
	var LicenseUsedBy uint
	err = database.DBB.QueryRow(context.Background(), "SELECT length, level, used_by FROM licenses WHERE license = $1  AND app_id = $2", rUser.License, rUser.AppID).Scan(&LicenseLength, &LicenseLevel, &LicenseUsedBy)
	if err != nil || LicenseUsedBy != 0 {
		c.JSON(200, models.Error{Message: "License invalid"})
		return
	}

	//Check if email is available
	var email string
	err = database.DBB.QueryRow(context.Background(), "SELECT email FROM app_users WHERE email = $1  AND app_id = $2", rUser.Email, rUser.AppID).Scan(&email)
	if err == nil {
		c.JSON(200, models.Error{Message: "Email not available"})
		return
	}

	//Check if username is available
	var username string
	err = database.DBB.QueryRow(context.Background(), "SELECT username FROM app_users WHERE username = $1 AND app_id = $2", rUser.Username, rUser.AppID).Scan(&username)
	if err == nil {
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

	rUser.Expiry = time.Now().Local().Add(time.Hour * 24 * time.Duration(LicenseLength))
	rUser.Level = LicenseLevel
	//Add user to DB
	database.DB.Create(&rUser)

	//Set the license to used
	database.DBB.QueryRow(context.Background(), "UPDATE licenses SET used_by = $1 WHERE license = $2", rUser.ID, rUser.License).Scan(&email)
}

func AppLogin(c *gin.Context) {
	var rUser models.AppUser
	c.ShouldBindJSON(&rUser)

	// err := rUser.Validate()
	// if err != nil {
	// 	return
	// }
	c.JSON(200, gin.H{"error": "logged"})

}
