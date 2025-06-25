package handler

import (
	"backend/exceptions"
	"backend/model"
	"backend/repository"
	"backend/service"
	"encoding/json"
	"net/http"
)

type UnidadeHandler struct {
	service *service.UnidadeService
}

func NewUnidadeHandler(service *service.UnidadeService) *UnidadeHandler {
	return &UnidadeHandler{service: service}
}

func (handler *UnidadeHandler) ListarUnidade(w http.ResponseWriter, r *http.Request) {
	cnesUnidade := r.URL.Query().Get("cnes")

	if cnesUnidade == "" {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	unidade, err := handler.service.ListarUnidade(&ctx, cnesUnidade)
	if err != nil {

		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(unidade)
}

func (handler *UnidadeHandler) CadastrarUnidade(w http.ResponseWriter, r *http.Request) {
	var cadastroUnidadeRequisicao model.UnidadeSaude

	if err := json.NewDecoder(r.Body).Decode(&cadastroUnidadeRequisicao); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := handler.service.CadastrarUnidade(&ctx, &cadastroUnidadeRequisicao); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (handler *UnidadeHandler) ListarLaboratorio(w http.ResponseWriter, r *http.Request) {
	cnesLaboratorio := r.URL.Query().Get("cnes")

	if cnesLaboratorio == "" {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	laboratorio, err := handler.service.ListarLaboratorio(&ctx, cnesLaboratorio)
	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(laboratorio)
}

func (handler *UnidadeHandler) CadastrarLaboratorio(w http.ResponseWriter, r *http.Request) {
	var cadastroLaboratorioRequisicao model.Laboratorio

	if err := json.NewDecoder(r.Body).Decode(&cadastroLaboratorioRequisicao); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := handler.service.CadastrarLaboratorio(&ctx, &cadastroLaboratorioRequisicao); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (handler *UnidadeHandler) ExisteUnidade(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	cnes := r.URL.Query().Get("cnes")
	if cnes == "" {
		http.Error(w, "CNES não fornecido", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.service.ExisteUnidade(&ctx, cnes)
	if err != nil {
		if err == repository.ErroUnidadeNaoEncontrada {
			http.Error(w, "Unidade não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro ao buscar unidade", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (handler *UnidadeHandler) ExisteLaboratorio(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	cnes := r.URL.Query().Get("cnes")
	if cnes == "" {
		http.Error(w, "CNES não fornecido", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.service.ExisteUnidadeLab(&ctx, cnes)
	if err != nil {
		if err == repository.ErroUnidadeNaoEncontrada {
			http.Error(w, "Unidade não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro ao buscar unidade", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
