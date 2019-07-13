package handlers

import (
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/renderings"
	"github.com/labstack/echo"
	"net/http"
)

// CreateUser handler
func CreateUser(c echo.Context) error {
	user := new(models.User)
	resp := renderings.UserResponse{}
	if err := c.Bind(user); err != nil {
		resp.Email = user.Email
		resp.Message = "Unable to bind request"
		return c.JSON(http.StatusBadRequest, resp)
	}
	err := user.Create()
	if err != nil {
		resp.Email = user.Email
		resp.Message = "Unable to create user"
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp.Email = user.Email
	resp.Message = "User created!"
	return c.JSON(http.StatusCreated, resp)
}
