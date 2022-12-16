package services

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "123@gmail.com",
		password: "123123123",
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
