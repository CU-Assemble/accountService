package helpers

import (
	"example/CUAccountService/initializers"
	"example/CUAccountService/models"
)

func IsStudentIdUnique(sid string) bool { //haven't try this yet
	var user models.User
	result := initializers.DB.Where("student_id = ?", sid).First(&user)
	if result.Error != nil {
		return false
	}
	return true
}
