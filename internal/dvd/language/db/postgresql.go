package db

import (
	"context"
	"fmt"
	"github.com/Hajymuhammet03/internal/dvd/language"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/Hajymuhammet03/pkg/postgresql"
)

type repository struct {
	db     postgresql.Client
	logger *logging.Logger
}

func NewRepository(db postgresql.Client, logger *logging.Logger) language.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) AddLanguage(ctx context.Context, dto language.LanguageDTO) (language.UUID, error) {
	var id language.UUID

	if dto.UUID == "" {
		q := `INSERT INTO language (name) VALUES ($1) RETURNING uuid`
		err := r.db.QueryRow(ctx, q, dto.Name).Scan(&id.UUID)
		if err != nil {
			fmt.Println("Error adding language: ", err)
			return id, err
		}
	} else {
		q := `UPDATE language SET name = $1 WHERE uuid = $2 RETURNING uuid`
		err := r.db.QueryRow(ctx, q, dto.Name, dto.UUID).Scan(&id.UUID)
		if err != nil {
			fmt.Println("Error updating language: ", err)
			return id, err
		}
	}
	return id, nil
}
