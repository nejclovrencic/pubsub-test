package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"lovrencic.com/pubsub/server"
	"os"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		fmt.Println(err)
	}

	server.StartServer()
}
