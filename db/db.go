package db

import (
	"15min-city/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	dsn := "xminutecity:xmc2024@tcp(pwk.its.ac.id:3306)/xminutecity?charset=utf8mb4&parseTime=True&loc=Local"
	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if dbErr != nil {
		log.Fatalf("failed to connect to database: %v", dbErr)
	}
}

// func handleDBConnection() {
// 	dsn := "root:@tcp(127.0.0.1:3306)/15min-city?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info), // Ini akan menampilkan log semua query SQL
// 	})
// 	if dbErr != nil {
// 		log.Fatalf("failed to connect to database: %v", dbErr)
// 	}
// }

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
