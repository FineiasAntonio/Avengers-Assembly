package model

type Laboratorio struct {
	CNES       string `json:"cnes" db:"cnes"`
	Nome       string `json:"nome" db:"nome"`
	CNPJ       string `json:"cnpj" db:"cnpj"`
	EnderecoID string `json:"-" db:"endereco"`
	Contato    string `json:"contato" db:"contato"`

	Endereco Endereco `json:"endereco"`
}
