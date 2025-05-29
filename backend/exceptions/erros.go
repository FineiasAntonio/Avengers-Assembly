package exceptions

import "errors"

var (
	ErroCredenciaisInvalidas = errors.New("Credenciais inválidas")
	ErroNaoAutorizado        = errors.New("Não autorizado")
	ErroRequisicaoInvalida   = errors.New("Requisição inválida")

	ErroConexaoPostgres = errors.New("Erro ao se conectar com o PostgreSQL")
	ErroConexaoMongo    = errors.New("Erro ao conectar com o MongoDB")
)
