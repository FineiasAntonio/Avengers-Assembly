package model

type Usuario struct {
	Nome      *string `json:"nome" db:"nome"`
	CPF       *string `json:"cpf" db:"cpf"`
	Email     string  `json:"email" db:"email"`
	Telefone  *string `json:"telefone" db:"telefone"`
	Registro  *string `json:"registro" db:"registro"`
	Senha     *string `json:"-" db:"senha"`
	Permissao *[]int  `json:"permissao" db:"permissao"`
}

type Permissao int

const (
	CRIA_EXAME Permissao = iota
	CRIA_PACIENTE
	CRIA_LAUDO
	GESTAO
	ADMINISTRADOR
)
