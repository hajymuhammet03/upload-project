package category

import "time"

type AddCategory struct {
	UUID       string    `json:"uuid"`
	NameTm     string    `json:"name_tm"`
	NameEn     string    `json:"name_en"`
	NameRu     string    `json:"name_ru"`
	LastUpdate time.Time `json:"last_update"`
}
