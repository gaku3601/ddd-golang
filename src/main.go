package main

import (
	"fmt"

	"github.com/gaku3601/ddd-golang/src/interfaces"
)

func main() {
	router := interfaces.Router()

	if err := router.Run(); err != nil {
		fmt.Print("server start error")
	}
}
