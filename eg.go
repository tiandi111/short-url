package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Generator struct {
	index	int
	pool 	map[int]string
}

func main() {
	g := Generator{0, make(map[int]string, 0)}
	
	r := gin.Default()
	
	r.LoadHTMLGlob("templates/*")

	// Main page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title": "Main Website",
		})
	})

	r.GET("/:index", func(c *gin.Context) {
		index, _:= strconv.Atoi( c.Param("index") )
		c.Redirect(http.StatusMovedPermanently, g.pool[index])
		fmt.Println("Redirect to: %s", g.pool[index])
	})

	r.POST("/:index/create", func(c *gin.Context) {
		url := c.PostForm("url")
		short := g.createURL(url)
		c.String(http.StatusOK, short)
		fmt.Println(short)
	})

	r.Run(":8080")
}

func (g *Generator) createURL(url string) string{
	i := g.index
	g.pool[i] = url
	postfix := strconv.Itoa(i)
	g.index++
	return "http://54.196.113.135:8080/"+postfix
}
