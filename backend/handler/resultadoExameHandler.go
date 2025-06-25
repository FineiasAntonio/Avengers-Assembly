package handler

import (
	"backend/model"
	"backend/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResultadoExameHandler struct {
	service *service.ResultadoExameService
}

func NewResultadoExameHandler(service *service.ResultadoExameService) *ResultadoExameHandler {
	return &ResultadoExameHandler{
		service: service,
	}
}

func (handler *ResultadoExameHandler) CadastrarResultadoExame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var requisicao model.ResultadoExame

	err := json.NewDecoder(r.Body).Decode(&requisicao)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = handler.service.CadastrarResultadoExame(&ctx, &requisicao)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (handler *ResultadoExameHandler) BuscarResultadoExamePorProtocolo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	protocoloExame := r.URL.Query().Get("protocolo")
	if protocoloExame == "" {
		http.Error(w, "Protocolo do exame é obrigatório", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	resultadoExame, err := handler.service.BuscarResultadoExamePorProtocolo(&ctx, protocoloExame)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Resultado do exame não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultadoExame)
}
