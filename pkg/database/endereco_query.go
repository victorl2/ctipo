package database

import "ctipo/pkg/models"

const getPacientePorIdQuery = "SELECT * FROM endereco where id = $1"

func GetEnderecoPorId(id int) *models.Endereco {

}
