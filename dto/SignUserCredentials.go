package dto

type SignUserCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
