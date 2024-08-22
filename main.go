package main

import (
	"15min-city/db"
	"15min-city/handler"
)

func main() {
	// Inisialisasi database
	db.InitializeDB()
	handler.App()
}
