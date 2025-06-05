package service

import (
	"backend/dto"
	"backend/model"
	"backend/repository"
	"backend/util"
	"context"
	"time"
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

func (s *AgendamentoService) ConsultarHorariosOcupados(ctx *context.Context, data time.Time) (*[]dto.HorariosOcupados, error) {
	dataFormatada := data.Format("2006-01-02")

	horariosOcupados, err := s.agendamentoRepository.ConsultarHorariosOcupados(ctx, dataFormatada)
	if err != nil {
		return nil, err
	}

	return horariosOcupados, nil
}
