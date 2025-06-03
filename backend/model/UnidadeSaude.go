package model

type UnidadeSaude struct {
	CNES       string `json:"cnes" db:"cnes"`
	Nome       string `json:"nome" db:"nome"`
	CNPJ       string `json:"cnpj" db:"cnpj"`
	EnderecoID string `json:"-" db:"endereco"`
	Telefone   string `json:"telefone" db:"telefone"`

	Endereco Endereco `json:"endereco"`
}
