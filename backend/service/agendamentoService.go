package service

import (
	"backend/dto"
	"backend/model"
	"backend/repository"
	"backend/util"
	"context"
)

type AgendamentoService struct {
	agendamentoRepository *repository.AgendamentoRepository
}

func NewAgendamentoService(agendamentoRepository *repository.AgendamentoRepository) *AgendamentoService {
	return &AgendamentoService{
		agendamentoRepository: agendamentoRepository,
	}
}

func (s *AgendamentoService) AgendarExame(ctx *context.Context, agendamento *model.AgendamentoExame) error {
	agendamento.Protocolo = util.GerarId(10)

	if err := s.agendamentoRepository.AgendarExame(ctx, agendamento); err != nil {
		return err
	}
	return nil
}

func (s *AgendamentoService) ConsultarHorariosOcupados(ctx *context.Context, data string, cnes string) (*[]dto.HorariosOcupados, error) {
	horariosOcupados, err := s.agendamentoRepository.ConsultarHorariosOcupados(ctx, data, cnes)
	if err != nil {
		return nil, err
	}

	return horariosOcupados, nil
}
