package app

import (
	"io/ioutil"
	"net/http"

	"github.com/esimov/microservice-demo/app/controller"
	"github.com/esimov/microservice-demo/app/middleware"
	"github.com/esimov/microservice-demo/config"
	"github.com/gin-gonic/gin"
)

func (s *Server) InitRoutes(config *config.Config) error {
	s.Route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to our platform.")
	})

	s.Route.POST("/login", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.Login(s.DB, config, c)
	})
	s.Route.GET("/users", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.GetUsers(s.DB, config, c)
	})
	s.Route.POST("/users/add", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.CreateUser(s.DB, config, c)
	})
	s.Route.PATCH("/users/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.UpdateUser(s.DB, config, c)
	})

	company := s.Route.Group("/company")
	company.Use(middleware.JwtAuth(config))
	company.POST("/create", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.CreateCompany(s.DB, config, c)
		body, _ := ioutil.ReadAll(c.Request.Body)
		// Kafka event test
		{
			s.Send("company", body)
			s.Receive("company")
		}

	})
	company.PATCH("/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.UpdateCompany(s.DB, config, c)
	})
	company.DELETE("/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.DeleteCompany(s.DB, config, c)
	})
	company.GET("/:id", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		controller.GetCompany(s.DB, config, c)
	})

	s.Route.Run()

	return nil
}
