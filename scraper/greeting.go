package scraper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(c *gin.Context) {

	res := make(map[string]string)
	res["sup"] = "🤘🚀"
	c.JSON(http.StatusOK, res)

}
