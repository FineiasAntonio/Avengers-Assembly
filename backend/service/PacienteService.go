package service

import (
	"backend/model"
	"backend/repository"
	"context"
)

type PacienteService struct {
	repository *repository.PacienteRepository
}

func NewPacienteService(repo *repository.PacienteRepository) *PacienteService {
	return &PacienteService{
		repository: repo,
	}
}

func (s *PacienteService) GetPacienteByCartaoSUS(ctx *context.Context, cartaoSUS string) (*model.Paciente, error) {
	paciente, err := s.repository.GetPacienteByCartaoSUS(ctx, cartaoSUS)
	if err != nil {
		return nil, err
	}
	return paciente, nil
}

func (s *PacienteService) CadastrarPaciente(ctx *context.Context, paciente *model.Paciente) error {
	if err := s.repository.CadastrarPaciente(ctx, paciente); err != nil {
		return err
	}
	return nil
}
