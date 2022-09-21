package main

import (
	// "errors"
	"net/http"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	StudentId string    `json:"studentId"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Birthdate time.Time `json:"birthdate"`
	Faculty   string    `json:"faculty"`
	Tel       string    `json:"tel"`
	Email     string    `json:"email"`
}

var users = []user{
	{StudentId: "6230000121", Name: "Jojo joeyboy", Nickname: "abdul", Birthdate: time.Date(2510, time.January, 1, 7, 0, 0, 0, time.Local), Faculty: "Engineering", Tel: "0801234567", Email: "jojo@gmail.com"},
	{StudentId: "6230000221", Name: "Jojo jostars", Nickname: "labul", Birthdate: time.Date(2510, time.February, 2, 7, 0, 0, 0, time.Local), Faculty: "Engineering", Tel: "0811234567", Email: "joja@gmail.com"},
	{StudentId: "6230000321", Name: "Jojo joystik", Nickname: "saber", Birthdate: time.Date(2510, time.March, 3, 7, 0, 0, 0, time.Local), Faculty: "Engineering", Tel: "0831234567", Email: "joje@gmail.com"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser user
	fmt.Println("before bind")
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	fmt.Println(newUser.Birthdate)
	users = append(users, newUser) //will insert to db soon ... SOON!!!
	c.IndentedJSON(http.StatusCreated, newUser)

}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/user", createUser)
	router.Run("localhost:8080")
}
