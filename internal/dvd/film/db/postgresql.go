package db

import (
	"github.com/Hajymuhammet03/internal/dvd/film"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/Hajymuhammet03/pkg/postgresql"
)

type repository struct {
	db     postgresql.Client
	logger *logging.Logger
}

func NewRepository(db postgresql.Client, logger *logging.Logger) film.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
