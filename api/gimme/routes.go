package gimme

import (
	"Intersect/scraper"

	"github.com/gin-gonic/gin"
)

// Routes for /gimme
func Routes(group *gin.RouterGroup) {

	group.GET("", scraper.Greet)
	//
	group.GET("/get", scraper.GetImgs)

}
