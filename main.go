package main

import (
	"io/ioutil"
	"os"

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

	r.GET("/env", func(c *gin.Context) {
		val, ok := os.LookupEnv("SOMETHING")
		if !ok {
			c.JSON(500, gin.H{
				"error": "Env SOMETHING not found",
			})
		} else {
			c.JSON(200, gin.H{
				"message": val,
			})
		}
	})

	r.Run()
}
