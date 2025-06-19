package handler

import (
	"backend/service"
	"encoding/json"
	"net/http"
)

type CentralAnaliseHandler struct {
	CentralAnaliseServico *service.CentralAnaliseService
}

func NewCentralAnaliseHandler(s *service.CentralAnaliseService) *CentralAnaliseHandler {
	return &CentralAnaliseHandler{
		CentralAnaliseServico: s,
	}
}

func (handler *CentralAnaliseHandler) PegarQuantidadePacientes(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	funcao := r.URL.Query().Get("funcao")
	if funcao == "" {
		http.Error(w, "Parametro não fornecido", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	var Objeto interface{}

	if funcao == "idade" {
		qtdPacientesPorIdade, err := handler.CentralAnaliseServico.PegarQtdPacientesPorIdade(&ctx)

		if err != nil {
			http.Error(w, "Erro ao pegar pacientes por idade", http.StatusInternalServerError)
			return
		}

		if qtdPacientesPorIdade == nil {
			http.Error(w, "Erro ao encontrar pacientes por idade", http.StatusBadRequest)
			return
		}
		Objeto = qtdPacientesPorIdade

	} else if funcao == "raca" {
		qtdPacientesPorRaca, err := handler.CentralAnaliseServico.PegarQtdPacientesPorRaca(&ctx)

		if err != nil {
			http.Error(w, "Erro ao pegar pacientes por raca", http.StatusInternalServerError)
			return
		}

		if qtdPacientesPorRaca == nil {
			http.Error(w, "Erro ao encontrar pacientes por raca", http.StatusBadRequest)
			return
		}

		Objeto = qtdPacientesPorRaca

	} else if funcao == "escolaridade" {
		qtdPacientesPorEscolaridade, err := handler.CentralAnaliseServico.PegarQtdPacientesPorEscolaridade(&ctx)

		if err != nil {
			http.Error(w, "Erro ao pegar pacientes por escolaridade", http.StatusInternalServerError)
			return
		}

		if qtdPacientesPorEscolaridade == nil {
			http.Error(w, "Erro ao encontrar pacientes por escolaridade", http.StatusBadRequest)
			return
		}

		Objeto = qtdPacientesPorEscolaridade

	} else if funcao == "padrao" {
		qtdPacientes, err := handler.CentralAnaliseServico.PegarQtdPacientes(&ctx)

		if err != nil {
			http.Error(w, "Erro ao pegar pacientes", http.StatusInternalServerError)
			return
		}

		if qtdPacientes == nil {
			http.Error(w, "Erro ao encontrar paciente", http.StatusBadRequest)
			return
		}

		Objeto = qtdPacientes
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Objeto)
}
