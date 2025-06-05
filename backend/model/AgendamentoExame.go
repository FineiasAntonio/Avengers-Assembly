package model

import "time"

type AgendamentoExame struct {
	Protocolo    string    `json:"protocolo" db:"protocolo"`
	Unidade      string    `json:"unidade" db:"unidade"`
	Paciente     string    `json:"paciente" db:"paciente"`
	Profissional string    `json:"profissional" db:"profissional"`
	Data         time.Time `json:"data" db:"data"`
}
