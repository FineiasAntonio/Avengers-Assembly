package model

import "time"

type ResultadoExame struct {
	Protocolo                       string    `json:"protocolo" bson:"protocolo"`
	DataRequisicaoRecebido          *time.Time `json:"data_requisicao_recebido" bson:"data_requisicao_recebido"`
	RejeicaoAmostra                 *[]string    `json:"rejeicao_amostra" bson:"rejeicao_amostra"`
	AdequabilidadeMaterial          *[]string    `json:"adequabilidade_material" bson:"adequabilidade_material"`
	EpiteliosAmostra				*[]string    `json:"epitelios_amostra" bson:"epitelios_amostra"`
	LimitesNormalidade              *bool      `json:"limites_normalidade" bson:"limites_normalidade"`
	AlteracoesCelulares             *[]string    `json:"alteracoes_celulares" bson:"alteracoes_celulares"`
	Microbiologia                   *[]string    `json:"microbiologia" bson:"microbiologia"`
	CelulasAtipicasEscamosas        *bool    `json:"celulas_atipicas_escamosas" bson:"celulas_atipicas_escamosas"`
	CelulasAtipicasGlandulares      *bool    `json:"celulas_atipicas_glandulares" bson:"celulas_atipicas_glandulares"`
	CelulasAtipicasOrigemIndefinida *bool    `json:"celulas_atipicas_origem_indefinida" bson:"celulas_atipicas_origem_indefinida"`
	AtipiasEmCelulasEscamosas       *[]string    `json:"atipias_em_celulas_escamosas" bson:"atipias_em_celulas_escamosas"`
	AtipiasCelulasGlandulares       *[]string    `json:"atipias_celulas_glandulares" bson:"atipias_celulas_glandulares"`
	OutrasNeoplasiasMalignas        *string    `json:"outras_neoplasias_malignas" bson:"outras_neoplasias_malignas"`
	PresencaCelulasEndometriais     *bool    `json:"presenca_celulas_endometriais" bson:"presenca_celulas_endometriais"`
	Observacoes                     *string    `json:"observacoes" bson:"observacoes"`
	ScreeningCitotecnico            *string    `json:"screening_citotecnico" bson:"screening_citotecnico"`
	Responsavel                     *Usuario   `json:"responsavel" bson:"responsavel"`
	DataResultado                   *time.Time `json:"data_resultado" bson:"data_resultado"`
}
