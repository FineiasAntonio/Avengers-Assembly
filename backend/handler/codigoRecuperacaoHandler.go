package handler

import (
	"backend/exceptions"
	"backend/model"
	"backend/service"
	"backend/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type CodigoHandler struct {
	CodigoService  *service.CodigoService
	UsuarioService *service.UsuarioService
}

func NewCodigoHandler(sc *service.CodigoService, su *service.UsuarioService) *CodigoHandler {
	return &CodigoHandler{
		CodigoService:  sc,
		UsuarioService: su,
	}
}

func (h *CodigoHandler) ConfirmarCodigo(w http.ResponseWriter, r *http.Request) {
	type CodigoRecuperacaoAgora struct {
		Codigo     string `bson:"codigo"`
		Credencial string `bson:"credencial"`
	}
	var codigoRec CodigoRecuperacaoAgora

	if err := json.NewDecoder(r.Body).Decode(&codigoRec); err != nil {
		http.Error(w, exceptions.ErroRequisicaoInvalida.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	codigoSalvo, credencialsalva, err := h.CodigoService.ConfirmarCodigo(ctx, codigoRec.Codigo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if codigoRec.Codigo != codigoSalvo {
		http.Error(w, "Código inválido!", http.StatusBadRequest)
		return
	}

	if codigoRec.Credencial != credencialsalva {
		http.Error(w, "Código inválido!", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (handler *CodigoHandler) EnviarEmailParaUsuarioRecuperarSenha(w http.ResponseWriter, r *http.Request) {
	parametro := r.URL.Query().Get("credencial")
	if parametro == "" {
		http.Error(w, "Credencial não fornecido", http.StatusBadRequest)
		return
	}

	var usuario *model.Usuario
	ctx := r.Context()
	var err error

	if len(parametro) == 11 {
		var cpf string
		cpf = parametro
		usuario, err = handler.UsuarioService.GetUsuarioByCPF(&ctx, cpf)

		if usuario == nil {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}

	} else {
		var registro string
		registro = parametro
		usuario, err = handler.UsuarioService.GetUsuarioByRegistro(&ctx, registro)

		if usuario == nil {
			http.Error(w, "Usuário não encontrado", http.StatusNotFound)
			return
		}
	}

	if err != nil {
		http.Error(w, exceptions.ErroInterno.Error(), http.StatusInternalServerError)
		return
	}

	email := usuario.Email
	id := usuario.Registro
	codigo, err := handler.CodigoService.GerarCodigo(ctx, id)
	mensagemSubj := "Código recuperação de senha"
	mensagem := fmt.Sprintf("Seu código de recuperação de senha é: %s", codigo)
	if err = util.EnviarEmail(mensagemSubj, mensagem, *email); err != nil {
		fmt.Print(err)
		http.Error(w, "Erro ao enviar email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
