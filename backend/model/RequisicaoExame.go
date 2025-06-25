package model

import "time"

type RequisicaoExame struct {
	Protocolo                string     `json:"protocolo" db:"protocolo"`
	PacienteID               string     `json:"paciente_id" db:"paciente"`
	MotivoExame              *string    `json:"motivo_exame" db:"motivoexame"`
	FezExamePreventivo       *bool      `json:"fez_exame_preventivo" db:"fezexamepreventivo"`
	AnoUltimoExame           *string    `json:"ano_ultimo_exame" db:"anoultimoexame"`
	UsaDIU                   *string    `json:"usa_diu" db:"usadiu"`
	EstaGravida              *string    `json:"esta_gravida" db:"estagravida"`
	UsaAnticoncepcional      *string    `json:"usa_anticoncepcional" db:"usaanticoncepcional"`
	UsaHormonioMenopausa     *string    `json:"usa_hormonio_menopausa" db:"usahormoniomenopausa"`
	FezRadioterapia          *string    `json:"fez_radioterapia" db:"fezradioterapia"`
	DataUltimaMenstruacao    *time.Time `json:"data_ultima_menstruacao" db:"dataultimamenstruacao"`
	SangramentoAposRelacoes  *string    `json:"sangramento_apos_relacoes" db:"sangramentoaposrelacoes"`
	SangramentoAposMenopausa *string    `json:"sangramento_apos_menopausa" db:"sangramentoaposmenopausa"`
	InspecaoColo             *string    `json:"inspecao_colo" db:"inspecaocolo"`
	SinaisDST                *bool      `json:"sinais_dst" db:"sinaisdst"`
	DataColeta               *time.Time `json:"data_coleta" db:"datacoleta"`
	ResponsavelRegistro      string     `json:"responsavel_registro" db:"responsavel"`
	ResultadoID              *string    `db:"resultado"`
	Status                   string     `json:"status" db:"status"`

	Paciente       Paciente       `json:"paciente"`
	Responsavel    Usuario        `json:"responsavel"`
	ResultadoExame ResultadoExame `json:"ressultado_exame"`
}

type Status string

const (
	SALVO         Status = "SALVO"
	AGUARDANDO    Status = "AGUARDANDO"
	LAUDO_EMITIDO Status = "LAUDO_EMITIDO"
)

type MensagensQuandoPassarTempoRetorno struct {
	Protocolo  string    `bson:"protocolo"`
	PacienteID string    `bson:"paciente_id"`
	DataEnvio  time.Time `bson:"data_envio"`
}
