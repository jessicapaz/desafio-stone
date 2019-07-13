package models

import (
	"github.com/jessicapaz/desafio-stone/config"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create creates a user on database
func (user *User) Create() error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	stmt := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := config.GetDB().Query(stmt, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
