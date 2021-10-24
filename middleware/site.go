package middleware

import (
	"astroauth-api/database"
	"astroauth-api/models"
	"context"
	"fmt"

	"github.com/gin-gonic/gin/binding"

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

//uses different middleware for the site to check if the session user id matches the owned by field for the app
func CheckAppSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			AppID string `json:"app_id" `
		}

		var r Request
		c.ShouldBindBodyWith(&r, binding.JSON)

		if r.AppID == "" || r.AppID == " " {
			c.JSON(404, models.Error{Message: "app_id cannot be blank"})
			c.Abort()
			return
		}
		//c.ShouldBindBodyWith is used instead of c.shouldbindjson, as it can redeclare the body in the next function
		c.ShouldBindBodyWith(&r, binding.JSON)

		var AppID string
		var OwnedBy uint
		err := database.DBB.QueryRow(context.Background(), "SELECT app_id, owned_by FROM apps WHERE app_id = $1", r.AppID).Scan(&AppID, &OwnedBy)
		if err != nil {
			c.JSON(404, models.Error{Message: "Application not found"})
			c.Abort()
			return
		}

		var se string
		fmt.Println(c.MustGet("userID"))
		err2 := database.DBB.QueryRow(context.Background(), "SELECT app_id FROM apps WHERE owned_by = $1", c.MustGet("userID")).Scan(&se)
		if err2 != nil {
			c.JSON(401, models.Error{Message: "Unauthorized"})
			c.Abort()
			return
		}

	}
}
