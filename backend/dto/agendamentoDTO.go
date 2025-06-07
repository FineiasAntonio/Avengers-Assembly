package dto

import "time"

type HorariosOcupados struct {
	Profissional string    `json:"profissional"`
	Data         time.Time `json:"data"`
}
