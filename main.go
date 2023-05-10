package main

import (
	"net/http"

	"github.com/esimov/xm/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")

	})

	model.Create()

	r.Run()
}
