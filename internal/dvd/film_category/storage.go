package film_category

import "context"

type Repository interface {
	AddFilmCategory(ctx context.Context, dto FilmCategoryReq) (UUID, error)
	GetFilmCategory(ctx context.Context, dto PaginationDTO) ([]GetFilmCategory, int64, error)
	GetFilmCategoryID(ctx context.Context, id UUID) (GetFilmCategory, error)
	DeleteFilmCategory(ctx context.Context, id UUID) error
}
