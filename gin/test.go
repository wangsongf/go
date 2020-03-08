package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})

	//http://localhost:8081  postdata: name=tom
	router.POST("/", postHome)

	router.Run(":8082")
}

func postHome(c *gin.Context) {
	uName := c.PostForm("name")
	c.JSON(200, gin.H{
		"say": "Hello " + uName,
	})
}