package handler

import (
	"backend/exceptions"
	"backend/model"
	"backend/service"
	"encoding/json"
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
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.PacienteServico.CadastrarPaciente(r.Context(), &paciente); err != nil {
		http.Error(w, "Erro ao Cadastrar Paciente"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
