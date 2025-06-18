package repository

import (
	"backend/database"
	"backend/model"
	"context"
	"fmt"
)

type UsuarioRepository struct {
	db *database.PostgresClient
}

func NewUsuarioRepository(db *database.PostgresClient) *UsuarioRepository {
	return &UsuarioRepository{
		db: db,
	}
}

func (r *UsuarioRepository) GetUsuarioByCPF(ctx *context.Context, cpf string) (*model.Usuario, error) {
	query := `SELECT * FROM usuario WHERE cpf = $1`
	var usuario model.Usuario
	err := r.db.DB.QueryRowContext(*ctx, query, cpf).Scan(
		&usuario.Registro,
		&usuario.Nome,
		&usuario.CPF,
		&usuario.Email,
		&usuario.Telefone,
		&usuario.Senha,
		&usuario.UnidadeSaudeCNES,
		&usuario.LaboratorioCNES,
		&usuario.Permissao,
		&usuario.PrimeiroAcesso,
	)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) GetUsuarioByRegistro(ctx *context.Context, registro string) (*model.Usuario, error) {
	query := `SELECT * FROM usuario WHERE registro = $1`
	var usuario model.Usuario
	err := r.db.DB.QueryRowContext(*ctx, query, registro).Scan(
		&usuario.Registro,
		&usuario.Nome,
		&usuario.CPF,
		&usuario.Email,
		&usuario.Telefone,
		&usuario.Senha,
		&usuario.UnidadeSaudeCNES,
		&usuario.LaboratorioCNES,
		&usuario.Permissao,
		&usuario.PrimeiroAcesso,
	)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) CadastrarUsuario(ctx *context.Context, usuario *model.Usuario) error {
	query := `INSERT INTO usuario (nome, cpf, email, telefone, registro, senha, permissao, primeiroacesso) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.DB.ExecContext(*ctx, query, usuario.Nome, usuario.CPF, usuario.Email, usuario.Telefone, usuario.Registro, usuario.Senha, usuario.Permissao, usuario.PrimeiroAcesso)
	return err
}

func (r *UsuarioRepository) AlterarSenha(ctx *context.Context, cpf string, novaSenha string) error {
	query := `UPDATE usuario SET senha = $1, primeiroacesso = false WHERE cpf = $2`
	_, err := r.db.DB.ExecContext(*ctx, query, novaSenha, cpf)
	return err
}

func (r *UsuarioRepository) AlterarInformacao(ctx *context.Context, cpf, campo, novoValor string) error {
	if campo != "email" && campo != "telefone" {
		return fmt.Errorf("Campo não permitido %s", campo)
	}
	query := fmt.Sprintf(`UPDATE usuario SET %s = $1 WHERE cpf = $2`, campo)
	_, err := r.db.DB.ExecContext(*ctx, query, novoValor, cpf)
	return err
}
