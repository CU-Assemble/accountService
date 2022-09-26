package service

type LoginService interface {
	LoginUser(studentId string, password string) bool
}
type loginInformation struct {
	studentId  string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		studentId:    "6231313021",
		password: "1234",
	}
}

func (info *loginInformation) LoginUser(studentId string, password string) bool {
	return info.studentId == studentId && info.password == password
}