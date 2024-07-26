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

func (r *repository) GetLanguage(ctx context.Context, search string) ([]language.Language, error) {

	var languages []language.Language
	q := `SELECT uuid, name, last_update FROM language WHERE name ILIKE $1||'%'`
	rows, err := r.db.Query(ctx, q, search)
	if err != nil {
		fmt.Println("Error getting language: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var language language.Language
		err := rows.Scan(&language.UUID, &language.Name, &language.LastUpdate)
		if err != nil {
			fmt.Println("Error getting rows language: ", err)
		}

		languages = append(languages, language)
	}

	return languages, nil
}

func (r *repository) GetLanguageID(ctx context.Context, id string) (language.UUID, error) {
	var l language.UUID
	q := `SELECT uuid FROM language WHERE uuid = $1`
	err := r.db.QueryRow(ctx, q, id).Scan(&l.UUID)
	if err != nil {
		fmt.Println("Error getting language id: ", err)
	}
	return l, nil
}

func (r *repository) DeleteLanguage(ctx context.Context, id string) error {

	q := `DELETE FROM language WHERE uuid = $1`
	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		fmt.Println("Error deleting language: ", err)
	}
	return err
}
