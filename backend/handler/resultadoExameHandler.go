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
