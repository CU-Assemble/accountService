package main

import (
	"example/CUAccountService/controllers"
	"example/CUAccountService/initializers"

	"github.com/gin-gonic/gin"

	/////
	"example/CUAccountService/middleware"
	"example/CUAccountService/service"
	"net/http"
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

	////new add
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(jwtService)

	r.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	v1 := r.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		})
	}
	r.Run()
	/////
}
