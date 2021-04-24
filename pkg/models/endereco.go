package models

type Endereco struct {
	Id              int    `json:"id" db:"id"`
	Bairro          string `json:"bairro" db:"bairro"`
	CEP             string `json:"cep" db:"cep"`
	Cidade          string `json:"cidade" db:"cidade"`
	Complemento     string `json:"complemento" db:"complemento"`
	Numero          string `json:"numero" db:"numero"`
	PontoReferencia string `json:"ponto_referencia" db:"ponto_referencia"`
	Rua             string `json:"rua" db:"rua"`
}
