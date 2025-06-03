package service

import (
	"backend/auth"
	"backend/dto"
	"backend/model"
	"backend/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

	senhaHash, err := bcrypt.GenerateFromPassword([]byte("000"), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao gerar hash da senha: " + err.Error())
	}

	paciente.PrimeiroAcesso = true
	paciente.Senha = string(senhaHash)

	if err = s.repository.CadastrarPaciente(ctx, paciente); err != nil {
		return errors.New("erro ao cadastrar paciente: " + err.Error())
	}
	return nil
}

func (s *PacienteService) AlterarSenha(ctx *context.Context, requisicaoNovaSenha *dto.RequisicaoNovaSenha) error {
	pacienteLogado := (*ctx).Value("usuarioAutenticado").(*auth.Claims)
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(requisicaoNovaSenha.NovaSenha), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao gerar hash da senha: " + err.Error())
	}

	err = s.repository.AlterarSenha(ctx, pacienteLogado.CPF, string(senhaHash))

	if err != nil {
		return errors.New("erro ao alterar senha: " + err.Error())
	}

	return nil
}
