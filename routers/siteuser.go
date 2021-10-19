package routers

import (
	"astroauth-api/database"
	"astroauth-api/middleware"
	"astroauth-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SiteUserRouter(router *gin.Engine) {
	siteuser := router.Group("/site")
	{
		siteuser.POST("/register", SiteRegister)
		siteuser.POST("/login", SiteLogin)
		siteuser.POST("/protected", middleware.SessionMiddleware(), ProtectedRoute)

	}
}

func ProtectedRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "SESSION ENDPOINT",
	})
}

func SiteRegister(c *gin.Context) {
	var rUser models.User

	c.ShouldBindJSON(&rUser)

	//Check if email is available
	if err := database.DB.Where("email=?", rUser.Email).First(&rUser).Error; err == nil {
		c.JSON(200, gin.H{
			"message": "Email not available",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rUser.Password), 8)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "password hash error",
		})
		panic("HASH ERROR")
	}
	rUser.Password = string(hashedPassword)

	//Add user to DB
	database.DB.Create(&rUser)
}

func SiteLogin(c *gin.Context) {
	var rUser models.User
	c.ShouldBindJSON(&rUser)

	//Check user credentials
	var DBUser models.User

	//If any information is incorrect, "incorrect email address or password" will be returned.

	//Check if email exists
	if err := database.DB.Where("email=?", rUser.Email).First(&DBUser).Error; err != nil {
		c.JSON(200, gin.H{
			"message": "email inco",
		})
		return
	}

	//Check password
	err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(rUser.Password))
	if err != nil { //Password does not match!
		c.JSON(200, gin.H{
			"message": "password incorrect",
		})
		return
	}

	//If email and password is correct
	session, _ := database.Store.Get(c.Request, "session")
	session.Values["userID"] = DBUser.ID

	session.Save(c.Request, c.Writer)
	c.JSON(200, gin.H{
		"message": "logged in",
	})
}
