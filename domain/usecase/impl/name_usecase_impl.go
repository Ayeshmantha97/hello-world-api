package impl

import (
	"context"
	"hello-world-api/domain/usecase"
	"strings"
	"unicode"
)

type nameUseCase struct {
}

func NewNameUseCase() usecase.NameUseCaseInterface {
	return &nameUseCase{}
}

func (n *nameUseCase) GetName(ctx context.Context, input string) (isValid bool, response string) {
	if input == "" {
		return false, response
	}

	firstChar := rune(strings.ToUpper(input)[0])

	if !unicode.IsLetter(firstChar) {
		return false, response
	}

	if firstChar >= 'A' && firstChar <= 'M' {
		return true, "Hello " + input
	}

	return false, response
}
