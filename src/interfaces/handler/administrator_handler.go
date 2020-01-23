package handler

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gaku3601/ddd-golang/src/application"
	"github.com/gin-gonic/gin"
)

type AdministratorHandler interface {
	CreateAdministrator(ctx *gin.Context)
}

type administratorHandler struct {
	application.AdministratorUseCase
}

func NewAdministratorHandler(handler application.AdministratorUseCase) AdministratorHandler {
	return &administratorHandler{handler}
}

func (u *administratorHandler) CreateAdministrator(ctx *gin.Context) {
	token := ctx.MustGet("token")
	claims := token.(*jwt.Token).Claims.(jwt.MapClaims)
	fmt.Println(claims["cognito:username"])
	fmt.Println(claims["custom:role"])
	createUserDto := &application.CreateAdministratorDto{}
	if err := ctx.BindJSON(createUserDto); err != nil {
		ctx.String(400, err.Error())
	}

	if err := u.AdministratorUseCase.RegisterAdministrator(ctx, createUserDto); err != nil {
		ctx.String(400, err.Error())
	} else {
		ctx.String(200, "success")
	}
}
