package film_category

type FilmCategoryReq struct {
	UUID       string `json:"uuid"`
	CategoryID string `json:"category_id"`
}

type PaginationDTO struct {
	Limit     int64  `json:"limit"`
	Page      int64  `json:"page"`
	Type      string `json:"type"`
	Search    string `json:"search"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
