package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	r.GET("/file", func(c *gin.Context) {
		bs, err := ioutil.ReadFile(".gitignore")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": string(bs),
			})
		}
	})

	r.Run()
}
