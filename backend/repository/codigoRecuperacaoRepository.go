package repository

import (
	"backend/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CodigoRepository struct {
	collection *mongo.Collection
}

func NewCodigoRepository(db *mongo.Database) *CodigoRepository {
	return &CodigoRepository{
		collection: db.Collection("codigos_recuperacao"),
	}
}

func (repo *CodigoRepository) SalvarCodigo(ctx context.Context, codigo model.CodigoRecuperacao) error {
	_, err := repo.collection.InsertOne(ctx, codigo)
	return err
}

func (repo *CodigoRepository) BuscarCodigo(ctx context.Context, codigo string) (*model.CodigoRecuperacao, error) {
	var result model.CodigoRecuperacao
	err := repo.collection.FindOne(ctx, bson.M{"codigo": codigo}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
