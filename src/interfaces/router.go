package interfaces

import (
	"github.com/gaku3601/ddd-golang/src/application"
	"github.com/gaku3601/ddd-golang/src/infrastructure"
	"github.com/gaku3601/ddd-golang/src/interfaces/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	userHandler := handler.NewUserHandler(application.NewUserUseCase(&infrastructure.UserRepository{}))
	router.GET("/", userHandler.CreateUser)
	return router
}
