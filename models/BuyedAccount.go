package models

import (
	"gorm.io/gorm"
)

type BuyedAccount struct {
	gorm.Model
	Link  string
	Price float32
}
