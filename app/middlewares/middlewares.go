package middlewares

import (
	"net/http"

	"github.com/esimov/xm/auth"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateToken(config, c.Request)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
