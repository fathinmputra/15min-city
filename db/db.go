package db

import (
	"15min-city/entity"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

// InitializeDB initializes the database connection and creates tables
func InitializeDB() {
	handleDBConnection()
	// createTable()
}

// handleDBConnection establishes a connection to the database
func handleDBConnection() {
	// Load environment variables from .env file

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=" + os.Getenv("DB_CHARSET") + "&parseTime=" + os.Getenv("DB_PARSE_TIME") + "&loc=" + os.Getenv("DB_LOC")
	_, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if dbErr != nil {
		log.Fatalf("failed to connect to database: %v", dbErr)
	}
}

// createTable creates tables in the database using AutoMigrate
func createTable() {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Dataset{},
		&entity.Corridor_Route{},
	)
	if err != nil {
		log.Fatalf("failed to create tables: %v", err)
	}
}

// GetDBInstance returns the database instance
func GetDBInstance() *gorm.DB {
	return db
}
