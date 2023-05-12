package middleware

import (
	"fmt"
	"net/http"

	"github.com/esimov/xm/auth"
	"github.com/esimov/xm/config"
	"github.com/gin-gonic/gin"
)

func JwtAuth(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateToken(config, c.Request)
		if err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("authorization error: %s", err.Error()))
			c.Abort()
			return
		}
		c.Next()
	}
}
