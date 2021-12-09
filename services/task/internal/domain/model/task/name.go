package task

import (
	"github.com/sk-develop/motivation-management/shared/errors"
)

type Name string

func NewName(value string) (*Name, error) {
	if len(value) == 0 {
		err := errors.NewValidationError("Name is required")

		return nil, err
	}

	if len(value) >= maxLength {
		err := errors.NewValidationError("Name must be 255 characters or less")

		return nil, err
	}
	title := Name(value)

	return &title, nil
}
