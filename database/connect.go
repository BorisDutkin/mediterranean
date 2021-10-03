package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect - setup database connection
func Connect() {
	connection, err := gorm.Open(sqlite.Open("./database/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection
	
	Migrate()
}