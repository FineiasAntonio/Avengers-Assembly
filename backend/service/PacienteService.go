package service

import (
	"backend/auth"
	"backend/dto"
	"backend/model"
	"backend/repository"
	"backend/util"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PacienteService struct {
	pacienteRepository *repository.PacienteRepository
	enderecoRepository *repository.EnderecoRepository
}

func NewPacienteService(pacienteRepository *repository.PacienteRepository, enderecoRepository *repository.EnderecoRepository) *PacienteService {
	return &PacienteService{
		pacienteRepository: pacienteRepository,
		enderecoRepository: enderecoRepository,
	}
}

func (s *PacienteService) GetPacienteByCartaoSUS(ctx *context.Context, cartaoSUS string) (*model.Paciente, error) {
	paciente, err := s.pacienteRepository.GetPacienteByCartaoSUS(ctx, cartaoSUS)
	if err != nil {
		return nil, errors.New("erro ao buscar paciente: " + err.Error())
	}
	return paciente, nil
}

func (s *PacienteService) GetPacienteByCPF(ctx *context.Context, cpf string) (*model.Paciente, error) {
	paciente, err := s.pacienteRepository.GetPacienteByCartaoCPF(ctx, cpf)
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

	var enderecoId string
	enderecoId, err = s.enderecoRepository.CadastrarEndereco(ctx, paciente.Endereco)
	prontuario := util.GerarProntuario()
	paciente.Prontuario = prontuario

	if err != nil {
		return err
	}
	paciente.EnderecoID = enderecoId
	if err = s.pacienteRepository.CadastrarPaciente(ctx, paciente); err != nil {
		return err
	}

	return nil
}

func (s *PacienteService) AlterarSenha(ctx *context.Context, requisicaoNovaSenha *dto.RequisicaoNovaSenha) error {
	pacienteLogado := (*ctx).Value("usuarioAutenticado").(*auth.Claims)
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(requisicaoNovaSenha.NovaSenha), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao gerar hash da senha: " + err.Error())
	}

	err = s.pacienteRepository.AlterarSenha(ctx, pacienteLogado.CPF, string(senhaHash))

	if err != nil {
		return errors.New("erro ao alterar senha: " + err.Error())
	}

	return nil
}

func (s *PacienteService) PacienteToDTO(p *model.Paciente) *dto.PacienteDTO {
	return &dto.PacienteDTO{
		CartaoSUS:      p.CartaoSUS,
		Prontuario:     p.Prontuario,
		Nome:           p.Nome,
		NomeMae:        p.NomeMae,
		CPF:            p.CPF,
		DataNascimento: p.DataNascimento,
		Idade:          p.Idade,
		Raca:           p.Raca,
		Nacionalidade:  p.Nacionalidade,
		Escolaridade:   p.Escolaridade,
		Telefone:       p.Telefone,
		Endereco:       p.Endereco,
		Agenda:         p.Agenda,
	}
}

func (s *PacienteService) ExistePaciente(ctx *context.Context, cartao_sus string) error {
	existe, err := s.pacienteRepository.ExistePaciente(ctx, cartao_sus)
	if err != nil {
		return err
	}

	if !existe {
		return repository.ErroPacienteNaoEncontrado
	}

	return nil
}
