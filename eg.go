package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func main() {
	r := gin.Default()
	
	r.LoadHTMLGlob("templates/*")
	go GenerateID()

	cURL := NewDbClient().Database("test").Collection("URL")

	// Render main page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title": "Main Website",
		})
	})

	// TODO: Intercept invalid url
	// Render redirected page
	r.GET("/:index", func(c *gin.Context) {
		index := c.Param("index")
		url := FindLongURL(context.TODO(), cURL, bson.D{{"id64", index}})
		c.Redirect(http.StatusFound, url)
	})

	// Create new short-long url pair
	r.POST("/:index/create", func(c *gin.Context) {
		long := c.PostForm("url")
		short, err := CreateShortURL(context.TODO(), cURL, long)
		if err != nil {
			c.String(http.StatusOK, "Operation failed, try again.")
		} else {
			c.String(http.StatusOK, "%s", short)
		}
	})

	r.Run(":8080")
}


