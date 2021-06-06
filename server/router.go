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
	/// Cors
	router.Use(setCors())
	/// Declare Middleware
	authorized := router.Group("/")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("", scraper.Greet)
		searchRouter := authorized.Group("search")
		searchRouter.GET("", scraper.GetImgs)
	}
	return router
}

// Cors
func setCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*", "https://the-intersect.vercel.app"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
