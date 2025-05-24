package model

import "time"

type ResultadoExame struct {
	Protocolo                       string    `json:"protocolo" db:"protocolo"`
	DataRequisicaoRecebido          time.Time `json:"data_requisicao_recebido" db:"data_requisicao_recebido"`
	RejeicaoAmostra                 string    `json:"rejeicao_amostra" db:"rejeicao_amostra"`
	AdequabilidadeMaterial          string    `json:"adequabilidade_material" db:"adequabilidade_material"`
	LimitesNormalidade              bool      `json:"limites_normalidade" db:"limites_normalidade"`
	AlteracoesCelulares             string    `json:"alteracoes_celulares" db:"alteracoes_celulares"`
	Microbiologia                   string    `json:"microbiologia" db:"microbiologia"`
	CelulasAtipicasEscamosas        string    `json:"celulas_atipicas_escamosas" db:"celulas_atipicas_escamosas"`
	CelulasAtipicasGlandulares      string    `json:"celulas_atipicas_glandulares" db:"celulas_atipicas_glandulares"`
	CelulasAtipicasOrigemIndefinida string    `json:"celulas_atipicas_origem_indefinida" db:"celulas_atipicas_origem_indefinida"`
	AtipiasEmCelulasEscamosas       string    `json:"atipias_em_celulas_escamosas" db:"atipias_em_celulas_escamosas"`
	AtipiasCelulasGlandulares       string    `json:"atipias_celulas_glandulares" db:"atipias_celulas_glandulares"`
	OutrasNeoplasiasMalignas        string    `json:"outras_neoplasias_malignas" db:"outras_neoplasias_malignas"`
	PresencaCelulasEndometriais     string    `json:"presenca_celulas_endometriais" db:"presenca_celulas_endometriais"`
	Observacoes                     string    `json:"observacoes" db:"observacoes"`
	ScreeningCitotecnico            string    `json:"screening_citotecnico" db:"screening_citotecnico"`
	Responsavel                     *Usuario  `json:"responsavel" db:"responsavel"`
	DataResultado                   time.Time `json:"data_resultado" db:"data_resultado"`
}
