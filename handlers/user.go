package handlers

import (
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/renderings"
	"github.com/labstack/echo"
	"net/http"
)

// CreateUser handler
func (h *Handler) CreateUser(c echo.Context) error {
	user := new(models.User)
	resp := renderings.UserResponse{}
	e := renderings.ErrorResponse{}
	if err := c.Bind(user); err != nil {
		e.Errors = []string{"Unable to bind request"}
		return c.JSON(http.StatusUnprocessableEntity, e)
	}
	if err := Validate(user); len(err) != 0 {
		e.Errors = err
		return c.JSON(http.StatusBadRequest, e)
	}
	u, err := h.UserModel.Create(user)
	if err != nil {
		e.Errors = []string{"Unable to create user"}
		return c.JSON(http.StatusInternalServerError, e)
	}
	resp.ID = u.ID
	resp.Email = u.Email
	resp.Message = "User created"
	return c.JSON(http.StatusCreated, resp)
}
