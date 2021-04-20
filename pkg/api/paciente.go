package api

import (
	"context"
	"ctipo/pkg/database"
	"net/http"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

type Paciente struct {
	Nome           string
	CPF            string    `json:"cpf" db:"cpf"`
	Prontuario     string    `json:"prontuario" db:"prontuario"`
	DataNascimento time.Time `json:"data_nascimento" db:"data_nascimento"`
	Sexo           int       `json:"sexo" db:"sexo"`
	IdEndereco     *int      `db:"id_endereco"`
}

func GetPacienteHandler(c *gin.Context) {
	var prontuario = c.Param("prontuario")

	if paciente := GetPaciente(prontuario); paciente != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": paciente})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusPreconditionFailed, "data": nil, "info": "Não foi possível encontrar um paciente com o prontuario informado"})
	}
}

func BuscarPacientesHandler(c *gin.Context) {
	var query = c.Param("query")

	if pacientes := BuscarPacientes(query); pacientes != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": pacientes})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusPreconditionFailed, "data": nil, "info": "Não foi possível encontrar um paciente com o termo de busca informado"})
	}
}

const getPacienteQuery = "SELECT nome,cpf, prontuario, data_nascimento,sexo, id_endereco FROM Paciente where prontuario = $1"
const buscaPacientesQuery = "SELECT nome,cpf, prontuario, data_nascimento,sexo, id_endereco FROM Paciente WHERE LOWER(nome) LIKE CONCAT('%', LOWER($1), '%') OR prontuario LIKE CONCAT('%', $1, '%')"

func GetPaciente(prontuario string) *Paciente {
	dbConnection := database.GetDBConnection()

	var pacientes []*Paciente
	if err := pgxscan.Select(context.Background(), dbConnection, &pacientes, getPacienteQuery, prontuario); err != nil {
		log.Warnf("Failed on retrieving prices for backtest from the database. %v", err.Error())
		return nil
	}

	return pacientes[0]
}

//BuscarPacientes procura por todos os pacientes que possuem a queryBusca similar ao prontuario ou nome
//espacos podem ser representados pelo sinal +
func BuscarPacientes(queryBusca string) []*Paciente {
	dbConnection := database.GetDBConnection()
	queryBusca = strings.Replace(queryBusca, "+", " ", -1)
	var pacientes []*Paciente
	if err := pgxscan.Select(context.Background(), dbConnection, &pacientes, buscaPacientesQuery, queryBusca); err != nil {
		log.Warnf("Failed on retrieving prices for backtest from the database. %v", err.Error())
		return nil
	}

	return pacientes
}
