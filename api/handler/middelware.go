package handler

import (
	"errors"
	"net/http"

	"github.com/asadbekGo/market_system/config"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckPasswordMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		password := c.GetHeader("API-KEY")
		if password != config.SecureApiKey {
			c.AbortWithError(http.StatusForbidden, errors.New("The request requires an user authentication."))
			return
		}

		c.Next()
	}
}
