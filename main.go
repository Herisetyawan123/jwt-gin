package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heirsetyawan233/jwt-gin/controller"
	"github.com/heirsetyawan233/jwt-gin/initializers"
	"github.com/heirsetyawan233/jwt-gin/middleware"
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
	r.GET("/validate", middleware.AuthMiddleware, controller.Validate)
	r.GET("/logout", middleware.AuthMiddleware, controller.SignOut)

	r.Run()

}
