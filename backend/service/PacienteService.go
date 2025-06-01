package service

import (
	"backend/model"
	"backend/repository"
	"context"
	"errors"
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
		return nil, errors.New("erro ao buscar paciente: " + err.Error())
	}
	return paciente, nil
}

func (s *PacienteService) CadastrarPaciente(ctx *context.Context, paciente *model.Paciente) error {
	if err := s.repository.CadastrarPaciente(ctx, paciente); err != nil {
		return errors.New("erro ao cadastrar paciente: " + err.Error())
	}
	return nil
}
