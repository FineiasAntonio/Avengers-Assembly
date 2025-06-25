package service

import (
	"backend/model"
	"backend/repository"
	"context"
)

type ResultadoExameService struct {
	resultadoExameRepository  *repository.ResultadoExameRepository
	requisicaoExameRepository *repository.RequisicaoExameRepository
}

func NewResultadoExameService(
	resultadoExameRepository *repository.ResultadoExameRepository,
	requisicaoExameRepository *repository.RequisicaoExameRepository,
) *ResultadoExameService {
	return &ResultadoExameService{
		resultadoExameRepository:  resultadoExameRepository,
		requisicaoExameRepository: requisicaoExameRepository,
	}
}

func (service *ResultadoExameService) CadastrarResultadoExame(ctx *context.Context, requisicao *model.ResultadoExame) error {
	protocolo := requisicao.ProtocoloExame
	err := service.resultadoExameRepository.SalvarResultadoExame(ctx, requisicao)
	if err != nil {
		return err
	}

	err = service.requisicaoExameRepository.AtualizarStatusRequisicao(ctx, protocolo, string(model.LAUDO_EMITIDO))
	if err != nil {
		return err
	}

	return nil
}

func (service *ResultadoExameService) BuscarResultadoExamePorProtocolo(ctx *context.Context, protocoloExame string) (*model.ResultadoExame, error) {
	resultadoExame, err := service.resultadoExameRepository.BuscarResultadoExame(ctx, protocoloExame)
	if err != nil {
		return nil, err
	}
	return resultadoExame, nil
}
