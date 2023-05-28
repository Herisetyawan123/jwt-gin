package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heirsetyawan233/jwt-gin/controller"
	"github.com/heirsetyawan233/jwt-gin/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	r.GET("/validate", controller.Validate)

	r.Run()
	fmt.Println("Hallo dunia")
}
