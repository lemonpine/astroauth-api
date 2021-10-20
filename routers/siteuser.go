package routers

import (
	"astroauth-api/database"
	"astroauth-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SiteUserRouter(router *gin.Engine) {
	siteuser := router.Group("/site")
	{
		siteuser.POST("/register", SiteRegister)
		siteuser.POST("/login", SiteLogin)
	}
}

func SiteRegister(c *gin.Context) {
	var rUser models.SiteUser

	c.ShouldBindJSON(&rUser)

	//Check if email is available
	if err := database.DB.Where("email=?", rUser.Email).First(&rUser).Error; err == nil {
		c.JSON(200, models.Error{Message: "Email not available"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rUser.Password), 8)
	if err != nil {
		c.JSON(200, models.Error{Message: "Internal server error"})
		return
	}
	rUser.Password = string(hashedPassword)

	//Add user to DB
	database.DB.Create(&rUser)
}

func SiteLogin(c *gin.Context) {
	var rUser models.SiteUser
	c.ShouldBindJSON(&rUser)

	var DBUser models.SiteUser

	//Check if email exists
	if err := database.DB.Where("email=?", rUser.Email).First(&DBUser).Error; err != nil {
		c.JSON(200, models.Error{Message: "Email or password incorrect"})
		return
	}

	//Check password
	if err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(rUser.Password)); err != nil {
		c.JSON(200, models.Error{Message: "Email or password incorrect"})
		return
	}

	//If email and password is correct, send session
	session, _ := database.Store.Get(c.Request, "session")
	session.Values["userID"] = DBUser.ID

	session.Save(c.Request, c.Writer)
	c.JSON(200, gin.H{
		"message": "logged in",
	})
}
