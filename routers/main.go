package routers

import (
	"github.com/gin-gonic/gin"
)

/*
This is where the main router is declared.
Endpoints are grouped up into different files
*/

func InitializeRouter() {
	router := gin.New()

	//Site
	SiteUserRouter(router)
	LicenseRouter(router)
	AppRouter(router)

	//App
	AppUserRouter(router)

	router.Run(":8080")

}
