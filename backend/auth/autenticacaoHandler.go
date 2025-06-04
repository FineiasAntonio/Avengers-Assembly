package auth

import (
	"backend/dto"
	"backend/exceptions"
	"encoding/json"
	"errors"
	"net/http"
)

type AutenticacaoHandler struct {
	servicoAutenticacao *ServicoAutenticacao
}

func NewAutenticacaoHandler(servicoAutenticacao *ServicoAutenticacao) *AutenticacaoHandler {
	return &AutenticacaoHandler{servicoAutenticacao: servicoAutenticacao}
}

func (handler *AutenticacaoHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credenciais dto.CredenciaisUsuario
	if err := json.NewDecoder(r.Body).Decode(&credenciais); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()

	token, err := handler.servicoAutenticacao.AutenticarUsuario(&ctx, credenciais)
	if err != nil {
		if errors.Is(err, exceptions.ErroCredenciaisInvalidas) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
