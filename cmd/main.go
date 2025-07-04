package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shortURL/internal/db"
)


func main() {
	router := gin.Default()

	// Подгружаем HTML файлы
	router.Static("/static", "../public")
	router.LoadHTMLFiles("../public/index.html")

	// Обработчик для / запроса
	router.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK,"index.html", nil)
	})

	// Запрос сокращения ссылки
	router.POST("/short", func(c *gin.Context) {
		url := c.PostForm("link")
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
	})

	// Переадресация на оригинальную ссылку
	router.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			link, _ := db.GetLink(id)

			c.Redirect(http.StatusFound, link)
	})

	router.Run(":3456")
}
