package middleware

import (
	"astroauth-api/database"
	"astroauth-api/models"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AppBasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			AppID string `json:"app_id" `
		}

		var r Request
		c.ShouldBindBodyWith(&r, binding.JSON)

		//Check if basic auth authorization header is present
		username, password, err := c.Request.BasicAuth()
		if !err {
			c.JSON(401, models.Error{Message: "BasicAuth authorization header missing"})
			c.Abort()
			return
		}

		if username == "" || password == "" {
			c.JSON(200, models.Error{Message: "Username or password incorrect"})
			c.Abort()
			return
		}

		var ID uint
		var DBPassword string
		var UserExpiry time.Time
		err2 := database.DBB.QueryRow(context.Background(), "SELECT id, password, expiry FROM app_users WHERE username = $1 AND app_id = $2", username, r.AppID).Scan(&ID, &DBPassword, UserExpiry)

		if err2 != nil {
			c.JSON(200, models.Error{Message: "Username or password incorrect"})
			c.Abort()
			return
		}
		c.Set("UserID", ID)

		//Check password
		if err := bcrypt.CompareHashAndPassword([]byte(DBPassword), []byte(password)); err != nil {
			c.JSON(200, models.Error{Message: "Username or password incorrect"})
			c.Abort()
			return
		}

		//Check if user has expired
		if time.Now().After(UserExpiry) {
			c.JSON(200, models.Error{Message: "User expired"})
			c.Abort()
			return
		}

	}
}

func CheckApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			AppID string `json:"app_id" `
		}

		var r Request
		c.ShouldBindBodyWith(&r, binding.JSON)

		//c.ShouldBindBodyWith is used instead of c.shouldbindjson, as it can redeclare the body in the next function
		c.ShouldBindBodyWith(&r, binding.JSON)

		var AppName string
		err := database.DBB.QueryRow(context.Background(), "SELECT name FROM apps WHERE app_id = $1", r.AppID).Scan(&AppName)
		if err != nil {
			c.JSON(404, models.Error{Message: "Application not found"})
			c.Abort()
			return
		}
		//Pass app id into next function
		c.Set("AppID", r.AppID)
		c.Set("AppName", AppName)
	}
}
