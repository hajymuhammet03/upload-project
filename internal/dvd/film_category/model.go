package film_category

import "time"

type UUID struct {
	UUID string `json:"uuid"`
}

type GetFilmCategory struct {
	UUID       string    `json:"uuid"`
	CategoryID string    `json:"category_id"`
	LastUpdate time.Time `json:"last_update"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetFilmCategoryResult struct {
	FilmCategory []GetFilmCategory `json:"film_category"`
	Count        int64             `json:"count"`
}
