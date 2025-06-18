package service

import (
	"backend/dto"
	"backend/repository"
	"context"
	"errors"
)

type MapaService struct {
	repository *repository.MapaRepository
}

func NewMapaService(repo *repository.MapaRepository) *MapaService {
	return &MapaService{
		repository: repo,
	}
}

func (s *MapaService) PegarQtdPacientes(ctx *context.Context) (*dto.MapaPacientesDTO, error) {
	qtdPacientes, err := s.repository.PegarQtdPacientes(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes: " + err.Error())
	}
	return qtdPacientes, err
}

func (s *MapaService) PegarQtdPacientesPorIdade(ctx *context.Context) (*dto.MapaPacientesPorIdadeDTO, error) {
	qtdPacientesPorIdade, err := s.repository.PegarQtdPacientesPorIdade(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes por idade: " + err.Error())
	}
	return qtdPacientesPorIdade, err
}

func (s *MapaService) PegarQtdPacientesPorRaca(ctx *context.Context) (*dto.MapaPacientesPorRacaDTO, error) {
	qtdPacientesPorRaca, err := s.repository.PegarQtdPacientesPorRaca(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes por raça: " + err.Error())
	}
	return qtdPacientesPorRaca, err
}

func (s *MapaService) PegarQtdPacientesPorEscolaridade(ctx *context.Context) (*dto.MapaPacientesPorEscolaridadeDTO, error) {
	qtdPacientesPorEscolaridade, err := s.repository.PegarQtdPacientesPorEscolaridade(ctx)
	if err != nil {
		return nil, errors.New("Erro ao pegar quantidade de pacientes por escolaridade: " + err.Error())
	}
	return qtdPacientesPorEscolaridade, err
}
