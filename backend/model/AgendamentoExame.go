package model

import "time"

type AgendamentoExame struct {
	Protocolo    string    `json:"protocolo" db:"protocolo"`
	Paciente     string    `json:"paciente" db:"paciente"`
	Profissional string    `json:"profissional" db:"profissional"`
	Data         time.Time `json:"data" db:"data"`
}
