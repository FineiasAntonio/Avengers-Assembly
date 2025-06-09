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
	agendamentoHandler     *handler.AgendamentoHandler
	requisicaoExameHandler *handler.RequisicaoExameHandler
}

func NewRotas(
	autenticacaoMiddleware *middleware.MiddlewareAutenticacao,
	autenticacaoHandler *auth.AutenticacaoHandler,
	usuarioHandler *handler.UsuarioHandler,
	pacienteHandler *handler.PacienteHandler,
	agendamentoHandler *handler.AgendamentoHandler,
	requisicaoExameHandler *handler.RequisicaoExameHandler,
) *Router {
	return &Router{
		autenticacaoMiddleware: autenticacaoMiddleware,
		autenticacaoHandler:    autenticacaoHandler,
		usuarioHandler:         usuarioHandler,
		pacienteHandler:        pacienteHandler,
		agendamentoHandler:     agendamentoHandler,
		requisicaoExameHandler: requisicaoExameHandler,
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
	rotasProtegidas.HandleFunc("GET /api/getPaciente", r.pacienteHandler.GetPaciente)

	rotasProtegidas.HandleFunc("POST /api/agendamento", r.agendamentoHandler.AgendarExame)
	rotasProtegidas.HandleFunc("GET /api/agendamento", r.agendamentoHandler.ConsultarHorariosOcupados)

	rotasProtegidas.HandleFunc("POST /api/requisicaoExame", r.requisicaoExameHandler.CadastrarRequisicaoExame)
	rotasProtegidas.HandleFunc("GET /api/getRequisicaoExame", r.requisicaoExameHandler.GetRequisicaoExameByProtocolo)

	rotasComuns.Handle("/api/", handlersProtegidos)

	return rotasComuns
}
