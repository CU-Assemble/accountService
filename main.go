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

/* to be change
*	model structure : remove Bdate, add picture & password
*
 */

func main() {
	r := gin.Default() //router
	r.POST("/user", controllers.UserCreate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:sid", controllers.GetUserById)
	r.PUT("/user/:sid", controllers.UserUpdate)
	r.DELETE("/user/:sid", controllers.UserDelete)
	r.Run()
}
