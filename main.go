package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heirsetyawan233/jwt-gin/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
	fmt.Println("Hallo dunia")
}
