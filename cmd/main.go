package main

import (
	"ctipo/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	usuario := router.Group("/usuario")
	{
		usuario.POST("/login", api.LoginUsuario)

	}

	paciente := router.Group("/paciente")
	{
		paciente.GET("/:id", api.GetPaciente)
	}

	router.Run()
}
