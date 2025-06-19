package main

import (
	"backend/api"
	"backend/auth"
	"backend/config"
	"backend/database"
	"backend/handler"
	"backend/middleware"
	"backend/repository"
	"backend/service"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := config.CarregarConfiguracoesDatabase()

	conexaoPostgres, err := database.ConectarPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err.Error())
	}

	/*conexaoMongo, err := database.ConectarMongo(cfg.Mongo)*/
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conexaoPostgres.FecharConexaoPostgres()
	/*defer conexaoMongo.FecharConexaoMongo()*/

	chaveJwt := os.Getenv("SECRET_JWT_KEY")
	if chaveJwt == "" {
		chaveJwt = "CHAVE_ULTRA_SECRETA"
		log.Println("Chave JWT não encontrada, usando chave padrão")
	}

	enderecoRepositorio := repository.NewEnderecoRepository(conexaoPostgres)

	usuarioRepositorio := repository.NewUsuarioRepository(conexaoPostgres)
	usuarioServico := service.NewUsuarioService(usuarioRepositorio)
	usuarioHandler := handler.NewUsuarioHandler(usuarioServico)

	pacienteRepositorio := repository.NewPacienteRepository(conexaoPostgres)
	pacienteServico := service.NewPacienteService(pacienteRepositorio, enderecoRepositorio)
	pacienteHandler := handler.NewPacienteHandler(pacienteServico)

	agendamentoRepository := repository.NewAgendamentoRepository(conexaoPostgres)
	agendamentoService := service.NewAgendamentoService(agendamentoRepository)
	agendamentoHandler := handler.NewAgendamentoHandler(agendamentoService)

	requisicaoExameRepositorio := repository.NewRequisicaoExameRepository(conexaoPostgres)
	requisicaoExameServico := service.NewRequisicaoExameService(requisicaoExameRepositorio)
	requisicaoExameHandler := handler.NewRequisicaoExameHandler(requisicaoExameServico)

	unidadeRpositorio := repository.NewUnidadeRepository(conexaoPostgres)
	unidadeService := service.NewUnidadeService(unidadeRpositorio, enderecoRepositorio)
	unidadeHandler := handler.NewUnidadeHandler(unidadeService)

	centralAnaliseRepositorio := repository.NewCentralAnaliseRepository(conexaoPostgres)
	centralAnaliseServico := service.NewCentralAnaliseService(centralAnaliseRepositorio)
	centralAnaliseHandler := handler.NewCentralAnaliseHandler(centralAnaliseServico)

	autenticacaoServico := auth.NewAutenticacaoService(usuarioRepositorio, []byte(chaveJwt))
	autenticacaoMiddleware := middleware.NewAutenticacaoMiddleware(autenticacaoServico)
	autenticacaoHandler := auth.NewAutenticacaoHandler(autenticacaoServico)

	corsMiddleware := middleware.NewCORSMiddleware()

	rotas := api.NewRotas(
		autenticacaoMiddleware,
		autenticacaoHandler,
		usuarioHandler,
		pacienteHandler,
		agendamentoHandler,
		requisicaoExameHandler,
		unidadeHandler,
		centralAnaliseHandler,
	)
	handerRotas := corsMiddleware.LiberarCORS(rotas.SetupRotas())

	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8080"
	}

	log.Printf("Servidor rodando na porta %s", porta)
	if err := http.ListenAndServe(":"+porta, handerRotas); err != nil {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v", err)
	}
}
