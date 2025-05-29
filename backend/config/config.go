package config

import (
	"os"
	"time"
)

type DBConfig struct {
	Postgres PostgresConfig
	Mongo    MongoConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type MongoConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

func CarregarConfiguracoesDatabase() DBConfig {
	return DBConfig{
		Postgres: PostgresConfig{
			Host:     pegarVariavelAmbiente("POSTGRES_HOST", "localhost"),
			Port:     pegarVariavelAmbiente("POSTGRES_PORT", "5432"),
			User:     pegarVariavelAmbiente("POSTGRES_USER", "postgres"),
			Password: pegarVariavelAmbiente("POSTGRES_PASSWORD", "postgres"),
			DBName:   pegarVariavelAmbiente("POSTGRES_DB", "ccts_database"),
		},
		Mongo: MongoConfig{
			URI:      pegarVariavelAmbiente("MONGO_URI", "mongodb://localhost:27017"),
			Database: pegarVariavelAmbiente("MONGO_DB", "ccts_database"),
			Timeout:  10 * time.Second,
		},
	}
}

func pegarVariavelAmbiente(key, valorDefault string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return valorDefault
}
