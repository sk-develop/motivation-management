package task

import (
	"errors"
)

type Name string

func NewName(value string) (*Name, error) {
	if len(value) == 0 {
		return nil, errors.New("Name is required")
	}

	if len(value) > maxLength {
		return nil, errors.New("Name must be less than 255 characters")
	}
	title := Name(value)

	return &title, nil
}
