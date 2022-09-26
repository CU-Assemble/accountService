package controllers

import (
	"example/CUAccountService/dto"
	"example/CUAccountService/service"
	"example/CUAccountService/initializers"
	"example/CUAccountService/models"
	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(
	jWtService service.JWTService) LoginController {
	return &loginController{
		jWtService:   jWtService,
	}
}


func (controller *loginController) Login(c *gin.Context) string {
	var credential dto.LoginCredentials
	err := c.ShouldBind(&credential)
	c.Bind(&credential)
	if err != nil {
		return "no data found"
	}

 	var user models.User
 	initializers.DB.Where("student_id = ?", credential.StudentId).First(&user)
	// if result.Error != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User with this student ID was not found."})
	// 	return
	// }
	hash_pass, err := HashPassword(credential.Password)
	var hpass = hash_pass


// 	return result.studentId == body.studentId && result.password == hpass
	if (user.StudentId == credential.StudentId && user.Password == hpass) {
		return controller.jWtService.GenerateToken(credential.StudentId)

	}
	return ""
}