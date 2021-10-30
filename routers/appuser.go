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

	appuser.POST("/register", middleware.AppUserRegisterValidate(), middleware.CheckApp(), AppRegister)
	appuser.POST("/login", middleware.AppUserLoginValidate(), middleware.CheckApp(), middleware.AppBasicAuth(), AppLogin)

}

func AppRegister(c *gin.Context) {
	var User models.AppUser
	c.ShouldBindBodyWith(&User, binding.JSON)

	//get length to set expiry of user, checks if license exists, checks if its used
	var LicenseLength uint
	var LicenseLevel uint
	var LicenseUsedBy uint
	err := database.DBB.QueryRow(context.Background(), "SELECT length, level, used_by FROM licenses WHERE license = $1  AND app_id = $2", User.License, User.AppID).Scan(&LicenseLength, &LicenseLevel, &LicenseUsedBy)
	if err != nil || LicenseUsedBy != 0 {
		c.JSON(200, models.Error{Message: "License invalid"})
		return
	}

	//Check if email is available
	FindEmail, err := database.DBB.Exec(context.Background(), "SELECT email FROM app_users WHERE email = $1  AND app_id = $2", User.Email, User.AppID)
	if err == nil {
		c.JSON(500, models.Error{Message: "Internal server error"})
		return
	}
	if FindEmail.RowsAffected() != 0 {
		c.JSON(200, models.Error{Message: "Email not available"})
		return
	}

	//Check if username is available
	FindName, err := database.DBB.Exec(context.Background(), "SELECT username FROM app_users WHERE username = $1 AND app_id = $2", User.Username, User.AppID)
	if err == nil {
		c.JSON(500, models.Error{Message: "Internal server error"})
		return
	}
	if FindName.RowsAffected() != 0 {
		c.JSON(200, models.Error{Message: "Username not available"})
		return
	}

	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), 8)
	if err != nil {
		c.JSON(500, models.Error{Message: "Internal server error"})
		return
	}
	User.Password = string(hashedPassword)

	User.Expiry = time.Now().Local().Add(time.Hour * 24 * time.Duration(LicenseLength))
	User.Level = LicenseLevel

	//Add user to DB
	database.DB.Create(&User)

	//Set the license to used
	database.DBB.Exec(context.Background(), "UPDATE licenses SET used_by = $1 WHERE license = $2", User.ID, User.License)
}

func AppLogin(c *gin.Context) {

	var username string
	var level uint
	err := database.DBB.QueryRow(context.Background(), "SELECT email, username, level FROM app_users WHERE id = $1", c.MustGet("UserID")).Scan(&username, &level)
	if err != nil {
		c.JSON(500, models.Error{Message: "Internal server error"})
		return
	}

	c.JSON(200, gin.H{
		"user": gin.H{
			"username": username,
			"level":    level,
		},
		"app": gin.H{
			"name": c.MustGet("AppName"),
		},
	})
}
