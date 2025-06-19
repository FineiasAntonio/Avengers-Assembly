package database

import (
	"backend/config"
	"backend/model"
	_ "backend/model"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type PostgresClient struct {
	DB *sql.DB
}

func ConectarPostgres(cfg config.PostgresConfig) (*PostgresClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stringConexao := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", stringConexao)
	if err != nil {
		return nil, fmt.Errorf("Erro ao tentar conectar com o PostgreSQL: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Minute * 5)

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Erro ao tentar se comunicar com o PostgreSQL: %w", err)
	}

	log.Println("Conexão com o PostgreSQL estabelecida com sucesso")

	err = IniciarTabelas(db)
	if err != nil {
		return nil, err
	}

	var adminExiste bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM usuario WHERE registro = $1)", "ADMIN").Scan(&adminExiste)
	if err != nil {
		return nil, err
	}

	if !adminExiste {
		query := `
		INSERT INTO usuario (
			registro, nome, cpf, email, telefone, senha,
			permissao, primeiroacesso
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
		senhaHash, err := bcrypt.GenerateFromPassword([]byte("000"), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		_, err = db.Exec(
			query,
			"ADMIN",
			"Administrador",
			"00000000000",
			"admin@admin.br",
			"5500000000000",
			string(senhaHash),
			model.ADMINISTRADOR,
			false,
		)

		if err != nil {
			return nil, err
		}

		log.Println("Usuário administrador criado | login: ADMIN ; senha: 000")

	}

	return &PostgresClient{DB: db}, nil
}

func (c *PostgresClient) FecharConexaoPostgres() {
	if c.DB != nil {
		c.DB.Close()
	}
}

func IniciarTabelas(db *sql.DB) error {
	caminhoConfiguracaoSQL := os.Getenv("CAMINHO_CONFIGURACAO_SQL")
	if caminhoConfiguracaoSQL == "" {
		caminhoConfiguracaoSQL = "../../config/sql/iniciar.sql"
	}
	content, err := os.ReadFile(caminhoConfiguracaoSQL)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo SQL: %w", err)
	}

	queries := strings.Split(string(content), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("erro ao executar comando '%s': %w", query, err)
		}
	}

	return nil
}
