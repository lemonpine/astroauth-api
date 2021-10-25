package routers

import (
	"astroauth-api/database"
	"astroauth-api/models"
	"context"

	"github.com/gin-gonic/gin/binding"

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

	//Validate input
	err := rUser.Validate()
	if err != nil {
		c.JSON(200, gin.H{"error": err})
		return
	}

	//Check if email is available
	var email string
	err = database.DBB.QueryRow(context.Background(), "SELECT email FROM site_users WHERE email = $1", rUser.Email).Scan(&email)
	if err == nil {
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
	var r models.SiteUser
	c.ShouldBindBodyWith(&r, binding.JSON)

	//Validate input
	err := r.Validate()
	if err != nil {
		c.JSON(200, gin.H{"error": err})
		return
	}

	//Gets ID to set session , gets password to compare
	var id uint
	var password string
	err = database.DBB.QueryRow(context.Background(), "SELECT id, password FROM site_users WHERE email = $1", r.Email).Scan(&id, &password)
	if err != nil {
		c.JSON(200, models.Error{Message: "Email or password incorrect"})
		return
	}

	//Check password
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(r.Password)); err != nil {
		c.JSON(200, models.Error{Message: "Email or password incorrect"})
		return
	}

	//If email and password is correct, send session
	session, _ := database.Store.Get(c.Request, "session")
	session.Values["userID"] = id

	session.Save(c.Request, c.Writer)
	c.JSON(200, gin.H{
		"message": "logged in",
	})
}
