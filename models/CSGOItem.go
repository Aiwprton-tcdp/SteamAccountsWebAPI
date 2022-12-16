package models

import (
	"gorm.io/gorm"
)

type CSGOItem struct {
	gorm.Model
	Name    string
	ClassId uint
	AppId   uint
	Price   float32
	Count   uint
	BPrice  float32
	BCount  uint
	Rarity  string
	Img     string
}
