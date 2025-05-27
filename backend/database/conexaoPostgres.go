package database

import (
	"backend/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	Pool *pgxpool.Pool
}


type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func ConectarPostgres(cfg config.PostgresConfig) (*PostgresClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stringConexao := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	poolConfig, err := pgxpool.ParseConfig(stringConexao)
	if err != nil {
		return nil, fmt.Errorf("Erro ao tentar conectar com o PostgreSQL: %w", err)
	}

	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar a pool do PostgreSQL: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Erro ao tentar se comuncar com o PostgreSQL: %w", err)
	}

	log.Println("Conexão com o PostgreSQL estabelecida com sucesso")
	return &PostgresClient{Pool: pool}, nil
}

func (c *PostgresClient) FecharConexaoPostgres() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}
