package main

import (
	"fmt"

	"github.com/gaku3601/ddd-golang/src/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		user.User()
		ctx.String(200, "test")
	})

	if err := router.Run(); err != nil {
		fmt.Print("server start error")
	}
}
