package database

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// DB is a global variable for the database connection
var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	// Replace with your SQL Server connection string
	dsn := "sqlserver://CarSharingUser:CarUser@NP@localhost:1433?database=CarSharingDB"

	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to the database successfully!")
}
