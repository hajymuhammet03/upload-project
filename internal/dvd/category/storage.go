package category

import "context"

type Repository interface {
	AddCategory(ctx context.Context, dto AddCategory) (UUID, error)
	GetCategory(ctx context.Context, dto PaginationDTO) ([]GetCategory, int64, error)
}
