package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Routes used in the application
	api := router.Group("/api/")
	//api.Use(middleware.JWT())
	{
		api.GET("/paciente/:prontuario", GetPacienteHandler)
		api.GET("/paciente/busca/:query", BuscarPacientesHandler)
		api.GET("/usuario/:email", GetUsuario)
	}

	//public := router.Group("/public")
	//public.GET("/aa", GetPaciente)

	return router
}
