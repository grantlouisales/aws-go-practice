package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "MADE IT TO AWS!!!!",
		})
	})
	router.Run(":8081") // listents on :8080
}
