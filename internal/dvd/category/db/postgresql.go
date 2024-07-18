package db

import (
	"context"
	"fmt"
	"github.com/Hajymuhammet03/internal/dvd/category"
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
	var id category.UUID
	if dto.UUID == "" {
		q := `INSERT INTO category (name_tm, name_en, name_ru) VALUES ($1, $2, $3) RETURNING uuid`
		err := r.db.QueryRow(ctx, q, dto.NameTm, dto.NameEn, dto.NameRu).Scan(&id.UUID)
		if err != nil {
			fmt.Println("Add Category Error: ", err)
			return id, err
		}
	} else {
		q := `UPDATE category SET (name_tm, name_en, name_ru) = ($1, $2, $3) WHERE uuid = $4 RETURNING uuid`
		err := r.db.QueryRow(ctx, q, dto.NameTm, dto.NameEn, dto.NameRu, dto.UUID).Scan(&id.UUID)
		if err != nil {
			fmt.Println("Update Category Error: ", err)
			return id, err
		}
	}
	return id, nil
}

func (r *repository) GetCategory(ctx context.Context, dto category.PaginationDTO) ([]category.GetCategory, int64, error) {
	var (
		lang       string
		categories []category.GetCategory
		count      int64
	)

	switch dto.Type {
	case "TM":
		lang = "name_tm"
	case "EN":
		lang = "name_en"
	default:
		lang = "name_ru"
	}
	q := `
          SELECT
               c.uuid, c.name_tm, c.name_en, c.name_ru,
               c.last_update, c.created_at
          FROM category c
          WHERE ` + lang + ` ILIKE $1||'%'
         `

	if dto.StartDate != "0" && dto.StartDate != "" {
		q += `AND c.created_at >= '` + dto.StartDate + `'`
	}
	if dto.EndDate != "0" && dto.EndDate != "" {
		q += `AND c.created_at <= '` + dto.EndDate + `'`
	}

	q += `ORDER BY c.created_at LIMIT $2 OFFSET $3`

	row, err := r.db.Query(ctx, q, dto.Search, dto.Limit, dto.Limit*dto.Page)
	if err != nil {
		fmt.Println("Get Category Error: ", err)
	}

	defer row.Close()

	for row.Next() {
		var category category.GetCategory
		err := row.Scan(&category.UUID, &category.NameTm, &category.NameEn,
			&category.NameRu, &category.LastUpdate, &category.CreatedAt)
		if err != nil {
			fmt.Println("Get Category Error in for rows: ", err)
			return nil, 0, err
		}
		categories = append(categories, category)
	}

	q = `
        SELECT
             COUNT(c.uuid)
        FROM category c
        WHERE ` + lang + ` ILIKE $1||'%'
         `
	if dto.StartDate != "0" && dto.StartDate != "" {
		q += `AND c.created_at >= '` + dto.StartDate + `'`
	}
	if dto.EndDate != "0" && dto.EndDate != "" {
		q += `AND c.created_at <= '` + dto.EndDate + `'`
	}

	err = r.db.QueryRow(ctx, q, dto.Search).Scan(&count)
	if err != nil {
		fmt.Println("Get Category Count Error: ", err)
	}
	return categories, count, nil
}

func (r *repository) GetCategoryID(ctx context.Context, id category.UUID) (category.GetCategory, error) {
	var category category.GetCategory

	q := `
         SELECT 
              c.uuid, c.name_tm, c.name_en, c.name_ru,
              c.last_update, c.created_at
         FROM category c 
         WHERE c.uuid = $1
        `

	err := r.db.QueryRow(ctx, q, id.UUID).Scan(&category.UUID, &category.NameTm, &category.NameEn,
		&category.NameRu, &category.LastUpdate, &category.CreatedAt)

	if err != nil {
		fmt.Println("Get Category ID Error: ", err)
	}
	return category, nil
}

func (r *repository) DeleteCategory(ctx context.Context, id category.UUID) error {
	q := ` DELETE FROM category WHERE uuid = $1 `
	_, err := r.db.Exec(ctx, q, id.UUID)
	if err != nil {
		fmt.Println("Delete Category Error: ", err)
		return err
	}
	return nil
}
