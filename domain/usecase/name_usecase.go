package usecase

import "context"

type NameUseCaseInterface interface {
	GetName(ctx context.Context, input string) (isValid bool, response string)
}
