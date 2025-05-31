package api

import (
	"backend/auth"
	"backend/handler"
	"backend/middleware"
	"net/http"
)

type Router struct {
	autenticacaoMiddleware *middleware.MiddlewareAutenticacao
	autenticacaoHandler    *auth.AutenticacaoHandler
	usuarioHandler         *handler.UsuarioHandler
	pacienteHandler        *handler.PacienteHandler
}

func NewRotas(
	autenticacaoMiddleware *middleware.MiddlewareAutenticacao,
	autenticacaoHandler *auth.AutenticacaoHandler,
	usuarioHandler *handler.UsuarioHandler,
	pacienteHandler *handler.PacienteHandler,
) *Router {
	return &Router{
		autenticacaoMiddleware: autenticacaoMiddleware,
		autenticacaoHandler:    autenticacaoHandler,
		usuarioHandler:         usuarioHandler,
		pacienteHandler:        pacienteHandler,
	}
}

func (r *Router) SetupRotas() http.Handler {
	rotasComuns := http.NewServeMux()
	rotasComuns.HandleFunc("POST /api/auth/login", r.autenticacaoHandler.Login)

	rotasProtegidas := http.NewServeMux()

	handlersProtegidos := r.autenticacaoMiddleware.MiddlewareAutenticacao(rotasProtegidas)
	rotasProtegidas.HandleFunc("POST /api/usuario", r.usuarioHandler.CadastrarUsuario)
	rotasProtegidas.HandleFunc("PATCH /api/usuario", r.usuarioHandler.AlterarSenhaUsuario)

	rotasProtegidas.HandleFunc("POST /api/paciente", r.pacienteHandler.CadastrarPaciente)

	rotasComuns.Handle("/api/", handlersProtegidos)

	return rotasComuns
}
