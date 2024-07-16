package category

import "context"

type Repository interface {
	AddCategory(ctx context.Context, dto AddCategory) (UUID, error)
}
