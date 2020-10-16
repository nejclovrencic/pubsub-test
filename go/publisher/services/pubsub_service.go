package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
	"lovrencic.com/pubsub/db"
	"net/http"
	"strings"
)

func SubscribeToChannel(channel, clientUrl string) {
	db.SADD(strings.ToLower(channel), clientUrl)
}

func UnsubscribeFromChannel(channel, clientUrl string) {
	db.SREM(strings.ToLower(channel), clientUrl)
}

func PublishMessageToChannel(channel string, message interface{}) error {
	clientUrls, err := db.SMEMBERS(strings.ToLower(channel))

	if err != nil {
		return err
	}

	for _, url := range clientUrls {
		go sendRequest(url, gin.H{
			"message": message,
		})
	}

	return nil
}

func FetchRssAndSendMessage(url string) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	for _, item := range feed.Items {
		city := strings.Split(item.GUID, "_")[0]
		go PublishMessageToChannel(strings.ToLower(city), item.Description)
	}
}

func sendRequest(url string, body interface{}) {
	req, _ := json.Marshal(body)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(req))

	if err != nil {
		fmt.Println(err)
	}
}
