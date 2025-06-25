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
	unidadadeHandler       *handler.UnidadeHandler
	CentralAnaliseHandler  *handler.CentralAnaliseHandler
	resultadoExameHandler  *handler.ResultadoExameHandler
	codigoHandler          *handler.CodigoHandler
}

func NewRotas(
	autenticacaoMiddleware *middleware.MiddlewareAutenticacao,
	autenticacaoHandler *auth.AutenticacaoHandler,
	usuarioHandler *handler.UsuarioHandler,
	pacienteHandler *handler.PacienteHandler,
	agendamentoHandler *handler.AgendamentoHandler,
	requisicaoExameHandler *handler.RequisicaoExameHandler,
	unidadeHandler *handler.UnidadeHandler,
	CentralAnaliseHandler *handler.CentralAnaliseHandler,
	resultadoExameHandler *handler.ResultadoExameHandler,
	codigoHandler *handler.CodigoHandler,
) *Router {
	return &Router{
		autenticacaoMiddleware: autenticacaoMiddleware,
		autenticacaoHandler:    autenticacaoHandler,
		usuarioHandler:         usuarioHandler,
		pacienteHandler:        pacienteHandler,
		agendamentoHandler:     agendamentoHandler,
		requisicaoExameHandler: requisicaoExameHandler,
		unidadadeHandler:       unidadeHandler,
		CentralAnaliseHandler:  CentralAnaliseHandler,
		resultadoExameHandler:  resultadoExameHandler,
		codigoHandler:          codigoHandler,
	}
}

func (r *Router) SetupRotas() http.Handler {
	rotasComuns := http.NewServeMux()
	rotasComuns.HandleFunc("POST /api/auth/login", r.autenticacaoHandler.Login)
	rotasComuns.HandleFunc("POST /api/codigo/email", r.codigoHandler.EnviarEmailParaUsuarioRecuperarSenha)
	rotasComuns.HandleFunc("POST /api/codigo", r.codigoHandler.ConfirmarCodigo)
	rotasComuns.HandleFunc("PATCH /api/usuario/esqueceuSenha", r.usuarioHandler.AlterarSenhaUsuarioEsqueceuSenha)

	rotasProtegidas := http.NewServeMux()

	handlersProtegidos := r.autenticacaoMiddleware.MiddlewareAutenticacao(rotasProtegidas)
	rotasProtegidas.HandleFunc("POST /api/usuario", r.usuarioHandler.CadastrarUsuario)
	rotasProtegidas.HandleFunc("PATCH /api/usuario", r.usuarioHandler.AlterarSenhaUsuario)
	rotasProtegidas.HandleFunc("PATCH /api/usuario/alterarInf", r.usuarioHandler.AlterarInformacao)
	rotasProtegidas.HandleFunc("GET /api/usuario", r.usuarioHandler.GetUsuario)
	rotasProtegidas.HandleFunc("HEAD /api/usuario", r.usuarioHandler.ExisteUsuario)

	rotasProtegidas.HandleFunc("POST /api/paciente", r.pacienteHandler.CadastrarPaciente)
	rotasProtegidas.HandleFunc("GET /api/paciente", r.pacienteHandler.GetPaciente)
	rotasProtegidas.HandleFunc("HEAD /api/paciente", r.pacienteHandler.ExisteRequisicaoExame)

	rotasProtegidas.HandleFunc("POST /api/agendamento", r.agendamentoHandler.AgendarExame)
	rotasProtegidas.HandleFunc("GET /api/agendamento", r.agendamentoHandler.ConsultarHorariosOcupados)

	rotasProtegidas.HandleFunc("POST /api/requisicaoExame", r.requisicaoExameHandler.CadastrarRequisicaoExame)
	rotasProtegidas.HandleFunc("GET /api/requisicaoExame", r.requisicaoExameHandler.GetRequisicaoExameByProtocolo)
	rotasProtegidas.HandleFunc("HEAD /api/requisicaoExame", r.requisicaoExameHandler.ExisteRequisicaoExame)

	rotasProtegidas.HandleFunc("POST /api/requisicaoExame/resultado", r.resultadoExameHandler.CadastrarResultadoExame)
	rotasProtegidas.HandleFunc("GET /api/requisicaoExame/resultado", r.resultadoExameHandler.BuscarResultadoExamePorProtocolo)

	rotasProtegidas.HandleFunc("GET /api/unidade", r.unidadadeHandler.ListarUnidade)
	rotasProtegidas.HandleFunc("POST /api/unidade", r.unidadadeHandler.CadastrarUnidade)
	rotasProtegidas.HandleFunc("GET /api/laboratorio", r.unidadadeHandler.ListarLaboratorio)
	rotasProtegidas.HandleFunc("POST /api/laboratorio", r.unidadadeHandler.CadastrarLaboratorio)
	rotasProtegidas.HandleFunc("HEAD /api/unidade", r.unidadadeHandler.ExisteUnidade)
	rotasProtegidas.HandleFunc("HEAD /api/laboratorio", r.unidadadeHandler.ExisteLaboratorio)

	rotasProtegidas.HandleFunc("GET /api/graficos", r.CentralAnaliseHandler.PegarQuantidadePacientes)
	rotasProtegidas.HandleFunc("GET /api/mapa", r.CentralAnaliseHandler.PegarQuantidadePacientesPorRegiao)

	rotasComuns.Handle("/api/", handlersProtegidos)

	return rotasComuns
}
