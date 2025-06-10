package dto

type UsuarioDTO struct {
	Registro         string  `json:"registro" db:"registro"`
	Nome             string  `json:"nome" db:"nome"`
	CPF              string  `json:"cpf" db:"cpf"`
	Email            *string `json:"email" db:"email"`
	Telefone         string  `json:"telefone" db:"telefone"`
	UnidadeSaudeCNES *string `json:"-" db:"unidadesaude"`
	LaboratorioCNES  *string `json:"-" db:"laboratorio"`
}
