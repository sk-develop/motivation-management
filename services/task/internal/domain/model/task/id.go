package task

import (
	"errors"
)

type ID int

func NewID(i int) (*ID, error) {
	if i < minLength {
		return nil, errors.New("ID must be an integer greater than or equal to 1")
	}
	id := ID(i)

	return &id, nil
}
