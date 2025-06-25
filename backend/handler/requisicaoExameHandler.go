package handler

import (
	"backend/exceptions"
	"backend/model"
	"backend/repository"
	"backend/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequisicaoExameHandler struct {
	RequisicaoExameServico *service.RequisicaoExameService
}

func NewRequisicaoExameHandler(res *service.RequisicaoExameService) *RequisicaoExameHandler {
	RequisicaoExameHandler := RequisicaoExameHandler{RequisicaoExameServico: res}
	return &RequisicaoExameHandler
}

func (handler *RequisicaoExameHandler) CadastrarRequisicaoExame(w http.ResponseWriter, r *http.Request) {
	var requisicaoExame model.RequisicaoExame
	if err := json.NewDecoder(r.Body).Decode(&requisicaoExame); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	protocolo, err := handler.RequisicaoExameServico.CadastrarRequisicaoExame(&ctx, &requisicaoExame)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Erro ao cadastrar requisição exame: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(protocolo)
}

func (handler *RequisicaoExameHandler) GetRequisicaoExameByProtocolo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	protocolo := r.URL.Query().Get("protocolo")
	if protocolo == "" {
		http.Error(w, "Protocolo não fornecido", http.StatusBadRequest)
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(requisicaoExame)
}

func (handler *RequisicaoExameHandler) ExisteRequisicaoExame(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	protocolo := r.URL.Query().Get("protocolo")
	if protocolo == "" {
		http.Error(w, "Protocolo não fornecido", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.RequisicaoExameServico.ExisteRequisicaoExame(&ctx, protocolo)

	if err != nil {
		if err == repository.ErroRequisicaoExameNaoEncontrada {
			http.Error(w, "Requisição de exame não encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro ao verificar requisição de exame: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
