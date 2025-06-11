package service

import (
	"backend/model"
	"backend/repository"
	"context"
	"errors"
)

type UnidadeService struct {
	unidadeRepository  *repository.UnidadeRepository
	enderecoRepository *repository.EnderecoRepository
}

func NewUnidadeService(unidadeRepository *repository.UnidadeRepository, enderecoRepository *repository.EnderecoRepository) *UnidadeService {
	return &UnidadeService{unidadeRepository: unidadeRepository, enderecoRepository: enderecoRepository}
}

func (service *UnidadeService) ListarUnidade(ctx *context.Context, cnes string) (*model.UnidadeSaude, error) {
	unidade, err := service.unidadeRepository.ListarUnidade(ctx, cnes)
	if err != nil {
		return nil, errors.New("erro ao listar unidade")
	}

	return unidade, nil
}

func (service *UnidadeService) CadastrarUnidade(ctx *context.Context, requisicao *model.UnidadeSaude) error {

	enderecoId, err := service.enderecoRepository.CadastrarEndereco(ctx, requisicao.Endereco)

	if err != nil {
		return errors.New("erro ao cadastrar endereco")
	}

	requisicao.EnderecoID = enderecoId

	if err = service.unidadeRepository.CadastrarUnidade(ctx, requisicao); err != nil {
		return errors.New("erro ao cadastrar unidade")
	}

	return nil
}
