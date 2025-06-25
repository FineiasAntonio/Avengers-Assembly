package service

import (
	"backend/model"
	"backend/repository"
	"context"
)

type ResultadoExameService struct {
	resultadoExameRepository *repository.ResultadoExameRepository
}

func NewResultadoExameService(resultadoExameRepository *repository.ResultadoExameRepository) *ResultadoExameService {
	return &ResultadoExameService{resultadoExameRepository: resultadoExameRepository}
}

func (service *ResultadoExameService) CadastrarResultadoExame(ctx *context.Context, requisicao *model.ResultadoExame) error {
	
	err := service.resultadoExameRepository.SalvarResultadoExame(ctx, requisicao)
	if err != nil {
		return err
	}
	return nil
}
