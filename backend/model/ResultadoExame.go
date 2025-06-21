package model

import "time"

type ResultadoExame struct {
	ProtocoloExame    string     `json:"protocolo_exame" bson:"protocolo_exame"`
	Laboratorio       string     `json:"laboratorio" bson:"laboratorio"`
	DataExameRecebido *time.Time `json:"data_exame_recebido" bson:"data_exame_recebido"`

	AvaliacaoPreAnalitica  AvaliacaoPreAnalitica  `json:"avaliacao_pre_analitica" bson:"avaliacao_pre_analitica"`
	AdequabilidadeMaterial AdequabilidadeMaterial `json:"adequabilidade_material" bson:"adequabilidade_material"`
	DiagnosticoDescritivo  DiagnosticoDescritivo  `json:"diagnostico_descritivo" bson:"diagnostico_descritivo"`
	Microbiologia          Microbiologia          `json:"microbiologia" bson:"microbiologia"`

	OutrasNeoplasiasMalignas    *string    `json:"outras_neoplasias_malignas" bson:"outras_neoplasias_malignas"`
	PresencaCelulasEndometriais bool       `json:"presenca_celulas_endometriais" bson:"presenca_celulas_endometriais"`
	Observacoes                 string     `json:"observacoes" bson:"observacoes"`
	ScreeningCitotecnico        *string    `json:"screening_citotecnico" bson:"screening_citotecnico"`
	RegistroResponsavel         string     `json:"registro_responsavel" bson:"registro_responsavel"`
	DataEmissaoLaudo            *time.Time `json:"data_emissao_laudo" bson:"data_emissao_laudo"`

	Status StatusResultadoExame `json:"status" bson:"status"`
}

type AvaliacaoPreAnalitica struct {
	RejeicaoAmostra        []string `json:"rejeicao_amostra" bson:"rejeicao_amostra"`
	EpiteliosRepresentados []string `json:"epitelios_representados" bson:"epitelios_representados"`
}

type AdequabilidadeMaterial struct {
	Satisfatoria   bool     `json:"satisfatoria" bson:"satisfatoria"`
	Insatisfatoria []string `json:"insatisfatoria" bson:"insatisfatoria"`
}

type DiagnosticoDescritivo struct {
	DentroLimitesNormalidade bool     `json:"dentro_limites_normalidade" bson:"dentro_limites_normalidade"`
	AlteracoesBenignas       []string `json:"alteracoes_benignas" bson:"alteracoes_benignas"`
}

type Microbiologia struct {
	Microorganismos                 []string `json:"microorganismos" bson:"microorganismos"`
	CelulasAtipicasEscamosas        string   `json:"celulas_atipicas_escamosas" bson:"celulas_atipicas_escamosas"`
	CelulasAtipicasGlandulares      string   `json:"celulas_atipicas_glandulares" bson:"celulas_atipicas_glandulares"`
	CelulasAtipicasOrigemIndefinida string   `json:"celulas_atipicas_origem_indefinida" bson:"celulas_atipicas_origem_indefinida"`
	AtipiasEscamosas                []string `json:"atipias_escamosas" bson:"atipias_escamosas"`
	AtipiasGlandulares              []string `json:"atipias_glandulares" bson:"atipias_glandulares"`
}

type StatusResultadoExame string

var (
	StatusPendente Status = "PENDENTE"
	StatusSalvo    Status = "SALVO"
	StatusEmitido  Status = "EMITIDO"
)
