package model

import "time"

type Paciente struct {
	CartaoSUS      string    `json:"cartao_sus" db:"cartaosus"`
	Prontuario     string    `json:"prontuario" db:"prontuario"`
	Nome           string    `json:"nome" db:"nome"`
	NomeMae        string    `json:"nome_mae" db:"nomemae"`
	CPF            string    `json:"cpf" db:"cpf"`
	DataNascimento time.Time `json:"data_nascimento" db:"datanascimento"`
	Idade          int       `json:"idade" db:"idade"`
	Raca           *string   `json:"raca" db:"raca"`
	Nacionalidade  string    `json:"nacionalidade" db:"nacionalidade"`
	Escolaridade   *string   `json:"escolaridade" db:"escolaridade"`
	Telefone       string    `json:"telefone" db:"telefone"`
	EnderecoID     string    `json:"-" db:"endereco"`
	Senha          string    `json:"-" db:"senha"`
	PrimeiroAcesso bool      `json:"primeiro_acesso" db:"primeiroacesso"`

	Endereco Endereco            `json:"endereco"`
	Agenda   *[]AgendamentoExame `json:"agenda"`
}
