package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
)

const PORT = 8080

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	log.Printf("Listening on: https://localhost:%d\n", PORT)
	r.Run(fmt.Sprintf(":%d", PORT))
}
