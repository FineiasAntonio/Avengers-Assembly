package model

import "time"

type CodigoRecuperacao struct {
	Codigo     string    `bson:"codigo"`
	Credencial string    `bson:"credencial"`
	CriadoEm   time.Time `bson:"criado_em"`
	ExpiraEm   time.Time `bson:"expira_em"`
}
