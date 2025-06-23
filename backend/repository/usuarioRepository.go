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
	var query string
	var args []interface{}

	if usuario.UnidadeSaudeCNES != nil {
		query = `INSERT INTO usuario (nome, cpf, email, telefone, registro, senha, permissao, primeiroacesso, unidadesaude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
		args = []interface{}{usuario.Nome, usuario.CPF, usuario.Email, usuario.Telefone, usuario.Registro, usuario.Senha, usuario.Permissao, usuario.PrimeiroAcesso, *usuario.UnidadeSaudeCNES}
	} else if usuario.LaboratorioCNES != nil {
		query = `INSERT INTO usuario (nome, cpf, email, telefone, registro, senha, permissao, primeiroacesso, laboratorio) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
		args = []interface{}{usuario.Nome, usuario.CPF, usuario.Email, usuario.Telefone, usuario.Registro, usuario.Senha, usuario.Permissao, usuario.PrimeiroAcesso, *usuario.LaboratorioCNES}
	} else {
		query = `INSERT INTO usuario (nome, cpf, email, telefone, registro, senha, permissao, primeiroacesso) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		args = []interface{}{usuario.Nome, usuario.CPF, usuario.Email, usuario.Telefone, usuario.Registro, usuario.Senha, usuario.Permissao, usuario.PrimeiroAcesso}
	}

	_, err := r.db.DB.ExecContext(*ctx, query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (r *UsuarioRepository) AlterarSenha(ctx *context.Context, credencial string, novaSenha string) error {
	var query string

	if len(credencial) == 11 {
		query = `UPDATE usuario SET senha = $1, primeiroacesso = false WHERE cpf = $2`

	} else {
		query = `UPDATE usuario SET senha = $1, primeiroacesso = false WHERE registro = $2`
	}
	_, err := r.db.DB.ExecContext(*ctx, query, novaSenha, credencial)
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
