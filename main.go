package main

import (
	"fmt"
	"log"

	"github.com/esimov/xm/app"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.GetConfig("app")
	if err != nil {
		log.Fatalf("Error getting the environment variables, %v", err)
	}
	fmt.Println("Config:", config)

	server := &app.Server{
		Route: gin.New(),
	}

	err = server.Init(config)
	if err != nil {
		log.Fatalf("Server error, %v", err)
	}
}
