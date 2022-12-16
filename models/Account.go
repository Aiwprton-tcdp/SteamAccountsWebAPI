package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserId       int
	Login        string `gorm:"uniqueIndex:login;not null;size:256" binding:"required"`
	Password     string `gorm:"not null;size:256" binding:"required"`
	SharedSecret string `gorm:"default:null;size:55"`
	Email        string `gorm:"default:null;size:256"`
	SteamId      string `gorm:"default:null"`
	Nickname     string `gorm:"default:null;size:256"`
	Prime        *bool
	Blocked      *bool
	Balance      *float32
	Lvl          *int16
	CSGORank     *int   `json:"rank"`
	FriendCode   string `gorm:"default:null;size:10"`
	InFarm       *bool
}
