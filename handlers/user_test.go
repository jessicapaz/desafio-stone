package handlers

import (
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type UserModel struct{}

func (u *UserModel) Create(user *models.User) (models.User, error) {
	return models.User{
		ID:    1,
		Email: "j@mail.com",
	}, nil
}

func (u *UserModel) ByEmail(email string) (models.User, error) {
	return models.User{
		ID:    1,
		Email: "j@mail.com",
	}, nil
}

func TestCreateUser(t *testing.T) {
	t.Run("returns a created user", func(t *testing.T) {
		e := echo.New()
		userJSON := `{"email":"j@mail.com","password":"123456"}`
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		u := &UserModel{}
		h := NewHandler(u, nil, nil)

		var want = `{"message":"User created","id":1,"email":"j@mail.com"}`
		if assert.NoError(t, h.CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, want+"\n", rec.Body.String())
		}
	})
	t.Run("returns a 400 status code if password is empty", func(t *testing.T) {
		e := echo.New()
		userJSON := `{"email":"j@mail.com"}`
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		u := &UserModel{}
		h := NewHandler(u, nil, nil)

		var want = `{"message":"password must not be empty"}`
		if assert.NoError(t, h.CreateUser(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, want+"\n", rec.Body.String())
		}
	})
}
