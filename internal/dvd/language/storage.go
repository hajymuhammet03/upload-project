package language

import "context"

type Repository interface {
	AddLanguage(ctx context.Context, dto LanguageDTO) (UUID, error)
	GetLanguage(ctx context.Context, search string) ([]Language, error)
	GetLanguageID(ctx context.Context, id string) (UUID, error)
	DeleteLanguage(ctx context.Context, id string) error
}
