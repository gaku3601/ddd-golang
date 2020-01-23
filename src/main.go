package main

import (
	"fmt"

	"github.com/gaku3601/ddd-golang/src/interfaces"
)

func main() {
	router, err := interfaces.Router()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := router.Run(); err != nil {
		fmt.Print("server start error")
		return
	}
}
