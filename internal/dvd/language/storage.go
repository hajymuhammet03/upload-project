package language

import "context"

type Repository interface {
	AddLanguage(ctx context.Context, dto LanguageDTO) (UUID, error)
}
