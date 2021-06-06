package server

import (
	"Intersect/scraper"

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
	/// Declare Middleware
	authorized := router.Group("/")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("", scraper.Greet)
		searchRouter := authorized.Group("search")
		searchRouter.GET("", scraper.GetImgs)
	}
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.

	// /gimme routes

	return router
}
