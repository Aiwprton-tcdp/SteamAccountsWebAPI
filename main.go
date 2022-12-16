package main

import (
	"log"
	"sawa/app"
	"sawa/models"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	models.Initialize()
	app.Initialize()
}
