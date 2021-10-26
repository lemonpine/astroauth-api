package routers

import (
	"astroauth-api/database"
	"astroauth-api/middleware"
	"astroauth-api/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func LicenseRouter(router *gin.Engine) {
	license := router.Group("/site/app")

	license.Use(middleware.CheckSession())
	{
		license.POST("/licenses", middleware.AddLicenseValidate(), middleware.CheckAppSite(), AddLicense)
		license.DELETE("/licenses", DeleteLicense)
		license.GET("/licenses", GetLicenses)
	}
}

func AddLicense(c *gin.Context) {
	var License models.License
	c.ShouldBindBodyWith(&License, binding.JSON)

	database.DB.Create(&License)
	c.JSON(200, gin.H{"license": License.License})
}

func DeleteLicense(c *gin.Context) {

}

func GetLicenses(c *gin.Context) {

}
