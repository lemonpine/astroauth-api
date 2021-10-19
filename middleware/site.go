package middleware

import (
	"astroauth-api/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "SESSION MIDDLEWARE",
		})
		session, err := database.Store.Get(c.Request, "session")

		if err != nil {
			c.Abort()
			fmt.Println("no sess")
			return
		}

		if session.Values["userID"] == nil {
			c.Abort()
		}
		c.JSON(200, gin.H{
			"uid": session.Values["userID"],
		})
	}
}
