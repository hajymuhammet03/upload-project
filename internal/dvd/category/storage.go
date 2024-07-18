package category

import "context"

type Repository interface {
	AddCategory(ctx context.Context, category AddCategory) (UUID, error)
	GetCategory(ctx context.Context, dto PaginationDTO) ([]GetCategory, int64, error)
	GetCategoryID(ctx context.Context, id UUID) (GetCategory, error)
	DeleteCategory(ctx context.Context, id UUID) error
}
