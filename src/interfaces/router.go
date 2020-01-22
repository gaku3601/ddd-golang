package interfaces

import (
	"github.com/gaku3601/ddd-golang/src/application"
	"github.com/gaku3601/ddd-golang/src/infrastructure"
	"github.com/gaku3601/ddd-golang/src/interfaces/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:9999",
	}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))
	userHandler := handler.NewAdministratorHandler(application.NewAdministratorUseCase(&infrastructure.AdministratorRepository{}))
	r.POST("/users", userHandler.CreateAdministrator)
	return r
}
