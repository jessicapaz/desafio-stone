package models

import (
	"github.com/jessicapaz/desafio-stone/config"
)

// User model
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create creates a user on database
func (user *User) Create() error {
	stmt := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := config.GetDB().Query(stmt, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
