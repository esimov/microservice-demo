package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/esimov/xm/app"
	"github.com/esimov/xm/app/model"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	Start()
	r.Run()
}

func Start() {
	config, err := config.GetConfig("app")
	if err != nil {
		log.Fatalf("Error getting the environment variables, %v", err)
	}
	fmt.Println("Config:", config)

	server := &app.Server{}
	err = server.Init(config)
	if err != nil {
		log.Fatalf("Server error, %v", err)
	}

	err = model.Load(server.DB)
	if err != nil {
		log.Fatalf("Server error, %v", err)
	}
}
