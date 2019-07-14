package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserModelImpl describes all methods of a UserModelImpl
type UserModelImpl interface {
	Create(user *User) (User, error)
	ByEmail(email string) (User, error)
}

type UserModel struct {
	db *sql.DB
}

// NewUserModel creates a new UserModel
func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

// Validate validates user data
func (user *User) Validate() error {
	if user.Password == "" {
		return errors.New("password must not be empty")
	}
	if user.Email == "" {
		return errors.New("email must not be empty")
	}
	return nil
}

// Create creates a user on database
func (u *UserModel) Create(user *User) (User, error) {
	newUser := User{}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	stmt := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *;"
	result := u.db.QueryRow(stmt, user.Email, user.Password)
	if err := result.Scan(&newUser.ID, &newUser.Email, &newUser.Password); err != nil {
		return newUser, err
	}
	return newUser, nil
}

// ByEmail gets user by email
func (u *UserModel) ByEmail(email string) (User, error) {
	user := User{}
	stmt := "SELECT email, password FROM users WHERE email=$1"
	result := u.db.QueryRow(stmt, email)
	if err := result.Scan(&user.Email, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}
