package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	os.Setenv("AUTHORIZATION_STRING", "h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs")
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		authHeader := os.Getenv("AUTHORIZATION_STRING")
		requiredAuth := c.Request.Header.Get("Authorization")
		if requiredAuth != authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized peasant 😃"})
		}
		c.Next()
		// middleware
	}
}
