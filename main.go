package main

import (
	"log"

	"github.com/esimov/xm/app"
	"github.com/esimov/xm/config"
	"github.com/esimov/xm/kafka"
	"github.com/gin-gonic/gin"
)

func main() {
	brokerList := []string{"kafkahost1:9092"}
	config, err := config.GetConfig("app")
	if err != nil {
		log.Fatalf("Error getting the environment variables, %v", err)
	}

	server := &app.Server{
		Route:    gin.New(),
		Producer: kafka.Producer(brokerList),
		Consumer: kafka.Consumer(brokerList),
	}

	err = server.Init(config)
	if err != nil {
		log.Fatalf("Server error, %v", err)
	}
}
