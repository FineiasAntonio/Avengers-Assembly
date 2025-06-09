package handler

import (
	"backend/exceptions"
	"backend/model"
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

func (handler *PacienteHandler) GetPaciente(w http.ResponseWriter,
	r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	parametro := r.URL.Query().Get("parametro")
	if parametro == "" {
		http.Error(w, "Cartão do SUS não fornecido", http.StatusBadRequest)
		return
	}

	var paciente *model.Paciente
	ctx := r.Context()
	var err error

	if len(parametro) == 11 {
		var cpf string
		cpf = parametro
		paciente, err = handler.pacienteServico.GetPacienteByCartaoSUS(&ctx, cpf)

	} else {
		var cartaoSUS string
		cartaoSUS = parametro
		paciente, err = handler.pacienteServico.GetPacienteByCartaoSUS(&ctx, cartaoSUS)
	}

	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	pacienteDTO := handler.pacienteServico.PacienteToDTO(paciente)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pacienteDTO)
}
