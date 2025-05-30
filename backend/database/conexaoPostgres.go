package database

import (
	"backend/config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Driver PostgreSQL
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

	// Configurações da pool de conexões
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Verifica a conexão
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Erro ao tentar se comunicar com o PostgreSQL: %w", err)
	}

	log.Println("Conexão com o PostgreSQL estabelecida com sucesso")
	return &PostgresClient{DB: db}, nil
}

func (c *PostgresClient) FecharConexaoPostgres() {
	if c.DB != nil {
		c.DB.Close()
	}
}
