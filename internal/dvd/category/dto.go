package category

import "time"

type AddCategory struct {
	UUID       string    `json:"uuid"`
	NameTm     string    `json:"name_tm"`
	NameEn     string    `json:"name_en"`
	NameRu     string    `json:"name_ru"`
	LastUpdate time.Time `json:"last_update"`
}

type PaginationDTO struct {
	Limit     int64  `json:"limit"`
	Page      int64  `json:"page"`
	Type      string `json:"type"`
	Search    string `json:"search"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
