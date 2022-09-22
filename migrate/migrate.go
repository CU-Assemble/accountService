package main

import (
	"example/CUAccountService/initializers"
	"example/CUAccountService/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
