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
		protocolo, paciente, motivoexame, fezexamepreventido, anoultimoexame, usadiu,
		estagravida, usaanticoncepcional, usahormoniomonopausa, fezradioterapia,
		dataultimamenstruacao, sangramentoaposrelacoes, sangramentoaposmenopausa,
		inspecaocolo, sinaisdst, datacoleta, responsavel, resultado, status) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`,
		re.Protocolo, re.PacienteID, re.MotivoExame, re.FezExamePreventivo,
		re.AnoUltimoExame, re.UsaDIU, re.EstaGravida,
		re.UsaAnticoncepcional, re.UsaHormonioMenopausa,
		re.FezRadioterapia, re.DataUltimaMenstruacao, re.SangramentoAposRelacoes,
		re.SangramentoAposMenopausa, re.InspecaoColo, re.SinaisDST, re.DataColeta,
		re.ResponsavelRegistro, re.ResultadoID, re.Status)

	if err != nil {
		return fmt.Errorf("Erro ao Cadastrar Requisição Exame: %v", err)
	}
	return nil
}
