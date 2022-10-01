package main

import (
  "github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	log.Println("Listening on: https://localhost:8080")
	r.Run(":8080")
}