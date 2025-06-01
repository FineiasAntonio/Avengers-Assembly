package service

import (
	"backend/model"
	"backend/repository"
	"context"
	"errors"
)

type RequisicaoExameService struct {
	repository *repository.RequisicaoExameRepository
}

func NewRequisicaoExameService(re *repository.RequisicaoExameRepository) *RequisicaoExameService {
	RequisicaoExameService := RequisicaoExameService{repository: re}
	return &RequisicaoExameService
}

func (r *RequisicaoExameService) CadastrarRequisicaoExame(ctx *context.Context,
	requisicaoExame *model.RequisicaoExame) error {
	if err := r.repository.CadastrarRequisicaoExame(ctx, requisicaoExame); err != nil {
		return errors.New("erro ao cadastrar requisicao exame: " + err.Error())
	}
	return nil
}

func (r *RequisicaoExameService) GetRequisicaoExameByProtocolo(ctx *context.Context,
	protocolo string) (*model.RequisicaoExame, error) {
	requisicaoExame, err := r.repository.GetRequisicaoExameByProtocolo(ctx, protocolo)

	if err != nil {
		return nil, errors.New("erro ao buscar requisicao exame: " + err.Error())
	}
	return requisicaoExame, nil
}
