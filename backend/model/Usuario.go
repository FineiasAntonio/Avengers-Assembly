package model

type Usuario struct {
	Registro         string  `json:"registro" db:"registro"`
	Nome             string  `json:"nome" db:"nome"`
	CPF              string  `json:"cpf" db:"cpf"`
	Email            *string `json:"email" db:"email"`
	Telefone         string  `json:"telefone" db:"telefone"`
	Senha            string  `json:"-" db:"senha"`
	UnidadeSaudeCNES *string `json:"unidade_saude_cnes" db:"unidadesaude"`
	LaboratorioCNES  *string `json:"laboratorio_cnes" db:"laboratorio"`
	Permissao        string  `json:"permissao" db:"permissao"`
	PrimeiroAcesso   bool    `json:"primeiro_acesso" db:"primeiroacesso"`

	UnidadeSaude *UnidadeSaude `json:"unidade_saude"`
	Laboratorio  *Laboratorio  `json:"laboratorio"`
}

type Permissao string

const (
	ACESSO_ATENDIMENTO Permissao = "ACESSO_ATENDIMENTO"
	ACESSO_EXAMES      Permissao = "ACESSO_EXAMES"
	ACESSO_LABORATORIO Permissao = "ACESSO_LABORATORIO"
	GESTAO             Permissao = "GESTAO"
	ADMINISTRADOR      Permissao = "ADMINISTRADOR"
)
