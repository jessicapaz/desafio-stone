package handlers

import (
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/renderings"
	"github.com/jessicapaz/desafio-stone/services"
	"github.com/labstack/echo"
	"net/http"
)

type loginHandler struct {
	LoginService services.LoginService
}

func NewLoginHandler(l services.LoginService) *loginHandler {
	return &loginHandler{l}
}

// Login handler
func (h loginHandler) Login(c echo.Context) error {
	user := new(models.User)
	resp := renderings.LoginResponse{}
	if err := c.Bind(user); err != nil {
		resp.Message = "Unable to bind request"
		return c.JSON(http.StatusBadRequest, resp)
	}
	token, err := h.LoginService.Login(user)
	if err != nil {
		return echo.ErrUnauthorized
	}
	resp.Message = "Success"
	resp.Token = token
	return c.JSON(http.StatusOK, resp)
}
