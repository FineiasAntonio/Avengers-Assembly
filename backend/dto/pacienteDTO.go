package dto

import (
	"backend/model"
	"time"
)

type PacienteDTO struct {
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

	Endereco model.Endereco            `json:"endereco"`
	Agenda   *[]model.AgendamentoExame `json:"agenda"`
}
