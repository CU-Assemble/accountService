package main

import (
	"example/CUAccountService/controllers"
	"example/CUAccountService/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default() //router
	r.POST("/user", controllers.UserCreate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:sid", controllers.GetUserById)
	r.PUT("/user/:sid", controllers.UserUpdate)
	// r.DELETE()
	r.Run()
}
