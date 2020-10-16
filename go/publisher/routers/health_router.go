package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthRouter(router *gin.RouterGroup) {
	router.POST("/liveness", handleLiveness)
}

func handleLiveness(c *gin.Context) {
	c.Status(http.StatusOK)
}
