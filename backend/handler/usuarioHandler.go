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
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (handler *UsuarioHandler) AlterarSenhaUsuarioEsqueceuSenha(w http.ResponseWriter, r *http.Request) {
	var novaSenha dto.RequisicaoNovaSenha
	if err := json.NewDecoder((r.Body)).Decode(&novaSenha); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	credencial := r.URL.Query().Get("credencial")
	if credencial == "" {
		http.Error(w, "Credencial não fornecida", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.usuarioServico.AlterarSenha(&ctx, novaSenha, credencial)
	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (handler *UsuarioHandler) AlterarSenhaUsuario(w http.ResponseWriter, r *http.Request) {
	var novaSenha dto.RequisicaoNovaSenha
	if err := json.NewDecoder((r.Body)).Decode(&novaSenha); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	credencial := r.URL.Query().Get("credencial")
	if credencial == "" {
		http.Error(w, "Credencial não fornecida", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err := handler.usuarioServico.AlterarSenha(&ctx, novaSenha, credencial)
	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (handler *UsuarioHandler) AlterarInformacao(w http.ResponseWriter, r *http.Request) {
	var dto dto.UsuarioAlterarInformacaoDTO

	if err := json.NewDecoder((r.Body)).Decode(&dto); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	cpf := r.URL.Query().Get("cpf")
	if cpf == "" {
		http.Error(w, "CPF não fornecido", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := handler.usuarioServico.AlterarInformacao(&ctx, cpf, &dto); err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (handler *UsuarioHandler) GetUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	parametro := r.URL.Query().Get("registro")
	if parametro == "" {
		http.Error(w, "Parametro não fornecido", http.StatusBadRequest)
		return
	}

	var usuario *model.Usuario
	ctx := r.Context()
	var err error

	if len(parametro) == 11 {
		var cpf string
		cpf = parametro
		usuario, err = handler.usuarioServico.GetUsuarioByCPF(&ctx, cpf)

	} else {
		var registro string
		registro = parametro
		usuario, err = handler.usuarioServico.GetUsuarioByRegistro(&ctx, registro)
	}

	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	usuarioDTO := handler.usuarioServico.UsuarioToDTO(usuario)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuarioDTO)
}
