package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	
	r.LoadHTMLGlob("templates/*")

	db := NewDataBase()

	// Render main page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title": "Main Website",
		})
	})

	// Render redirected page
	r.GET("/:index", func(c *gin.Context) {
		index := c.Param("index")
		url := db.GetLongURL(index)
		c.Redirect(http.StatusMovedPermanently, url)
	})

	// Create new short-long url pair
	r.POST("/:index/create", func(c *gin.Context) {
		long := c.PostForm("url")
		short := CreateShortURL(db, long)
		c.String(http.StatusOK, short)
		fmt.Println(short)
	})

	r.Run(":8080")
}


