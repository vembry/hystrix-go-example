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

	r.Run("localhost:3001")
}
