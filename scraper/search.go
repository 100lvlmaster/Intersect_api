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
	// Requesting a url for html
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// search query
	pexelsQuery := strings.Replace(searchString, "-", "%20", -1)
	stocSnapQuery := strings.Replace(searchString, "-", "+", -1)
	//
	c.Visit("https://unsplash.com/s/" + searchString)
	c.Visit("https://burst.shopify.com/photos/search?utf8=%E2%9C%93&q=" + searchString + "&button=")
	c.Visit("https://www.pexels.com/search/" + pexelsQuery + "/")
	c.Visit("https://www.flickr.com/search/?text=" + pexelsQuery)
	c.Visit("http://www.google.com/images?q=" + stocSnapQuery)
	c.Visit("https://stocksnap.io/search/" + stocSnapQuery)
	//
	return Images{
		Count: len(array),
		Data:  array}
}

type Images struct {
	Count int      `json:"counts"`
	Data  []string `json:"data"`
}
