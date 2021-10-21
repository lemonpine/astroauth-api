package middleware

import (
	"astroauth-api/database"
	"astroauth-api/models"

	"github.com/gin-gonic/gin"
)

/*
	Can be applied to any endpoint
	Session cookie must be passed, if nil or invalid function will abort/return
	if cookie is valid, userid will be extracted then passed into the handler
*/
func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := database.Store.Get(c.Request, "session")
		if err != nil {
			c.JSON(401, models.Error{Message: "Unauthorized"})
			c.Abort()
			return
		}

		if session.Values["userID"] == nil {
			c.JSON(401, models.Error{Message: "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", session.Values["userID"])
	}
}
