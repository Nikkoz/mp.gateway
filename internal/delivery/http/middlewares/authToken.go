package middlewares

import (
	"errors"
	"github.com/Nikkoz/mp.gateway/internal/configs"
	"github.com/Nikkoz/mp.gateway/internal/delivery/http/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Auth(config configs.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			error.SetError(c, http.StatusUnauthorized, errors.New("no Authorization header provided"))
			c.Abort()

			return
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		if token != config.Token {
			error.SetError(c, http.StatusUnauthorized, errors.New("unauthorized"))
			c.Abort()

			return
		}

		c.Next()
	}
}
