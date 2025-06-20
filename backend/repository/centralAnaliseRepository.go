package repository

import (
	"backend/database"
	"backend/dto"
	"context"
)

type CentralAnaliseRepository struct {
	db *database.PostgresClient
}

func NewCentralAnaliseRepository(db *database.PostgresClient) *CentralAnaliseRepository {
	return &CentralAnaliseRepository{
		db: db,
	}
}

func (r *CentralAnaliseRepository) PegarQtdPacientes(ctx *context.Context) (*dto.GraficoPacientesDTO, error) {
	var qtdPaciente dto.GraficoPacientesDTO
	query := `SELECT COUNT(*) FROM paciente`

	if err := r.db.DB.QueryRowContext(*ctx, query).Scan(
		&qtdPaciente.QuantidadePacientes,
	); err != nil {
		return nil, err
	}
	return &qtdPaciente, nil
}

func (r *CentralAnaliseRepository) PegarQtdPacientesPorIdade(ctx *context.Context) (*dto.GraficoPacientesPorIdadeDTO, error) {
	var qtdPacIdade dto.GraficoPacientesPorIdadeDTO
	query := `SELECT
				COUNT(*),
				SUM(CASE WHEN idade >= 25 AND idade < 30 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 30 AND idade < 40 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 40 AND idade < 50 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 50 AND idade < 60 THEN 1 ELSE 0 END),
				SUM(CASE WHEN idade >= 60 AND idade < 65 THEN 1 ELSE 0 END)
				FROM paciente
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

func (r *CentralAnaliseRepository) PegarQtdPacientesPorRaca(ctx *context.Context) (*dto.GraficoPacientesPorRacaDTO, error) {
	var qtdPacRaca dto.GraficoPacientesPorRacaDTO
	query := `SELECT 
				SUM(CASE WHEN raca = 'Branca' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'Negra' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'Parda' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'Amarela' THEN 1 ELSE 0 END),
				SUM(CASE WHEN raca = 'Indigena' THEN 1 ELSE 0 END)
				FROM paciente
	`
	err := r.db.DB.QueryRowContext(*ctx, query).Scan(
		&qtdPacRaca.Branca,
		&qtdPacRaca.Negra,
		&qtdPacRaca.Parda,
		&qtdPacRaca.Amarela,
		&qtdPacRaca.Indigena,
	)

	if err != nil {
		return nil, err
	}
	return &qtdPacRaca, nil
}

func (r *CentralAnaliseRepository) PegarQtdPacientesPorEscolaridade(ctx *context.Context) (*dto.GraficoPacientesPorEscolaridadeDTO, error) {
	var qtdPacEsc dto.GraficoPacientesPorEscolaridadeDTO
	query := `SELECT
				SUM(CASE WHEN escolaridade ILIKE '%analfabeto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%fundamental%' AND escolaridade ILIKE '%incompleto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%fundamental%' AND escolaridade ILIKE '%completo%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%médio%' AND escolaridade ILIKE '%incompleto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%médio%' AND escolaridade ILIKE '%completo%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%superior%' AND escolaridade ILIKE '%incompleto%' THEN 1 ELSE 0 END),
				SUM(CASE WHEN escolaridade ILIKE '%superior%' AND escolaridade ILIKE '%completo%' THEN 1 ELSE 0 END)
				FROM paciente
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

func (r *CentralAnaliseRepository) PegarQtdPacientesPorRegiao(ctx *context.Context) (*[]dto.MapaPacientesPorRegiao, error) {
	var resultados []dto.MapaPacientesPorRegiao

	query := `
		SELECT e.bairro, COUNT(*)
		FROM paciente p
		INNER JOIN endereco e ON p.endereco = e.endereco_id
		GROUP BY e.bairro
	`

	rows, err := r.db.DB.QueryContext(*ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var qtdPacReg dto.MapaPacientesPorRegiao
		err := rows.Scan(&qtdPacReg.Bairro, &qtdPacReg.Quantidade)

		if err != nil {
			return nil, err
		}

		resultados = append(resultados, qtdPacReg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &resultados, nil
}
