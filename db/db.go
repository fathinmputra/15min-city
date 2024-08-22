package db

import (
	"15min-city/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	dbErr error
)

// InitializeDB initializes the database connection and creates tables
func InitializeDB() {
	handleDBConnection()
	createTable()
}

// handleDBConnection establishes a connection to the database
func handleDBConnection() {
	dsn := "root:@tcp(127.0.0.1:3306)/15min-city?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatalf("failed to connect to database: %v", dbErr)
	}
}

// createTable creates tables in the database using AutoMigrate
func createTable() {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Dataset{},
	)
	if err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}
}

// GetDBInstance returns the database instance
func GetDBInstance() *gorm.DB {
	return db
}
