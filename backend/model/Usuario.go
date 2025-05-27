package model

type Usuario struct {
	Nome      string `json:"nome" db:"nome"`
	CPF       string `json:"cpf" db:"cpf"`
	Email     string `json:"email" db:"email"`
	Telefone  string `json:"telefone" db:"telefone"`
	Registro  string `json:"registro" db:"registro"`
	Senha     string `json:"-" db:"senha"`
	Permissao string `json:"permissao" db:"permissao"`
}

type Permissao string

const (
	CRIA_EXAME    Permissao = "CRIA_EXAME"
	CRIA_PACIENTE Permissao = "CRIA_PACIENTE"
	CRIA_LAUDO    Permissao = "CRIA_LAUDO"
	GESTAO        Permissao = "GESTAO"
	ADMINISTRADOR Permissao = "ADMINISTRADOR"
)
