package model

import "time"

type RequisicaoExame struct {
	Protocolo                string     `json:"protocolo" db:"protocolo"`
	Paciente                 Paciente   `json:"paciente" db:"paciente"`
	MotivoExame              *string    `json:"motivo_exame" db:"motivoexame"`
	FezExamePreventivo       *bool      `json:"fez_exame_preventivo" db:"fezexamepreventivo"`
	AnoUltimoExame           *string    `json:"ano_ultimo_exame" db:"anoultimoexame"`
	UsaDIU                   *bool      `json:"usa_diu" db:"usadiu"`
	EstaGravida              *bool      `json:"esta_gravida" db:"estagravida"`
	UsaAnticoncepcional      *bool      `json:"usa_anticoncepcional" db:"usaanticoncepcional"`
	UsaHormonioMenopausa     *bool      `json:"usa_hormonio_menopausa" db:"usahormoniomenopausa"`
	FezRadioterapia          *bool      `json:"fez_radioterapia" db:"fezradioterapia"`
	DataUltimaMenstruacao    *time.Time `json:"data_ultima_menstruacao" db:"dataultimamenstruacao"`
	SangramentoAposRelacoes  *bool      `json:"sangramento_apos_relacoes" db:"sangramentoaposrelacoes"`
	SangramentoAposMenopausa *bool      `json:"sangramento_apos_menopausa" db:"sangramentoaposmenopausa"`
	InspecaoColo             *string    `json:"inspecao_colo" db:"inspecaocolo"`
	SinaisDST                *bool      `json:"sinais_dst" db:"sinaisdst"`
	DataColeta               *time.Time `json:"data_coleta" db:"datacoleta"`
	Responsavel              Usuario    `json:"responsavel" db:"responsavel"`
}
