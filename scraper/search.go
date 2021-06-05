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
	fmt.Print("the search string is")
	fmt.Print(c.Query("q"))
	res := getSearch(searchQuery)
	c.JSON(http.StatusOK, res)
}

func getSearch(searchQuery string) Images {
	searchString := strings.Replace(searchQuery, " ", "-", -1)
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0"
	c.AllowURLRevisit = true
	c.DisableCookies()
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
