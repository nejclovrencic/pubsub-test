package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"lovrencic.com/pubsub/db"
	"lovrencic.com/pubsub/routers"
	"lovrencic.com/pubsub/services"
	"lovrencic.com/pubsub/utils"
)

func StartServer() {
	db.Init(fmt.Sprintf("%s:%s", utils.Getenv("REDIS_HOST", "localhost"), utils.Getenv("REDIS_HOST", "6379")))
	server := gin.Default()

	routers.SetupRoute("/", routers.HealthRouter, server.Group("/health"))

	routers.SetupRoutesV1(server.Group("/api/v1"))
	go startRssFetching()

	server.Run(fmt.Sprintf(":%s", utils.Getenv("PORT", "3000")))
}

func startRssFetching() {
	c := cron.New()
	c.AddFunc("@every 1m", func() { services.FetchRssAndSendMessage("http://meteo.arso.gov.si/uploads/probase/www/observ/surface/text/en/observation_eu-capital_latest.rss") })
	c.Start()
}
