package service

import (
	"backend/dto"
	"backend/repository"
	"context"
	"errors"
)

type CentralAnaliseService struct {
	repository *repository.CentralAnaliseRepository
}

func NewCentralAnaliseService(repo *repository.CentralAnaliseRepository) *CentralAnaliseService {
	return &CentralAnaliseService{
		repository: repo,
	}
}

func (s *CentralAnaliseService) PegarQtdPacientes(ctx *context.Context) (*dto.GraficoPacientesDTO, error) {
	qtdPacientes, err := s.repository.PegarQtdPacientes(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes: " + err.Error())
	}
	return qtdPacientes, nil
}

func (s *CentralAnaliseService) PegarQtdPacientesPorIdade(ctx *context.Context) (*dto.GraficoPacientesPorIdadeDTO, error) {
	qtdPacientesPorIdade, err := s.repository.PegarQtdPacientesPorIdade(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes por idade: " + err.Error())
	}
	return qtdPacientesPorIdade, nil
}

func (s *CentralAnaliseService) PegarQtdPacientesPorRaca(ctx *context.Context) (*dto.GraficoPacientesPorRacaDTO, error) {
	qtdPacientesPorRaca, err := s.repository.PegarQtdPacientesPorRaca(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes por raça: " + err.Error())
	}
	return qtdPacientesPorRaca, nil
}

func (s *CentralAnaliseService) PegarQtdPacientesPorEscolaridade(ctx *context.Context) (*dto.GraficoPacientesPorEscolaridadeDTO, error) {
	qtdPacientesPorEscolaridade, err := s.repository.PegarQtdPacientesPorEscolaridade(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes por escolaridade: " + err.Error())
	}
	return qtdPacientesPorEscolaridade, nil
}

func (s *CentralAnaliseService) PegarQtdPacientesPorRegiao(ctx *context.Context) (*[]dto.MapaPacientesPorRegiao, error) {
	qtdPacientesPorRegiao, err := s.repository.PegarQtdPacientesPorRegiao(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes por região? " + err.Error())
	}
	return qtdPacientesPorRegiao, nil
}
