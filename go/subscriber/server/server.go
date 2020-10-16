package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"lovrencic.com/pubsub/routers"
	"lovrencic.com/pubsub/utils"
	"net/http"
	"os"
)

func StartServer() {
	server := gin.Default()

	routers.SetupRoute("/", routers.PubSubRouter, server.Group("/api/v1"))

	go sendRequest(utils.Getenv("SERVER_URL", "http://"), gin.H{
		"channel": os.Args[1],
		"clientUrl": fmt.Sprintf("http://localhost:%s/api/v1/message", utils.Getenv("PORT", "9000")),
	})

	server.Run(fmt.Sprintf(":%s", utils.Getenv("PORT", "9000")))
}


func sendRequest(url string, body interface{}) {
	req, _ := json.Marshal(body)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(req))

	if err != nil {
		fmt.Println(err)
	}
}