package main

import (
	"log"

	"github.com/esimov/microservice-demo/app"
	"github.com/esimov/microservice-demo/config"
	"github.com/esimov/microservice-demo/kafka"
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
