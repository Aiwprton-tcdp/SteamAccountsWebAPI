package main

import (
	"log"
	"saw/app"
	"saw/models"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	models.Initialize()
	app.Initialize()
}
