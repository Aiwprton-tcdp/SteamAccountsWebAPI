package models

import (
	"gorm.io/gorm"
)

type PasswordReset struct {
	gorm.Model
	Email string
	Token string
}
