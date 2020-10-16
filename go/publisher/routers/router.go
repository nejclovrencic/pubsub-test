package routers

import (
	"github.com/gin-gonic/gin"
)

type routeSetupFunction func(*gin.RouterGroup)

func SetupRoutesV1(router *gin.RouterGroup) {
	SetupRoute("/publisher", PubSubRouter, router)
}

func SetupRoute(path string, setup routeSetupFunction, router *gin.RouterGroup) {
	setup(router.Group(path))
}

