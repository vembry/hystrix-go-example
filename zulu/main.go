package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "zulu-pong",
		})
	})

	r.POST("/ping/:pathVariable", func(c *gin.Context) {
		type Body struct {
			SomeKeyValue1 map[string]string `json:"someKeyValue1"`
			SomeKeyValue2 string            `json:"someKeyValue2"`
		}
		type Header struct {
			SomeHeader string `header:"x-some-header"`
		}

		var body Body
		var header Header

		c.BindJSON(&body)
		c.BindHeader(&header)
		c.JSON(200, gin.H{
			"message": "zulu-pong",
			"header":  header,
			"body":    body,
		})
	})

	r.Run("localhost:3002")
}
