package service

import (
	"backend/model"
	"backend/repository"
	"backend/util"
	"context"
	"errors"
	"fmt"
	"time"
)

type CodigoService struct {
	repository *repository.CodigoRepository
}

func NewCodigoService(repository *repository.CodigoRepository) *CodigoService {
	return &CodigoService{
		repository: repository,
	}
}

func (s *CodigoService) GerarCodigo(ctx context.Context, credencial string) (string, error) {
	codigo := util.GerarCodigo()
	expira := time.Now().Add(30 * time.Minute)

	cod := model.CodigoRecuperacao{
		Codigo:     codigo,
		Credencial: credencial,
		CriadoEm:   time.Now(),
		ExpiraEm:   expira,
	}

	err := s.repository.SalvarCodigo(ctx, cod)
	if err != nil {
		return "", err
	}
	fmt.Print("salvou o codigo: ", cod)
	return codigo, nil
}

func (s *CodigoService) ConfirmarCodigo(ctx context.Context, codigo string) (string, string, error) {
	cod, err := s.repository.BuscarCodigo(ctx, codigo)
	if err != nil {
		return "", "", errors.New("código inválido")
	}

	if time.Now().After(cod.ExpiraEm) {
		return "", "", errors.New("código expirado")
	}

	return cod.Codigo, cod.Credencial, nil
}
