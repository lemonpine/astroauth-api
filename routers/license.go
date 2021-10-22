package routers

import (
	"astroauth-api/database"
	"astroauth-api/middleware"
	"astroauth-api/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
All routes for handling licenses within apps
*/

func LicenseRouter(router *gin.Engine) {
	license := router.Group("/site/app")

	license.Use(middleware.CheckSession(), middleware.CheckApp())
	{
		license.POST("/licenses", AddLicense)
		license.DELETE("/licenses", DeleteLicense)
		license.GET("/licenses", GetLicenses)
	}
}

func AddLicense(c *gin.Context) {
	var rLicense models.License
	c.ShouldBindBodyWith(&rLicense, binding.JSON)
	fmt.Println(rLicense.AppID)

	err := rLicense.Validate()
	if err != nil {
		c.JSON(200, gin.H{"error": err})
		return
	}

	//Validate user input

	database.DB.Create(&rLicense)
	c.JSON(200, gin.H{"license": rLicense.License})
}

func DeleteLicense(c *gin.Context) {

}

func GetLicenses(c *gin.Context) {

}
