package repository

import (
	"backend/database"
	"backend/model"
	"context"
	"database/sql"
	"fmt"
)

type RequisicaoExameRepository struct {
	db *database.PostgresClient
}

func NewRequisicaoExameRepository(db *database.PostgresClient) *RequisicaoExameRepository {
	RequisicaoExameRepository := RequisicaoExameRepository{db: db}
	return &RequisicaoExameRepository
}

func (r *RequisicaoExameRepository) GetRequisicaoExameByProtocolo(ctx *context.Context, protocolo string) (*model.RequisicaoExame, error) {
	var re model.RequisicaoExame

	query := "SELECT * FROM requisicao_exame JOIN paciente ON requisicao_exame.paciente = paciente.cartaosus JOIN endereco ON paciente.endereco = endereco.endereco_id JOIN usuario ON requisicao_exame.responsavel = usuario.registro WHERE protocolo = $1"

	row := r.db.DB.QueryRowContext(*ctx, query, protocolo)

	err := row.Scan(
		&re.Protocolo,
		&re.PacienteID,
		&re.MotivoExame,
		&re.FezExamePreventivo,
		&re.AnoUltimoExame,
		&re.UsaDIU,
		&re.EstaGravida,
		&re.UsaAnticoncepcional,
		&re.UsaHormonioMenopausa,
		&re.FezRadioterapia,
		&re.DataUltimaMenstruacao,
		&re.SangramentoAposRelacoes,
		&re.SangramentoAposMenopausa,
		&re.InspecaoColo,
		&re.SinaisDST,
		&re.DataColeta,
		&re.ResponsavelRegistro,
		&re.ResultadoID,
		&re.Status,
		&re.Paciente.CartaoSUS,
		&re.Paciente.Prontuario,
		&re.Paciente.Nome,
		&re.Paciente.NomeMae,
		&re.Paciente.CPF,
		&re.Paciente.DataNascimento,
		&re.Paciente.Idade,
		&re.Paciente.Raca,
		&re.Paciente.Nacionalidade,
		&re.Paciente.Escolaridade,
		&re.Paciente.Telefone,
		&re.Paciente.EnderecoID,
		&re.Paciente.Senha,
		&re.Paciente.PrimeiroAcesso,
		&re.Paciente.Endereco.EnderecoID,
		&re.Paciente.Endereco.Logradouro,
		&re.Paciente.Endereco.Numero,
		&re.Paciente.Endereco.Complemento,
		&re.Paciente.Endereco.Bairro,
		&re.Paciente.Endereco.CodMunicipio,
		&re.Paciente.Endereco.Municipio,
		&re.Paciente.Endereco.UF,
		&re.Paciente.Endereco.CEP,
		&re.Paciente.Endereco.PontoReferencia,
		&re.Responsavel.Registro,
		&re.Responsavel.Nome,
		&re.Responsavel.CPF,
		&re.Responsavel.Email,
		&re.Responsavel.Telefone,
		&re.Responsavel.Senha,
		&re.Responsavel.UnidadeSaudeCNES,
		&re.Responsavel.LaboratorioCNES,
		&re.Responsavel.Permissao,
		&re.Responsavel.PrimeiroAcesso,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Requisicao Exame Não Encontrada")
		}
		return nil, fmt.Errorf("Erro ao Buscar Requisiçao Exame %v", err)
	}
	return &re, nil
}

func (r *RequisicaoExameRepository) CadastrarRequisicaoExame(ctx *context.Context, re *model.RequisicaoExame) error {

	_, err := r.db.DB.ExecContext(*ctx,
		`INSERT INTO requisicao_exame (
		protocolo, paciente, motivoexame, fezexamepreventivo, anoultimoexame, usadiu,
		estagravida, usaanticoncepcional, usahormoniomenopausa, fezradioterapia,
		dataultimamenstruacao, sangramentoaposrelacoes, sangramentoaposmenopausa,
		inspecaocolo, sinaisdst, datacoleta, responsavel, resultado, status) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`,
		re.Protocolo, re.PacienteID, re.MotivoExame, re.FezExamePreventivo,
		re.AnoUltimoExame, re.UsaDIU, re.EstaGravida,
		re.UsaAnticoncepcional, re.UsaHormonioMenopausa,
		re.FezRadioterapia, re.DataUltimaMenstruacao, re.SangramentoAposRelacoes,
		re.SangramentoAposMenopausa, re.InspecaoColo, re.SinaisDST, re.DataColeta,
		re.ResponsavelRegistro, re.ResultadoID, re.Status)

	if err != nil {
		return fmt.Errorf("Erro ao Cadastrar Requisição Exame: %v", err)
	}
	return nil
}
