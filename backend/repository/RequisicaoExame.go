package repository

import (
	"backend/database"
	"backend/model"
	"context"
	"database/sql"
	"fmt"
)

type RequisicaoExameRepository struct {
	db *database.PostgresClient
}

func NewRequisicaoExameRepository(db *database.PostgresClient) *RequisicaoExameRepository {
	RequisicaoExameRepository := RequisicaoExameRepository{db: db}
	return &RequisicaoExameRepository
}

func (r *RequisicaoExameRepository) GetRequisicaoExameByProtocolo(
	ctx *context.Context, protocolo string) (*model.RequisicaoExame, error) {
	var re model.RequisicaoExame

	row := r.db.DB.QueryRowContext(*ctx,
		"SELECT * FROM requisicao_exame WHERE protocolo = $1", protocolo)

	err := row.Scan(&re.Protocolo, &re.Paciente, &re.MotivoExame, &re.FezExamePreventivo,
		&re.AnoUltimoExame, &re.UsaDIU, &re.EstaGravida, &re.UsaAnticoncepcional,
		&re.UsaHormonioMenopausa, &re.FezRadioterapia, &re.DataUltimaMenstruacao,
		&re.SangramentoAposRelacoes, &re.SangramentoAposMenopausa, &re.InspecaoColo,
		&re.SinaisDST, &re.DataColeta, &re.Responsavel)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Requisicao Exame Não Encontrada")
		}
		return nil, fmt.Errorf("Erro ao Buscar Requisiçao Exame %v", err)
	}
	return &re, nil
}

func (r *RequisicaoExameRepository) CadastrarRequisicaoExame(
	ctx *context.Context, re *model.RequisicaoExame) error {

	_, err := r.db.DB.ExecContext(*ctx,
		`INSERT INTO requisicao_exame (
		protocolo, paciente, motivo_exame, fez_exame_preventido, ano_ultimo_exame, usa_diu,
		esta_gravida, usa_anticoncepcional, usa_hormonio_monopausa, fez_radioterapia,
		data_ultima_menstruacao, sangramento_apos_relacoes, sangramento_apos_menopausa,
		inspecao_colo, sinais_dst, data_coleta, responsavel) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`,
		re.Protocolo, re.Paciente, re.MotivoExame, re.FezExamePreventivo,
		re.AnoUltimoExame, re.UsaDIU, re.EstaGravida,
		re.UsaAnticoncepcional, re.UsaHormonioMenopausa,
		re.FezRadioterapia, re.DataUltimaMenstruacao, re.SangramentoAposRelacoes,
		re.SangramentoAposMenopausa, re.InspecaoColo, re.SinaisDST, re.DataColeta,
		re.Responsavel)

	if err != nil {
		return fmt.Errorf("Erro ao Cadastrar Requisição Exame: %v", err)
	}
	return nil
}
