package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"shortURL/internal/db"
)

func isURL(link string) bool {
	u, err := url.Parse(link)
	if err != nil {
		return false
	}
	return (u.Opaque != "" || u.Host != "") && u.Scheme != ""
}

func main() {
	router := gin.Default()

	// Load HTML
	router.Static("/static", "../public")
	router.LoadHTMLFiles("../public/index.html")

	// / Page handler
	router.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK,"index.html", nil)
	})

	// Link shortener handler
	router.POST("/short", func(c *gin.Context) {
		url := c.PostForm("link")
		if isURL(url) {
			log.Println("получено: ", url)

			if url == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Пустая ссылка"})
				return
			}

			short, _ := db.AddLink(url)
			c.JSON(http.StatusOK, gin.H{
				"short": short,
				"link": url,
				})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "Bad link"})
		}
	})

	// Pereadresation
	router.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			link, _ := db.GetLink(id)

			c.Redirect(http.StatusFound, link)
	})

	router.Run(":4040")
}
