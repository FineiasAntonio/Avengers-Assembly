package repository

import (
	"backend/database"
	"backend/model"
	"context"
	"database/sql"
	"fmt"
)

type PacienteRepository struct {
	db *database.PostgresClient
}

func NewPacienteRepository(db *database.PostgresClient) *PacienteRepository {
	return &PacienteRepository{db: db}
}

func (p *PacienteRepository) GetPacienteByCartaoSUS(ctx *context.Context, cartaoSUS string) (*model.Paciente, error) {
	row := p.db.DB.QueryRowContext(*ctx, "SELECT * FROM paciente WHERE cartaosus = $1", cartaoSUS)

	var paciente model.Paciente

	err := row.Scan(
		&paciente.CartaoSUS,
		&paciente.Prontuario,
		&paciente.Nome,
		&paciente.NomeMae,
		&paciente.CPF,
		&paciente.DataNascimento,
		&paciente.Idade,
		&paciente.Raca,
		&paciente.Nacionalidade,
		&paciente.Escolaridade,
		&paciente.Telefone,
		&paciente.Endereco,
		&paciente.Senha,
		&paciente.PrimeiroAcesso,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Paciente não encontrado")
		}
		return nil, fmt.Errorf("Erro ao buscar paciente: %v", err)
	}

	return &paciente, nil
}
func (p *PacienteRepository) GetPacienteByCartaoCPF(ctx *context.Context, cpf string) (*model.Paciente, error) {
	row := p.db.DB.QueryRowContext(*ctx, "SELECT * FROM paciente WHERE cpf = $1", cpf)

	var paciente model.Paciente

	err := row.Scan(
		&paciente.CartaoSUS,
		&paciente.Prontuario,
		&paciente.Nome,
		&paciente.NomeMae,
		&paciente.CPF,
		&paciente.DataNascimento,
		&paciente.Idade,
		&paciente.Raca,
		&paciente.Nacionalidade,
		&paciente.Escolaridade,
		&paciente.Telefone,
		&paciente.Endereco,
		&paciente.Senha,
		&paciente.PrimeiroAcesso,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Paciente não encontrado")
		}
		return nil, fmt.Errorf("Erro ao buscar paciente: %v", err)
	}

	return &paciente, nil
}

func (p *PacienteRepository) CadastrarPaciente(ctx *context.Context, paciente *model.Paciente) error {
	_, err := p.db.DB.ExecContext(*ctx, `
		INSERT INTO paciente (
			cartaosus, prontuario, nome, nomemae, cpf,
			datanascimento, idade, raca, nacionalidade,
			escolaridade, telefone, endereco, senha,
			primeiroacesso
		) VALUES ($1, $2, $3, $4, $5,
			$6, $7, $8, $9,
			$10, $11, $12, $13, $14)
	`, paciente.CartaoSUS, paciente.Prontuario, paciente.Nome,
		paciente.NomeMae, paciente.CPF, paciente.DataNascimento,
		paciente.Idade, paciente.Raca, paciente.Nacionalidade,
		paciente.Escolaridade, paciente.Telefone, paciente.EnderecoID, paciente.Senha, paciente.PrimeiroAcesso)

	if err != nil {
		return fmt.Errorf("Erro ao cadastrar paciente: %v", err)
	}

	return nil
}
func (p *PacienteRepository) AlterarSenha(ctx *context.Context, cpf string, novaSenha string) error {
	query := `UPDATE paciente SET senha = $1, primeiroacesso = false WHERE cpf = $2`
	_, err := p.db.DB.ExecContext(*ctx, query, novaSenha, cpf)
	return err
}
