package handler

import (
	"backend/exceptions"
	"backend/model"
	"backend/service"
	"encoding/json"
	"net/http"
)

type RequisicaoExameHandler struct {
	RequisicaoExameServico *service.RequisicaoExameService
}

func NewRequisicaoExameHandler(res *service.RequisicaoExameService) *RequisicaoExameHandler {
	RequisicaoExameHandler := RequisicaoExameHandler{RequisicaoExameServico: res}
	return &RequisicaoExameHandler
}

func (handler *RequisicaoExameHandler) CadastrarRequisicaoExame(w http.ResponseWriter,
	r *http.Request) {
	var requisicaoExame model.RequisicaoExame
	if err := json.NewDecoder(r.Body).Decode(&requisicaoExame); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := handler.RequisicaoExameServico.CadastrarRequisicaoExame(&ctx,
		&requisicaoExame); err != nil {
		http.Error(w, "Erro ao cadastrar requisição exame: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (handler *RequisicaoExameHandler) GetRequisicaoExameByProtocolo(w http.ResponseWriter,
	r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	var protocolo string
	if err := json.NewDecoder(r.Body).Decode(&protocolo); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}
	var requisicaoExame *model.RequisicaoExame

	ctx := r.Context()

	requisicaoExame, err := handler.RequisicaoExameServico.GetRequisicaoExameByProtocolo(&ctx, protocolo)

	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(requisicaoExame)
}
