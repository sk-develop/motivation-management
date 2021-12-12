package task

import (
	"github.com/sk-develop/motivation-management/shared/errors"
)

type Name string

func NewName(value string) (*Name, error) {
	if len(value) == 0 {
		return nil, errors.ValidationError("Name is required")
	}

	if len(value) > maxLength {
		return nil, errors.ValidationError("Name must be less than 255 characters")
	}
	title := Name(value)

	return &title, nil
}
