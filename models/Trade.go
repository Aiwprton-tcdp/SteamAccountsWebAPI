package models

import (
	"gorm.io/gorm"
)

type Trade struct {
	gorm.Model
	SenderId uint
	TradeId  uint
	Link     string
	Status   uint
}
