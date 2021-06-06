package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	authHeader := "let-me-in"
	// os.Getenv("AUTHORIZATION_STRING")
	return func(c *gin.Context) {
		requiredAuth := c.Request.Header.Get("Authorization")
		if requiredAuth != authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "unauthorized peasant 😃"})
		}
		c.Next()
		// middleware
	}
}
