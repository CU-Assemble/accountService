package controllers

import (
	"example/CUAccountService/initializers"
	"example/CUAccountService/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {

	var body struct {
		StudentId string
		Name      string
		Nickname  string
		Birthdate time.Time
		Faculty   string
		Tel       string
		Email     string
	}

	c.Bind(&body)

	user := models.User{
		StudentId: body.StudentId,
		Name:      body.Name,
		Nickname:  body.Nickname,
		Birthdate: body.Birthdate,
		Faculty:   body.Faculty,
		Tel:       body.Tel,
		Email:     body.Email,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUserById(c *gin.Context) {
	sid := c.Param("sid")

	var user models.User
	result := initializers.DB.Where("student_id = ?", sid).First(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User with this student ID was not found."})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	// Get SID
	sid := c.Param("sid")

	// Get data from req body
	var body struct {
		StudentId string
		Name      string
		Nickname  string
		Birthdate time.Time
		Faculty   string
		Tel       string
		Email     string
	}
	c.Bind(&body)

	//find the user
	var user models.User
	result := initializers.DB.Where("student_id = ?", sid).First(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User with this student ID was not found."})
		return
	}

	//Update multiple column use Updates
	updatedResult := initializers.DB.Model(&user).Updates(models.User{
		StudentId: body.StudentId,
		Name:      body.Name,
		Nickname:  body.Nickname,
		Birthdate: body.Birthdate,
		Faculty:   body.Faculty,
		Tel:       body.Tel,
		Email:     body.Email,
	})

	if updatedResult.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request body for update"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User has been updated",
		"user":    user,
	})
}

func UserDelete(c *gin.Context) {
	// Get SID
	sid := c.Param("sid")

	initializers.DB.Where("student_id = ?", sid).Delete(&models.User{})

	c.JSON(200, gin.H{
		"message": "User has been deleted",
	})
}
