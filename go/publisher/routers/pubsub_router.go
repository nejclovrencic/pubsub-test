package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"lovrencic.com/pubsub/services"
	"lovrencic.com/pubsub/types"
	"net/http"
)

func PubSubRouter(router *gin.RouterGroup) {
	router.POST("/subscribe", subscribe)
	router.POST("/unsubscribe", unsubscribe)
	router.POST("/publish", publish)
}

func subscribe(c *gin.Context) {
	var body types.SubscribeRequest

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
			"error":   err.Error(),
		})
		return
	}

	services.SubscribeToChannel(body.Channel, body.ClientURL)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func unsubscribe(c *gin.Context) {
	var body types.SubscribeRequest

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
			"error":   err.Error(),
		})
		return
	}

	services.UnsubscribeFromChannel(body.Channel, body.ClientURL)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func publish(c *gin.Context) {
	var body types.PublishRequest

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters",
			"error":   err.Error(),
		})
		return
	}

	services.PublishMessageToChannel(body.Channel, body.Message)

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

