package scraper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"sup": "🤘🚀"})

}
