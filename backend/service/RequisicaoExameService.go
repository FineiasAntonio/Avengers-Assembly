package service

import (
	"backend/model"
	"backend/repository"
	"backend/util"
	"context"
	"errors"
	"fmt"
	"time"
)

type RequisicaoExameService struct {
	repository *repository.RequisicaoExameRepository
}

func NewRequisicaoExameService(re *repository.RequisicaoExameRepository) *RequisicaoExameService {
	RequisicaoExameService := RequisicaoExameService{repository: re}
	return &RequisicaoExameService
}

func (r *RequisicaoExameService) CadastrarRequisicaoExame(ctx *context.Context, requisicaoExame *model.RequisicaoExame) (string, error) {

	requisicaoExame.Status = string(model.AGUARDANDO)
	requisicaoExame.Protocolo = util.GerarId(10)

	if err := r.repository.CadastrarRequisicaoExame(ctx, requisicaoExame); err != nil {
		return "", errors.New("erro ao cadastrar requisicao exame: " + err.Error())
	}
	return requisicaoExame.Protocolo, nil
}

func (r *RequisicaoExameService) GetRequisicaoExameByProtocolo(ctx *context.Context, protocolo string) (*model.RequisicaoExame, error) {
	requisicaoExame, err := r.repository.GetRequisicaoExameByProtocolo(ctx, protocolo)

	if err != nil {
		return nil, errors.New("erro ao buscar requisicao exame: " + err.Error())
	}
	return requisicaoExame, nil
}

func (r *RequisicaoExameService) ExisteRequisicaoExame(ctx *context.Context, protocolo string) error {
	existe, err := r.repository.ExisteRequisicaoExame(ctx, protocolo)
	if err != nil {
		return err
	}

	if !existe {
		return repository.ErroRequisicaoExameNaoEncontrada
	}

	return nil
}

func (s *RequisicaoExameService) ProcessarLembretes() error {
	requisicoes, err := s.repository.BuscarRequisicoesComPaciente()
	if err != nil {
		return err
	}

	hoje := time.Now()

	for _, req := range requisicoes {
		if req.DataColeta == nil {
			continue
		}

		dataLimite := req.DataColeta.AddDate(1, 0, 0)
		if !hoje.After(dataLimite) {
			continue
		}

		enviou, err := s.repository.JaEnviouMsg(req.Protocolo)
		if err != nil {
			fmt.Printf("Erro ao verificar envio (protocolo %s): %v\n", req.Protocolo, err)
			continue
		}
		if enviou {
			continue
		}

		msg := fmt.Sprintf("Olá %s, notamos que seu último exame foi em %s. Recomendamos agendar um novo preventivo.",
			req.Paciente.Nome,
			req.DataColeta.Format("02/01/2006"),
		)

		telefone := util.FormatarTelefone(req.Paciente.Telefone)
		err = util.EnviarMensagemWaha(telefone, msg)
		if err != nil {
			fmt.Printf("Erro ao enviar mensagem para %s: %v\n", req.Paciente.Nome, err)
			continue
		}

		fmt.Printf("Mensagem enviada para %s (%s)\n", req.Paciente.Nome, req.Paciente.Telefone)

		err = s.repository.RegistrarEnvioMsg(req.Protocolo, req.Paciente.Prontuario)
		if err != nil {
			fmt.Printf("Erro ao registrar envio no Mongo para %s: %v\n", req.Paciente.Nome, err)
		}
	}

	return nil
}
