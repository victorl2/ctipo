package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Paciente struct {
	Nome       string
	CPF        string `json:"CPF"`
	Prontuario string `json:"Prontuario" gorm:"primary_key"`
}

func GetPaciente(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Paciente{
		Nome:       "administrador",
		CPF:        "178.347.121-98",
		Prontuario: "17394-3",
	}})
}
