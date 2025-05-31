package handler

import (
	"backend/dto"
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

func (handler *UsuarioHandler) CadastrarUsuario(w http.ResponseWriter, r *http.Request) {
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

func (handler *UsuarioHandler) AlterarSenhaUsuario(w http.ResponseWriter, r *http.Request) {
	var novaSenha dto.RequisicaoNovaSenha
	if err := json.NewDecoder((r.Body)).Decode(&novaSenha); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.usuarioServico.AlterarSenha(&ctx, novaSenha)
	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
