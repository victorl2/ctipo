package models

import "time"

type Paciente struct {
	Nome           string    `json:"nome" db:"nome"`
	CPF            string    `json:"cpf" db:"cpf"`
	Prontuario     string    `json:"prontuario" db:"prontuario"`
	DataNascimento time.Time `json:"data_nascimento" db:"data_nascimento"`
	Sexo           int       `json:"sexo" db:"sexo"`

	IdEndereco *int      `db:"id_endereco" json:"-"`
	Endereco   *Endereco `json:"endereco,omitempty"`
}
