package routers

import (
	"github.com/gin-gonic/gin"
)

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
