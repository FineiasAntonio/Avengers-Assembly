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

type UsuarioService struct {
	repository *repository.UsuarioRepository
}

func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		repository: repo,
	}
}

func (s *UsuarioService) CadastrarUsuario(ctx *context.Context, requisicao *model.Usuario) error {
	usuarioRequisicao := *requisicao

	senha, err := bcrypt.GenerateFromPassword([]byte("000"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usuarioRequisicao.Senha = string(senha)
	usuarioRequisicao.PrimeiroAcesso = true

	if usuarioRequisicao.Registro == "" {
		precisaRegistro := usuarioRequisicao.Permissao == string(model.ACESSO_EXAMES) || 
						  usuarioRequisicao.Permissao == string(model.ACESSO_LABORATORIO)
		
		if !precisaRegistro {
			usuarioRequisicao.Registro = util.GerarId(8)
		}
	}

	err = s.repository.CadastrarUsuario(ctx, &usuarioRequisicao)
	if err != nil {
		return errors.New("erro ao cadastrar usuário: " + err.Error())
	}

	return nil
}

func (s *UsuarioService) AlterarSenha(ctx *context.Context, requisicaoNovaSenha dto.RequisicaoNovaSenha) error {
	usuario := (*ctx).Value("usuarioAutenticado").(*auth.Claims)
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(requisicaoNovaSenha.NovaSenha), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao gerar hash da senha: " + err.Error())
	}

	err = s.repository.AlterarSenha(ctx, usuario.CPF, string(senhaHash))
	if err != nil {
		return errors.New("erro ao alterar senha: " + err.Error())
	}

	return nil
}

func (s *UsuarioService) AlterarInformacao(ctx *context.Context, cpf string, dto *dto.UsuarioAlterarInformacaoDTO) error {
	if err := s.repository.AlterarInformacao(ctx, cpf, dto.Campo, dto.NovoValor); err != nil {
		return err
	}
	return nil
}

func (s *UsuarioService) GetUsuarioByCPF(ctx *context.Context, cpf string) (*model.Usuario, error) {
	usuario, err := s.repository.GetUsuarioByCPF(ctx, cpf)
	if err != nil {
		return nil, errors.New("erro ao buscar usuario: " + err.Error())
	}
	return usuario, nil
}

func (s *UsuarioService) GetUsuarioByRegistro(ctx *context.Context, registro string) (*model.Usuario, error) {
	usuario, err := s.repository.GetUsuarioByRegistro(ctx, registro)
	if err != nil {
		return nil, errors.New("erro ao buscar usuario: " + err.Error())
	}
	return usuario, nil
}

func (s *UsuarioService) UsuarioToDTO(user *model.Usuario) *dto.UsuarioDTO {
	return &dto.UsuarioDTO{
		Registro:         user.Registro,
		Nome:             user.Nome,
		CPF:              user.CPF,
		Email:            user.Email,
		Telefone:         user.Telefone,
		UnidadeSaudeCNES: user.UnidadeSaudeCNES,
		LaboratorioCNES:  user.LaboratorioCNES,
	}
}
