package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func GetImgs(c *gin.Context) {
	searchQuery := c.Query("q")
	res := getSearch(searchQuery)
	c.JSON(http.StatusOK, res)
}

func getSearch(searchQuery string) Images {
	searchString := strings.Replace(searchQuery, " ", "-", -1)
	c := colly.NewCollector()

	array := []string{}

	// Find and visit all links
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if src != "" {
			array = append(array, e.Attr("src"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://unsplash.com/s/" + searchString)
	c.Visit("https://www.freeimages.com/search/" + searchString)
	return Images{
		Count: len(array),
		Data:  array}
}

type Images struct {
	Count int      `json:"counts"`
	Data  []string `json:"data"`
}
