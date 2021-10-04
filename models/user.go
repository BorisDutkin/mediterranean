package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User - user database struct
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `gorm:"unique"`
	Password string
}

// BeforeCreate - before create hook, will hash the password before the creation
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {	
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14); if err != nil {
		return err
	}	
	u.Password = string(hashed)
	return
}

// ValidatePassword - validate user password
func (u *User) ValidatePassword(password string) (err error) {	
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return
}