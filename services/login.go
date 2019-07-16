package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jessicapaz/desafio-stone/config"
	"github.com/jessicapaz/desafio-stone/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// LoginService describes all methods of a LoginService
type LoginImpl interface {
	Login(user *models.User) (token string, err error)
}

type jwtLoginClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginService struct{}

// NewLoginService creates a new loginService
func NewLoginService() *LoginService {
	return &LoginService{}
}

// Login service create a jwt token for a user
func (l *LoginService) Login(user *models.User) (token string, err error) {
	userModel := models.NewUserModel(config.GetDB())
	u, err := userModel.ByEmail(user.Email)
	if err != nil {
		return "", errors.New("User doesn't exist")
	}
	if u.Email != user.Email {
		return "", errors.New("Invalid login credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("Invalid login credentials")
	}
	token, err = createToken(user.Email)
	if err != nil {
		return "", errors.New("Unable to create token")
	}
	return token, nil
}

func createToken(email string) (token string, err error) {
	claims := &jwtLoginClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
	if err != nil {
		return token, err
	}
	return token, nil
}
