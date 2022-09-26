package service

type LoginService interface {
	LoginUser(studentId string, password string) bool
}
type loginInformation struct {
	studentId  string
	password string
}

// func StaticLoginService() LoginService {
// 	return &loginInformation{
// 		studentId:    "6231313021",
// 		password: "1234",
// 	}
// }


// func LoginUser(studentId string, password string) bool {

// 	var body struct {
// 		StudentId string
// 		Password  string
// 	}


// 	var user models.User
// 	result := initializers.DB.Where("student_id = ?", body.StudentId).First(&user)

// 	hash_pass, err := HashPassword(body.Password)
// 	var hpass = hash_pass

// 	if result.Error != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User with this student ID was not found."})
// 		return
// 	}

// 	return result.studentId == body.studentId && result.password == hpass
// }
