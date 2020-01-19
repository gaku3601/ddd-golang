package handler

import (
	"github.com/gaku3601/ddd-golang/src/application"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateUser(ctx *gin.Context)
}

type userHandler struct {
	application.UserUseCase
}

func NewUserHandler(handler application.UserUseCase) UserHandler {
	return &userHandler{handler}
}

func (u *userHandler) CreateUser(ctx *gin.Context) {
	if err := u.UserUseCase.RegisterUser(ctx, "uid", "ga"); err != nil {
		ctx.String(400, err.Error())
	} else {
		ctx.String(200, "success")
	}
}
