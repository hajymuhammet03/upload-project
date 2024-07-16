package db

import (
	"context"
	"fmt"
	"github.com/Hajymuhammet03/internal/admin/category"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/Hajymuhammet03/pkg/postgresql"
)

type repository struct {
	db     postgresql.Client
	logger *logging.Logger
}

func NewRepository(db postgresql.Client, logger *logging.Logger) category.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) AddCategory(ctx context.Context, dto category.AddCategory) (category.UUID, error) {
	var uuid category.UUID
	q := `INSERT INTO category (name_tm, name_en, name_ru) VALUES ($1, $2, $3)`
	err := r.db.QueryRow(ctx, q, dto.NameTm, dto.NameEn, dto.NameRu).Scan(&uuid)
	if err != nil {
		fmt.Println("Add Category Error: ", err)
	}

	return uuid, err
}
