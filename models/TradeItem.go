package models

import (
	"gorm.io/gorm"
)

type TradeItem struct {
	gorm.Model
	TradeId    uint
	AssetId    uint
	ClassId    uint
	InstanceId uint
	Price      float32
	Type       string
}
