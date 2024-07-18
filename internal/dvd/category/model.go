package category

import "time"

type UUID struct {
	UUID string `json:"uuid"`
}

type GetCategory struct {
	UUID       string    `json:"uuid"`
	NameTm     string    `json:"name_tm"`
	NameEn     string    `json:"name_en"`
	NameRu     string    `json:"name_ru"`
	LastUpdate time.Time `json:"last_update"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetCategoryResult struct {
	Category []GetCategory `json:"category"`
	Count    int64         `json:"count"`
}
