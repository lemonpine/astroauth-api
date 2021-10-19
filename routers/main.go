package routers

import (
	"github.com/gin-gonic/gin"
)

/*
This is where the main router is declared.
Endpoints are grouped up into different files
*/

var router *gin.Engine

func InitializeRouter() {
	router = gin.New()
	SiteUserRouter(router)
	router.Run(":8080")

}
