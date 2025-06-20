package service

import (
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

func (s *AgendamentoService) ConsultarHorariosOcupados(ctx *context.Context, data string, cnes string) (*map[string][]time.Time, error) {
	profissionais, horariosOcupados, err := s.agendamentoRepository.ConsultarHorariosOcupados(ctx, data, cnes)
	if err != nil {
		return nil, err
	}

	horariosPorProfissional := make(map[string][]time.Time)
	for _, profissional := range *profissionais {
		horariosPorProfissional[profissional] = []time.Time{}
	}

	for _, horario := range *horariosOcupados {
		if _, exists := horariosPorProfissional[horario.Profissional]; exists {
			horariosPorProfissional[horario.Profissional] = append(horariosPorProfissional[horario.Profissional], horario.Data)
		}
	}

	return &horariosPorProfissional, nil
}
