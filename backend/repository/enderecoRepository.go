package repository

import (
	"backend/database"
	"backend/model"
	"backend/util"
	"context"
	"fmt"
)

type EnderecoRepository struct {
	db *database.PostgresClient
}

func NewEnderecoRepository(db *database.PostgresClient) *EnderecoRepository {
	return &EnderecoRepository{db: db}
}

func (r *EnderecoRepository) CadastrarEndereco(ctx *context.Context, requisicao model.Endereco) (string, error) {

	idGerado := util.GerarId(10)

	_, err := r.db.DB.ExecContext(*ctx, `
		INSERT INTO endereco (
			endereco_id, logradouro, numero, complemento, 
		    bairro, municipio, uf, cep, pontoreferencia
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		idGerado, requisicao.Logradouro, requisicao.Numero, requisicao.Complemento,
		requisicao.Bairro, requisicao.Municipio, requisicao.UF, requisicao.CEP,
		requisicao.PontoReferencia)

	if err != nil {
		return "", fmt.Errorf("Erro ao cadastrar endereço: %v", err)
	}

	return idGerado, nil
}
