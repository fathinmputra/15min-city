package main

import (
	"15min-city/db"
	"15min-city/handler"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Inisialisasi database
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db.InitializeDB()
	handler.App()
}
