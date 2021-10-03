package database

import "github.com/BorisDutkin/mediterranean/models"

// Migrate - database models migrations
func Migrate() {
	DB.AutoMigrate(&models.User{})
}