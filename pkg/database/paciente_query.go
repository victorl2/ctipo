package database

import (
	"context"
	"ctipo/pkg/models"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/prometheus/common/log"
)

const getPacienteQuery = "SELECT nome,cpf, prontuario, data_nascimento,sexo, id_endereco FROM Paciente where prontuario = $1"
const buscaPacientesQuery = "SELECT nome,cpf, prontuario, data_nascimento,sexo, id_endereco FROM Paciente WHERE LOWER(nome) LIKE CONCAT('%', LOWER($1), '%') OR prontuario LIKE CONCAT('%', $1, '%')"

func GetPaciente(prontuario string) *models.Paciente {
	dbConnection := GetDBConnection()

	var pacientes []*models.Paciente
	if err := pgxscan.Select(context.Background(), dbConnection, &pacientes, getPacienteQuery, prontuario); err != nil {
		log.Warnf("Failed on retrieving prices for backtest from the database. %v", err.Error())
		return nil
	}

	return pacientes[0]
}

//BuscarPacientes procura por todos os pacientes que possuem a queryBusca similar ao prontuario ou nome
//espacos podem ser representados pelo sinal +
func BuscarPacientes(queryBusca string) []*models.Paciente {
	dbConnection := GetDBConnection()
	queryBusca = strings.Replace(queryBusca, "+", " ", -1)
	var pacientes []*models.Paciente
	if err := pgxscan.Select(context.Background(), dbConnection, &pacientes, buscaPacientesQuery, queryBusca); err != nil {
		log.Warnf("Failed on retrieving prices for backtest from the database. %v", err.Error())
		return nil
	}

	return pacientes
}
