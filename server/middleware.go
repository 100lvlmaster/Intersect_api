package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	authHeader := os.Getenv("AUTHORIZATION_STRING")
	return func(c *gin.Context) {
		requiredAuth := c.Request.Header.Get("Authorization")
		if requiredAuth != authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized peasant 😃"})
		}
		c.Next()
		// middleware
	}
}
