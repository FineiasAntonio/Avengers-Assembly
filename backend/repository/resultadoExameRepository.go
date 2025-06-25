package repository

import (
	"backend/database"
	"backend/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type ResultadoExameRepository struct {
	db *database.MongoClient
}

func NewResultadoExameRepository(db *database.MongoClient) *ResultadoExameRepository {
	resultadoExameRepository := ResultadoExameRepository{db: db}
	return &resultadoExameRepository
}

func (repository *ResultadoExameRepository) BuscarResultadoExame(ctx *context.Context, protocoloExame string) (*model.ResultadoExame, error) {
	resultadoExameCollection := repository.db.Database.Collection("resultado_exame")

	resultadoExame := &model.ResultadoExame{}
	if err := resultadoExameCollection.FindOne(*ctx, bson.M{"protocolo_exame": protocoloExame}).Decode(resultadoExame); err != nil {
		return nil, err
	}

	return resultadoExame, nil
}

func (repository *ResultadoExameRepository) SalvarResultadoExame(ctx *context.Context, resultadoExame *model.ResultadoExame) error {
	resultadoExameCollection := repository.db.Database.Collection("resultado_exame")

	_, err := resultadoExameCollection.InsertOne(*ctx, resultadoExame)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ResultadoExameRepository) EmitirResultadoExame(ctx *context.Context, protocoloExame string) error {
	resultadoExameCollection := repository.db.Database.Collection("resultado_exame")

	_, err := resultadoExameCollection.UpdateOne(
		*ctx,
		bson.M{"protocolo_exame": protocoloExame},
		bson.M{"$set": bson.M{"data_emissao_laudo": time.Now()}},
	)

	if err != nil {
		return err
	}

	return nil
}
