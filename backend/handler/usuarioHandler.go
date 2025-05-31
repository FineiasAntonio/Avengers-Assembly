package handler

import (
	"backend/exceptions"
	"backend/model"
	"backend/service"
	"encoding/json"
	"net/http"
)

type UsuarioHandler struct {
	usuarioServico *service.UsuarioService
}

func NewUsuarioHandler(usuarioServico *service.UsuarioService) *UsuarioHandler {
	return &UsuarioHandler{usuarioServico: usuarioServico}
}

func (handler *UsuarioHandler) CadastrarNovoUsuario(w http.ResponseWriter, r *http.Request) {
	var requisicaoCadastro model.Usuario
	if err := json.NewDecoder(r.Body).Decode(&requisicaoCadastro); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.usuarioServico.CadastrarUsuario(&ctx, &requisicaoCadastro)
	if err != nil {
		http.Error(w, "Erro ao cadastrar usuário: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
