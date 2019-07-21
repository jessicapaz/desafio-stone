package handlers

import (
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/renderings"
	"github.com/labstack/echo"
	"net/http"
)

// Login handler
func (h *Handler) Login(c echo.Context) error {
	user := new(models.User)
	resp := renderings.LoginResponse{}
	e := renderings.ErrorResponse{}
	if err := c.Bind(user); err != nil {
		e.Errors = []string{"Unable to bind request"}
		return c.JSON(http.StatusUnprocessableEntity, e)
	}
	token, err := h.LoginService.Login(user)
	if err != nil {
		return echo.ErrUnauthorized
	}
	resp.Message = "Success"
	resp.Token = token
	return c.JSON(http.StatusOK, resp)
}
