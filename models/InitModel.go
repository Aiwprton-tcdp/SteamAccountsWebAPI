package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, username, password, dbName, os.Getenv("DB_PORT"))
	// fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print("gorm Open error")
		fmt.Print(err)
	}
	if conn == nil {
		fmt.Print("conn is null")
	}

	db = conn
	/**
	db.Debug().AutoMigrate(
		&AccessToken{},
		&Account{},
		&BuyedAccount{},
		&CSGOItem{},
		&PasswordReset{},
		&Trade{},
		&TradeItem{},
		&User{},
	)
	/**/
}

func GetDB() *gorm.DB {
	return db
}
