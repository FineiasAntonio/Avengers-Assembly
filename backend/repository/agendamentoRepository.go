package repository

import (
	"backend/database"
	"backend/dto"
	"backend/model"
	"context"
)

type AgendamentoRepository struct {
	db *database.PostgresClient
}

func NewAgendamentoRepository(db *database.PostgresClient) *AgendamentoRepository {
	return &AgendamentoRepository{
		db: db,
	}
}

func (r *AgendamentoRepository) AgendarExame(ctx *context.Context, agendamento *model.AgendamentoExame) error {
	_, err := r.db.DB.ExecContext(
		*ctx,
		"INSERT INTO agendamento_exame (protocolo, unidade, paciente, profissional, data) VALUES ($1,$2,$3,$4,$5)",
		agendamento.Protocolo, agendamento.Unidade, agendamento.Paciente, agendamento.Profissional, agendamento.Data,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *AgendamentoRepository) ConsultarHorariosOcupados(ctx *context.Context, data string, cnes string) (*[]dto.HorariosOcupados, error) {
	resultado, err := r.db.DB.QueryContext(
		*ctx,
		"SELECT profissional.nome, agendamento.data FROM agendamento_exame agendamento JOIN usuario profissional ON agendamento.profissional = profissional.registro WHERE CAST(agendamento.data AS DATE) = $1 AND agendamento.unidade = $2 AND profissional.permissao = 'ACESSO_EXAMES'",
		data, cnes,
	)
	defer resultado.Close()

	if err != nil {
		return nil, err
	}

	var horariosOcupados []dto.HorariosOcupados

	for resultado.Next() {
		var horarioOcupado dto.HorariosOcupados
		if err := resultado.Scan(&horarioOcupado.Profissional, &horarioOcupado.Data); err != nil {
			return nil, err
		}
		horariosOcupados = append(horariosOcupados, horarioOcupado)
	}

	return &horariosOcupados, nil
}
