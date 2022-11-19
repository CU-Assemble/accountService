package main

import (
	"example/CUAccountService/controllers"
	"example/CUAccountService/initializers"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	/////
	"example/CUAccountService/middleware"
	"example/CUAccountService/service"
	"log"
	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func serviceRegistryWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	serviceID := "account-service2"
	port, _ := strconv.Atoi(getPort()[1:len(getPort())])
	address := "192.168.0.101"

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "account-service",
		Port:    port,
		Address: getHostname(),
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Consul check")
}

func getPort() (port string) {
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	return
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return
}

func main() {
	r := gin.Default() //router
	r.Use(CORSMiddleware())
	serviceRegistryWithConsul()
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
			ctx.JSON(200, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	r.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"check": "ok",
		})

	})
	v1 := r.Group("/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})
	}
	/////
	r.Run()

}
