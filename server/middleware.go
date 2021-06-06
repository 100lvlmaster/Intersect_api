package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	os.Setenv("AUTHORIZATION_STRING", "h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs")
	return func(c *gin.Context) {
		authHeader := os.Getenv("AUTHORIZATION_STRING")
		requiredAuth := c.Request.Header.Get("Authorization")
		if requiredAuth != authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized peasant 😃"})
		}
		c.Next()
		// middleware
	}
}
