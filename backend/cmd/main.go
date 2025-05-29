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
	_ "fmt"
	"log"
	"net/http"
	_ "net/http"
	"os"
)

func main() {
	cfg := config.CarregarConfiguracoesDatabase()

	conexaoPostgres, err := database.ConectarPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err.Error())
	}

	conexaoMongo, err := database.ConectarMongo(cfg.Mongo)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conexaoPostgres.FecharConexaoPostgres()
	defer conexaoMongo.FecharConexaoMongo()

	chaveJwt := os.Getenv("SECRET_JWT_KEY")
	if chaveJwt == "" {
		chaveJwt = "CHAVE_ULTRA_SECRETA"
		log.Println("Chave JWT não encontrada, usando chave padrão")
	}

	usuarioRepositorio := repository.NewUsuarioRepository(conexaoPostgres)
	usuarioServico := service.NewUsuarioService(usuarioRepositorio)
	usuarioHandler := handler.NewUsuarioHandler(usuarioServico)

	autenticacaoServico := auth.NewAutenticacaoService(usuarioRepositorio, []byte(chaveJwt))
	autenticacaoMiddleware := middleware.NewAutenticacaoMiddleware(autenticacaoServico)
	autenticacaoHandler := auth.NewAutenticacaoHandler(autenticacaoServico)

	corsMiddleware := middleware.NewCORSMiddleware()

	rotas := api.NewRotas(
		autenticacaoMiddleware,
		autenticacaoHandler,
		usuarioHandler,
	)
	handerRotas := corsMiddleware.LiberarCORS(rotas.SetupRotas())

	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8080"
	}

	log.Printf("Servidor rodando na porta %s...", porta)
	if err := http.ListenAndServe(":"+porta, handerRotas); err != nil {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v", err)
	}
}
