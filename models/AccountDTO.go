package models

import "gorm.io/gorm"

type AccountDTO struct {
	gorm.Model
	UserId     uint    `json:"user_id"`
	Login      string  `json:"login"`
	Password   string  `json:"password"`
	Email      string  `json:"email"`
	SteamId    string  `json:"steamId"`
	Nickname   string  `json:"nickname"`
	Prime      bool    `json:"prime"`
	Blocked    bool    `json:"blocked"`
	Balance    float32 `json:"balance"`
	Lvl        uint    `json:"lvl"`
	CSGORank   uint    `json:"rank"`
	FriendCode string  `json:"friend_code"`
	InFarm     bool    `json:"in_farm"`
}
