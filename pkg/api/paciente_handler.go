package api

import (
	"ctipo/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPacienteHandler(c *gin.Context) {
	var prontuario = c.Param("prontuario")

	if paciente := database.GetPaciente(prontuario); paciente != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": paciente})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusPreconditionFailed, "data": nil, "info": "Não foi possível encontrar um paciente com o prontuario informado"})
	}
}

func BuscarPacientesHandler(c *gin.Context) {
	var query = c.Param("query")

	if pacientes := database.BuscarPacientes(query); pacientes != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": pacientes})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusPreconditionFailed, "data": nil, "info": "Não foi possível encontrar um paciente com o termo de busca informado"})
	}
}
