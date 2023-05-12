package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/esimov/xm/app/controller"
	"github.com/esimov/xm/app/models"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Engine *gin.Engine
}

func (s *Server) Init(c *config.Config) error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.UserName, c.Password, c.HostName, c.Port, c.DB)
	s.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	err = models.Load(s.DB)
	if err != nil {
		log.Fatalf("Server error, %v", err)
	}

	s.InitRoutes(c)

	return nil
}

func (s *Server) InitRoutes(config *config.Config) {
	// Home Route
	s.Engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	//s.Engine.Use(middlewares.JwtAuthMiddleware(c))
	s.Engine.POST("/login", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		//ctx.JSON(code, obj)
		controller.Login(s.DB, config, ctx)
	})
	s.Engine.GET("/users", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		controller.GetUsers(s.DB, config, ctx)
	})
	s.Engine.POST("/users/add", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		controller.CreateUser(s.DB, config, ctx)
	})
	s.Engine.PATCH("/users/:id", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		controller.UpdateUser(s.DB, config, ctx)
	})
	s.Engine.GET("/company:id", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	s.Engine.Run()
}
