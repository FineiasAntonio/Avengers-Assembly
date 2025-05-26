package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func ConectarMongo(cfg MongoConfig) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	opcoesConexao := options.Client().ApplyURI(cfg.URI)

	client, err := mongo.Connect(ctx, opcoesConexao)
	if err != nil {
		return nil, fmt.Errorf("Erro ao se conectar com o MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Erro ao se comunicar com o MongoDB: %w", err)
	}

	db := client.Database(cfg.Database)

	log.Println("Conexão com o MongoDB estabelecida com sucesso")
	return &MongoClient{
		Client:   client,
		Database: db,
	}, nil
}

func (c *MongoClient) FecharConexaoMongo() {
	if c.Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = c.Client.Disconnect(ctx)
	}
}
