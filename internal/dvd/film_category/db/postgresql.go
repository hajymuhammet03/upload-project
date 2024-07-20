package db

import (
	"context"
	"fmt"
	"github.com/Hajymuhammet03/internal/dvd/film_category"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/Hajymuhammet03/pkg/postgresql"
)

type repository struct {
	db     postgresql.Client
	logger *logging.Logger
}

func NewRepository(db postgresql.Client, logger *logging.Logger) film_category.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) AddFilmCategory(ctx context.Context, dto film_category.FilmCategoryReq) (film_category.UUID, error) {
	var id film_category.UUID

	if dto.UUID == "" {
		q := `INSERT INTO film_category (category_id) VALUES ($1) RETURNING uuid`
		err := r.db.QueryRow(ctx, q, dto.CategoryID).Scan(&id.UUID)
		if err != nil {
			fmt.Println("Error adding film category: ", err)
			return id, err
		}
	} else {
		q := `UPDATE film_category SET category_id = $1 WHERE uuid = $2 RETURNING uuid`
		err := r.db.QueryRow(ctx, q, dto.CategoryID, dto.UUID).Scan(&id.UUID)
		if err != nil {
			fmt.Println("Error updating film category: ", err)
			return id, err
		}
	}
	return id, nil
}

func (r *repository) GetFilmCategory(ctx context.Context, dto film_category.PaginationDTO) ([]film_category.GetFilmCategory, int64, error) {
	var (
		films []film_category.GetFilmCategory
		count int64
	)

	q := `
         SELECT
             fc.uuid, c.uuid, fc.last_update, fc.created_at
         FROM film_category fc
             LEFT JOIN category c ON fc.category_id = c.uuid
        `

	if dto.StartDate != "0" && dto.StartDate != "" {
		q += `AND fc.created_at >= '` + dto.StartDate + `'`
	}

	if dto.EndDate != "0" && dto.EndDate != "" {
		q += `AND fc.created_at <= '` + dto.EndDate + `'`
	}

	q += `ORDER BY fc.created_at LIMIT $1 OFFSET $2`

	row, err := r.db.Query(ctx, q, dto.Limit, dto.Limit*dto.Page)
	if err != nil {
		fmt.Println("Error getting film category: ", err)
	}
	defer row.Close()

	for row.Next() {
		var film film_category.GetFilmCategory
		err := row.Scan(&film.UUID, &film.CategoryID, &film.LastUpdate, &film.CreatedAt)
		if err != nil {
			fmt.Println("Get Film Category Error in for rows: ", err)
			return nil, 0, err
		}
		films = append(films, film)
	}

	q = `SELECT COUNT(fc.uuid) FROM film_category fc`

	if dto.StartDate != "0" && dto.StartDate != "" {
		q += `AND fc.created_at >= '` + dto.StartDate + `'`
	}
	if dto.EndDate != "0" && dto.EndDate != "" {
		q += `AND fc.created_at <= '` + dto.EndDate + `'`
	}

	err = r.db.QueryRow(ctx, q).Scan(&count)
	if err != nil {
		fmt.Println("Get Film Category Count Error: ", err)
	}
	return films, count, nil
}

func (r *repository) GetFilmCategoryID(ctx context.Context, id film_category.UUID) (film_category.GetFilmCategory, error) {
	var film film_category.GetFilmCategory

	q := `SELECT
             fc.uuid, c.uuid, fc.last_update, fc.created_at
          FROM film_category fc
             LEFT JOIN category c ON fc.category_id = c.uuid
          WHERE fc.uuid = $1 
         `

	err := r.db.QueryRow(ctx, q, id.UUID).Scan(&film.UUID, &film.CategoryID, &film.LastUpdate, &film.CreatedAt)
	if err != nil {
		fmt.Println("Error getting film category id: ", err)
	}
	return film, nil
}

func (r *repository) DeleteFilmCategory(ctx context.Context, id film_category.UUID) error {
	q := `DELETE FROM film_category WHERE uuid = $1 `
	_, err := r.db.Exec(ctx, q, id.UUID)
	if err != nil {
		fmt.Println("Delete Film Category: ", err)
		return err
	}
	return nil
}
