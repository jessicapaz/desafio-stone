package handlers

import (
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/services"
)

type Handler struct {
	UserModel models.UserModelImpl
	LoginService services.LoginImpl
}

func NewHandler(u models.UserModelImpl, l services.LoginImpl) *Handler {
	return &Handler{
		UserModel: u,
		LoginService: l,
	}
}
