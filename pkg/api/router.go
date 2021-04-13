package api

import (
	"ctipo/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Routes used in the application
	api := router.Group("/api/")
	api.Use(middleware.JWT())
	{
		api.GET("/paciente/:id", GetPaciente)
		api.GET("/usuario/:email", GetUsuario)
	}

	return router
}
