package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `gorm:"unique"`
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {	
	return
}