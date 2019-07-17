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

type loginService struct{}

func (u *loginService) Login(user *models.User) (string, error) {
	return "secret", nil
}

func TestLogin(t *testing.T) {
	e := echo.New()
	loginJSON := `{"email":"j@mail.com","password":"123456"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	l := &loginService{}
	h := NewHandler(nil, l, nil)

	var want = `{"message":"Success","token":"secret"}`
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, want+"\n", rec.Body.String())
	}
}
