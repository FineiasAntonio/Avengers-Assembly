package handler

import (
	"backend/exceptions"
	"backend/model"
	"backend/service"
	"encoding/json"
	"net/http"
	"time"
)

type AgendamentoHandler struct {
	agendamentoService *service.AgendamentoService
}

func NewAgendamentoHandler(agendamentoService *service.AgendamentoService) *AgendamentoHandler {
	return &AgendamentoHandler{
		agendamentoService: agendamentoService,
	}
}

func (h *AgendamentoHandler) AgendarExame(w http.ResponseWriter, r *http.Request) {
	var agendamentoExame model.AgendamentoExame

	if err := json.NewDecoder(r.Body).Decode(&agendamentoExame); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	if err := h.agendamentoService.AgendarExame(&ctx, &agendamentoExame); err != nil {
		http.Error(w, "Erro ao agendar exame"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *AgendamentoHandler) ConsultarHorariosOcupados(w http.ResponseWriter, r *http.Request) {
	var dataConsultada time.Time

	if err := json.NewDecoder(r.Body).Decode(&dataConsultada); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	horariosOcupados, err := h.agendamentoService.ConsultarHorariosOcupados(&ctx, dataConsultada)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(horariosOcupados)
}
