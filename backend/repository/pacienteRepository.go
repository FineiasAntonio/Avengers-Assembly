package repository

import (
	"backend/database"
	"backend/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type PacienteRepository struct {
	db *database.PostgresClient
}

func NewPacienteRepository(db *database.PostgresClient) *PacienteRepository {
	return &PacienteRepository{db: db}
}

var (
	ErroPacienteNaoEncontrado = errors.New("Paciente não encontrado")
)

func (p *PacienteRepository) GetPacienteByCartaoSUS(ctx *context.Context, cartaoSUS string) (*model.Paciente, error) {
	row := p.db.DB.QueryRowContext(*ctx, "SELECT * FROM paciente JOIN endereco ON paciente.endereco = endereco.endereco_id WHERE cartaosus = $1", cartaoSUS)

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
		&paciente.EnderecoID,
		&paciente.Senha,
		&paciente.PrimeiroAcesso,
		&paciente.Endereco.EnderecoID,
		&paciente.Endereco.Logradouro,
		&paciente.Endereco.Numero,
		&paciente.Endereco.Complemento,
		&paciente.Endereco.Bairro,
		&paciente.Endereco.CodMunicipio,
		&paciente.Endereco.Municipio,
		&paciente.Endereco.UF,
		&paciente.Endereco.CEP,
		&paciente.Endereco.PontoReferencia,
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
	row := p.db.DB.QueryRowContext(*ctx, "SELECT * FROM paciente JOIN endereco ON paciente.endereco = endereco.endereco_id WHERE cpf = $1", cpf)

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
		&paciente.EnderecoID,
		&paciente.Senha,
		&paciente.PrimeiroAcesso,
		&paciente.Endereco.EnderecoID,
		&paciente.Endereco.Logradouro,
		&paciente.Endereco.Numero,
		&paciente.Endereco.Complemento,
		&paciente.Endereco.Bairro,
		&paciente.Endereco.CodMunicipio,
		&paciente.Endereco.Municipio,
		&paciente.Endereco.UF,
		&paciente.Endereco.CEP,
		&paciente.Endereco.PontoReferencia,
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

func (p *PacienteRepository) ExistePaciente(ctx *context.Context, cartao_sus string) (bool, error) {
	var existe bool
	query := "SELECT EXISTS(SELECT 1 FROM paciente WHERE cartaosus = $1)"
	err := p.db.DB.QueryRowContext(*ctx, query, cartao_sus).Scan(&existe)

	if err != nil {
		return false, err
	}

	return existe, nil
}
