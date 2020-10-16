package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lovrencic.com/pubsub/types"
	"net/http"
)

func PubSubRouter(router *gin.RouterGroup) {
	router.POST("/message", receiveMessage)
}

func receiveMessage(c *gin.Context) {
	var body types.MessageRequest

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println(body.Message)

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

