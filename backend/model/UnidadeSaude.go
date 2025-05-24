package model

type UnidadeSaude struct {
	CNES     *string   `json:"cnes" db:"cnes"`
	Nome     *string   `json:"nome" db:"nome"`
	Endereco *Endereco `json:"endereco" db:"endereco"`
	Telefone *string   `json:"telefone" db:"telefone"`
}
