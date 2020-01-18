package main

import (
	"fmt"

	"github.com/gaku3601/ddd-golang/src/application"
	"github.com/gaku3601/ddd-golang/src/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		userUseCase := application.NewUserUseCase(&infrastructure.UserRepository{})
		if err := userUseCase.RegisterUser(ctx, "uid", "g"); err != nil {
			ctx.String(400, err.Error())
		} else {
			ctx.String(200, "success")
		}
	})

	if err := router.Run(); err != nil {
		fmt.Print("server start error")
	}
}
