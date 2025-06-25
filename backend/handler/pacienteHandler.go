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

type PacienteHandler struct {
	pacienteServico *service.PacienteService
}

func NewPacienteHandler(pacienteService *service.PacienteService) *PacienteHandler {
	return &PacienteHandler{pacienteServico: pacienteService}
}

func (handler *PacienteHandler) CadastrarPaciente(w http.ResponseWriter, r *http.Request) {
	var paciente model.Paciente

	if err := json.NewDecoder(r.Body).Decode(&paciente); err != nil {
		fmt.Println(err.Error())
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := handler.pacienteServico.CadastrarPaciente(&ctx, &paciente); err != nil {
		http.Error(w, "Erro ao Cadastrar Paciente"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (handler *PacienteHandler) GetPaciente(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	parametro := r.URL.Query().Get("paciente")
	if parametro == "" {
		http.Error(w, "Parametro não fornecido", http.StatusBadRequest)
		return
	}

	var paciente *model.Paciente
	ctx := r.Context()
	var err error

	if len(parametro) == 11 {
		var cpf string
		cpf = parametro
		paciente, err = handler.pacienteServico.GetPacienteByCPF(&ctx, cpf)

	} else {
		var cartaoSUS string
		cartaoSUS = parametro
		paciente, err = handler.pacienteServico.GetPacienteByCartaoSUS(&ctx, cartaoSUS)
	}

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	pacienteDTO := handler.pacienteServico.PacienteToDTO(paciente)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pacienteDTO)
}

func (handler *PacienteHandler) ExisteRequisicaoExame(w http.ResponseWriter, r *http.Request) {
	if r.Method != "HEAD" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	cartao_sus := r.URL.Query().Get("cartao_sus")
	if cartao_sus == "" {
		http.Error(w, "Cartão SUS não fornecido", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.pacienteServico.ExistePaciente(&ctx, cartao_sus)

	if err != nil {
		if err == repository.ErroPacienteNaoEncontrado {
			http.Error(w, "Paciente não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro ao verificar paciente: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
