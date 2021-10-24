package routers

import (
	"astroauth-api/database"
	"astroauth-api/middleware"
	"astroauth-api/models"
	"context"

	"github.com/gin-gonic/gin"
)

/*
All routes for handling apps
*/

func AppRouter(router *gin.Engine) {
	appuser := router.Group("/site")

	appuser.Use(middleware.CheckSession())
	{
		appuser.POST("/app", CreateApp)
	}
}

func CreateApp(c *gin.Context) {
	var rApp models.App
	c.ShouldBindJSON(&rApp)

	var name string
	err := database.DBB.QueryRow(context.Background(), "SELECT name FROM apps WHERE name = $1", rApp.Name).Scan(&name)
	if err == nil {
		c.JSON(200, models.Error{Message: "Name not available"})
		return
	}
	rApp.OwnedBy = c.MustGet("userID").(uint)
	database.DB.Create(&rApp)
	c.JSON(200, rApp)
}
