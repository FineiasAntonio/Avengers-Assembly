package repository

import (
	"backend/database"
	"backend/model"
	"backend/util"
	"context"
	"errors"
)

type UnidadeRepository struct {
	db *database.PostgresClient
}

func NewUnidadeRepository(db *database.PostgresClient) *UnidadeRepository {
	return &UnidadeRepository{db: db}
}

var (
	ErroUnidadeNaoEncontrada = errors.New("Unidade não encontrado")
)

func (respository *UnidadeRepository) ListarUnidade(ctx *context.Context, cnes string) (*model.UnidadeSaude, error) {
	query := `SELECT * FROM unidade_saude JOIN endereco ON unidade_saude.endereco = endereco.endereco_id WHERE cnes = $1`
	row := respository.db.DB.QueryRowContext(
		*ctx,
		query,
		cnes,
	)

	unidade := &model.UnidadeSaude{}
	if err := row.Scan(
		&unidade.CNES,
		&unidade.Nome,
		&unidade.CNPJ,
		&unidade.EnderecoID,
		&unidade.Telefone,
		&unidade.Endereco.EnderecoID,
		&unidade.Endereco.Logradouro,
		&unidade.Endereco.Numero,
		&unidade.Endereco.Complemento,
		&unidade.Endereco.Bairro,
		&unidade.Endereco.CodMunicipio,
		&unidade.Endereco.Municipio,
		&unidade.Endereco.UF,
		&unidade.Endereco.CEP,
		&unidade.Endereco.PontoReferencia,
	); err != nil {
		return nil, err
	}

	return unidade, nil
}

func (repository *UnidadeRepository) CadastrarUnidade(ctx *context.Context, requisicao *model.UnidadeSaude) error {
	cnes := util.GerarId(10)
	query := `INSERT INTO unidade_saude (cnes, nome, cnpj, endereco, telefone) VALUES ($1, $2, $3, $4, $5)`

	_, err := repository.db.DB.ExecContext(
		*ctx,
		query,
		cnes, requisicao.Nome, requisicao.CNPJ, requisicao.EnderecoID, requisicao.Telefone,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UnidadeRepository) ListarLaboratorio(ctx *context.Context, cnes string) (*model.Laboratorio, error) {
	query := `SELECT * FROM laboratorio JOIN endereco ON laboratorio.endereco = endereco.endereco_id WHERE cnes = $1`

	row := repository.db.DB.QueryRowContext(*ctx, query, cnes)

	laboratorio := &model.Laboratorio{}
	if err := row.Scan(
		&laboratorio.CNES,
		&laboratorio.Nome,
		&laboratorio.CNPJ,
		&laboratorio.EnderecoID,
		&laboratorio.Telefone,
		&laboratorio.Endereco.EnderecoID,
		&laboratorio.Endereco.Logradouro,
		&laboratorio.Endereco.Numero,
		&laboratorio.Endereco.Complemento,
		&laboratorio.Endereco.Bairro,
		&laboratorio.Endereco.CodMunicipio,
		&laboratorio.Endereco.Municipio,
		&laboratorio.Endereco.UF,
		&laboratorio.Endereco.CEP,
		&laboratorio.Endereco.PontoReferencia,
	); err != nil {
		return nil, err
	}

	return laboratorio, nil
}

func (repository *UnidadeRepository) CadastrarLaboratorio(ctx *context.Context, requisicao *model.Laboratorio) error {
	cnes := util.GerarId(10)
	query := `INSERT INTO laboratorio (cnes, nome, cnpj, endereco, telefone) VALUES ($1, $2, $3, $4, $5)`

	_, err := repository.db.DB.ExecContext(
		*ctx,
		query,
		cnes, requisicao.Nome, requisicao.CNPJ, requisicao.EnderecoID, requisicao.Telefone,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository *UnidadeRepository) ExisteUnidade(ctx *context.Context, cnes string) (bool, error) {
	var existe bool
	query := "SELECT EXISTS(SELECT 1 FROM unidade_saude WHERE cnes = $1)"
	err := repository.db.DB.QueryRowContext(*ctx, query, cnes).Scan(&existe)

	if err != nil {
		return false, err
	}

	return existe, nil
}

func (repository *UnidadeRepository) ExisteUnidadeLab(ctx *context.Context, cnes string) (bool, error) {
	var existe bool
	query := "SELECT EXISTS(SELECT 1 FROM laboratorio WHERE cnes = $1)"
	err := repository.db.DB.QueryRowContext(*ctx, query, cnes).Scan(&existe)

	if err != nil {
		return false, err
	}

	return existe, nil
}
