package repository

import (
	"backend/database"
	"backend/dto"
	"context"
)

type MapaRepository struct {
	db *database.PostgresClient
}

func NewMapaRepository(db *database.PostgresClient) *MapaRepository {
	return &MapaRepository{
		db: db,
	}
}

func (r *MapaRepository) PegarQtdPacientes(ctx *context.Context) (*dto.MapaPacientesDTO, error) {
	var qtdPaciente dto.MapaPacientesDTO
	query := `SELECT COUNT(*) FROM pacientes`

	if err := r.db.DB.QueryRowContext(*ctx, query).Scan(
		&qtdPaciente,
	); err != nil {
		return nil, err
	}
	return &qtdPaciente, nil
}

func (r *MapaRepository) PegarQtdPacientesPorIdade(ctx *context.Context) (*dto.MapaPacientesPorIdadeDTO, error) {
	var qtdPacIdade dto.MapaPacientesPorIdadeDTO
	query := `SELECT
				COUNT(*),
				SUM(CASE WHEN idade >= 25 AND idade < 30 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 30 AND idade < 40 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 40 AND idade < 50 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 50 AND idade < 60 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 60 AND idade < 65 THEN 1 ELSE 0 END)
				FROM pacientes
	`
	err := r.db.DB.QueryRowContext(*ctx, query).Scan(
		&qtdPacIdade.Total,
		&qtdPacIdade.Qtd25a30,
		&qtdPacIdade.Qtd30a40,
		&qtdPacIdade.Qtd40a50,
		&qtdPacIdade.Qtd50a60,
		&qtdPacIdade.Qtd60a65,
	)

	if err != nil {
		return nil, err
	}
	return &qtdPacIdade, nil
}

func (r *MapaRepository) PegarQtdPacientesPorRaca(ctx *context.Context) (*dto.MapaPacientesPorRacaDTO, error) {
	var qtdPacRaca dto.MapaPacientesPorRacaDTO
	query := `SELECT 
				SUM(CASE WHEN raca = 'branca' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'preta' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'parda' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'amarela' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'indigena' THEN 1 ELSE 0 END)
				FROM pacientes
	`
	err := r.db.DB.QueryRowContext(*ctx, query).Scan(
		&qtdPacRaca.Branca,
		&qtdPacRaca.Preta,
		&qtdPacRaca.Parda,
		&qtdPacRaca.Amarela,
		&qtdPacRaca.Indigena,
	)

	if err != nil {
		return nil, err
	}
	return &qtdPacRaca, nil
}

func (r *MapaRepository) PegarQtdPacientesPorEscolaridade(ctx *context.Context) (*dto.MapaPacientesPorEscolaridadeDTO, error) {
	var qtdPacEsc dto.MapaPacientesPorEscolaridadeDTO
	query := `SELECT
				SUM(CASE WHEN escolaridade ILIKE '%analfabeto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%fundamental%' AND escolaridade ILIKE '%incompleto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%fundamental%' AND escolaridade ILIKE '%completo%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%medio%' AND escolaridade ILIKE '%incompleto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%medio%' AND escolaridade ILIKE '%completo%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%superior%' AND escolaridade ILIKE '%incompleto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%superior%' AND escolaridade ILIKE '%completo%' THEN 1 ELSE 0 END)
				FROM pacientes
	`
	err := r.db.DB.QueryRowContext(*ctx, query).Scan(
		&qtdPacEsc.Analfabeta,
		&qtdPacEsc.FundamentalIncompleto,
		&qtdPacEsc.FundamentalCompleto,
		&qtdPacEsc.MedioIncompleto,
		&qtdPacEsc.MedioCompleto,
		&qtdPacEsc.SuperiorIncompleto,
		&qtdPacEsc.SuperiorCompleto,
	)

	if err != nil {
		return nil, err
	}
	return &qtdPacEsc, nil
}
