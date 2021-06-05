package server

import (
	"Intersect/api/gimme"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter : Function with routes
func NewRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Gin and CORS Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	// /gimme routes
	router.GET("", func(c *gin.Context) {
		res := map[string]string{"hello": "wrong route tho 🙁"}
		c.JSON(http.StatusOK, res)
	})
	gimmeRouter := router.Group("search")
	gimme.Routes(gimmeRouter)

	return router
}
