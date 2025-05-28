package repository

import (
	"backend/database"
	"backend/model"
)

type UsuarioRepository struct {
	db *database.PostgresClient
}

func NewUsuarioRepository(db *database.PostgresClient) *UsuarioRepository {
	return &UsuarioRepository{
		db: db,
	}
}

func (r *UsuarioRepository) GetUsuarioByCPF(cpf string) (*model.Usuario, error) {
	query := `SELECT * FROM usuarios WHERE cpf = $1`
	var usuario model.Usuario
	err := r.db.DB.QueryRow(query, cpf).Scan(&usuario.CPF, &usuario.Nome, &usuario.Email, &usuario.Telefone)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) GetUsuarioByRegistro(registro string) (*model.Usuario, error) {
	query := `SELECT * FROM usuarios WHERE registro = $1`
	var usuario model.Usuario
	err := r.db.DB.QueryRow(query, registro).Scan(&usuario.CPF, &usuario.Nome, &usuario.Email, &usuario.Telefone)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepository) CadastrarUsuario(usuario *model.Usuario) error {
	query := `INSERT INTO usuarios (nome, cpf, email, telefone, registro, senha, permissao, primeiroacesso) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.DB.Exec(query, usuario.Nome, usuario.CPF, usuario.Email, usuario.Telefone, usuario.Registro, usuario.Senha, usuario.Permissao, usuario.PrimeiroAcesso)
	return err
}

func (r *UsuarioRepository) AlterarSenha(cpf string, novaSenha string) error {
	query := `UPDATE usuarios SET senha = $1, primeiroacesso = false WHERE cpf = $2`
	_, err := r.db.DB.Exec(query, novaSenha, cpf)
	return err
}
