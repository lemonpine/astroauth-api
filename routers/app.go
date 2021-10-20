package routers

import (
	"astroauth-api/database"
	"astroauth-api/middleware"
	"astroauth-api/models"

	"github.com/gin-gonic/gin"
)

func AppRouter(router *gin.Engine) {
	appuser := router.Group("/site")

	appuser.Use(middleware.CheckSession())
	{
		appuser.POST("/app", CreateApp)
	}
}

func CreateApp(c *gin.Context) {
	var app models.App
	c.ShouldBindJSON(&app)

	if err := database.DB.Where("name=?", app.Name).First(&app).Error; err == nil {
		c.JSON(200, models.Error{Message: "Application name not available"})
		return
	}
	app.OwnedBy = c.MustGet("userID").(uint)
	database.DB.Create(&app)
	c.JSON(200, app)
}
