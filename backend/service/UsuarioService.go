package service

import (
	"backend/auth"
	"backend/model"
	"backend/repository"
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

	err = s.repository.CadastrarUsuario(ctx, &usuarioRequisicao)
	if err != nil {
		return errors.New("erro ao cadastrar usuário: " + err.Error())
	}

	return nil
}

func (s *UsuarioService) AlterarSenha(ctx *context.Context, novaSenha string) error {
	usuario := (*ctx).Value("usuarioAutenticado").(*auth.Claims)
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(novaSenha), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("erro ao gerar hash da senha: " + err.Error())
	}

	err = s.repository.AlterarSenha(ctx, usuario.CPF, string(senhaHash))
	if err != nil {
		return errors.New("erro ao alterar senha: " + err.Error())
	}

	return nil
}
