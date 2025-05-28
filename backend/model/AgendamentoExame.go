package model

import "time"

type AgendamentoExame struct {
	Protocolo    string    `json:"protocolo" db:"protocolo"`
	Paciente     Paciente  `json:"paciente" db:"paciente"`
	Profissional Usuario   `json:"profissional" db:"profissional"`
	Data         time.Time `json:"data" db:"data"`
}
